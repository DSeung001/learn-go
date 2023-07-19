package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gochatapp/model"
	"log"
	"strings"
	"time"
)

// RegisterNewUser : 신규 유저 자장
func RegisterNewUser(username, password string) error {
	// redis-cli
	// SYNTAX: SET key value
	// SET username password
	// register new username:password key-value pair
	err := redisClient.Set(context.Background(), username, password, 0).Err()
	if err != nil {
		log.Println("error while adding new user", err)
		return err
	}

	// redis-cli
	// SYNTAX : SADD key value
	// SADD users username
	err = redisClient.SAdd(context.Background(), userSetKey(), username).Err()
	if err != nil {
		log.Println("error while adding user in set", err)
		// redis-cli
		// SYNTAX: DEL key
		// DEL username
		// drop the registered user
		redisClient.Del(context.Background(), username)

		return err
	}
	return nil
}

// IsUserExist : 유저 중복 체크
func IsUserExist(username string) bool {
	// redis-cli
	// SYNTAX: SISMEMBER key value
	// SISMEMBER users username
	return redisClient.SIsMember(context.Background(), userSetKey(), username).Val()
}

// IsUSerAuthentic : 로그인 체크
func IsUSerAuthentic(username, password string) error {
	// redis-cli
	// SYNTAX: GET key
	// GET username
	p := redisClient.Get(context.Background(), username).Val()

	if !strings.EqualFold(p, password) {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

// UpdateContactList : Contact List 추가 및 수정
func UpdateContactList(username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

	// redis-cli SCORE is always float or int
	// SYNTAX: ZADD key SCORE MEMBER
	// ZADD contracts:username 1661360942123 contact
	err := redisClient.ZAdd(context.Background(),
		contactListZKey(username),
		zs,
	).Err()

	if err != nil {
		log.Println("error while updating contact list. username: ", username, "contact:", contact, err)
		return err
	}
	return nil
}

// CreateChat : 채팅 생성
func CreateChat(c *model.Chat) (string, error) {
	chatKey := chatKey()
	fmt.Println("chat key", chatKey)

	by, _ := json.Marshal(c)

	// redis-cli
	// SYNTAX: JSON.SET key & json_in_string
	// JSON.SET chat#1661360942123 $ '{"from"L"sum", "to":"earth", "message":"good morning!"}'
	res, err := redisClient.Do(
		context.Background(),
		"JSON_SET",
		chatKey,
		"$",
		string(by),
	).Result()

	if err != nil {
		log.Println("error while setting chat json", err)
		return "", err
	}

	log.Println("chat successfully set", res)

	// add contacts to both user's contact list
	err = UpdateContactList(c.From, c.To)
	if err != nil {
		log.Println("error while updating contact list of", c.From)
	}

	err = UpdateContactList(c.To, c.From)
	if err != nil {
		log.Println("error while updating contact list of", c.To)
	}

	return chatKey, nil
}

// CreateFetchChatBetweenIndex : Redisearch 모듈 사용을 위한 index 추출
func CreateFetchChatBetweenIndex() {
	res, err := redisClient.Do(context.Background(),
		"FT.CREATE",
		chatIndex(),
		"ON", "JSON",
		"PREFIX", "1", "chat#",
		"SCHEMA", "$.from", "AS", "from", "TAG",
		"$.to", "AS", "to", "TAG",
		"$.timestamp", "AS", "timestamp", "NUMERIC", "SORTABLE",
	).Result()

	fmt.Println(res, err)
}

// FetchChatBetween : 채팅방 내용 갱신
func FetchChatBetween(username1, username2, fromTS, toTS string) ([]model.Chat, error) {
	// redis-cli
	// SYNTAX: FT.SEARCH index query
	// FT.SEARCH id#chats '@from:{user2|user1} @to:{user2|user1} @timestamp:[0 +inf]'
	query := fmt.Sprintf("@from:{%s|%s} @to:{%s|%s} @timestamp:[%s %s]",
		username1, username2, username1, username2, fromTS, toTS)

	res, err := redisClient.Do(context.Background(),
		"FT.SEARCH",
		chatIndex(),
		query,
		"SORTBY", "timestamp", "DESC",
	).Result()

	if err != nil {
		return nil, err
	}

	// deserialize redis data to map
	data := Deserialize(res)

	// deserialize data map to chat
	chats := DeserializeChat(data)
	return chats, nil
}

// FetchContactList : 로그인한 유저로 연결 리스트 가져오기
func FetchContactList(username string) ([]model.ContactList, error) {
	zRangeArg := redis.ZRangeArgs{
		Key:   contactListZKey(username),
		Start: 0,
		Stop:  -1,
		Rev:   true,
	}

	// redis-cli
	// SYNTAX: ZRANGE key from_index to_index REV WITHSCORES
	res, err := redisClient.ZRangeArgsWithScores(context.Background(), zRangeArg).Result()

	if err != nil {
		log.Println("error while fetching contact list. username:", username, err)
		return nil, err
	}

	contactList := DeserializeContactList(res)
	return contactList, nil
}
