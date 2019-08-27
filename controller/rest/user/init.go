package user

import "github.com/ha666/gin_demo/service"

var (
	userService *service.UserService
)

func Init() {
	userService = service.NewUserService()
}
