package Contents

import (
    "../../struct/DB"
    "../../utils/sha256"

    "github.com/labstack/echo"

    "net/http"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Client(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")

        // TODO クエリの単一化
        user := new(DB.Auth)
        db.Where("user_id = ?", userid).First(&user)
        auth := new(DB.AuthCode)
        db.Where("user_id = ?", userid).First(&auth)

        hash := sha256.Sha256Sum([]byte(userid + user.PW + auth.Code))

        if code != hash || code == "" {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        data := struct {
            User string
            Code string
        } {
            userid,
            code,
        }

        return c.Render(http.StatusOK, "index", data)
    }
}

func Admin(c echo.Context) error {
    return c.File("html/admin.html")
}
