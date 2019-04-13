package User

import (
    "../../struct/DB"
    "../../utils/sha256"
    "../../utils/redirect"

    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func GenerateAuthCode(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        code := new(DB.AuthCode)
        code.UserId = userid
        //TODO ランダム生成(env?sha256?)
        code.Code = "code"

        //TODO updateと切り替え
        db.Create(&code)

        return c.HTML(http.StatusOK, code.Code)
    }
}

func Login(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        form_data := new(DB.AuthCode)
        if err := c.Bind(form_data); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        // TODO クエリの単一化
        user := new(DB.Auth)
        db.Where("user_id = ?", form_data.UserId).First(&user)
        code := new(DB.AuthCode)
        db.Where("user_id = ?", form_data.UserId).Last(&code)

        auth := sha256.Sha256Sum([]byte(user.UserId + user.PW + code.Code))

        if form_data.Code == auth {
            return c.HTML(http.StatusOK, "code")
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

        team := DB.Team{ UserId: user.UserId, Team: "yellow" }
        db.Create(&team)

        // TODO cを含めてランダム生成&DBへ
        return c.HTML(http.StatusOK, "code")
    }
}

func Delete(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

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
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

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
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

        join := new(DB.IsJoin)
        db.Where("user_id = ?", userid).First(&join)

        join.IsJoin = false

        db.Save(&join)
        return c.HTML(http.StatusOK, "ok")
    }
}

func IsJoins(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.FormValue("userid")
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

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
        code := c.FormValue("code")
        if !redirect.IsAccurateCode(userid, code, db) {
            return c.Redirect(http.StatusTemporaryRedirect, "/")
        }

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
