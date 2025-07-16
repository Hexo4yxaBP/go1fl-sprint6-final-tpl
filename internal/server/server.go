package server

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/config"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type MyServer struct {
	Logger *log.Logger
	Server *http.Server
}

// инициируем сервер
func InitServer(l *log.Logger) (*MyServer, error) {
	//загружаем конфигурацию сервера из конфига
	if err := config.Load("server.config"); err != nil {
		log.Fatal(err)
	}

	r := http.NewServeMux()

	r.HandleFunc("/", handlers.GetIndexPage)
	r.HandleFunc("/upload", handlers.PostFileForm)

	srv := &http.Server{
		Addr:         config.Server.Addr,
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	mySrv := &MyServer{
		Logger: l,
		Server: srv,
	}

	return mySrv, nil

}
