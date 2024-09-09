#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

function launch_orderers() {
  push_fn "Launching orderers"

  apply_template manifests/sarb/sarb-orderer1.yaml $SARB_NS
  apply_template manifests/sarb/sarb-orderer2.yaml $SARB_NS
  apply_template manifests/sarb/sarb-orderer3.yaml $SARB_NS

  kubectl -n $SARB_NS rollout status deploy/sarb-orderer1
  kubectl -n $SARB_NS rollout status deploy/sarb-orderer2
  kubectl -n $SARB_NS rollout status deploy/sarb-orderer3

  pop_fn
}

function launch_peers() {
  push_fn "Launching peers"

  apply_template manifests/absa/absa-peer1.yaml $ABSA_NS
  apply_template manifests/absa/absa-peer2.yaml $ABSA_NS
  apply_template manifests/crossborder/crossborder-peer1.yaml $CROSSBORDER_NS
  apply_template manifests/crossborder/crossborder-peer2.yaml $CROSSBORDER_NS

  kubectl -n $ABSA_NS rollout status deploy/absa-peer1
  kubectl -n $ABSA_NS rollout status deploy/absa-peer2
  kubectl -n $CROSSBORDER_NS rollout status deploy/crossborder-peer1
  kubectl -n $CROSSBORDER_NS rollout status deploy/crossborder-peer2

  pop_fn
}

# Each network node needs a registration, enrollment, and MSP config.yaml
function create_node_local_MSP() {
  local node_type=$1
  local org=$2
  local node=$3
  local csr_hosts=$4
  local ns=$5
  local id_name=${org}-${node}
  local id_secret=${node_type}pw
  local ca_name=${org}-ca

  # Register the node admin
  rc=0
  fabric-ca-client  register \
    --id.name       ${id_name} \
    --id.secret     ${id_secret} \
    --id.type       ${node_type} \
    --url           https://${ca_name}.${DOMAIN}:${NGINX_HTTPS_PORT} \
    --tls.certfiles $TEMP_DIR/cas/${ca_name}/tlsca-cert.pem \
    --mspdir        $TEMP_DIR/enrollments/${org}/users/${RCAADMIN_USER}/msp \
    || rc=$?        # trap error code from registration without exiting the network driver script"

  if [ $rc -eq 1 ]; then
    echo "CA admin was (probably) previously registered - continuing"
  fi

  # Enroll the node admin user from within k8s.  This will leave the certificates available on a volume share in the
  # cluster for access by the nodes when launching in a container.
  cat <<EOF | kubectl -n ${ns} exec deploy/${ca_name} -i -- /bin/sh

  set -x
  export FABRIC_CA_CLIENT_HOME=/var/hyperledger/fabric-ca-client
  export FABRIC_CA_CLIENT_TLS_CERTFILES=/var/hyperledger/fabric/config/tls/ca.crt

  fabric-ca-client enroll \
    --url https://${id_name}:${id_secret}@${ca_name} \
    --csr.hosts ${csr_hosts} \
    --mspdir /var/hyperledger/fabric/organizations/${node_type}Organizations/${org}.example.com/${node_type}s/${id_name}.${org}.example.com/msp

  # Create local MSP config.yaml
  echo "NodeOUs:
    Enable: true
    ClientOUIdentifier:
      Certificate: cacerts/${org}-ca.pem
      OrganizationalUnitIdentifier: client
    PeerOUIdentifier:
      Certificate: cacerts/${org}-ca.pem
      OrganizationalUnitIdentifier: peer
    AdminOUIdentifier:
      Certificate: cacerts/${org}-ca.pem
      OrganizationalUnitIdentifier: admin
    OrdererOUIdentifier:
      Certificate: cacerts/${org}-ca.pem
      OrganizationalUnitIdentifier: orderer" > /var/hyperledger/fabric/organizations/${node_type}Organizations/${org}.example.com/${node_type}s/${id_name}.${org}.example.com/msp/config.yaml
EOF
}

function create_orderer_local_MSP() {
  local org=$1
  local orderer=$2
  local csr_hosts=${org}-${orderer}

  create_node_local_MSP orderer $org $orderer $csr_hosts $SARB_NS
}

function create_peer_local_MSP() {
  local org=$1
  local peer=$2
  local ns=$3
  local csr_hosts=localhost,${org}-${peer},${org}-peer-gateway-svc

  create_node_local_MSP peer $org $peer $csr_hosts ${ns}
}

function create_local_MSP() {
  push_fn "Creating local node MSP"

  create_orderer_local_MSP sarb orderer1
  create_orderer_local_MSP sarb orderer2
  create_orderer_local_MSP sarb orderer3

  create_peer_local_MSP absa peer1 $ABSA_NS
  create_peer_local_MSP absa peer2 $ABSA_NS

  create_peer_local_MSP crossborder peer1 $CROSSBORDER_NS
  create_peer_local_MSP crossborder peer2 $CROSSBORDER_NS

  pop_fn
}

function network_up() {

  # Kube config
  init_namespace
  init_storage_volumes
  load_org_config

  # Service account permissions for the k8s builder
  if [ "${CHAINCODE_BUILDER}" == "k8s" ]; then
    apply_k8s_builder_roles
    apply_k8s_builders
  fi

  # Network TLS CAs
  init_tls_cert_issuers

  # Network ECert CAs
  launch_ECert_CAs
  enroll_bootstrap_ECert_CA_users

  # Test Network
  create_local_MSP

  launch_orderers
  launch_peers
}

function stop_services() {
  push_fn "Stopping Fabric services"
  for ns in $SARB_NS $ABSA_NS $CROSSBORDER_NS; do
    kubectl -n $ns delete ingress --all
    kubectl -n $ns delete deployment --all
    kubectl -n $ns delete pod --all
    kubectl -n $ns delete service --all
    kubectl -n $ns delete configmap --all
    kubectl -n $ns delete cert --all
    kubectl -n $ns delete issuer --all
    kubectl -n $ns delete secret --all
  done

  pop_fn
}

function scrub_org_volumes() {
  push_fn "Scrubbing Fabric volumes"
  for org in sarb absa crossborder; do
    # clean job to make this function can be rerun
    local namespace_variable=${org^^}_NS
    kubectl -n ${!namespace_variable} delete jobs --all

    # scrub all pv contents
    kubectl -n ${!namespace_variable} create -f manifests/${org}/${org}-job-scrub-fabric-volumes.yaml
    kubectl -n ${!namespace_variable} wait --for=condition=complete --timeout=60s job/job-scrub-fabric-volumes
    kubectl -n ${!namespace_variable} delete jobs --all
  done
  pop_fn
}

function network_down() {

  set +e
  for ns in $SARB_NS $ABSA_NS $CROSSBORDER_NS; do
    kubectl get namespace $ns > /dev/null
    if [[ $? -ne 0 ]]; then
      echo "No namespace $ns found - nothing to do."
      return
    fi
  done
  set -e

  stop_services
  scrub_org_volumes

  delete_namespace

  rm -rf $PWD/build
}
