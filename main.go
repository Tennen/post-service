package main

import (
	_ "post-service/routers"
	"github.com/astaxie/beego"
)

func main() {
    beego.Run()
}