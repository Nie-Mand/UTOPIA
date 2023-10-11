package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
    SECRET = "key_nye_west"
)
func main() {
    e := echo.New()
    
    e.OPTIONS("/", func(c echo.Context) error {

        agent := c.Request().Header.Get("User-Agent")

        if agent != "KANYE" {
            return c.String(http.StatusBadRequest, "You ain't my agent")
        }

        xff := c.Request().Header.Get("X-Forwarded-For")
        if xff != "XO-TION" {
            return c.String(http.StatusBadRequest, "You ain't from XO-TION")
        }

        
        bars, err := c.Cookie("bars")
        if err != nil {
            return c.String(http.StatusBadRequest, "Okay, but still I only serve artists with valid bars, but you got none in your jar!")
        } else {
            if len(bars.Value) < 64 {
                return c.String(http.StatusBadRequest, "Too low, you need Mo'. Make like Sixty Fou', I might let you in the crib bro..")
            } else {
                return c.String(http.StatusOK, "You THRIVE FOR SUCCESS! AND I AM THE KEY! SPELL IT ITS : " + SECRET + "\n")
            }
        }
    })

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "You never get, but you still have options")
    })

    e.Logger.Fatal(e.Start(":8001"))
}