package models

import (
    "github.com/google/uuid"

    "time"
)

type TVShow struct {
    BaseMedia
    Name          string    `gorm:"type:varchar(255);not null"`
    OriginalName  string    `gorm:"type:varchar(255)"`
    FirstAirDate  time.Time `gorm:"type:date"`
    OriginCountry string    `gorm:"type:varchar(255)"`
    UserID        uuid.UUID `gorm:"type:uuid"`
    User          User      `gorm:"foreignKey:UserID"`
}
