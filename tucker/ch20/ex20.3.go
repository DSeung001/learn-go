package main

import (
	"ch20.com/fedex"
)

func SendBook2(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	//sender := &koreaPost.PostSender{}
	// 당연히 koreaPost는 fedex와 같은 구조체가 아니여서 에러
	//SendBook2("어린 왕자", sender)
	//SendBook2("그리스인 조르바", sender)

}
