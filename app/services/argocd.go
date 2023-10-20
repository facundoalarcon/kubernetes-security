package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/facundoalarcon/kubernetes-security/models"
	"io"
	"net/http"
)


type ArgoCDClient struct {
	Url string `json:"domain"`
	Token string `json:"token"`
}

func NewArgoCDClient(argoUrl, username, password string) (cli ArgoCDClient, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	sessionRequest := models.SessionRequest{
		Username: username,
		Password: password,
	}
	url := fmt.Sprintf("%s/api/v1/session", argoUrl)
	fmt.Println(url)
	body, err := json.Marshal(sessionRequest)
	if err != nil {
		fmt.Println("Error on parse")
		return cli, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error on request")
		return cli, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error on request, expected 200 got %v\n",resp.StatusCode )
		return  cli, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not conver body in bytes", err)
		return cli, err
	}

	err = json.Unmarshal(bodyBytes, &cli)
	if err != nil {
		fmt.Println("Error in read body from response", err)
		return cli, err
	}

	cli = ArgoCDClient{
		Url: argoUrl,
		Token: cli.Token,
	}

	return cli, err
}

func (argocd ArgoCDClient) AddCluster(clusterApi, clusterArgoName, IdPToken, CAClusterData string) (err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	addClusterReq := models.AddClusterRequest{
		Server: clusterApi,
		Name:   clusterArgoName,
		Config: models.ClusterConfig{
			BearerToken: IdPToken,
			TLSClientConfig: models.TLSClientConfig{
				Insecure: false,
				CAData:   CAClusterData,
			},
		},
	}

	url := fmt.Sprintf("%s/api/v1/clusters?upsert=true",  argocd.Url)
	fmt.Println(url)
	body, err := json.Marshal(addClusterReq)
	if err != nil {
		fmt.Println("Error on parse")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error on build request")
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cookie", fmt.Sprintf("argocd.token=%s", argocd.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error on request, expected 200 got %v\n",resp.StatusCode )
		return  err
	}

	return err
}