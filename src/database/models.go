package database

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	BaseModel
	UUID           string `gorm:"unique;not null"`
	Username       string `gorm:"unique;not null"`
	HashedPassword string `gorm:"not null"`
	Email          string `gorm:"unique;not null"`
	IsActive       bool   `gorm:"default:true"`
	IsAdmin        bool   `gorm:"default:false"`
}

type Proxy struct {
	BaseModel
	Address    string `gorm:"not null"`
	Port       int    `gorm:"not null"`
	StillAlive bool   `gorm:"default:true"`
}
