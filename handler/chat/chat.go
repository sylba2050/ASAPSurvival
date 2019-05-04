package Chat

import (
    "../../struct/DB"
    "../../utils/redirect"

    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Create(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.FormValue("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        chat := new(DB.Chat)
        if err := c.Bind(chat); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return c.NoContent(http.StatusBadRequest)
        }

        db.Create(&chat)

        return c.NoContent(http.StatusOK)
    }
}

func Read(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.FormValue("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        team := new(DB.Team)
        db.Where("user_id = ?", userid).First(&team)

        chat := []DB.Chat{}
        db.Limit(20).
            Where("target_user_id = ?", userid).
            Or("target_team = ?", team.Team).
            Order("created_at desc").
            Find(&chat)

        return c.JSON(http.StatusOK, chat)
    }
}
