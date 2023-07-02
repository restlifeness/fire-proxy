package database

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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
	Address    string `gorm:"not null" json:"address"`
	Port       int    `gorm:"not null" json:"port"`
	StillAlive bool   `gorm:"default:true" json:"still_alive"`
}

func (Proxy) TableName() string {
	return "proxy"
}
