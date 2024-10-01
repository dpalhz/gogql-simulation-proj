package utils

import (
	"notes/backend/services"

	"gorm.io/gorm"
)

// InitServices menginisialisasi semua layanan yang diperlukan dan mengembalikan struct Services.
func InitServices(db *gorm.DB) *services.Services {
	userService := services.CreateUserService(db)
	noteService := services.CreateNoteService(db)

	return &services.Services{
		UserService: userService,
		NoteService: noteService,
		// Add Other Services here
	}
}
