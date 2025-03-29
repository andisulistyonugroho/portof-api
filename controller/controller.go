package controller

import "database/sql"

type DuckHandler struct {
	DB *sql.DB
}