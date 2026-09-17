package PaymentGateway
import "fmt"


type PaymentGateway interface {
	Checkout(amount float64) error
}

func ProcessPayment(gateway PaymentGateway, amount float64) {
	process := gateway.Checkout(amount)
	if process != nil {
		fmt.Println("Payment failed:", process)
	}
	fmt.Println("Payment successfully...!")
}
