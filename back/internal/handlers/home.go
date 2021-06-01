package handlers

import "github.com/gin-gonic/gin"

func Home(context *gin.Context)  {
	context.JSONP(200 , gin.H{"message":"ok"})
}
