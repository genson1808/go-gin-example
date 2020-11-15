package v1

import (
	"github.com/ROGGER1808/go-gin-example/models"
	"github.com/ROGGER1808/go-gin-example/pkg/e"
	"github.com/ROGGER1808/go-gin-example/pkg/logging"
	"github.com/ROGGER1808/go-gin-example/pkg/setting"
	"github.com/ROGGER1808/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/unknwon/com"
	"net/http"
)

// Get multiple article tags
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// Add article tag
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	err := validation.Errors{
		"name":      validation.Validate(name, validation.Required, validation.Length(1, 100)),
		"state":     validation.Validate(state, validation.Required, validation.In(0, 1)),
		"createdBy": validation.Validate(createdBy, validation.Required, validation.Length(1, 100)),
	}.Filter()

	code := e.INVALID_PARAMS


	if err == nil {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}else {
		logging.Info(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// Edit article tag
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	err := validation.Errors{
		"id": validation.Validate(id, validation.Required, validation.Min(1)),
		"state":     validation.Validate(state, validation.In(0, 1)),
		"modifiedBy": validation.Validate(modifiedBy, validation.Required, validation.Length(1, 100)),
		"name":      validation.Validate(name, validation.Required, validation.Length(1, 100)),
	}.Filter()

	code := e.INVALID_PARAMS
	//var dt interface{}
	if err == nil {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			data["name"] = name
			data["state"] = state

			models.EditTag(id, data)
		}
		//dt = make(map[string]string)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
		logging.Info(err)
		//dt = err
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// Delete article tag
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	err := validation.Validate(id, validation.Required, validation.Min(1))

	code := e.INVALID_PARAMS
	if err == nil {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}else {
		logging.Info(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}
