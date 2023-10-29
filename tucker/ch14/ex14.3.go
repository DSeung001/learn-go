package main

import "fmt"

// 포인터 사용 이유 : 메모리 절약 + 데이터 복사

type Data struct {
	value int
	data  [200]int
}

func ChangeData1(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangeData2(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData1(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	ChangeData2(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	// 아래 데이터들은 주소가 할당되지 않은 각각의 인스턴스로 즉 인스턴스 3개다
	// - 인스턴스는 메모리에 생성된 데이터의 실체이다
	// - 포인터를 이용해서 인스턴스를 가리키게 할 수 있다
	// - 함수 호출 시 포인터 인수를 통해 인스턴스를 입력받고 그 값을 변형할 수 있다
	// - 쓸모 없어진 인스턴스는 가비지 컬랙터가 지워준다-
	// var data1 Data
	// var data2 Data = data1
	// var data3 Data = data1
}
