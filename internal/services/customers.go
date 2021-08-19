package services

import (
	"github.com/Freeline95/GoCrud/internal/models"
	"github.com/Freeline95/GoCrud/internal/repositories"
)

type CustomersService struct {
	repository repositories.CustomersRepositoryInterface
}

func NewCustomersService (repository repositories.CustomersRepositoryInterface) CustomersService {
	return CustomersService{
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