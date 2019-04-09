package DB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type Auth struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    PW string `json:"pw" form:"pw" query:"pw"`
    Status string `json:"status" form:"status" query:"status"`
}

type AuthCode struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Code string `json:"code" form:"code" query:"code"`
}

type IsJoin struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    IsJoin bool `json:"is_join" form:"is_join" query:"is_join"`
}

type IsSurvival struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    IsSurvival bool `json:"is_survival" form:"is_survival" query:"is_survival"`
}

type Team struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Team bool `json:"team" form:"team" query:"team"`
}
