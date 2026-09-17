package Mock

import "fmt"
// import "payment/PaymentGateWay"
	
type Mock struct {
	baseURL string
	apiKey  string
}

func (m *Mock) Checkout(amount float64) error {
	fmt.Println("Mock payment succesfull ");
	return nil
}

func NewMock(apikey string, baseURL string) *Mock {
	return &Mock{
		apiKey:  apikey,
		baseURL: baseURL,
	}
}


