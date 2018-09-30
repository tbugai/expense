package main

import (
  "fmt"
  "os"

  "expense/models"
  "expense/controllers"

  "github.com/joho/godotenv"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
  err = godotenv.Load()
  if err != nil {
    fmt.Print(err)
  }

  dbUser := os.Getenv("db_user")
  dbPass := os.Getenv("db_pass")
  dbHost := os.Getenv("db_host")
  dbName := os.Getenv("db_name")

  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)

  db, err = gorm.Open("postgres", dbUri)

  if err != nil {
    panic("failed to connect to database")
  }

  db.AutoMigrate(&models.Expense{})

  defer db.Close()

  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/expense", controllers.GetExpenses(db))
  e.POST("/expense", controllers.CreateExpense(db))
  e.PUT("/expense/:id", controllers.UpdateExpense(db))
  e.DELETE("/expense/:id", controllers.DestroyExpense(db))

  e.Logger.Fatal(e.Start(":8000"))
}
