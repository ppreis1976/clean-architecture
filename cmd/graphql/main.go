package main

import (
	"clean-architecture/graph"
	"clean-architecture/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8081"

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		*orderRepository,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
