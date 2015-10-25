package main

import (
	"github.com/ckeyer/LogC/conf"
	"github.com/ckeyer/LogC/routers"
)

func main() {
	c, err := conf.GetConfig()
	if err != nil {
		panic(err)
	}
	r := routers.Init()

	_, _ = c, r
}
