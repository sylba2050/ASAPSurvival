package Contents

import (
    "github.com/labstack/echo"
)

func Client(c echo.Context) error {
    return c.File("html/index.html")
}

func Admin(c echo.Context) error {
    return c.File("html/admin.html")
}
