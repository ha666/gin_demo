package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/controller/middleware"
	"github.com/ha666/gin_demo/controller/rest/proj"
	"github.com/ha666/gin_demo/controller/rest/user"
	"github.com/ha666/gin_demo/initial/config"
	"github.com/ha666/gin_demo/model"
)

// 启动服务
func Start() {

	//初始化Service
	{
		user.Init()
		proj.Init()
	}
	//endregion

	//region 初始化gin
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
	r := gin.New()
	r.Use(gin.Recovery())
	//endregion

	//region 添加路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, model.Response{Code: "not found", Message: "Page not found"})
	})
	v1 := r.Group("/v1")
	v1.Use(middleware.LimitMiddleware(), middleware.CORSMiddleware())
	{

		//region REST接口
		restNode := v1.Group("rest")
		restNode.Use()
		{

			//region 项目表
			projNode := restNode.Group("projs")
			{
				projNode.GET("", proj.GetProjList)
				projNode.GET(":id", proj.GetProjInfo)
				projNode.POST("", proj.InsertProjInfo)
				projNode.PUT(":id", proj.UpdateProjInfo)
			}
			//endregion

			//region 用户信息表
			userNode := restNode.Group("users")
			{
				userNode.GET("", user.GetUserList)
				userNode.GET(":id", user.GetUserInfo)
				userNode.POST("", user.InsertUserInfo)
				userNode.PUT(":id", user.UpdateUserInfo)
			}
			//endregion

			//region 其它
			{
			}
			//endregion

		}
		//endregion

	}
	//endregion

	//region 启动api服务
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Config.Port))
	//endregion

}
