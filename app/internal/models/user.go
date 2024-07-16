package models

import (
    "github.com/google/uuid"

)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    FirstName string    `gorm:"type:varchar(255);not null"`
    LastName  string    `gorm:"type:varchar(255);not null"`
    Username  string    `gorm:"type:varchar(255);not null;uniqueIndex"`
    Email     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
    Password  string    `gorm:"type:varchar(255);not null"`
}
