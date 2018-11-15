package payu

type requestInterface interface {
	getPath() string
	getData() []byte
	getMethod() string
}
