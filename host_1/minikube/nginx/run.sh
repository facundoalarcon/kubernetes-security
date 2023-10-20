#!/bin/bash

docker run \
--name nginx \
-p ${MINIKUBE_IP}:${MINIKUBE_IP} \
-v ~/.minikube/profiles/minikube/client.key:/etc/nginx/certs/minikube-client.key \
-v ~/.minikube/profiles/minikube/client.crt:/etc/nginx/certs/minikube-client.crt \
-v ./nginx.conf:/etc/nginx/nginx.conf \
nginx:alpine3.18
