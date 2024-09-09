#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

function init_namespace() {
  local namespaces=$(echo "$SARB_NS $ABSA_NS $CROSSBORDER_NS" | xargs -n1 | sort -u)
  for ns in $namespaces; do
    push_fn "Creating namespace \"$ns\""
    kubectl create namespace $ns || true
    pop_fn
  done
}

function delete_namespace() {
  local namespaces=$(echo "$SARB_NS $ABSA_NS $CROSSBORDER_NS" | xargs -n1 | sort -u)
  for ns in $namespaces; do
    push_fn "Deleting namespace \"$ns\""
    kubectl delete namespace $ns || true
    pop_fn
  done
}

function init_storage_volumes() {
  push_fn "Provisioning volume storage"

  # Both KIND and k3s use the Rancher local-path provider.  In KIND, this is installed
  # as the 'standard' storage class, and in Rancher as the 'local-path' storage class.
  if [ "${CLUSTER_RUNTIME}" == "kind" ]; then
    export STORAGE_CLASS="standard"

  elif [ "${CLUSTER_RUNTIME}" == "k3s" ]; then
    export STORAGE_CLASS="local-path"

  else
    echo "Unknown CLUSTER_RUNTIME ${CLUSTER_RUNTIME}"
    exit 1
  fi

  cat manifests/pvc-fabric-sarb.yaml | envsubst | kubectl -n $SARB_NS create -f - || true
  cat manifests/pvc-fabric-absa.yaml | envsubst | kubectl -n $ABSA_NS create -f - || true
  cat manifests/pvc-fabric-crossborder.yaml | envsubst | kubectl -n $CROSSBORDER_NS create -f - || true

  pop_fn
}

function load_org_config() {
  push_fn "Creating fabric config maps"

  kubectl -n $SARB_NS delete configmap sarb-config || true
  kubectl -n $ABSA_NS delete configmap absa-config || true
  kubectl -n $CROSSBORDER_NS delete configmap crossborder-config || true

  kubectl -n $SARB_NS create configmap sarb-config --from-file=config/sarb
  kubectl -n $ABSA_NS create configmap absa-config --from-file=config/absa
  kubectl -n $CROSSBORDER_NS create configmap crossborder-config --from-file=config/crossborder

  pop_fn
}

function apply_k8s_builder_roles() {
  push_fn "Applying k8s chaincode builder roles"

  apply_template manifests/fabric-builder-role.yaml $ABSA_NS
  apply_template manifests/fabric-builder-rolebinding.yaml $ABSA_NS

  pop_fn
}

function apply_k8s_builders() {
  push_fn "Installing k8s chaincode builders"

  apply_template manifests/absa/absa-install-k8s-builder.yaml $ABSA_NS
  apply_template manifests/crossborder/crossborder-install-k8s-builder.yaml $CROSSBORDER_NS

  kubectl -n $ABSA_NS wait --for=condition=complete --timeout=60s job/absa-install-k8s-builder
  kubectl -n $CROSSBORDER_NS wait --for=condition=complete --timeout=60s job/crossborder-install-k8s-builder

  pop_fn
}
