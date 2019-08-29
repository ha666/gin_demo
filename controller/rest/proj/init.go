package proj

import (
	"github.com/ha666/gin_demo/service"
)

var (
	projService *service.ProjService
)

func Init() {
	projService = service.NewProjService()
}
