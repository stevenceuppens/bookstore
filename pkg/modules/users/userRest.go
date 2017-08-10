/*******************************************************************************
* Package and license  ...
*******************************************************************************/

package users

import "github.com/gin-gonic/gin"
import "net/http"

/*******************************************************************************
* Imports ...
*******************************************************************************/

/*******************************************************************************
* Type definitions ...
*******************************************************************************/

// UserREST ...
type UserREST struct {
	api UserAPI
}

/*******************************************************************************
* Init methods ...
*******************************************************************************/

/*******************************************************************************
* Private methods ...
*******************************************************************************/

// findAll ...
func (u *UserREST) findAllHandler(c *gin.Context) {
	users, err := u.api.FindAll(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Sorry...")
	}

	c.JSON(http.StatusOK, users)
}

/*******************************************************************************
* Public methods ...
*******************************************************************************/

// NewUserREST ...
func NewUserREST(api UserAPI) *UserREST {
	return &UserREST{api}
}

// RegisterRoutes ...
func (u *UserREST) RegisterRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", u.findAllHandler)
	}
}
