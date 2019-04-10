package Contents

import (
    _ "../../utils/redirect"

    "github.com/labstack/echo"

    _ "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Client(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.File("html/index.html")
    }
}

func Admin(c echo.Context) error {
    return c.File("html/admin.html")
}
