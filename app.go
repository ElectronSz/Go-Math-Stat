package main

//imports
import (
	"net/http"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

//go client
var client *mongo.Client

// User struct
type User struct {
    Username  string `json:"username" form:"username" query:"username"`
    Password string  `json:"password" form:"password" query:"password"`
}

//e.GET("/", homeRoute)
func homeRoute(c echo.Context) error {

	user := &User{Username: "Faye",Password: "P@ssw0rd"}

	return c.JSON(http.StatusOK,user)
}

// e.POST("/login", userLogin
func userLoginRoute(c echo.Context) (err error) {
	// Get name and email
	user := new(User)
	if err = c.Bind(user); err != nil {
		return
	}

  return c.JSON(http.StatusOK, user)

}


//main function
func main() {

	fmt.Println("Simple interest calculation")
    p := 5000.0
    r := 10.0
    t := 1.0
    si := auth.Calculate(p, r, t)
    fmt.Println("Simple interest is", si)
	// e := echo.New()

	//register routes
	// e.GET("/", homeRoute)
	// e.POST("/login", userLoginRoute)

	//start server
	// e.Logger.Fatal(e.Start(":1323"))
}
