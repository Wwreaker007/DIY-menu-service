package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetPostgressDBConnector() (*sql.DB, error) {

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PSWD"), os.Getenv("DB_NAME"))

	// Get Postgres DB connection
	db, err := sql.Open("postgres", url)
	if(err != nil) {
		fmt.Println("unable to connect to postgresDB : ", err.Error(), url)
		return nil, err
	}

	// Ping to check the connection with the DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: " + err.Error())
		return nil, err
	}
	fmt.Println("SUCCESFULLY CONNECTED TO DB")

	// // Run migrations scripts for order table setup
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	fmt.Println("unable to create orders schema : ", err.Error())
	// 	return nil, err
	// }

	// // Specify the migration scripts path for running scripts
	// m, err := migrate.NewWithDatabaseInstance("file://db/migrations/sql", "postgres", driver)
	// if err != nil {
	// 	fmt.Println("unable to run instance with scheme : ", err.Error())
	// 	return nil, err
	// }

	// // Run migrations
	// err = m.Up()
	// if err != nil {
	// 	fmt.Println("unable to run migration : ", err.Error())
	// 	return nil, err
	// }

	return db, nil
}

/*
	1. Load all the environment variables.
	2. Get DB connection for orders table.
	3. Open a connection to the port for starting the GRPC server
	4. Start the GRPC server registering the service implementation
*/
func main () {
	db, err := GetPostgressDBConnector()
	if err != nil {
		log.Fatalln("unable to load env file : ", err.Error())
	}
	grpcServer := NewGrpcServer("tcp", ":9001", db)
	err = grpcServer.Start()
	if err != nil {
		fmt.Println(err.Error())
		panic(0)
	}
}