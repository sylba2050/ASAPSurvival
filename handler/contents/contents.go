package Contents

import (
    "github.com/labstack/echo"

    "net/http"
)

func Client(c echo.Context) error {
    userid := c.Param("userid")

    data := struct {
        User string
    } {
        userid,
    }

    return c.Render(http.StatusOK, "index", data)
}

func Admin(c echo.Context) error {
    return c.File("html/admin.html")
}
