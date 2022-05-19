package handlers

import (
	"els/internal/model"
	"els/internal/utls"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchResponsePosts struct {
	Results []*model.Post `json:"results,omitempty"`
	Took    uint          `json:"took,omitempty"`
	Total   uint          `json:"total,omitempty"`
}

type SearchResponseReplies struct {
	Results []*model.StatusReply `json:"results,omitempty"`
	Took    uint                 `json:"took,omitempty"`
	Total   uint                 `json:"total,omitempty"`
}

type SearchResponseStatuses struct {
	Results []*model.Status `json:"results,omitempty"`
	Took    uint            `json:"took,omitempty"`
	Total   uint            `json:"total,omitempty"`
}

type SearchResponseUsers struct {
	Results []*model.User `json:"results,omitempty"`
	Took    uint          `json:"took,omitempty"`
	Total   uint          `json:"total,omitempty"`
}

const (
	maxLimit = 200
)

// SearchPost .
func SearchPost(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utls.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	offset, _ := strconv.Atoi(c.Query("o"))

	limit, _ := strconv.Atoi(c.Query("l"))
	if limit >= maxLimit {
		limit = maxLimit
	} else if limit == 0 {
		limit = 10
	}

	posts, took, total, err := model.SearchByPost(keyword, offset, limit)
	if err != nil {
		log.Println(err)
		utls.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")

		return
	}

	res := &SearchResponsePosts{
		Results: posts,
		Took:    took,
		Total:   total,
	}

	c.JSON(200, res)
}

func SearchReplies(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utls.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	offset, _ := strconv.Atoi(c.Query("o"))

	limit, _ := strconv.Atoi(c.Query("l"))
	if limit >= maxLimit {
		limit = maxLimit
	} else if limit == 0 {
		limit = 10
	}

	replies, took, total, err := model.SearchByReply(keyword, offset, limit)
	if err != nil {
		log.Println(err)
		utls.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")

		return
	}

	res := &SearchResponseReplies{
		Results: replies,
		Took:    took,
		Total:   total,
	}

	c.JSON(200, res)
}

func SearchStatuses(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utls.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	offset, _ := strconv.Atoi(c.Query("o"))

	limit, _ := strconv.Atoi(c.Query("l"))
	if limit >= maxLimit {
		limit = maxLimit
	} else if limit == 0 {
		limit = 10
	}

	statuses, took, total, err := model.SearchByStatus(keyword, offset, limit)
	if err != nil {
		log.Println(err)
		utls.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")

		return
	}

	res := &SearchResponseStatuses{
		Results: statuses,
		Took:    took,
		Total:   total,
	}

	c.JSON(200, res)
}

func SearchUsers(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utls.ErrorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}

	offset, _ := strconv.Atoi(c.Query("o"))

	limit, _ := strconv.Atoi(c.Query("l"))
	if limit >= maxLimit {
		limit = maxLimit
	} else if limit == 0 {
		limit = 10
	}

	users, took, total, err := model.SearchByUsers(keyword, offset, limit)
	if err != nil {
		log.Println(err)
		utls.ErrorResponse(c, http.StatusBadRequest, "Something went wrong")

		return
	}

	res := &SearchResponseUsers{
		Results: users,
		Took:    took,
		Total:   total,
	}

	c.JSON(200, res)
}
