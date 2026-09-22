package main



func main() {
	// buffered channel
	unbufferedChannel := make(chan string, 3)

	func() {
		fileStatus := "Sending 1 to unbuffered channel"
		unbufferedChannel <- fileStatus
	}()
	func() {
		fileStatus := "Sending 2 to unbuffered channel"
		unbufferedChannel <- fileStatus
	}()
	func() {
		fileStatus := "Sending 3 to unbuffered channel"
		unbufferedChannel <- fileStatus
	}()

	//deadlock will occur as no one is receiving from the channel in 4th iteration
	for range 4 {
		fileStatus := <-unbufferedChannel
		println(fileStatus)
	}
}