package model

import (
	"bytes"
	"context"
	"els/internal/repository/elastic"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Join      time.Time `json:"join"`
	About     string    `json:"about"`
	Group     string    `json:"group"`
	Status    string    `json:"status"`
	Age       string    `json:"age"`
	Birthday  string    `json:"birthday"`
	Sex       string    `json:"sex"`
	City      string    `json:"city"`
	Interests string    `json:"interests"`
	Site      string    `json:"site"`
	ICQ       string    `json:"icq"`
	Jabber    string    `json:"jabber"`
	Byond     string    `json:"byond"`
	Discord   string    `json:"discord"`
	Skype     string    `json:"skype"`
}

func SearchByUsers(keyword string, offset, limit int) ([]*User, uint, uint, error) {
	elasticClient := elastic.GetDB()

	query := fmt.Sprintf(`
	{
		"from": %d,
		"size": %d,
		"query": {
		  "query_string": {
			"query": "%s",
			"fuzziness": "AUTO",
			"minimum_should_match": "2",
			"fields": [
		      "username",
			  "about",
			  "group",
			  "status",
			  "sex",
			  "city",
			  "interests",
			  "site",
			  "icq",
			  "jabber",
			  "byond",
			  "discord",
			  "skype"
			]
		  }
		}
	}`, offset, limit, keyword)

	buff := bytes.NewBufferString(query)

	res, err := elasticClient.Search(
		elasticClient.Search.WithContext(context.Background()),
		elasticClient.Search.WithIndex("users"),
		elasticClient.Search.WithBody(buff),
		elasticClient.Search.WithTrackTotalHits(true),
		elasticClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, 0, 0, err
	}

	defer res.Body.Close()

	if res.IsError() {
		var err error

		var e map[string]interface{}

		if er := json.NewDecoder(res.Body).Decode(&e); er != nil {
			err = er
		} else {
			err = fmt.Errorf("%s", e["error"].(map[string]interface{})["reason"])
		}

		return nil, 0, 0, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, 0, err
	}

	total := uint(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	took := uint(result["took"].(float64))

	users := []*User{}

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		user := &User{}

		source := hit.(map[string]interface{})["_source"]

		b, _ := json.Marshal(source)
		json.Unmarshal(b, user)

		users = append(users, user)
	}

	return users, took, total, nil
}
