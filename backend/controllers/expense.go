package controllers

import (
  "fmt"
  "strconv"
  "net/http"

  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"

  "expense/backend/models"
)

type H map[string]interface{}

func GetExpenses(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    var allExpenses []models.Expense
    db.Find(&allExpenses)
    return c.JSON(http.StatusOK, allExpenses)
  }
}

func CreateExpense(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {

    expense := new(models.Expense)

    if err := c.Bind(expense); err != nil {
      fmt.Print(err)
      return c.JSON(http.StatusBadRequest, "bad request")
    }

    db.Create(&expense)

    return c.JSON(http.StatusCreated, expense)
  }
}

func UpdateExpense(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {

    expense := new(models.Expense)
    db.First(&expense, c.Param("id"))

    if err := c.Bind(expense); err != nil {
      fmt.Print(err)
      return c.JSON(http.StatusBadRequest, expense)
    }

    db.Save(&expense)

    return c.JSON(http.StatusOK, expense)
  }
}

func DestroyExpense(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    expense := new(models.Expense)

    if db.First(&expense, id).RecordNotFound() {
      return c.JSON(http.StatusNotFound, "not found")
    }

    db.Delete(&expense)

    return c.JSON(http.StatusOK, H{
      "deleted": id,
    })
  }
}

