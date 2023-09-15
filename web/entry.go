package web

import "testCourse/conf"

func Test() {
	router := SetupRouter()
	err := router.Run(":" + conf.Config.Port)
	if err != nil {
		return
	}
}
