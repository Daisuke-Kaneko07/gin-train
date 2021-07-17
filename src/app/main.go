package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ping",
        })
    })
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })
    // ポートを設定しています。
    r.Run(":3000")
}