package payu

type requestInterface interface {
	getPath() string
	getData() interface{}
	getMethod() string
}
