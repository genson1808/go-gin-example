package routers

import (
	"github.com/ROGGER1808/go-gin-example/middleware/jwt"
	"github.com/ROGGER1808/go-gin-example/pkg/setting"
	"github.com/ROGGER1808/go-gin-example/routers/api"
	v1 "github.com/ROGGER1808/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// Get a list of tags
		apiv1.GET("/tags", v1.GetTags)
		// New label
		apiv1.POST("/tags", v1.AddTag)
		// Update the specified label
		apiv1.PUT("/tags/:id", v1.EditTag)
		// Delete the specified label
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// Get list of articles
		apiv1.GET("/articles", v1.GetArticles)
		// Get a single article
		apiv1.GET("/articles/:id", v1.GetArticle)
		// New article
		apiv1.POST("/articles", v1.AddArticle)
		// Update the specified article
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// Delete the specified article
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
