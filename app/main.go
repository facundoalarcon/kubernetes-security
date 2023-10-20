package main

import (
    "context"
    "fmt"
    "github.com/Nerzal/gocloak/v13"
    "github.com/facundoalarcon/kubernetes-security/services"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Could not load .env file")
        return
    }

    keycloakURL := os.Getenv("IDP_URL")
    realmName := os.Getenv("REALM")
    clientID := os.Getenv("CLIENT_ID")
    username := os.Getenv("IDP_USERNAME")
    password := os.Getenv("IDP_PASSWORD")

    // GoCloak Client
    client := gocloak.NewClient(keycloakURL)

    // KeyCloak Auth
    token, err := client.Login(
        context.Background(),
        clientID,
        "",
        realmName,
        username,
        password,
    )
    if err != nil {
        fmt.Printf("KeyCloak: Auth Error: %v\n", err)
        return
    }

    // user Token
    //fmt.Printf("token: %s\n", token.IDToken)

    argocdURL := os.Getenv("ARGOCD_URL")
    argocdUser := os.Getenv("ARGOCD_USERNAME")
    argocdPassword := os.Getenv("ARGOCD_PASSWORD")
    clusterApi := os.Getenv("CLUSTER_API")
    clusterArgoName := os.Getenv("CLUSTER_NAME")
    clusterCA := os.Getenv("B64_CLUSTER_CA")

    // argocd basic auth
    argocdClient, err := services.NewArgoCDClient(argocdURL, argocdUser, argocdPassword)
    if err != nil {
        fmt.Printf("Argocd: auth error: %v\n", err)
        return
    }
    err = argocdClient.AddCluster(clusterApi, clusterArgoName, token.IDToken, clusterCA)
    if err != nil {
        fmt.Printf("Argocd: could not join cluster: %v\n", err)
        return
    }

    fmt.Println("Cluster Joined!")
}
