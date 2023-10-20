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