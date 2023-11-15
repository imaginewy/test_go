package articleservice

import (
	"gin_blog/pkg/gredis"
	cacheservice "gin_blog/service/cache_service"
)

type Article struct {
	ID int
}

func (a *Article) ExistByID() (bool, error) {
	article := cacheservice.Article{ID: a.ID}
	exist := gredis.Exists(article.GetArticleKey())

	return exist, nil
}
//git@github.com:imaginewy/test_go.git
//https://github.com/users/set_protocol?protocol_selector=ssh&protocol_type=push
//https://github.com/imaginewy/test_go.git

func (a *Article) Get() (string, error) {
	article := cacheservice.Article{ID: a.ID}
	result, err := gredis.Get(article.GetArticleKey())
	return string(result), err

}
