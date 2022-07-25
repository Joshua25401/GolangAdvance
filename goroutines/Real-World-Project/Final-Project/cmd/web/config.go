package main

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/joshuaryandafres/golang/goroutines/Real-World-Project/Final-Project/data"
	"log"
	"sync"
)

type Config struct {
	Session   *scs.SessionManager
	DB        *sql.DB
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
	WaitGroup *sync.WaitGroup
	Models    data.Models
}
