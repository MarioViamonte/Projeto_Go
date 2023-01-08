package routes

import controllers "github.com/MarioViamonte/APIGo.git/Controllers"

func ConfigRoutes(router *gin.Engine) *gin.Engine{
	main := router.Group("api/v1")
	{
		books:= main.Group("books")
		{
			books.GET("/", controllers.ShowBook) 
		}
	}
	return router
}