package main

import (
	_ "blogserver/routers"
	_ "blogserver/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
