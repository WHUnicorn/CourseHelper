package web

import "testCourse/conf"

func Test() {
	router := SetupRouter()
	err := router.Run("127.0.0.1:" + conf.Config.Port)
	if err != nil {
		return
	}
}
