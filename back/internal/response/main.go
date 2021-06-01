package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessMessage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}
func Ok(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}

func ErrorResponse(c *gin.Context, data interface{}, code int) {
	c.JSON(code, gin.H{"errors": data})
}

func InternalServerError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "error": data})
}
func Unauthorized(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": data})
}

func Forbidden(c *gin.Context, data interface{}) {
	c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden", "error": data})
}

func UnexpectedError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected Error", "error": data})
}

func BadRequest(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"error": data })
}

func ValidationError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": data})
}

func PaginationResponse(c *gin.Context, data interface{}, total int64, limit int) {
	SuccessResponse(c, gin.H{"data": data, "total": total, "per_page": limit})
}

func NotFound(c *gin.Context) {
	ErrorResponse(c, gin.H{"message": "not found"}, http.StatusNotFound)
}
