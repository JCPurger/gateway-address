package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gateway-address/config"
	"github.com/gateway-address/routes"
	"github.com/gorilla/mux"
)

func serverGetPort() (int, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err == nil {
		return port, nil
	}

	cfgFile, err := config.LoadConfig()
	if err != nil {
		return 0, err
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		return 0, err
	}
	return cfg.Server.Port, nil
}

func StartServer(mux *mux.Router) {
	port, err := serverGetPort()
	if err != nil {
		fmt.Printf("%s", err)
	}
	address := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Printf("Listening on %s\n", address)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		panic(err)
	}
}

func StartServerV1(mux *mux.Router) {
}

func RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/user", routes.UserMethodController)
}

func GetMuxV1() *mux.Router {
	mux := mux.NewRouter()
	return mux
}
