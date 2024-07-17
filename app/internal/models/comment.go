package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Comment struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserID    uuid.UUID `gorm:"type:uuid;not null"`
    MediaID   int       `gorm:"type:int;not null"`
    MediaType string    `gorm:"type:varchar(255);not null"`
    Content   string    `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
    Likes     int       `gorm:"type:int;default:0"`

    User      User      `gorm:"foreignKey:UserID"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
    c.ID = uuid.New()
    c.CreatedAt = time.Now()
    c.UpdatedAt = time.Now()
    return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
    c.UpdatedAt = time.Now()
    return
}
