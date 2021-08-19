package repository

import (
	"github.com/jmoiron/sqlx"
)

type BaseRepository struct{
	Db *sqlx.DB
}