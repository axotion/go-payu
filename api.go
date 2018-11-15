package payu

const (
	production = "https://secure.payu.com"
	sandbox    = "https://secure.snd.payu.com"
)

type api struct {
	production   bool
	clientID     int
	clientSecret string
	httpClient   httpClientInterface
	accessToken  string
}

// New instance of PayU api...
func New(clientID int, clientSecret string, production bool) *api {
	httpClient := &buildInClient{}
	return &api{production: production, httpClient: httpClient, clientID: clientID, clientSecret: clientSecret}
}

func (api *api) Authorize() error {

	obtainJwt := "dsds"

	api.httpClient.SetAuthorization(obtainJwt)
	return nil
}

func (api *api) CreateOrder(orderDto createOrderDto) (string, error) {
	return "", nil
}

func (api *api) RefundAmount(externalOrderID string, amount int) error {
	return nil
}

func (api *api) ValidateWebhook(data map[string]string) error {
	return nil
}
