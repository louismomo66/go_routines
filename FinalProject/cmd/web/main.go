package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"subscription/data"
	"sync"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	//connect to the database
	db := initDB()
	db.Ping()
	//create sessions
	session := initSession()
	//setup loggs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//create channels

	//create waitgroup
	wg := sync.WaitGroup{}
	//set up the application config
	app := Config{
		Session:       session,
		DB:            db,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Wait:          &wg,
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	//set up mail
	app.Mailer = app.createMail()
	go app.listenForMail()
	//listen for web connections
	go app.listenForShutdown()
	//listen for errors
	go app.listenForErrors()
	// listen for web connection
	app.serve()
}
