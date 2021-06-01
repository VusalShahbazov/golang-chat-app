package api

import (
	"back/internal/handlers"
	"back/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (apiServer *Server) InitRoutes() {

	apiServer.Engine = gin.Default()


	apiServer.Engine.Use(CORSMiddleware())

	//apiServer.Engine.GET("/", handlers.Home)

	api := apiServer.Engine.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.POST("/logout", middleware.Auth, handlers.Logout)
			auth.GET("/me", middleware.Auth, handlers.Me)
		}
		room := api.Group("rooms" )
		{
			room.GET("", middleware.Auth, handlers.GetRooms)
			room.POST("", middleware.Auth, handlers.StoreRoom)
			room.GET(":id", middleware.Auth,  handlers.GetRoom)
			room.PUT(":id", middleware.Auth,  handlers.UpdateRoom)
			room.DELETE(":id", middleware.Auth,  handlers.DeleteRoom)
			room.GET(":id/messages", middleware.Auth,  handlers.GetMessages)
			room.POST(":id/messages", middleware.Auth, handlers.SendMessage)
		}
	}


	apiServer.Engine.GET("/ws" , func(context *gin.Context) {
		ServeWs(context.Writer , context.Request)
	})


}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}