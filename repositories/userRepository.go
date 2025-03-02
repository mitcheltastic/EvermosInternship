package repositories

import (
	"errors"

	"github.com/mitcheltastic/EvermosInternship/config"
	"github.com/mitcheltastic/EvermosInternship/models"
)

// GetUserByID fetches user details
func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user details
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// DeleteUser performs soft delete
func DeleteUser(userID uint) error {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}
	return config.DB.Delete(&user).Error
}
