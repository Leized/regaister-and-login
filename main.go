package main

import (
	_ "Project/boot"
	_ "Project/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
