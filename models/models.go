package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the User model for GORM
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"` // ID unik pengguna
	Name      string    `json:"name"`                                                         // Nama pengguna
	Username  string    `json:"username"`                                                    // Nama pengguna
	Password  string    `json:"password"`                                                    // Kata sandi
	Notes     []*Note   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"notes,omitempty"` // Relasi ke catatan
	CreatedAt time.Time `json:"createdAt"`                                                  // Tanggal pembuatan pengguna
}

// Note represents the Note model for GORM
type Note struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"` // ID unik catatan
	Title     string    `json:"title"`                                                         // Judul catatan
	Body      string    `json:"body"`                                                          // Isi catatan
	CreatedAt time.Time `json:"createdAt"`                                                    // Tanggal pembuatan catatan
	UserID    uuid.UUID `json:"userId"`                                                       // Foreign key untuk relasi ke pengguna
}
