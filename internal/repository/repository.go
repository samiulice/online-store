package repository

import "online_store/internal/models"

type DatabaseRepo interface {

	GetDate(id int) (models.Date, error)
	InsertTransaction(tnx models.Transaction) (int, error)
	InsertOrder(order models.Order) (int, error)
	InsertCustomer(customer models.Customer) (int, error)

	//User
	GetUserbyUserName(userName string) (models.User, error)
	UpdatePasswordByUserID(id, newPassword string) (error)

	//Token
	InsertToken(t *models.Token, u models.User)error
	GetUserbyToken(token string) (*models.User, error)
}