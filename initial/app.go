package initial

import "github.com/ha666/golibs"

var (
	ServerIP string
)

func initApp() {
	ServerIP = golibs.GetCurrentIntranetIP()
}
