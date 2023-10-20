#!/bin/bash

cp certs_minikube/ssl/ca.pem /usr/local/share/ca-certificates/

update-ca-certificates
