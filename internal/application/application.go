package application

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/DexScen/api.calculator/pkg/calculator"
	"github.com/joho/godotenv"
)

type Expression struct {
	Expression string `json:"expression"`
}

type Error struct {
	Error string `json:"error"`
}

type Config struct {
	Address string
}

type Result struct {
	Result float64 `json:"result"`
}

var (
	internalError = Error{
		Error: "Internal server error",
	}
	expNotValidError = Error{
		Error: "Expression is not valid",
	}
)

func ConfigFromEnv() *Config {
	config := new(Config)

	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	var exists bool
	config.Address, exists = os.LookupEnv("PORT")
	if !exists {
		log.Printf("No PORT in confin file")
		config.Address = "8080"
	}

	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		final, _ := json.Marshal(internalError)
		w.Write(final)
		return
	}

	exp := &Expression{}
	if err = json.Unmarshal(reqBytes, &exp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		final, _ := json.Marshal(internalError)
		w.Write(final)
		return
	}

	res := Result{}
	if res.Result, err = calculator.Calc(exp.Expression); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		final, _ := json.Marshal(expNotValidError)
		w.Write(final)
		return
	}

	final, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		final, _ := json.Marshal(internalError)
		w.Write(final)
		return
	}

	w.Write(final)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", Handler)
	return http.ListenAndServe(":"+a.config.Address, nil)
}
