package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("Korea Post sends %v parcel\n", parcel)
}
