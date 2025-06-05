package main

import "tasks.com/app"

func main() {
	application := app.BuildApp()
	application.Run()
}
