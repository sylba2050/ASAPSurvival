package Contents

import (
    "../../utils/redirect"

    "github.com/labstack/echo"

    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Client(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        return c.File("html/index.html")
    }
}

func Admin(c echo.Context) error {
    return c.File("html/admin.html")
}
