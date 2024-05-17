package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type Info struct {
	URL          string
	TimeRequest  string
	TimeResponse string
	CodeResponse int
	Duration     string
}
