#!/bin/sh
# Before run:
# 1. run the script inside ./cert folder
# 2. define HOST_IP,HOST_ARGO,YOUR_ISSUER_URL,YOUR_USERNAME_CLAIM,YOUR_CLIENT_ID,YOUR_GROUP_CLAIM env vars
# 3. trust in the CA

VM_DRIVER="${VM_DRIVER:=docker}" 

mkdir -p ~/.minikube/files/var/lib/minikube/certs/ && \
cp -a ./cert/ssl/* ~/.minikube/files/var/lib/minikube/certs/

mkdir -p ~/.minikube/files/etc
cat << EOF > ~/.minikube/files/etc/hosts 
127.0.0.1       localhost
255.255.255.255	broadcasthost
::1             localhost
${HOST_IP}   keycloak.local minikube-user
${HOST_ARGO}  argocd.local
EOF


minikube start --v=5  --apiserver-port="${MINIKUBE_PORT}" \
--extra-config=apiserver.authorization-mode=Node,RBAC \
--extra-config=apiserver.oidc-issuer-url="${YOUR_ISSUER_URL}" \
--extra-config=apiserver.oidc-username-claim="${YOUR_USERNAME_CLAIM}" \
--extra-config=apiserver.oidc-ca-file=/var/lib/minikube/certs/ca.pem \
--extra-config=apiserver.oidc-client-id="${YOUR_CLIENT_ID}" \
--extra-config=apiserver.oidc-groups-claim="${YOUR_GROUP_CLAIM}"

export MINIKUBE_IP=$(minikube ip)

if [[ -z "${MINIKUBE_IP}" ]]; then
  echo "Please ensure you have env var 'MY_MINIKUBE_IP' defined before running."
  echo "Perhaps you forgot to source your SECRET.source_me file"
  exit 1
fi
