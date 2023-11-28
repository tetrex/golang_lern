package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// ShowAccount godoc
	// @Summary      Show an account
	// @Description  get string by ID
	// @Tags         accounts
	// @Accept       json
	// @Produce      json
	// @Success      200  {object}  String
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /hello [get]
	e.GET("/hello", helloHandler("hello world", 01))

	// group
	group1 := e.Group("/g1")

	group1.GET("/root", helloHandler("root-g1", 36))
	group1.GET("/root-2", helloHandler("r2-g1", 03))

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
// func hello(text string, c echo.Context) error {
// 	return c.String(http.StatusOK, text)
// }

func helloHandler(text string, someNumber int) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, text+" --- "+strconv.Itoa(someNumber))
	}
}
