package backend

import (
  "fmt"
  "os"
  "strconv"

  "expense/backend/models"
  "expense/backend/controllers"

  "github.com/joho/godotenv"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Server struct {
  Port int64
  Echo *echo.Echo
  DB *gorm.DB
}

func (s Server) Serve() {
  err = godotenv.Load()
  if err != nil {
    fmt.Print(err)
  }

  dbUser := os.Getenv("db_user")
  dbPass := os.Getenv("db_pass")
  dbHost := os.Getenv("db_host")
  dbName := os.Getenv("db_name")

  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPass)

  s.DB, err = gorm.Open("postgres", dbUri)

  if err != nil {
    panic("failed to connect to database")
  }

  s.DB.AutoMigrate(&models.Expense{})

  defer s.DB.Close()

  s.Echo = echo.New()

  s.Echo.Use(middleware.Logger())
  s.Echo.Use(middleware.Recover())

  s.Echo.GET("/api/expenses", controllers.GetExpenses(s.DB))
  s.Echo.POST("/api/expenses", controllers.CreateExpense(s.DB))
  s.Echo.PUT("/api/expenses/:id", controllers.UpdateExpense(s.DB))
  s.Echo.DELETE("/api/expenses/:id", controllers.DestroyExpense(s.DB))

  s.Echo.Logger.Fatal(s.Echo.Start(":" + strconv.FormatInt(s.Port, 10)))
}
