package User

import (
    "../../struct/DB"

    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"

)

func Create(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(DB.Auth)
        if err := c.Bind(user); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        status := c.Get("status")
        if status == "admin" {
            user.Status = "admin"
        } else {
            user.Status = "client"
        }

        isUsedUserId := new(DB.Auth)
        db.Where("user_id = ?", user.UserId).First(&isUsedUserId)
        if isUsedUserId.UserId == user.UserId {
            return c.HTML(http.StatusOK, "NG")
        }
        //TODO トランザクション
        db.Create(&user)

        survival := DB.IsSurvival{ UserId: user.UserId, IsSurvival: false }
        db.Create(&survival)

        join := DB.IsJoin{ UserId: user.UserId, IsJoin: false }
        db.Create(&join)

        return c.HTML(http.StatusOK, "ok")
    }
}

func AdminCreate(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Delete(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        //TODO トランザクション
        user := new(DB.Auth)
        db.Where("user_id = ?", userid).First(&user)
        db.Delete(&user)

        survival:= new(DB.IsSurvival)
        db.Where("user_id = ?", userid).First(&survival)
        db.Delete(&survival)

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)
        db.Delete(&join)

        return c.HTML(http.StatusOK, "ok")
    }
}

func Join(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Get("userid")
        //TODO useridのヌル判定

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)

        join.IsJoin = true

        db.Save(&join)
        return c.HTML(http.StatusOK, "ok")
    }
}

func DontJoin(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Get("userid")
        //TODO useridのヌル判定

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)

        join.IsJoin = false

        db.Save(&join)
        return c.HTML(http.StatusOK, "ok")
    }
}

func IsJoins(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        u := []DB.IsJoin{}

        db.Where("is_join = ?", true).Find(&u)

        return c.JSON(http.StatusOK, u)
    }
}
