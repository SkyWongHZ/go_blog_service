package dao

import (
	"fmt"

	"github.com/go-programming-tour-book/blog-service/internal/model"
)

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	fmt.Println("appKey", appKey, "appSecret", appSecret)
	return auth.Get(d.engine)
}
