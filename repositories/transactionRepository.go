package repositories

import (
	"github.com/mitcheltastic/EvermosInternship/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}
