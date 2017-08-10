/*******************************************************************************
* Package and license  ...
*******************************************************************************/

/*******************************************************************************
* Imports ...
*******************************************************************************/

package main

import (
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
	"github.com/stevenceuppens/bookstore/pkg/modules/books"
	"github.com/stevenceuppens/bookstore/pkg/modules/users"
)

/*******************************************************************************
* Vars  ...
*******************************************************************************/

// Default ENV
var (
	Port = ":3000"
	Mock = false
)

/*******************************************************************************
* Type definitions ...
*******************************************************************************/

/*******************************************************************************
* Init methods ...
*******************************************************************************/

// Capture ENV
func init() {
	if p := os.Getenv("PORT"); p != "" {
		Port = ":" + p
	}
	if m := os.Getenv("Mock"); m != "" {
		Mock = true
	}
}

/*******************************************************************************
* Private methods ...
*******************************************************************************/
func main() {

	router := gin.Default()

	var userAPI users.UserAPI
	var bookAPI books.BookAPI

	if Mock == false {

		// connect to DB
		session, err := mgo.Dial("server1.example.com,server2.example.com")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		// load apis
		userAPI = users.NewUserAPIMongo(session)
		bookAPI = books.NewBookAPIMongo(session)
	} else {

		// load mocks
		userAPI = users.NewUserAPIMock()
		bookAPI = books.NewBookAPIMock()
	}

	userREST := users.NewUserREST(userAPI)
	userREST.RegisterRoutes(router)

	bookREST := books.NewBookREST(bookAPI)
	bookREST.RegisterRoutes(router)

	router.Run(Port)
}

/*******************************************************************************
* Public methods ...
*******************************************************************************/
