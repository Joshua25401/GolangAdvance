package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/joshuaryandafres/golang/goroutines/Real-World-Project/Final-Project/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// Connect to db
	db := initializeDB()

	// Create session
	session := initSession()

	// Create logger
	infoLog := log.New(os.Stdout, "INFO :", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)

	// Create channel

	// Create WaitGroup
	wg := sync.WaitGroup{}

	// Set application config
	app := Config{
		Session:   session,
		DB:        db,
		InfoLog:   infoLog,
		ErrorLog:  errorLog,
		WaitGroup: &wg,
		Models:    data.New(db),
	}

	// Set up mail

	// Listen for signals (SIGINT, SIGTERM)
	go app.listenForShutdown()

	// Listen for web connection
	app.serve()
}

func initializeDB() *sql.DB {
	conn := connectDB()
	if conn == nil {
		log.Panic("Can't connect to database!")
	}
	return conn
}

func connectDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...", err)
		} else {
			log.Println("Connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Println("Backing off for 1 second!")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initSession() *scs.SessionManager {
	// Register model to redis cache
	gob.Register(data.User{})

	// Set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	// Pool of redis connection
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
		MaxIdle: 10,
	}
	return redisPool
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// Perform any cleanup goroutine task
	app.InfoLog.Println("Running cleanup task...")

	// Block until WaitGroup is Empty
	app.WaitGroup.Wait()

	app.InfoLog.Println("Closing channels and shutting down application...")
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("Starting web server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
