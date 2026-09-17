package main

//module name: "payment" which has defined in the go.mod file. so all the packages are imported with the module name "payment"
import "payment/PaymentGateWay"
import "payment/PaymentChannel"
import "payment/Mock"

func main() {
	payPal := PaymentChannel.NewPayPal("sk1234343kdjfldsjf", "https://api.paypal.com")
	PaymentGateway.ProcessPayment(payPal, 100.0)

	stripe := PaymentChannel.NewStripe("sk1234343kdjfldsjf", "https://api.stripe.com")
	PaymentGateway.ProcessPayment(stripe, 200.0)


	// Mock payment gateway
	mock := Mock.NewMock("mockapikey", "https://mockapi.com")
	PaymentGateway.ProcessPayment(mock, 50.0)

}
