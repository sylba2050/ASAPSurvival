package Survival

import (
    "../../struct/DB"

    _ "os"
    _ "fmt"
    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"

)

func IsSurvivals(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
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

func Resporn(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Get("userid")
        //TODO useridのヌル判定

        survival := new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)

        survival.IsSurvival = true

        db.Save(&survival)
        return c.HTML(http.StatusOK, "ok")
    }
}

func Dead(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Get("userid")
        //TODO useridのヌル判定

        survival := new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)

        survival.IsSurvival = false

        db.Save(&survival)
        return c.HTML(http.StatusOK, "ok")
    }
}
