package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions(userId int) ([]models.Transaction, error)
	GetTransaction(transactionId int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("buyer_id=?", userId).Preload("Product").Preload("Buyer").Preload("Seller").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(transactionId int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Product").Preload("Buyer").Preload("Seller").First(&transaction, transactionId).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Product").Preload("Buyer").Preload("Seller").Create(&transaction).Error

	return transaction, err
}
