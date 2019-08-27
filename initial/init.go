package initial

import (
	"github.com/ha666/gin_demo/dao"
	"github.com/ha666/gin_demo/initial/config"
)

func init() {
	initLog()
	config.InitConfig()
	initApp()
	initId()
	dao.Init()
}
