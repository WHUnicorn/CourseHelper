package data

import (
	"fmt"
	"testCourse/conf"
	"testCourse/utils"
)

var Plan *Node
var MyCourses []PersonalCourse

func init() {
	Plan = ReadTrainingPlan("./trainingPlans/cs.yaml")
	MyCourses = getPersonalCourses(conf.Config.Cookie)
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

type DisplayTable struct {
	DisplayName       string
	SelectedCourses   []string
	UnselectedCourses []string
	DemandTotalScore  string
	CurCredit         string
	AnotherCredit     string
}

// GetValidCourse 获取该节点下所有已修课程
func GetValidCourse(node Node) (res DisplayTable) {
	courses := make([]Course, 0)
	resA := make([]string, 0)
	resB := make([]string, 0)
	GetNodeCourse(node, &courses)
	sumA := 0.0
	sumB := 0.0
	for _, course := range courses {
		myCourses := make([]PersonalCourse, len(MyCourses))
		copy(myCourses, MyCourses)
		if validLearnt(&myCourses, course) {
			resA = append(resA, "第 "+course.Semester+" 学期: "+course.Name+" "+fmt.Sprint(course.Credit))
			sumA += course.Credit
		} else {
			resB = append(resB, "第 "+course.Semester+" 学期: "+course.Name+" "+fmt.Sprint(course.Credit))
			sumB += course.Credit
		}
	}

	res.DisplayName = node.DisplayName
	res.SelectedCourses = resA
	res.UnselectedCourses = resB
	res.DemandTotalScore = fmt.Sprint(node.DemandScore)
	res.CurCredit = fmt.Sprint(sumA)
	res.AnotherCredit = fmt.Sprint(sumB)
	return
}

// GetNodeCourse 获取该节点下的所有课程
func GetNodeCourse(node Node, courses *[]Course) {
	if node.Courses != nil && len(node.Courses) > 0 {
		*courses = append(*courses, node.Courses...)
		return
	}
	if node.ChildrenNode == nil || len(node.ChildrenNode) == 0 {
		utils.Debug("该节点无子节点！")
		return
	}
	for _, n := range node.ChildrenNode {
		GetNodeCourse(n, courses)
	}
}
