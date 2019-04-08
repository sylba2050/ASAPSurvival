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
    "html/template"
    "io"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"

)

type Renderer struct {
        templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
        return r.templates.ExecuteTemplate(w, name, data)
}

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

    e.Renderer = &Renderer{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
    e.Static("/js", "js")
    e.Static("/css", "css")

    a := e.Group("/admin")
    a.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        if username == "admin" && password == "admin" {
            c.Set("userid", username)
            return true, nil
        }
        return false, nil
    }))

    a.GET("", Contents.Admin)
    a.POST("/evolution/:userid", User.Evolution(db))

    e.File("/", "html/login.html")
    e.File("/create", "html/create.html")

    e.GET("/client/:userid", Contents.Client)

    e.POST("/login", User.Login(db))
    e.POST("/create", User.Create(db))
    e.POST("/delete/:userid", User.Delete(db))

    e.GET("/survival", Survival.IsSurvivals(db))
    e.GET("/join", User.IsJoins(db))

    e.GET("/survival/:userid", Survival.IsSurvivalMe(db))
    e.GET("/join/:userid", User.IsJoinMe(db))

    e.POST("/join/:userid", User.Join(db))
    e.POST("/dontjoin/:userid", User.DontJoin(db))
    e.POST("/resporn/:userid", Survival.Resporn(db))
    e.POST("/dead/:userid", Survival.Dead(db))

    e.Start(":8080")
}
