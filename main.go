package main

import (
    "./handler/user"
    "./handler/survival"
    "./handler/contents"
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
    e.POST("/login", User.Login(db))
    e.POST("/delete/:userid", User.Delete(db))

    e.File("/login", "html/login.html")
    e.File("/css/login.css", "css/login.css")
    e.File("/js/login.js", "js/login.js")

    e.File("/create", "html/create.html")
    e.File("/css/create.css", "css/create.css")
    e.File("/js/create.js", "js/create.js")

    a := e.Group("/admin")
    a.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "admin" && password == "admin" {
            c.Set("userid", username)
            return true, nil
        }
        return false, nil
    }))

    a.GET("", Contents.Admin)
    a.GET("/survival", Survival.IsSurvivals(db))
    a.POST("/evolution/:userid", User.Evolution(db))
    a.GET("/join", User.IsJoins(db))

    cli := e.Group("/client")
    cli.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "client" && password == "client" {
            c.Set("userid", username)
            return true, nil
        }
        return false, nil
    }))

    cli.GET("", Contents.Client)

    cli.GET("/survival", Survival.IsSurvivalMe(db))
    cli.GET("/join", User.IsJoinMe(db))

    cli.POST("/join", User.Join(db))
    cli.POST("/dontjoin", User.DontJoin(db))
    cli.POST("/resporn", Survival.Resporn(db))
    cli.POST("/dead", Survival.Dead(db))

    e.Start(":8080")
}
