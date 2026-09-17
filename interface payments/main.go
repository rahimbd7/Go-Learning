package main
import (
	"fmt"
	
)

type PaymentGateway interface {
	checkout(amount float64) error
}

/*Stripe payment gateway*/
type Stripe struct{
	baseURL string
	apiKey string
}

func (s *Stripe) checkout(amount float64) error {
	fmt.Println("Processing payment of", amount, "through Stripe")
	return nil
}


/*PayPal payment gateway*/
type PayPal struct{
	baseURL string
	apiKey string
}
func (p *PayPal) checkout(amount float64) error {
	fmt.Println("Processing payment of", amount, "through PayPal")
	return nil
}

func newPayPal(apikey string, baseURL string) *PayPal {
	return &PayPal{
		apiKey: apikey,
		baseURL: baseURL,
	}
}
func processPayment(gateway PaymentGateway, amount float64) {
	process := gateway.checkout(amount)
	if process != nil {
		fmt.Println("Payment failed:", process)
	}
	fmt.Println("Payment successfully...!")
}

func main() {
	payPal := newPayPal("sk1234343kdjfldsjf", "https://api.paypal.com")
	processPayment(payPal, 100.0);

}
