# Rocking Kubernetes Security: Safeguarding Hundreds of Clusters in a Dynamic Multi-Cloud Environment

Having an external Identity Provider (IdP) not only offers a broader range of security features, including Multi-Factor Authentication (MFA), but also enables us to maintain a provider-agnostic cloud architecture. This, in turn, empowers us to automate large-scale processes. In this context, we will outline the necessary configurations within the cluster to integrate with an external IdP. Additionally, we will demonstrate how to leverage this integration to automate the process of cluster joining with ArgoCD. Following the cluster joining process, we will proceed to deploy security tools within the cluster for enhanced security.

## Lab architecture

![lab](https://github.com/facundoalarcon/kubernetes-security/blob/main/img/lab.png)