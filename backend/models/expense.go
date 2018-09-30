package models

import (
  "github.com/jinzhu/gorm"
)

type Expense struct {
  gorm.Model
  Description string `json:"description"`
  Amount float32 `json:"amount"`
}

