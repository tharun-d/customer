package main

import (
	"customer/common"
	"customer/common/mysql"
	cfg "customer/config"
	vault "customer/config/vault"
	Handlerhttp "customer/delivery/http"
	"customer/repository"
	"customer/service"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var config cfg.Config
var dbConn *sql.DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	config, err = getConfig()

	if err != nil {
		log.Println("Error get config")
		panic(err)
	}
	dbConn, err = getDB()
	if err != nil {
		log.Println("Error connecting db")
		panic(err)
	}
}

func getConfig() (cfg.Config, error) {
	config, err := vault.GetConfig("customer")
	if err != nil {
		log.Println(err)
		return config, err
	}
	return config, err
}

func getDB() (*sql.DB, error) {
	dbCa := config.GetBinary(`database.ca`)
	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)
	maxOpen := config.GetInt(`database.max.open`)
	maxIdle := config.GetInt(`database.max.idle`)
	maxLifetime := config.GetInt(`database.max.lifetime`)

	dbConfig := mysql.Config{
		Host:        dbHost,
		Port:        dbPort,
		User:        dbUser,
		Password:    dbPass,
		Name:        dbName,
		MaxOpen:     int(maxOpen),
		MaxIdle:     int(maxIdle),
		MaxLifetime: int(maxLifetime),
		CA:          dbCa,
		Location:    "Asia/Jakarta",
		ParseTime:   true,
	}

	dbConn, err := mysql.DB(dbConfig)
	if err != nil {
		log.Println(err)
		return dbConn, err
	}
	return dbConn, err
}

func serveHTTP(addr string, customerService service.CustomerService) error {
	Handlerhttp.NewCustomerHandler(customerService)
	log.Println("http server started. Listening on port: ", addr)
	if err := http.ListenAndServe(addr, Handlerhttp.DefaultHandler(http.HandlerFunc(common.Serve))); err != nil {
		return err
	}
	return nil
}

func main() {
	cr := repository.NewMariaCustomerRepository(dbConn)
	cu := service.NewCustomerService(cr)
	httpPort := config.GetString("server.address.http")
	go serveHTTP(httpPort, cu)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	//l.Close()
	log.Println("All server stopped!")
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Print(err)
		}
	}()
}
