package main

import (
	"testCourse/conf"
)

var plan *Node
var res []PersonalCourse

func init() {
	plan = ReadTrainingPlan("./data/cs.yaml")
	res = getPersonalCourses(conf.Config.Cookie)
}
func main() {

}
