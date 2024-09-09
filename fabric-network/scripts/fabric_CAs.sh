#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

function launch_ECert_CAs() {
  push_fn "Launching Fabric CAs"

  apply_template manifests/sarb/sarb-ca.yaml $SARB_NS
  apply_template manifests/absa/absa-ca.yaml $ABSA_NS
  apply_template manifests/crossborder/crossborder-ca.yaml $CROSSBORDER_NS

  kubectl -n $SARB_NS rollout status deploy/sarb-ca
  kubectl -n $ABSA_NS rollout status deploy/absa-ca
  kubectl -n $CROSSBORDER_NS rollout status deploy/crossborder-ca

  # todo: this papers over a nasty bug whereby the CAs are ready, but sporadically refuse connections after a down / up
  sleep 5

  pop_fn
}

# experimental: create TLS CA issuers using cert-manager for each org.
function init_tls_cert_issuers() {
  push_fn "Initializing TLS certificate Issuers"

  # Create a self-signing certificate issuer / root TLS certificate for the blockchain.
  # TODO : Bring-Your-Own-Key - allow the network bootstrap to read an optional ECDSA key pair for the TLS trust root CA.
  kubectl -n $SARB_NS apply -f manifests/root-tls-cert-issuer.yaml
  kubectl -n $SARB_NS wait --timeout=30s --for=condition=Ready issuer/root-tls-cert-issuer
  kubectl -n $ABSA_NS apply -f manifests/root-tls-cert-issuer.yaml
  kubectl -n $ABSA_NS wait --timeout=30s --for=condition=Ready issuer/root-tls-cert-issuer
  kubectl -n $CROSSBORDER_NS apply -f manifests/root-tls-cert-issuer.yaml
  kubectl -n $CROSSBORDER_NS wait --timeout=30s --for=condition=Ready issuer/root-tls-cert-issuer

  # Use the self-signing issuer to generate three Issuers, one for each org.
  kubectl -n $SARB_NS apply -f manifests/sarb/sarb-tls-cert-issuer.yaml
  kubectl -n $ABSA_NS apply -f manifests/absa/absa-tls-cert-issuer.yaml
  kubectl -n $CROSSBORDER_NS apply -f manifests/crossborder/crossborder-tls-cert-issuer.yaml

  kubectl -n $SARB_NS wait --timeout=30s --for=condition=Ready issuer/sarb-tls-cert-issuer
  kubectl -n $ABSA_NS wait --timeout=30s --for=condition=Ready issuer/absa-tls-cert-issuer
  kubectl -n $CROSSBORDER_NS wait --timeout=30s --for=condition=Ready issuer/crossborder-tls-cert-issuer

  pop_fn
}

function enroll_bootstrap_ECert_CA_user() {
  local org=$1
  local ns=$2

  # Determine the CA information and TLS certificate
  CA_NAME=${org}-ca
  CA_DIR=${TEMP_DIR}/cas/${CA_NAME}
  mkdir -p ${CA_DIR}

  # Read the CA's TLS certificate from the cert-manager CA secret
  echo "retrieving ${CA_NAME} TLS root cert"
  kubectl -n $ns get secret ${CA_NAME}-tls-cert -o json \
    | jq -r .data.\"ca.crt\" \
    | base64 -d \
    > ${CA_DIR}/tlsca-cert.pem

  # Enroll the root CA user
  fabric-ca-client enroll \
    --url https://${RCAADMIN_USER}:${RCAADMIN_PASS}@${CA_NAME}.${DOMAIN}:${NGINX_HTTPS_PORT} \
    --tls.certfiles $TEMP_DIR/cas/${CA_NAME}/tlsca-cert.pem \
    --mspdir $TEMP_DIR/enrollments/${org}/users/${RCAADMIN_USER}/msp
}

function enroll_bootstrap_ECert_CA_users() {
  push_fn "Enrolling bootstrap ECert CA users"

  enroll_bootstrap_ECert_CA_user sarb $SARB_NS
  enroll_bootstrap_ECert_CA_user absa $ABSA_NS
  enroll_bootstrap_ECert_CA_user crossborder $CROSSBORDER_NS

  pop_fn
}
