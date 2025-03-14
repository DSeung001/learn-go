package test

import (
	"bytes"
	"encoding/json"
	_struct "grpc_vs_rest/struct"
	"net/http"
	"testing"
)

func BenchmarkRestRegisterParallel(b *testing.B) {
	user := _struct.User{
		LastName:  "Kim",
		FirstName: "Minji",
		Phone:     "010-1234-5678",
		Email:     "minji@example.com",
		Gender:    "F",
		BirthDate: "1990-01-01",
		Username:  "minji90",
	}
	data, _ := json.Marshal(user)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
			if err != nil {
				b.Error(err)
			}
			resp.Body.Close()
		}
	})
}
