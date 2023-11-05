package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
