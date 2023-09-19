package data

import (
	"encoding/json"
	"fmt"
	"os"
	"testCourse/conf"
	"testCourse/utils"
	"time"
)

var Plan *Node
var MaxDepth int
var MyCourses PersonalCourses

func init() {
	Plan, MaxDepth = ReadTrainingPlan(conf.Config.DatafilePath)
	MyCourses.Courses = getPersonalCourses(conf.Config.Cookie)
	filePath := "./resources/myCourse.json"
	if MyCourses.Courses != nil {
		MyCourses.Date = time.Now()
		MyCourses.IsOutOfDate = false
		// 获取数据，更新文件
		writeToFile(&MyCourses, filePath)
		return
	}

	// cookie 失效，未获取数据，尝试读文件
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		utils.Error("cookie 已失效且本地无缓存，请重新获取 cookie...")
		return
	}

	// 读文件并解析
	err = json.Unmarshal(bytes, &MyCourses)
	if err != nil {
		utils.Error("解析本地文件失败，请删除该文件！")
		return
	}
	MyCourses.IsOutOfDate = true // 标记为过时

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

// GetSubNode 根据名字获取该节点的子节点
func GetSubNode(node Node, displayName ...string) *Node {
	if len(displayName) == 0 {
		return &node
	}
	for _, n := range node.ChildrenNode {
		if n.DisplayName == displayName[0] {
			return GetSubNode(n, displayName[1:]...)
		}
	}
	return nil
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
		myCourses := make([]PersonalCourse, len(MyCourses.Courses))
		copy(myCourses, MyCourses.Courses)
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
	res.AnotherCredit = fmt.Sprint(node.DemandScore - sumA)
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

func writeToFile(obj interface{}, filePath string) {
	myCoursesJson, _ := json.Marshal(obj)
	err := os.WriteFile(filePath, myCoursesJson, 0666)
	if err != nil {
		utils.Error(err)
		return
	}
}
