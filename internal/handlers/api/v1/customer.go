package v1

import (
	"database/sql"
	"encoding/json"
	"github.com/Freeline95/GoCrud/internal/models"
	"github.com/Freeline95/GoCrud/internal/repositories"
	"github.com/Freeline95/GoCrud/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"strconv"
)

type CustomersHandler struct {
	service services.CustomersService
	validator *validator.Validate
	render *render.Render
}

func NewCustomersHandler(service services.CustomersService, validator *validator.Validate, render *render.Render) CustomersHandler {
	return CustomersHandler{
		service: service,
		validator: validator,
		render: render,
	}
}

type JsonError struct {
	ErrorMessage string `json:"error"`
}

func (ch *CustomersHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()

	limit := params.Get("limit")
	intLimit, err := strconv.ParseInt(limit, 10, 0)
	if err != nil {
		intLimit = repositories.DEFAULT_LIMIT
		err = nil
	}

	offset := params.Get("offset")
	intOffset, err := strconv.ParseInt(offset, 10, 0)
	if err != nil {
		intOffset = repositories.DEFAULT_OFFSET
		err = nil
	}

	filters := map[string]interface{}{
		"searchString": params.Get("searchString"),
		"limit": intLimit,
		"offset": intOffset,
	}

	customers, err := ch.service.GetCustomersByFilters(filters)

	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusInternalServerError, JsonError{"Error while getting customers. Try again"})

		return
	}

	ch.render.JSON(writer, http.StatusOK, customers)

	return
}

func (ch *CustomersHandler) getOne(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	idInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{"Id must be int"})
		
		return
	}
	
	customer, err := ch.service.GetCustomer(uint(idInt))

	if err == sql.ErrNoRows {
		ch.render.JSON(writer, http.StatusOK, struct{}{})

		return
	}
	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusInternalServerError, JsonError{"Error while getting customer. Try again"})

		return
	}

	ch.render.JSON(writer, http.StatusOK, customer)

	return
}

func (ch *CustomersHandler) create(writer http.ResponseWriter, request *http.Request) {
	customer := models.Customer{}

	err := json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{err.Error()})

		return
	}

	err = ch.validator.Struct(customer)
	if err != nil {
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{err.Error()})

		return
	}

	err = ch.service.Create(&customer)
	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusInternalServerError, JsonError{"Error while customer creating. Try again"})

		return
	}

	ch.render.JSON(writer, http.StatusOK, struct{}{})

	return
}

func (ch *CustomersHandler) update(writer http.ResponseWriter, request *http.Request) {
	var customer models.Customer

	params := mux.Vars(request)
	id := params["id"]
	idInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{"Id must be int"})

		return
	}

	customer.Id = uint(idInt)

	err = json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{err.Error()})

		return
	}

	err = ch.validator.Struct(customer)
	if err != nil {
		ch.render.JSON(writer, http.StatusBadRequest, JsonError{err.Error()})

		return
	}

	err = ch.service.Update(&customer)
	if err != nil {
		log.Println(err)
		ch.render.JSON(writer, http.StatusInternalServerError, JsonError{"Error while customer updateing. Try again"})

		return
	}

	ch.render.JSON(writer, http.StatusOK, struct{}{})

	return
}

func (ch *CustomersHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/all/", ch.getAll).Methods("GET")
	router.HandleFunc("/api/customer/{id:[0-9]+}/", ch.getOne).Methods("GET")
	router.HandleFunc("/api/customer/", ch.create).Methods("POST")
	router.HandleFunc("/api/customer/{id:[0-9]+}/", ch.update).Methods("PUT")
}