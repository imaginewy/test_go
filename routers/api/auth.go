package api

import (
	"gin_blog/models"
	"gin_blog/pkg/app"
	"gin_blog/pkg/e"
	"gin_blog/pkg/gredis"
	"gin_blog/pkg/util"
	"net/http"
	"time"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ginc := app.Gin{C: c}
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		token, err := util.GenerateToken(username, password)
		if err == nil {

			//isExist := models.CheckAuth(username, password)
			isExist := gredis.Exists(token)
			if isExist {
				data["token"] = token
				code = e.SUCCESS
			} else {
				isExistA := models.CheckAuth(username, password)
				if !isExistA {
					code = e.ERROR_AUTH
				} else {
					gredis.Set(token, a, 3*time.Hour)
					data["token"] = token
					code = e.SUCCESS
				}

			}
		} else {
			code = e.ERROR_AUTH_TOKEN
		}
	} else {
		app.MarkErr(valid.Errors)
	}

	ginc.Response(http.StatusOK, code, data)
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})
}
