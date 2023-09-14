package main

import (
	"testCourse/data"
	"testCourse/utils"
	"testCourse/web"
)

func main() {

	utils.Info(data.GetValidCourse(data.Plan.ChildrenNode[2]))
	web.Test()

}
