package main

import "github.com/DSeung001/learn-go/tucker/ch20/fedex"

func SendBook1(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	SendBook1("어린 왕자", sender)
	SendBook1("그리스인 조르바", sender)
}
