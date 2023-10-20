# Cluster Joiner

This is an application that joins clusters to ArgoCD

Create an .env file with the following data:
```
IDP_URL=idp_url
REALM=ream
CLIENT_ID=clientid
IDP_USERNAME=user
IDP_PASSWORD=password
ARGOCD_URL=argo_url
ARGOCD_USERNAME=argousr
ARGOCD_PASSWORD=argopassw
CLUSTER_API=remote_cluster_url
CLUSTER_NAME=argo_cluster_name
B64_CLUSTER_CA=b64_cluster_ca
```

Add the client secret if it is required for the application, or if [PKCE](https://oauth.net/2/pkce/) is not used.

## execution
```sh
go mod init xxx # or your repo
go mod tidy
go run ./...
```