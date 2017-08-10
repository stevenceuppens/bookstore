/*******************************************************************************
* Package and license  ...
*******************************************************************************/

package books

import "github.com/gin-gonic/gin"
import "net/http"

/*******************************************************************************
* Imports ...
*******************************************************************************/

/*******************************************************************************
* Type definitions ...
*******************************************************************************/

// BookREST ...
type BookREST struct {
	api BookAPI
}

/*******************************************************************************
* Init methods ...
*******************************************************************************/

/*******************************************************************************
* Private methods ...
*******************************************************************************/

// findAll ...
func (u *BookREST) findAllHandler(c *gin.Context) {

	books, err := u.api.FindAll(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Sorry...")
	}

	// if books is nill
	if books == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, books)
}

/*******************************************************************************
* Public methods ...
*******************************************************************************/

// NewBookREST ...
func NewBookREST(api BookAPI) *BookREST {
	return &BookREST{api}
}

// RegisterRoutes ...
func (u *BookREST) RegisterRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", u.findAllHandler)
		// other handlers...
	}
}
