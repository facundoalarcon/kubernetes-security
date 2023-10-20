# Testing The Config
To confirm that the OPA policy that you implemented works, you can test it out with the two Kubernetes Manifests below.

The Manifest with the latest tag won’t work because you created a policy in the previous step to ensure that latest tags cannot be used. The deployment itself will deploy, but the Pods won’t come online.

The Manifest below will work and the Pods will come online because a container image version number is specified.
```bash
# FAILS
kubectl apply -f bad-policy.yaml
# DEPLOY
kubectl apply -f policy-ok.yaml
```

# Policies description
* *bad-policy.yaml missconfigurations:* 
Result
```
Warning: would violate PodSecurity "restricted:latest": allowPrivilegeEscalation != false (container "nginxdeployment" must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (container "nginxdeployment" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or container "nginxdeployment" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "nginxdeployment" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")
```
The error message indicates that the nginx-deployment deployment violates the following security policies:

* **allowPrivilegeEscalation != false**: This means that the policy requires securityContext.allowPrivilegeEscalation to be set to false to prevent privilege escalation in containers. To resolve this, you need to add a securityContext section in the container and configure allowPrivilegeEscalation: false.

* **unrestricted capabilities** (container "nginxdeployment" must set securityContext.capabilities.drop=["ALL"]): This policy demands that securityContext.capabilities.drop be configured to reduce the container's privileges. You should add securityContext and configure capabilities.drop in the container to comply with this policy.

* **runAsNonRoot != true**: This indicates that securityContext.runAsNonRoot is expected to be set to true to ensure that the container runs as a non-root user. You should configure runAsNonRoot: true in the container.

* **seccompProfile (pod or container "nginxdeployment" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")**: This pertains to Seccomp profile settings. You need to configure securityContext.seccompProfile.type in the container to be either "RuntimeDefault" or "Localhost" to comply with this requirement."

However, it's important to note that in this specific case, the *nginx:1.23.1* container is intended to run with root privileges. This is because it requires those privileges for certain tasks. Therefore, although Gatekeeper's security policy is configured to limit various security aspects, such as privilege escalation and container capabilities, it's necessary for the container to run with root privileges.

That said, to address Gatekeeper's security policy while still allowing the *nginx:1.23.1* container to operate as intended, you may consider using an alternative container image, such as *nginxinc/nginx-unprivileged:1.24*. This image is designed to run without requiring root privileges. By doing so, you can maintain compliance with Gatekeeper's security policies without encountering container crashes due to security restrictions.



# Resources
* [blog](https://dev.to/thenjdevopsguy/writing-your-first-kubernetes-opa-policy-with-gatekeeper-145a)
