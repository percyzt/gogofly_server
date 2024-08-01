package cmd

import (
	"fmt"
	"gogofly/conf"
	"gogofly/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clear() {
	fmt.Println("---------------Clear-----------------")
}
