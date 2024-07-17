package models

import (
    "github.com/google/uuid"

)

type BaseMedia struct {
    ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    OriginalID      int       `gorm:"uniqueIndex;not null"`
    OriginalLanguage string    `gorm:"type:varchar(255)"`
    Overview        string    `gorm:"type:text"`
    Popularity      float64   `gorm:"type:float"`
    VoteAverage     float64   `gorm:"type:float"`
    VoteCount       int       `gorm:"type:int"`
    GenreIDs        string    `gorm:"type:text"`  // Armazene como JSON ou CSV
    BackdropPath    string    `gorm:"type:varchar(255)"`
    PosterPath      string    `gorm:"type:varchar(255)"`
    IsAdult         bool      `gorm:"default:false"`
}
