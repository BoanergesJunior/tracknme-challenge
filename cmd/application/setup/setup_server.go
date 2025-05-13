package setup

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func SetupServer() *http.Server {
	mux := http.NewServeMux()

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      mux,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Minute * 170,
	}
}
