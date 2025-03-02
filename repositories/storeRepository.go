package repositories

import (
	"github.com/mitcheltastic/EvermosInternship/models"
	"gorm.io/gorm"
)

type StoreRepository struct {
	DB *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{DB: db}
}

func (r *StoreRepository) GetAllStores() ([]models.Store, error) {
	var stores []models.Store
	err := r.DB.Find(&stores).Error
	return stores, err
}
