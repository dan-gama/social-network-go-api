package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// ApiPort porta da aplicação
	ApiPort = ""

	// ConnectionString conexão com o banco de dados
	ConnectionString = ""

	// ApiHost endereço da aplicação
	ApiHost = ""
)

// Init função que inicializa o config
func Init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort = os.Getenv("API_PORT")
	ApiHost = os.Getenv("API_HOST")
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
	)
}
