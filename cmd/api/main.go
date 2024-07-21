package main

import (
	"clean-architecture/internal/handlers"
	"clean-architecture/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)

	err = orderRepository.DeleteAll()
	if err != nil {
		log.Fatalf("Failed to delete orders: %v", err)
	}
	fmt.Println("Delete sample ORDERS")

	//err = orderRepository.InsertSampleData()
	//if err != nil {
	//	log.Fatalf("Failed to insert sample orders: %v", err)
	//}
	//fmt.Println("Insert sample ORDERS")

	router := setupRouter(*orderRepository)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting graphql on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func setupDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/clean")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createOrdersTable(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database setup complete")
	return db, nil
}

func setupRouter(orderRepository repository.OrderRepository) *mux.Router {
	orderHandler := handlers.NewOrderHandler(orderRepository)
	router := mux.NewRouter()
	router.HandleFunc("/orders", orderHandler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods(http.MethodPost)

	return router
}

func createOrdersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		PedidoID VARCHAR(255) PRIMARY KEY,
		Cliente VARCHAR(255),
		Vendedor VARCHAR(255),
		Produto VARCHAR(255),
		Quantidade INT,
		PrecoUnitario FLOAT,
		PrecoTotal FLOAT,
		StatusPedido VARCHAR(255),
		DataEntrega DATE
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
