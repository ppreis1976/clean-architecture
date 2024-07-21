package main

import (
	"clean-architecture/internal/pb"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
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

	orderService := service.NewOrderService(*orderRepository)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
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
