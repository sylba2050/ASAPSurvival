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

func GenerateAuthCode(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(DB.Auth)
        if err := c.Bind(user); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        auth := new(DB.Auth)
        db.Where("user_id = ?", user.UserId).First(&auth)

        if user.UserId == auth.UserId && user.PW == auth.PW {
            // TODO ランダム生成&DBへ
            return c.HTML(http.StatusOK, "code")
        } else {
            return c.HTML(http.StatusUnauthorized, "NG")
        }
    }
}

func Login(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(DB.AuthCode)
        if err := c.Bind(user); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        auth := new(DB.AuthCode)
        db.Where("user_id = ?", user.UserId).First(&auth)

        if user.UserId == auth.UserId && user.Code == auth.Code {
            return c.NoContent(http.StatusOK)
        } else {
            return c.HTML(http.StatusUnauthorized, "NG")
        }
    }
}

func Create(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(DB.Auth)
        if err := c.Bind(user); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }
        if user.UserId == "" || user.PW == "" {
            return c.HTML(http.StatusBadRequest, "NG")
        }

        user.Status = "client"

        isUsedUserId := new(DB.Auth)
        db.Where("user_id = ?", user.UserId).First(&isUsedUserId)
        if isUsedUserId.UserId == user.UserId {
            return c.HTML(http.StatusBadRequest, "NG")
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
        userid := c.Param("userid")

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)

        join.IsJoin = true

        db.Save(&join)
        return c.HTML(http.StatusOK, "ok")
    }
}

func DontJoin(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

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

        db.Table("is_joins").
            Select("is_joins.is_join, is_joins.user_id").
            Joins("left join auths on auths.user_id = is_joins.user_id").
            Where("is_joins.is_join = ? AND auths.status = ?", true, "client").
            Scan(&u)

        return c.JSON(http.StatusOK, u)
    }
}

func IsJoinMe(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)

        return c.JSON(http.StatusOK, join)
    }
}

func Evolution(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        user := new(DB.Auth)
        db.Where("user_id = ?", userid).First(&user)

        user.Status = "admin"

        db.Save(&user)
        return c.HTML(http.StatusOK, "ok")
    }
}
