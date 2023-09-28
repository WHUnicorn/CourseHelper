package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testCourse/data"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("resources/template/*")
	router.Static("/resources/static", "./resources/static")
	router.GET("/aa", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
		})
	})

	major := data.Plan.DisplayName

	// 定义根节点为专业课
	rootNode := data.Plan.ChildrenNode[2]
	//maxDepth := data.MaxDepth - 1

	info := ""
	if data.MyCourses.IsOutOfDate {
		info = "您的cookie已过期，当前数据为上次 (" +
			data.MyCourses.Date.Format("2006/01/02--15:01") + ") 的缓存，如您选课情况有变，请更新cookie"
	}

	// 第一级路由
	router.GET("/", func(context *gin.Context) {
		rootDisplayTable := data.GetValidCourse(rootNode)
		// 遍历子节点并获取信息
		//displayTable := make([]data.DisplayTable, 0)
		childDisplayNames := make([]string, 0)
		if rootNode.ChildrenNode != nil {
			for _, n := range rootNode.ChildrenNode {
				// 变量必须始终存在，才能解析到页面上，页面会随着变量的变化而变化
				//displayTable = append(displayTable, data.GetValidCourse(n))
				childDisplayNames = append(childDisplayNames, n.DisplayName)
			}
		}
		//utils.Debug(data.GetSubNode(rootNode, "专业必修课程", "本专业必修课程"))

		context.HTML(http.StatusOK, "display.gohtml", gin.H{
			"lastPath":          "/",
			"major":             major,
			"info":              info,
			"displayNames":      childDisplayNames, // 下属节点
			"displayName":       rootDisplayTable.DisplayName,
			"selectedCourses":   rootDisplayTable.SelectedCourses,
			"curCredit":         rootDisplayTable.CurCredit,
			"totalCredit":       rootDisplayTable.DemandTotalScore,
			"anotherCredit":     rootDisplayTable.AnotherCredit,
			"unselectedCourses": rootDisplayTable.UnselectedCourses,
		})
	})

	// 第二级路由
	if rootNode.ChildrenNode == nil || len(rootNode.ChildrenNode) == 0 {
		return router
	}
	router.GET("/:table1", func(context *gin.Context) {
		table1 := context.Param("table1")
		curNode := data.GetSubNode(rootNode, table1)
		if curNode == nil {
			context.String(404, "NotFound!")
			return
		}
		curDisplayTable := data.GetValidCourse(*curNode)
		// 遍历子节点并获取名字
		childDisplayNames := make([]string, 0)
		if curNode.ChildrenNode != nil {
			for _, n := range curNode.ChildrenNode {
				childDisplayNames = append(childDisplayNames, table1+"/"+n.DisplayName)
			}
		}

		context.HTML(http.StatusOK, "display.gohtml", gin.H{
			"lastPath":          "/",
			"info":              info,
			"displayNames":      childDisplayNames,
			"displayName":       curDisplayTable.DisplayName,
			"selectedCourses":   curDisplayTable.SelectedCourses,
			"curCredit":         curDisplayTable.CurCredit,
			"totalCredit":       curDisplayTable.DemandTotalScore,
			"anotherCredit":     curDisplayTable.AnotherCredit,
			"unselectedCourses": curDisplayTable.UnselectedCourses,
		})
	})

	// 第三组路由
	router.GET("/:table1/:table2", func(context *gin.Context) {
		table1 := context.Param("table1")
		table2 := context.Param("table2")

		curNode := data.GetSubNode(*data.GetSubNode(rootNode, table1), table2)
		if curNode == nil {
			context.String(404, "NotFound!")
			return
		}
		curDisplayTable := data.GetValidCourse(*curNode)
		// 遍历子节点并获取名字
		childDisplayNames := make([]string, 0)
		if curNode.ChildrenNode != nil {
			for _, n := range curNode.ChildrenNode {
				childDisplayNames = append(childDisplayNames, table1+"/"+table2+"/"+n.DisplayName)
			}
		}

		context.HTML(http.StatusOK, "display.gohtml", gin.H{
			"lastPath":          "/" + table1,
			"info":              info,
			"displayNames":      childDisplayNames,
			"displayName":       curDisplayTable.DisplayName,
			"selectedCourses":   curDisplayTable.SelectedCourses,
			"curCredit":         curDisplayTable.CurCredit,
			"totalCredit":       curDisplayTable.DemandTotalScore,
			"anotherCredit":     curDisplayTable.AnotherCredit,
			"unselectedCourses": curDisplayTable.UnselectedCourses,
		})

	})

	router.GET("/:table1/:table2/:table3", func(context *gin.Context) {
		table1 := context.Param("table1")
		table2 := context.Param("table2")
		table3 := context.Param("table3")

		curNode := data.GetSubNode(*data.GetSubNode(
			*data.GetSubNode(rootNode, table1), table2),
			table3)
		if curNode == nil {
			context.String(404, "NotFound!")
			return
		}
		curDisplayTable := data.GetValidCourse(*curNode)
		// 遍历子节点并获取名字
		childDisplayNames := make([]string, 0)
		if curNode.ChildrenNode != nil {
			for _, n := range curNode.ChildrenNode {
				childDisplayNames = append(childDisplayNames,
					table1+"/"+table2+"/"+table3+"/"+n.DisplayName)
			}
		}

		context.HTML(http.StatusOK, "display.gohtml", gin.H{
			"lastPath":          "/" + table1 + "/" + table2,
			"info":              info,
			"displayNames":      childDisplayNames,
			"displayName":       curDisplayTable.DisplayName,
			"selectedCourses":   curDisplayTable.SelectedCourses,
			"curCredit":         curDisplayTable.CurCredit,
			"totalCredit":       curDisplayTable.DemandTotalScore,
			"anotherCredit":     curDisplayTable.AnotherCredit,
			"unselectedCourses": curDisplayTable.UnselectedCourses,
		})
	})

	//router.GET("/:table/:table2", func(context *gin.Context) {
	//	param := context.Param("table")
	//	for _, e := range displayTable {
	//		if e.DisplayName == param {
	//			context.HTML(http.StatusOK, "display.gohtml", gin.H{
	//				"displayNames":      displayNames,
	//				"displayName":       e.DisplayName,
	//				"selectedCourses":   e.SelectedCourses,
	//				"curCredit":         e.CurCredit,
	//				"totalCredit":       e.DemandTotalScore,
	//				"anotherCredit":     e.AnotherCredit,
	//				"unselectedCourses": e.UnselectedCourses,
	//			})
	//		}
	//	}
	//
	//})

	return router
}
