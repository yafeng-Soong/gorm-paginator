# 简介
A gorm paginator, based on generic of go 1.18</br>
基于Go 1.18泛型对Gorm分页的简单封装
# 安装
```go get -u github.com/yafeng-Soong/gorm-paginator```
# 使用
内置的分页结构体
```go
type Page[T any] struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Total       int64 `json:"total"`
	Pages       int64 `json:"pages"`
	Data        []T   `json:"data"`
}
```
分页查询函数SelectPages封装为了Page的结构体方法。使用时只需要提前设置好当前页数、分页大小，再传入设置好查询条件的gorm.DB结构体指针即可，如：
```go
func test() {
    // 先构建查询条件
	query := db.Where("Continent = ? and IndepYear > ?", "Asia", 1900)
    // 构建Page结构体，设置当前页数和分页大小
	page := paginator.Page[Country]{CurrentPage: 1, PageSize: 15}
	page.SelectPages(query) // 执行分页查询
}
```
# 完整示例
```go
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
	log.Println(p.CurrentPage)
	log.Println(p.PageSize)
	log.Println(p.Total)
	log.Println(p.Pages)
	log.Println(p.Data)
}
```

