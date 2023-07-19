package redisrepo

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gochatapp/model"
	"log"
)

type Document struct {
	ID      string `json:"_id"`
	Payload []byte `json:"payload"`
	Total   int64  `json:"total"`
}

// Deserialize : Redisearch 의 결과값인 []interface{}을 Document Struct 로 변환
func Deserialize(res interface{}) []Document {

	// interface 타입에 따른 스위칭
	switch v := res.(type) {
	case []interface{}:
		if len(v) > 1 {
			total := len(v) - 1
			var docs = make([]Document, 0, total/2)

			for i := 1; i <= total; i = i + 2 {
				// Type Assertion = a.(T) 일때 a가 T 타입에 속하는 지 체크
				arrOfValues := v[i+1].([]interface{})
				value := arrOfValues[len(arrOfValues)-1].(string)

				doc := Document{
					ID:      v[i].(string),
					Payload: []byte(value),
					Total:   v[0].(int64),
				}
				docs = append(docs, doc)
			}
			return docs
		}
	default:
		log.Printf("different response type otherthan []interface{}. type: %T", res)
		return nil
	}
	return nil
}

// DeserializeChat : Document를  model.Chat 배열로 변환
func DeserializeChat(docs []Document) []model.Chat {
	chats := []model.Chat{}
	for _, doc := range docs {
		var c model.Chat
		json.Unmarshal(doc.Payload, &c)

		c.ID = doc.ID
		chats = append(chats, c)
	}
	return chats
}

// DeserializeContactList : model.ContractList 로 변환
func DeserializeContactList(contacts []redis.Z) []model.ContactList {
	contactList := make([]model.ContactList, 0, len(contacts))

	for _, contact := range contacts {
		contactList = append(contactList, model.ContactList{
			Username:     contact.Member.(string),
			LastActivity: int64(contact.Score),
		})
	}

	return contactList
}
