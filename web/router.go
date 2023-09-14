package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testCourse/data"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*")
	router.Static("/static", "./static")
	router.GET("/aa", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
		})
	})

	node := data.Plan.ChildrenNode[2]

	displayTable := make([]data.DisplayTable, 0)

	displayNames := make([]string, 0)
	for _, n := range node.ChildrenNode {
		// 变量必须始终存在，才能解析到页面上，页面会随着变量的变化而变化
		displayTable = append(displayTable, data.GetValidCourse(n))
		displayNames = append(displayNames, n.DisplayName)
	}

	rootDisplayTable := data.GetValidCourse(node)
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "display.gohtml", gin.H{
			"displayNames":      displayNames,
			"displayName":       rootDisplayTable.DisplayName,
			"selectedCourses":   rootDisplayTable.SelectedCourses,
			"curCredit":         rootDisplayTable.CurCredit,
			"totalCredit":       rootDisplayTable.DemandTotalScore,
			"anotherCredit":     rootDisplayTable.AnotherCredit,
			"unselectedCourses": rootDisplayTable.UnselectedCourses,
		})
	})

	router.GET("/:table", func(context *gin.Context) {
		param := context.Param("table")
		for _, e := range displayTable {
			if e.DisplayName == param {
				context.HTML(http.StatusOK, "display.gohtml", gin.H{
					"displayNames":      displayNames,
					"displayName":       e.DisplayName,
					"selectedCourses":   e.SelectedCourses,
					"curCredit":         e.CurCredit,
					"totalCredit":       e.DemandTotalScore,
					"anotherCredit":     e.AnotherCredit,
					"unselectedCourses": e.UnselectedCourses,
				})
			}
		}

	})

	return router
}
