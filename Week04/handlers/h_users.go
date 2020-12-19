package handlers

import (
	"Go-000/Week04/model"

	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("user", userAll)

	groupApi.POST("user", userCreateUserOfRole)
	groupApi.PATCH("user", userUpdate)

}

//All
func userAll(c *gin.Context) {
	mdl := model.User{}
	query := &model.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//CreateUserOfRole
func userCreateUserOfRole(c *gin.Context) {
	var mdl model.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.CreateUserOfRole()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

//Update
func userUpdate(c *gin.Context) {
	var mdl model.User
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
