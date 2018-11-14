package payu

import (
	"net/http"
)

type buildInClient struct {
	accessToken string
	client      *http.Client
}

func NewHttpClient(accessToken string) httpClientInterface {
	client := &http.Client{}
	return &buildInClient{accessToken: accessToken, client: client}
}

func (client *buildInClient) SetAuthorization(accessToken string) {
	client.accessToken = accessToken
}

func (client buildInClient) DoRequest(request requestInterface) interface{} {
	return nil
}
