package services

import (
	"fmt"
	"notes/backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type UserService struct {
	DB *gorm.DB
}


func CreateUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}


func (s *UserService) AuthenticateUser(username string, password string) (*models.User, error) {
	var user models.User

	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	return &user, nil
}

func (s *UserService) RegisterUser(name, username, password string) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}
	user := &models.User{
		Name:     name,
		Username: username,
		Password: string(hashedPassword),
	}
	if err := s.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}


func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (s *UserService) GetUserNotes(userID string) ([]*models.Note, error) {
	var notes []*models.Note

	if err := s.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch notes: %v", err)
	}

	return notes, nil
}