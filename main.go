package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
    "os"
)

// init関数でログの初期設定を行う
func init() {
    logrus.SetLevel(logrus.InfoLevel)

    file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logrus.Fatal(err)
    }

    logrus.SetOutput(file)
    logrus.SetFormatter(&logrus.JSONFormatter{})
}

// 環境変数を読み込み、PORTが設定されているか確認する
func loadEnv() string {
    err := godotenv.Load()
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "error": err.Error(),
        }).Fatal(".envファイルの読み込みエラー")
    }

    port := os.Getenv("PORT")
    if port == "" {
        logrus.Fatal("PORT環境変数が設定されていません")
    }

    return port
}

// Ginエンジンの設定を行う
func setupRouter() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    r := gin.New()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    return r
}

func main() {
    port := loadEnv()
    r := setupRouter()
    r.Run(":" + port)
}
