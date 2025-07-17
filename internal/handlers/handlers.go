package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/config"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetIndexPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	wd, _ := os.Getwd()
	fmt.Printf("PATH!!!!!!!!!!!!!! %s", wd)
	path, err := filepath.Abs("./../index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	htmlContent, err := os.ReadFile(path)

	//автотест не проходит при обработке ошибки ниже, хотя локально всё выполняется корректно
	if err != nil {
		fmt.Printf("ERRRRRRRRRRRRRRRRRRRRRRORRRRRRRRRRRRRRRRRRRRRR!!!!!! %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlContent)
}

func PostFileForm(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(config.Server.MultipartFormMaxMemory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formFile, formFileHeader, err := r.FormFile("myFile")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer formFile.Close()

	var formFileData []byte = make([]byte, formFileHeader.Size)
	n, err := formFile.Read(formFileData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(n)
	convertedData := service.MorseConvert(string(formFileData))
	os.WriteFile(time.Now().UTC().String()+filepath.Ext(formFileHeader.Filename), []byte(convertedData), 0755)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(convertedData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
