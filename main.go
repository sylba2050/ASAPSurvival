package main

import (
    "./handler/user"
    "./handler/survival"
    "./struct/DB"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    _ "os"
    _ "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"

)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())

    db, err := gorm.Open("sqlite3", "DB/main.sqlite3")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&DB.Auth{})
    db.AutoMigrate(&DB.IsSurvival{})
    db.AutoMigrate(&DB.IsJoin{})
    db.AutoMigrate(&DB.Team{})

    e.POST("/create", User.Create(db))
    e.POST("/delete/:userid", User.Delete(db))

    a := e.Group("/admin")
    a.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "admin" && password == "admin" {
            c.Set("userid", username)
            c.Set("status", "admin")
            return true, nil
        }
        return false, nil
    }))

    a.File("", "html/admin.html")
    a.GET("/survival", Survival.IsSurvivals(db))

    cli := e.Group("/client")
    cli.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "admin" && password == "admin" {
            c.Set("userid", username)
            return true, nil
        }
        return false, nil
    }))

    cli.File("", "html/index.html")

    cli.GET("/isjoins", User.IsJoins(db))
    cli.POST("/join", User.Join(db))
    cli.POST("/dontjoin", User.DontJoin(db))
    cli.POST("/resporn", Survival.Resporn(db))
    cli.POST("/dead", Survival.Dead(db))

    e.Start(":8080")
}
