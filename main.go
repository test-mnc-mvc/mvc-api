package main

import (
	_config "mvc/config"
	_routes "mvc/routes"
)

func main() {
	_config.InitDB()
	e := _routes.New()
	e.Logger.Fatal(e.Start(":3000"))
}
