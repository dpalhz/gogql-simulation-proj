package services

import (
	"fmt"
	"notes/backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteService struct {
	DB *gorm.DB
}

// NewNoteService creates a new NoteService
func CreateNoteService(db *gorm.DB) *NoteService {
	return &NoteService{DB: db}
}

// AddNote adds a new note to the database
func (s *NoteService) AddNote(title, body string, userID uuid.UUID) (*models.Note, error) {
	note := &models.Note{
		Title:  title,
		Body:   body,
		UserID: userID,
	}
	if err := s.DB.Create(note).Error; err != nil {
		return nil, fmt.Errorf("failed to add note: %w", err)
	}
	return note, nil
}

// UpdateNote updates an existing note in the database
func (s *NoteService) UpdateNote(id string, title *string, body *string) (*models.Note, error) {
	note := &models.Note{}
	if err := s.DB.First(note, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("note not found: %w", err)
	}

	if title != nil {
		note.Title = *title
	}
	if body != nil {
		note.Body = *body
	}

	if err := s.DB.Save(note).Error; err != nil {
		return nil, fmt.Errorf("failed to update note: %w", err)
	}
	return note, nil
}

// DeleteNote deletes a note from the database
func (s *NoteService) DeleteNote(id string) (bool, error) {
	if err := s.DB.Delete(&models.Note{}, id).Error; err != nil {
		return false, fmt.Errorf("failed to delete note: %w", err)
	}
	return true , nil
}

// GetNotes retrieves all notes from the database
func (s *NoteService) GetNotes() ([]*models.Note, error) {
	var notes []*models.Note
	if err := s.DB.Find(&notes).Error; err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}
	return notes, nil
}

func (s *NoteService) GetNoteByID(id string) (*models.Note, error) {
	note := &models.Note{}
	if err := s.DB.First(note, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("note not found: %w", err)
	}
	return note, nil
}


func (s *NoteService) GetUserByIDNote(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	
	return &user, nil
}