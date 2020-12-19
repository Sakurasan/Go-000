package handlers

import (
	"Go-000/Week04/model"
	"errors"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func jsonError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(200, gin.H{"code": 0, "msg": msg})
}
func jsonData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"code": 1, "data": data})
}
func jsonPagination(c *gin.Context, list interface{}, total uint, query *model.PaginationQuery) {
	c.JSON(200, gin.H{"code": 1, "data": list, "total": total, "offset": query.Offset, "limit": query.Limit})
}
func jsonSuccess(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "msg": "success"})
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func parseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, errors.New("id must be an unsigned int")
	}
	return uint(parseId), nil
}

func enableCorsMiddleware() {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
}
