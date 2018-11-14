package payu

type httpClientInterface interface {
	DoRequest(request requestInterface) interface{}
	SetAuthorization(accessToken string)
}
