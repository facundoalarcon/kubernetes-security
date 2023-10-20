# Exposing an Application with Nginx in Minikube and Connecting a Keycloak Container to the Minikube Network

**[Minikube](https://minikube.sigs.k8s.io/docs/start/)** is a tool that allows you to set up and run a single-node Kubernetes cluster on a local machine. It's useful for developing and testing applications in local Kubernetes environments.

To expose an application through an Nginx container in Minikube, you should follow these steps:

1. Start Minikube using the `minikube.sh` script.

2. Configure Nginx to redirect traffic to Minikube using a script named `nginx/setup.sh`.

3. Deploy an Nginx container within Minikube using a script named `nginx/run.sh`.

To connect a Keycloak container to the Minikube network, you need to ensure that the container runs within the same Docker network namespace as Minikube. This allows them to communicate at the network level within Docker.

The **Docker network** is a component that facilitates communication between containers within a Docker system. By default, each container is connected to a Docker bridge network, enabling them to communicate with one another. Connecting a Keycloak container to the Minikube network means that containers in both Minikube and Keycloak can exchange data through the Docker network. This is particularly useful for distributed applications and microservices systems.

## After scripts executions

```
# check containers
docker ps

# connect nginx in minikube networks
kubectl networks ls # check the minikube network id

docker network connect [minikube_network_id] [nginx_container_id]
```

## Cluster Authentication
[kubelogin](https://github.com/int128/kubelogin) is a kubectl plugin for Kubernetes OpenID Connect (OIDC) authentication, also known as kubectl oidc-login. With this plugin, we will use the [PKCE flow](https://oauth.net/2/pkce/), a secure authentication method recommended for distributed clients.

```bash
kubectl oidc-login setup \
--oidc-issuer-url=$YOUR_ISSUER_URL\
--oidc-client-id=$YOUR_CLIENT_ID \
--oidc-use-pkce

kubectl config set-credentials oidc \
          --exec-api-version=client.authentication.k8s.io/v1beta1 \
          --exec-command=kubectl \
          --exec-arg=oidc-login \
          --exec-arg=get-token \
          --exec-arg=--oidc-issuer-url=$YOUR_ISSUER_URL \
          --exec-arg=--oidc-client-id=$YOUR_CLIENT_ID \
          --exec-arg=--oidc-use-pkce

kubectl create clusterrolebinding oidc-cluster-admin --clusterrole=cluster-admin --group=$YOUR_GROUP

kubectl --user=oidc get nodes

# if you want
kubectl config set-context --current --user=oidc

```