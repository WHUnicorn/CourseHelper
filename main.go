package main

import (
	"testCourse/conf"
	"testCourse/utils"
	"testCourse/web"
)

var plan *Node
var myCourses []PersonalCourse

func init() {
	plan = ReadTrainingPlan("./data/cs.yaml")
	myCourses = getPersonalCourses(conf.Config.Cookie)
}

func validLearnt(arr *[]PersonalCourse, elem Course) bool {
	for i, e := range *arr {

		if e.CourseIdentifier == elem.Identifier {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
			return true
		}
	}
	return false
}

// 获取该节点下的所有课程
func getNodeCourse(node Node, courses *[]Course) {
	if node.Courses != nil && len(node.Courses) > 0 {
		*courses = append(*courses, node.Courses...)
		return
	}
	if node.ChildrenNode == nil || len(node.ChildrenNode) == 0 {
		utils.Debug("该节点无子节点！")
		return
	}
	for _, n := range node.ChildrenNode {
		getNodeCourse(n, courses)
	}
}

func main() {
	utils.Info(plan.ChildrenNode[2].DisplayName)

	sumAll := 0.0
	for _, node := range plan.ChildrenNode[2].ChildrenNode {
		utils.Info(node.DisplayName)
		courses := make([]Course, 0)
		getNodeCourse(node, &courses)
		sum := 0.0
		for _, course := range courses {
			// 传入切片时，如果内部对切片进行操作，数组长度参数不变，但底层存储的数据会发生改变！
			if validLearnt(&myCourses, course) {
				utils.Debug(course.Name, " ", course.Credit)
				sum += course.Credit
			}
		}
		utils.Info("以上共计 ", sum, " 学分", " required: ", node.DemandScore)
		sumAll += sum

	}

	utils.Warning("以上共计 ", sumAll, " 学分", " required: ", plan.ChildrenNode[2].DemandScore)

	//utils.Error(courseMap)
	web.Test()

}
