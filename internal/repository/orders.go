package repository

import (
	"clean-architecture/internal/model"
	"database/sql"
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

//type Order interface {
//	FindAll() ([]model.Order, error)
//	CreateOrder(model.OrderDTO) (model.Order, error)
//	InsertSampleData() error
//	DeleteAll() error
//}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) FindAll() ([]*model.Order, error) {
	rows, err := r.db.Query("SELECT * FROM orders ORDER BY PedidoID")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.PedidoID, &order.Cliente, &order.Vendedor, &order.Produto, &order.Quantidade, &order.PrecoUnitario, &order.PrecoTotal, &order.StatusPedido, &order.DataEntrega)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func randomStatusPedido() model.StatusPedido {
	statusPedidos := []model.StatusPedido{model.PAGO, model.EMPROCESSAMENTO, model.AGUARDANDOPAGAMENTO, model.ENVIADO, model.FINALIZADO}
	return statusPedidos[gofakeit.Number(0, len(statusPedidos)-1)]
}

func (r *OrderRepository) InsertSampleData() error {
	sampleOrders := make([]model.Order, 10)
	for i := 0; i < 10; i++ {
		precoUnitario := gofakeit.Float64Range(10.0, 1000.0)
		quantidade := gofakeit.Number(1, 99)
		sampleOrders[i] = model.Order{
			PedidoID:      uuid.NewString(),
			Cliente:       gofakeit.Name(),
			Vendedor:      gofakeit.Name(),
			Produto:       gofakeit.ProductName(),
			Quantidade:    quantidade,
			PrecoUnitario: precoUnitario,
			PrecoTotal:    precoUnitario * float64(quantidade),
			StatusPedido:  randomStatusPedido(),
			DataEntrega:   gofakeit.Date().Format("2006-01-02"),
		}
	}

	for _, order := range sampleOrders {
		_, err := r.db.Exec(`INSERT INTO orders (PedidoID, Cliente, Vendedor, Produto, Quantidade, PrecoUnitario, PrecoTotal, StatusPedido, DataEntrega) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			order.PedidoID, order.Cliente, order.Vendedor, order.Produto, order.Quantidade, order.PrecoUnitario, order.PrecoTotal, order.StatusPedido, order.DataEntrega)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) DeleteAll() error {
	_, err := r.db.Exec(`DELETE FROM orders`)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) CreateOrder(orderDTO model.OrderDTO) (model.Order, error) {
	isValid, validStatusPedido := IsValidStatusPedido(orderDTO.StatusPedido)
	if !isValid {
		return model.Order{}, fmt.Errorf("status do pedido inválido. Status Validos: %s", validStatusPedido)
	}

	order := model.Order{
		PedidoID:      uuid.NewString(),
		Cliente:       orderDTO.Cliente,
		Vendedor:      orderDTO.Vendedor,
		Produto:       orderDTO.Produto,
		Quantidade:    orderDTO.Quantidade,
		PrecoUnitario: orderDTO.PrecoUnitario,
		PrecoTotal:    orderDTO.PrecoUnitario * float64(orderDTO.Quantidade),
		StatusPedido:  orderDTO.StatusPedido,
		DataEntrega:   orderDTO.DataEntrega,
	}

	_, err := r.db.Exec(`INSERT INTO orders (PedidoID, Cliente, Vendedor, Produto, Quantidade, PrecoUnitario, PrecoTotal, StatusPedido, DataEntrega) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		order.PedidoID, order.Cliente, order.Vendedor, order.Produto, order.Quantidade, order.PrecoUnitario, order.PrecoTotal, order.StatusPedido, order.DataEntrega)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func IsValidStatusPedido(status model.StatusPedido) (bool, string) {
	validStatus := []model.StatusPedido{
		model.AGUARDANDOPAGAMENTO,
		model.PAGO,
		model.EMPROCESSAMENTO,
		model.ENVIADO,
		model.FINALIZADO,
		model.INVALIDO,
	}

	for _, v := range validStatus {
		if status == v {
			return true, ""
		}
	}

	validStatusStr := make([]string, len(validStatus))
	for i, status := range validStatus {
		validStatusStr[i] = string(status)
	}

	return false, strings.Join(validStatusStr, ", ")
}
