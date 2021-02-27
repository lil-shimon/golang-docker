package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//受け取り用の構造体
type Message struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}

//レスポンス用の構造体
type Response struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
	Status string `json:"status"`
}

func main() {
	e := echo.New()
	e.GET("/user", show)
	e.POST("/user", post)
	e.POST("/send", sendMessage)

	e.Logger.Fatal(e.Start(":8080"))
}

//show api
func show(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

//post api
func post(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}
	r := new(Response)
	r.Name = m.Name
	r.Email= m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}
