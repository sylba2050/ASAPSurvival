package Survival

import (
    "../../struct/DB"
    "../../utils/redirect"

    _ "os"
    _ "fmt"
    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"

)

func IsSurvivals(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.FormValue("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        u := []DB.IsSurvival{}

        db.Table("is_survivals").
            Select("is_survivals.is_survival, is_survivals.user_id").
            Joins("left join is_joins on is_joins.user_id = is_survivals.user_id").
            Joins("left join auths on auths.user_id = is_survivals.user_id").
            Where("is_survivals.is_survival = ? AND is_joins.is_join = ? AND auths.status = ?", true, true, "client").
            Scan(&u)

        return c.JSON(http.StatusOK, u)
    }
}

func IsSurvivalMe(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        survival := new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)

        return c.JSON(http.StatusOK, survival)
    }
}

func Resporn(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        survival := new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)

        survival.IsSurvival = true

        db.Save(&survival)
        return c.HTML(http.StatusOK, "ok")
    }
}

func Dead(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        survival := new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)

        survival.IsSurvival = false

        db.Save(&survival)
        return c.HTML(http.StatusOK, "ok")
    }
}
