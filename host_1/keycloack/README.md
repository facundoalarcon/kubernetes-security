# Deploy Keycloak
Keycloak is an open-source solution for identity and access management, used to add authentication, authorization, and user management capabilities to applications and services. It provides features such as user authentication, role-based authorization, two-factor authentication, and support for security standards like OAuth 2.0 and OpenID Connect.

You can deploy Keycloak in a Docker container to simplify its management and scalability. This allows you to run Keycloak in an isolated environment and make it easier to manage, ensuring that it's available for your applications and services that require authentication and authorization.

## Steps
1. run ./cert/createCA.sh file
2. trust the CA
3. run ./keycloak.sh

## Relateds links
* [dex](https://github.com/konveyor/minikube_dex_oidc_example)