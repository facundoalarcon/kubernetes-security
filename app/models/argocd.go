package models

type SessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddClusterRequest struct {
	Server string         `json:"server"`
	Name   string         `json:"name"`
	Config ClusterConfig `json:"config"`
}

type ClusterConfig struct {
	BearerToken     string          `json:"bearerToken"`
	TLSClientConfig TLSClientConfig `json:"tlsClientConfig"`
}

type TLSClientConfig struct {
	Insecure bool   `json:"insecure"`
	CAData   string `json:"caData"`
}

