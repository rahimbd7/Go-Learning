package PaymentChannel
import (
	"fmt"
	// "payment/PaymentGateWay"
)



/*Stripe payment gateway*/
type Stripe struct{
	baseURL string
	apiKey string
}

func (s *Stripe) Checkout(amount float64) error {
	fmt.Println("Stripe Processing payment of", amount, "through Stripe")
	return nil
}

func NewStripe(apikey string, baseURL string) *Stripe {
	return &Stripe{
		apiKey: apikey,
		baseURL: baseURL,
	}
}


/*PayPal payment gateway*/
type PayPal struct{
	baseURL string
	apiKey string
}
func (p *PayPal) Checkout(amount float64) error {
	fmt.Println("Paypal Processing payment of", amount, "through PayPal")
	return nil
}

func NewPayPal(apikey string, baseURL string) *PayPal {
	return &PayPal{
		apiKey: apikey,
		baseURL: baseURL,
	}
}
