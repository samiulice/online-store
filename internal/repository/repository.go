package repository

import "online_store/internal/models"

type DatabaseRepo interface {

	GetDate(id int) (models.Date, error)
	InsertTransaction(tnx models.Transaction) (int, error)
	InsertOrder(order models.Order) (int, error)
	InsertCustomer(customer models.Customer) (int, error)
}