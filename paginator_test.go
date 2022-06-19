package paginator_test

import (
	"log"
	"testing"

	paginator "github.com/yafeng-Soong/gorm-paginator"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Country struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Continent string `json:"continent"`
	Region    string `json:"region"`
	IndepYear int    `json:"indepYear"`
}

func (c *Country) TableName() string {
	return "country"
}

func Test_page(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connect error")
	}
	query := db.Where("Continent = ? and IndepYear > ?", "Asia", 1900)
	p := paginator.Page[Country]{CurrentPage: 1, PageSize: 15}
	p.SelectPages(query)
	log.Println(p)
}
