package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type MyServer struct {
	Logger *log.Logger
	Server *http.Server
}

// инициируем сервер
func InitServer(l *log.Logger) (*MyServer, error) {
	//загружаем конфигурацию сервера из конфига
	/*if err := config.Load("server.config"); err != nil {
		log.Fatal(err)
	}*/

	r := http.NewServeMux()

	r.HandleFunc("/", handlers.GetIndexPage)
	r.HandleFunc("/upload", handlers.PostFileForm)

	srv := &http.Server{
		Addr:         ":8080", //config.Server.Addr,
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,  //config.Server.ReadTimeout,
		WriteTimeout: 10 * time.Second, //config.Server.WriteTimeout,
		IdleTimeout:  15 * time.Second, //config.Server.IdleTimeout,
	}

	mySrv := &MyServer{
		Logger: l,
		Server: srv,
	}

	return mySrv, nil

}
