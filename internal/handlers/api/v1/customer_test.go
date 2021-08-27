package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Freeline95/GoCrud/internal/models"
	"github.com/Freeline95/GoCrud/internal/repositories/mocks"
	"github.com/Freeline95/GoCrud/internal/services"
	"github.com/Freeline95/GoCrud/pkg/customTypes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var customersStorage = []models.Customer {
	{
		Id: 1,
		FirstName: "FirstName1",
		LastName: "LastName1",
		Gender: "Male",
		Email: "email1@mail.ru",
		BirthDate: customTypes.YmdTime{time.Date(2001, time.Month(2), 21, 0, 0, 0, 0, time.UTC)},
		Address: "Address1",
	},
	{
		Id: 2,
		FirstName: "FirstName2",
		LastName: "LastName2",
		Gender: "Female",
		Email: "email2@mail.ru",
		BirthDate: customTypes.YmdTime{time.Date(2002, time.Month(2), 21, 0, 0, 0, 0, time.UTC)},
		Address: "Address2",
	},
}

var youngCustomer = models.Customer{
	Id: 1,
	FirstName: "FirstName3",
	LastName: "LastName3",
	Gender: "Male",
	Email: "email3@mail.ru",
	BirthDate: customTypes.YmdTime{time.Date(2020, time.Month(2), 21, 0, 0, 0, 0, time.UTC)},
	Address: "Address3",
}

func TestGetAll(t *testing.T) {
	filters := map[string]interface{}{
		"searchString": "",
		"limit": int64(5),
		"offset": int64(0),
	}

	urlString := fmt.Sprintf(
		"/api/v1/customer/all?searchString=%s&limit=%d&offset=%d",
		filters["searchString"],
		filters["limit"],
		filters["offset"],
	)

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	customersRepositoryMock.On("GetAllByFilters", filters).Return(customersStorage, nil).Once()

	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)

	request := httptest.NewRequest(http.MethodGet, urlString, nil)
	responseRecorder := httptest.NewRecorder()

	customersHandler.getAll(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var actualCustomers []models.Customer

	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &actualCustomers)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, customersStorage, actualCustomers)
}

func TestGetAllFail(t *testing.T) {
	filters := map[string]interface{}{
		"searchString": "",
		"limit":        int64(5),
		"offset":       int64(0),
	}

	urlString := fmt.Sprintf(
		"/api/v1/customer/all?searchString=%s&limit=%d&offset=%d",
		filters["searchString"],
		filters["limit"],
		filters["offset"],
	)

	request := httptest.NewRequest(http.MethodGet, urlString, nil)

	responseRecorder := httptest.NewRecorder()

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	customersRepositoryMock.On("GetAllByFilters", filters).Return([]models.Customer{}, errors.New("error while getting from db")).Once()

	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)
	customersHandler.getAll(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	jsonError := JsonError{}

	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &jsonError)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, JsonError{ErrorMessage:"Error while getting customers. Try again"}, jsonError)
}

func TestGetOne(t *testing.T) {
	expectedCustomer := customersStorage[0]

	urlString := fmt.Sprintf("/api/v1/customer/%d", expectedCustomer.Id)

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	customersRepositoryMock.On("GetOneById", expectedCustomer.Id).Return(expectedCustomer, nil).Once()

	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)

	request := httptest.NewRequest(http.MethodGet, urlString, nil)
	responseRecorder := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprintf("%d", expectedCustomer.Id),
	}
	request = mux.SetURLVars(request, vars)

	customersHandler.getOne(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var actualCustomer models.Customer

	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &actualCustomer)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedCustomer, actualCustomer)
}

func TestCreate(t *testing.T) {
	expectedCustomer := customersStorage[0]

	data, err := json.Marshal(expectedCustomer)

	if err != nil {
		t.Error(err)
	}

	urlString := "/api/v1/customer/"

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	customersRepositoryMock.On("Create", &expectedCustomer).Return(nil).Once()

	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)

	request := httptest.NewRequest(http.MethodPost, urlString, strings.NewReader(string(data)))
	responseRecorder := httptest.NewRecorder()

	customersHandler.create(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "{}", string(body))
}

func TestCreateBirthDateValidationFailed(t *testing.T) {
	data, err := json.Marshal(youngCustomer)

	if err != nil {
		t.Error(err)
	}

	urlString := "/api/v1/customer/"

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)

	request := httptest.NewRequest(http.MethodPost, urlString, strings.NewReader(string(data)))
	responseRecorder := httptest.NewRecorder()

	customersHandler.create(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	jsonError := JsonError{}

	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &jsonError)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, JsonError{ErrorMessage:"Key: 'Customer.BirthDate' Error:Field validation for 'BirthDate' failed on the 'adult' tag"}, jsonError)
}

func TestUpdate(t *testing.T) {
	expectedCustomer := customersStorage[1]

	data, err := json.Marshal(expectedCustomer)

	if err != nil {
		t.Error(err)
	}

	urlString := fmt.Sprintf("/api/v1/customer/%d", expectedCustomer.Id)

	customersRepositoryMock := &mocks.CustomersRepositoryInterface{}
	customersRepositoryMock.On("UpdateOne", &expectedCustomer).Return(nil).Once()

	service := services.NewCustomersService(customersRepositoryMock)
	validator := services.NewValidator()
	rend := render.New()

	customersHandler := NewCustomersHandler(service, validator, rend)

	request := httptest.NewRequest(http.MethodPut, urlString, strings.NewReader(string(data)))

	vars := map[string]string{
		"id": fmt.Sprintf("%d", expectedCustomer.Id),
	}
	request = mux.SetURLVars(request, vars)

	responseRecorder := httptest.NewRecorder()

	customersHandler.update(responseRecorder, request)
	resp := responseRecorder.Result()

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "{}", string(body))
}