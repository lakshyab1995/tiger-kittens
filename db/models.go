package db

import (
	"time"
)

type User struct {
	ID       string `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

type Tiger struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	DateOfBirth string `gorm:"not null"`
}

func (Tiger) TableName() string {
	return "tiger"
}

type Coordinates struct {
	Lat float64 `gorm:"not null"`
	Lon float64 `gorm:"not null"`
}

type Sighting struct {
	ID          uint `gorm:"primaryKey"`
	TigerID     int
	Tiger       Tiger       `gorm:"foreignKey:TigerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Timestamp   time.Time   `gorm:"not null"`
	Coordinates Coordinates `gorm:"embedded"`
	ImageURL    string
}

type CoordinatesInput struct {
	Lat float64 `gorm:"not null"`
	Lon float64 `gorm:"not null"`
}
