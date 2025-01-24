package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Login struct {
    User string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func postLoginHandler(context *gin.Context) {
    var login Login
    err := context.ShouldBind(&login)
    if err != nil {
        context.JSON(
            http.StatusBadRequest,
            gin.H {
                "error": err.Error(),
            },
        )
        return
    }

    if login.User != "admin" || login.Password != "Aa@123456" {
        context.JSON(
            http.StatusUnauthorized,
            gin.H {
                "error": "Unauthorized",
            },
        )
        return
    }

    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "You are logged in",
        },
    )
}

func main() {
    engine := gin.Default()
    engine.POST("/login", postLoginHandler)
    engine.Run()
}
