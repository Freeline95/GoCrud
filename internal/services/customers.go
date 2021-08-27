package services

import (
	"github.com/Freeline95/GoCrud/internal/models"
	"github.com/Freeline95/GoCrud/internal/repositories"
)

type CustomersServiceInterface interface {
	GetCustomersByFilters (filters map[string]interface{}) ([]models.Customer, error)
	GetCustomer (id uint) (models.Customer, error)
	Create (customer *models.Customer) error
	Update (customer *models.Customer) error
}

type CustomersService struct {
	repository repositories.CustomersRepositoryInterface
}

func NewCustomersService (repository repositories.CustomersRepositoryInterface) CustomersServiceInterface {
	return &CustomersService{
		repository: repository,
	}
}

func (service *CustomersService) GetCustomersByFilters (filters map[string]interface{}) ([]models.Customer, error) {
	return service.repository.GetAllByFilters(filters)
}

func (service *CustomersService) GetCustomer (id uint) (models.Customer, error) {
	return service.repository.GetOneById(id)
}

func (service *CustomersService) Create (customer *models.Customer) error {
	return service.repository.Create(customer)
}

func (service *CustomersService) Update (customer *models.Customer) error {
	return service.repository.UpdateOne(customer)
}