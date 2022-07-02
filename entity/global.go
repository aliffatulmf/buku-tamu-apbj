package entity

import (
	"time"

	"gorm.io/gorm"
)

type Destination struct {
	ID              uint           `gorm:"primaryKey"`
	DestinationName string         `gorm:"index:,unique;not null"`
	Consultations   []Consultation `gorm:"foreignKey:DestinationID"`
	Pokjas          []Pokja        `gorm:"foreignKey:DestinationID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Consultation struct {
	ID               uint   `gorm:"primaryKey"`
	ConsultationName string `gorm:"index:,unique;not null"`
	DestinationID    uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type Pokja struct {
	ID            uint   `gorm:"primaryKey"`
	PokjaName     string `gorm:"index:,unique;not null"`
	Status        bool
	DestinationID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Pemda struct {
	ID           uint   `gorm:"primaryKey"`
	PemdaName    string `gorm:"not null"`
	Phone        string
	SkpdOpd      string `gorm:"not null"`
	AgencyID     uint   `gorm:"not null"`
	Destination  string
	Consultation string
	Pokja        string
	Image        string `gorm:"not null"`
	Verified     bool   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Provider struct {
	ID           uint   `gorm:"primaryKey"`
	ProviderName string `gorm:"index,unique;not null"`
	Phone        string
	Company      string `gorm:"not null"`
	Description  string
	Destination  string `gorm:"not null"`
	Consultation string
	Pokja        string
	Image        string `gorm:"not null"`
	Verified     bool   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Agency struct {
	ID              uint   `gorm:"primaryKey"`
	AgencyName      string `gorm:"index:,unique;not null"`
	AgencyEmail     string
	AgencyTelephone string
	Pemdas          []Pemda `gorm:"foreignKey:AgencyID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
