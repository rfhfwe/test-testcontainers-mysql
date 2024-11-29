package repo

import (
	"context"

	"gorm.io/gorm"

	"test_container/models"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) FindOne(ctx context.Context, id int64) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	return &customer, err
}

func (r *CustomerRepo) Create(ctx context.Context, customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepo) Update(ctx context.Context, customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Delete(&models.Customer{}, id).Error
}

func (r *CustomerRepo) FindAll(ctx context.Context) ([]*models.Customer, error) {
	var customers []*models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}
