package redirect

import (
    "../../struct/DB"
    "../sha256"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)


func IsAccurateCode (userid, code string, db *gorm.DB) bool {
    // TODO クエリの単一化
    user := new(DB.Auth)
    db.Where("user_id = ?", userid).First(&user)
    auth := new(DB.AuthCode)
    db.Where("user_id = ?", userid).Last(&auth)

    hash := sha256.Sha256Sum([]byte(userid + user.PW + auth.Code))

    return  code == hash && code != ""
}
