package models

import (
    "github.com/google/uuid"

    "time"
)

type Movie struct {
    BaseMedia
    Title         string    `gorm:"type:varchar(255);not null"`
    OriginalTitle string    `gorm:"type:varchar(255)"`
    ReleaseDate   time.Time `gorm:"type:date"`
    Video         bool      `gorm:"default:false"`
    UserID        uuid.UUID `gorm:"type:uuid"`
    User          User      `gorm:"foreignKey:UserID"`
}
