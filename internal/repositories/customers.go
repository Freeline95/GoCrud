package repositories

import (
	"errors"
	"fmt"
	"github.com/Freeline95/GoCrud/internal/models"
	"github.com/Freeline95/GoCrud/pkg/repository"
	"log"
)

const DEFAULT_LIMIT = 10

type CustomersRepositoryInterface interface {
	GetAllByFilters(filters map[string]interface{}) ([]models.Customer, error)
	GetOneById(id uint) (models.Customer, error)
	Create (customer *models.Customer) error
	UpdateOne (customer *models.Customer) error
}

type CustomersRepository struct {
	BaseRepository repository.BaseRepository
}

func NewCustomersRepository(baseRepository repository.BaseRepository) CustomersRepositoryInterface {
	return &CustomersRepository{
		baseRepository,
	}
}

func (rep *CustomersRepository) GetAllByFilters(filters map[string]interface{}) ([]models.Customer, error) {
	var queryString string
	paramsList := map[string]interface{}{}
	customers := make([]models.Customer, 0)

	queryString = "SELECT * FROM customers WHERE true=true"

	if searchString, ok := filters["searchString"]; ok {
		queryString = queryString + " AND (first_name ILIKE :search_string OR last_name ILIKE :search_string)"
		paramsList["search_string"] = "%" + fmt.Sprintf("%v", searchString) + "%"
	}

	queryString = queryString + " LIMIT :limit"
	if limit, ok := filters["limit"]; ok {
		paramsList["limit"] = limit
	} else {
		paramsList["limit"] = 10
	}

	if offset, ok := filters["limit"]; ok {
		queryString = queryString + " OFFSET :offset"
		paramsList["offset"] = offset
	}

	result, err := rep.BaseRepository.Db.NamedQuery(queryString, paramsList)

	if err != nil {
		return customers, err
	}

	defer result.Close()

	for result.Next() {
		var customer models.Customer

		err = result.StructScan(&customer)

		if err != nil {
			return customers, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (rep *CustomersRepository) GetOneById(id uint) (models.Customer, error) {
	var customer models.Customer

	err := rep.BaseRepository.Db.Get(
		&customer,
	"SELECT * FROM customers WHERE id = $1",
		[]interface{}{id}...
	)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (rep *CustomersRepository) Create (customer *models.Customer) error {
	queryString := "INSERT INTO customers (first_name, last_name, birth_date, gender, email, address) " +
		"VALUES (:first_name, :last_name, :birth_date, :gender, :email, :address)"

	_, err := rep.BaseRepository.Db.NamedExec(
		queryString,
		customer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rep *CustomersRepository) UpdateOne (customer *models.Customer) error {
	queryString := "UPDATE customers " +
		"SET first_name = :first_name, " +
		"last_name = :last_name, " +
		"birth_date = :birth_date, " +
		"gender = :gender, " +
		"email = :email, " +
		"address = :address " +
		"WHERE id = :id"

	result, err := rep.BaseRepository.Db.NamedExec(
		queryString,
		customer,
	)
	if err != nil {
		return err
	}

	rowsCountAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())

		return err
	}
	if rowsCountAffected < 1 {
		err = errors.New("Customer was not updated")

		return err
	}

	return nil
}