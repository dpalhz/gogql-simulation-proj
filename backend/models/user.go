package models

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`                                                        
	Username  string    `json:"username"`                                               
	Password  string    `json:"password"`                                               
	Notes     []*Note   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"notes,omitempty"` 
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`                                            
}
