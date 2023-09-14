package web

func Test() {
	router := SetupRouter()
	err := router.Run(":12345")
	if err != nil {
		return
	}
}
