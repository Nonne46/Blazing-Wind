package model

import (
	"bytes"
	"context"
	"els/internal/repository/elastic"
	"encoding/json"
	"fmt"
)

type StatusReply struct {
	StatusID       uint64 `json:"status_id"`
	StatusUsername string `json:"status_username"`
	StatusText     string `json:"status_text"`
	ReplyID        uint64 `json:"reply_id"`
	ReplyUsername  string `json:"reply_username"`
	ReplyText      string `json:"reply_text"`
}

func SearchByReply(keyword string, offset, limit int) ([]*StatusReply, uint, uint, error) {
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
		      "status_text",
			  "reply_text"
			]
		  }
		}
	}`, offset, limit, keyword)

	buff := bytes.NewBufferString(query)

	res, err := elasticClient.Search(
		elasticClient.Search.WithContext(context.Background()),
		elasticClient.Search.WithIndex("status_replies"),
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

	replies := []*StatusReply{}

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		reply := &StatusReply{}

		source := hit.(map[string]interface{})["_source"]

		b, _ := json.Marshal(source)
		json.Unmarshal(b, reply)

		replies = append(replies, reply)
	}

	return replies, took, total, nil
}
