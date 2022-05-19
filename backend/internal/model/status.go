package model

import (
	"bytes"
	"context"
	"els/internal/repository/elastic"
	"encoding/json"
	"fmt"
	"time"
)

type Status struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func SearchByStatus(keyword string, offset, limit int) ([]*Status, uint, uint, error) {
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
			  "text"
			]
		  }
		}
	}`, offset, limit, keyword)

	buff := bytes.NewBufferString(query)

	res, err := elasticClient.Search(
		elasticClient.Search.WithContext(context.Background()),
		elasticClient.Search.WithIndex("statuses"),
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

	statuses := []*Status{}

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		status := &Status{}

		source := hit.(map[string]interface{})["_source"]

		b, _ := json.Marshal(source)
		json.Unmarshal(b, status)

		statuses = append(statuses, status)
	}

	return statuses, took, total, nil
}
