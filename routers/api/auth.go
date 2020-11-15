package api

import (
	"github.com/ROGGER1808/go-gin-example/models"
	"github.com/ROGGER1808/go-gin-example/pkg/e"
	"github.com/ROGGER1808/go-gin-example/pkg/logging"
	"github.com/ROGGER1808/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (a auth) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required, validation.Length(1, 50)),
		validation.Field(&a.Password, validation.Required, validation.Length(1, 50)))
}

func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	err := a.Validate()

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if err == nil {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			}else {
				data["token"] = token

				code = e.SUCCESS
			}
		}else {
			code = e.ERROR_AUTH
		}
	}else {
		logging.Info(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}