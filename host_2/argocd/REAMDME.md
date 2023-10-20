# ArgoCD Host
## Setup minikube as host_1
Setup minikube as host_1
## Deploy ArgoCD
**[ArgoCD](https://argo-cd.readthedocs.io/en/stable/getting_started/)** is an open-source declarative, GitOps continuous delivery tool for Kubernetes. It is designed to help developers and operations teams manage and automate the deployment and configuration of applications on Kubernetes clusters. ArgoCD is particularly well-suited for cloud-native and containerized applications.

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# modify load balancer to ingress
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

# get initial password
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

# use minikube tunnel
minikube tunnel
```

## Extract argocd crt and key
```bash
kubectl get secret  -n argocd
kubectl get secret argocd-secret -n argocd -o jsonpath="{.data['tls\.crt']}" | base64 -d > argocd.crt

kubectl get secret argocd-secret -n argocd -o jsonpath="{.data['tls\.key']}" | base64 -d > argocd.key
```
## Check the enpoint and remplace in nginx config file
```bash
export PROXY_IP=$(minikube service -n argocd argocd-server --url | head -1)
echo $PROXY_IP # use this IP and port in Nginx
```
## Run container for reverse proxy in the same network of minikube
```bash
docker network ls

# 443 is the port defined on nginx.conf, use -dif you want run this command as background background
ARGOCD_PORT=${ARGOCD_PORT=443}

docker run  --network [minikube_network_id] \
--name nginx \
-p $ARGOCD_PORT:$ARGOCD_PORT \
-v ./argocd.key:/etc/nginx/certs/minikube-client.key \
-v ./argocd.crt:/etc/nginx/certs/minikube-client.crt \
-v ./nginx.conf:/etc/nginx/nginx.conf \
nginx:alpine3.18
```
## Setup your hosts file

in `/etc/hosts`
```
${HOST_KEYCLOAK_MINIKUBE}   keycloak.local minikube-user
${HOST_IP}  argocd.local
```