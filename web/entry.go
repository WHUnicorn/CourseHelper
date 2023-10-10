package web

import "testCourse/setup"

func Test() {
	router := SetupRouter()
	err := router.Run("127.0.0.1:" + setup.Config.Port)
	if err != nil {
		return
	}
}
