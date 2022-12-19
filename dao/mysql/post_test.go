package mysql

import (
	"testing"

	"github.com/Olixn/GoWebLearn/setting"

	"github.com/Olixn/GoWebLearn/models"
)

func init() {
	dbCfg := setting.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "555555555",
		Content:     "888888888888",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("create post into musql failed,err: %v\n", err)
	}
	t.Logf("pass")
}
