package initial

import (
	"github.com/ha666/golibs"
	"log"
)

var (
	IDNode *golibs.Node
	err    error
)

func initId() {
	IDNode, err = golibs.NewNode(golibs.Unix() % 1024)
	if err != nil {
		log.Fatalf("【initId】加载节点出错:%s", err.Error())
		return
	}
	//logs.Info("【initId】已加载节点")
}
