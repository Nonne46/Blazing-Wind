package repository

import (
	"els/internal/repository/elastic"
	"els/internal/repository/pg"

	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

func GetPGDB() *gorm.DB {
	return pg.GetDB()
}

func GetElastic() *elasticsearch.Client {
	return elastic.GetDB()
}
