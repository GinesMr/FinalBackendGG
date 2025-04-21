package main

import (
	"awesomeProject/Config"
	"awesomeProject/Routes"
)

func main() {
	Config.CreateConfig()
	r := Routes.GetRoutes()
	r.Run(":8080")
}
