package v1

import (
	"fmt"
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

// Get multiple articles
func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var err error
	var errors = make(map[string]interface{})

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		err = validation.Validate(state, validation.In(0, 1))
		if err == nil {
			maps["state"] = state
		}else {
			errors["state"] = err
		}
	}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		if !models.ExistTagById(tagId) {
			err = fmt.Errorf("can't exist Tag_id = %d", tagId)
		}else {
		err = validation.Validate(state, validation.Min(1))
		}
		if err == nil {
			maps["tag_id"] = tagId
		}else {
			errors["tag_id"] = err
		}
	}

	code := e.INVALID_PARAMS
	if errors == nil {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for f, er := range errors{

			logging.Info(f, er.(error).Error())
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// Get a single article
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	err := validation.Validate(id, validation.Required, validation.Min(1))

	code := e.INVALID_PARAMS
	var data interface{}
	if err == nil {
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		logging.Info(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// New Article
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	errors := validation.Errors{
		"tag_id":     validation.Validate(tagId, validation.Min(1), validation.Required),
		"title":      validation.Validate(title, validation.Required),
		"desc":       validation.Validate(desc, validation.Required),
		"content":    validation.Validate(content, validation.Required),
		"created_by": validation.Validate(createdBy, validation.Required),
		"state":      validation.Validate(state, validation.In(0, 1)),
	}.Filter()

	code := e.INVALID_PARAMS
	if errors == nil {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		logging.Info(errors)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// Modify article
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	state := -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		validation.Validate(state, validation.In(0, 1))
	}
	errors := validation.Errors{
		"id":          validation.Validate(id, validation.Required, validation.Min(1)),
		"tag_id":      validation.Validate(tagId, validation.Min(1)),
		"title":       validation.Validate(title, validation.Length(1, 100)),
		"desc":        validation.Validate(desc, validation.Length(1, 255)),
		"content":     validation.Validate(content, validation.Length(5, 65535)),
		"modified_by": validation.Validate(modifiedBy, validation.Required, validation.Length(1, 100)),
	}.Filter()

	code := e.INVALID_PARAMS
	if err == nil && errors == nil {
		if models.ExistArticleById(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}

	}else {
		if err != nil {
			logging.Info(err)
		}
		if errors != nil{
			logging.Info(errors)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// Delete article
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	err := validation.Validate(id, validation.Required, validation.Min(1))

	code := e.INVALID_PARAMS
	if err == nil {
		if models.ExistArticleById(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
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
