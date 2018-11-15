package payu

import (
	"bytes"
	"net/http"
	"time"
)

type buildInClient struct {
	accessToken string
	client      *http.Client
}

func NewHttpClient(accessToken string) httpClientInterface {
	client := &http.Client{Timeout: time.Second * 10}
	return &buildInClient{accessToken: accessToken, client: client}
}

func (client *buildInClient) SetAuthorization(accessToken string) {
	client.accessToken = accessToken
}

func (client buildInClient) DoRequest(request requestInterface) interface{} {

	httpRequest, err := http.NewRequest(request.getMethod(), request.getPath(), bytes.NewBuffer(request.getData()))

	if err != nil {
		panic(err)
	}

	resp, err := client.client.Do(httpRequest)

	return nil
}
