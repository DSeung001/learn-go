package main

import (
 	"github.com/DSeung001/learn-go/tucker/ch20/fedex"
	"github.com/DSeung001/learn-go/tucker/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

type SendBook(name string, sender Sender3) {
	sender.Send(name)
}

func main()  {
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}