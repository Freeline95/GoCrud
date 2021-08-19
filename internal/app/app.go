package app

import (
	v1 "github.com/Freeline95/GoCrud/internal/handlers/api/v1"
	"github.com/Freeline95/GoCrud/internal/handlers/web"
	"github.com/Freeline95/GoCrud/internal/repositories"
	"github.com/Freeline95/GoCrud/internal/services"
	"github.com/Freeline95/GoCrud/pkg/database"
	"github.com/Freeline95/GoCrud/pkg/repository"
	"github.com/Freeline95/GoCrud/pkg/server"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Run() error {
	var err error

	// Database initialize
	db, err := dbInit()
	if err != nil {
		return err
	}
	defer db.Close()

	// Router initialize
	router := mux.NewRouter()

	// Handlers initialize
	handlersInit(db, router)
	http.Handle("/", router)

	// Server start
	err = server.Start(os.Getenv("SERVER_PORT"))
	if err != nil {
		return err
	}

	return err
}

func dbInit() (*sqlx.DB, error) {
	dbHost, dbUser, dbPassword, dbName, dbPort :=
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT")

	dbPortInt, err := strconv.ParseInt(dbPort, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	return database.Init(dbHost, dbUser, dbPassword, dbName, dbPortInt)
}

func handlersInit(db *sqlx.DB, router *mux.Router) {
	baseRepository := repository.BaseRepository{
		Db: db,
	}

	apiHandlersInit(baseRepository, router)
	webHandlersInit(router)
	staticHandlerInit()
}

func apiHandlersInit(baseRepository repository.BaseRepository, router *mux.Router) {
	customersRepository := repositories.NewCustomersRepository(baseRepository)
	customersService := services.NewCustomersService(customersRepository)
	myValidator := services.NewValidator()
	customersHandler := v1.NewCustomersHandler(customersService, myValidator, render.New())

	customersHandler.RegisterRoutes(router)
}

func webHandlersInit(router *mux.Router) {
	mainHandler := web.NewMainHandler()

	mainHandler.RegisterRoutes(router)
}

func staticHandlerInit() {
	assetsDir, err := filepath.Abs("./web")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir(assetsDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}