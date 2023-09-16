package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testCourse/utils"
)

type Node struct {
	Depth        int      `json:"depth"`
	DisplayName  string   `json:"displayName"`
	DemandScore  float64  `json:"demandScore"`
	ChildrenNode []Node   `json:"childrenNode"`
	Courses      []Course `json:"courses"`
}
type Course struct {
	Identifier  string  `json:"identifier"`
	Name        string  `json:"name"`
	Credit      float64 `json:"credit"`
	Semester    string  `json:"semester"`
	Description string  `json:"description"`
}

func isNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 32)
	return err == nil
}
func countNodeDepth(str string) (depth int, isCourse bool) {
	trimmedStr := strings.TrimLeftFunc(str, func(r rune) bool {
		return r == ' ' || r == '\t'
	})
	isCourse = false
	if len(trimmedStr) > 0 && trimmedStr[0] == '-' {
		isCourse = true
	}
	return (len(str) - len(trimmedStr)) >> 1, isCourse
}

func trimAll(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return r == '"' || r == '-' || r == ' ' || r == '\t' || r == ':'
	})

}

func ReadTrainingPlan(filePath string) *Node {
	data, err := os.Open(filePath)
	if err != nil {
		utils.Error("打开培养方案文件出错: ", err)
		return nil
	}
	defer func(data *os.File) {
		err := data.Close()
		if err != nil {
			utils.Error(err)
		}
	}(data)

	var dataNode Node
	var tempNode *Node

	// 逐行读取文件内容并解析
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if temp := trimAll(line); len(temp) > 0 && temp[0] == '#' {
			// 注释
			continue
		}
		depth, ok := countNodeDepth(line)
		// 叶子
		if ok {
			courseInfo := strings.Split(trimAll(line), " ")
			if len(courseInfo) == 0 {
				utils.Error("未知错误！")
			}

			var newCourse Course
			// 课头号貌似都是13位, 保险起见取前七位...
			if msg := courseInfo[0]; len(msg) > 6 && isNumeric(msg[0:7]) {
				isFinished := false
				for i, elem := range courseInfo {
					if isFinished {
						break
					}
					switch i {
					case 0:
						newCourse.Identifier = elem
					case 1:
						newCourse.Name = elem
					default:
						if isNumeric(elem) {
							credit, _ := strconv.ParseFloat(elem, 64)
							newCourse.Credit = credit
							isFinished = true
						} else {
							newCourse.Name += " " + elem
						}
					}
				}
				for i := len(courseInfo) - 1; i >= 0; i-- {
					isMatch, err := regexp.MatchString("[0-9]+(-[0-9])?", courseInfo[i])
					if err != nil {
						utils.Error(err)
					}
					if isMatch {
						newCourse.Semester = courseInfo[i]
						break
					}
				}

			} else {
				newCourse.Description = trimAll(line)
			}
			//utils.Info(newCourse)

			tempNode.Courses = append(tempNode.Courses, newCourse)
		} else {
			parts := strings.Split(trimAll(line), " ")
			var score float64
			var name string
			switch len(parts) {
			case 1:
				name = parts[0]
				score = -1
			case 2:
				name = parts[0]
				if isNumeric(parts[1]) {
					score, _ = strconv.ParseFloat(parts[1], 64)
				} else {
					score = -1
				}
			default:
				utils.Error("解析出错！")
			}
			if depth != 0 {
				newNode := Node{
					Depth:        depth,
					DisplayName:  name,
					DemandScore:  score,
					ChildrenNode: make([]Node, 0),
					Courses:      nil,
				}
				tempNode = &dataNode
				for i := 0; i < depth-1; i++ {
					// depth == 1 时, tempNode 应当为根节点
					if len(tempNode.ChildrenNode) > 0 {
						tempNode = &tempNode.ChildrenNode[len(tempNode.ChildrenNode)-1]
					} else {
						utils.Error("子节点断言出错了！")
					}
				}
				tempNode.ChildrenNode = append(tempNode.ChildrenNode, newNode)
				tempNode = &tempNode.ChildrenNode[len(tempNode.ChildrenNode)-1]
			} else {
				dataNode.Depth = 0
				dataNode.DisplayName = name
				dataNode.DemandScore = score
				dataNode.ChildrenNode = make([]Node, 0)
				dataNode.Courses = nil
			}

		}
	}

	// scanner 异常结束
	if err := scanner.Err(); err != nil {
		utils.Error("Error reading from file:", err)
	}

	return &dataNode

}
