package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"` 
    Title     string    `json:"title"`                                                         
    Body      string    `json:"body"`                                                         
    CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`                             
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`                            
    UserID    uuid.UUID `json:"userId"`                                                       
}
