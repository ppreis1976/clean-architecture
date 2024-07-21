package service

import (
	"clean-architecture/internal/model"
	"clean-architecture/internal/pb"
	"clean-architecture/internal/repository"
	"context"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	OrderDB repository.OrderRepository
}

func NewOrderService(orderDB repository.OrderRepository) *OrderService {
	return &OrderService{OrderDB: orderDB}
}

func (o *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	statusPedido := model.StatusPedido(in.StatusPedido)
	orderDTO := model.OrderDTO{
		Cliente:       in.Cliente,
		Vendedor:      in.Vendedor,
		Produto:       in.Produto,
		Quantidade:    int(in.Quantidade),
		PrecoUnitario: in.PrecoUnitario,
		PrecoTotal:    in.PrecoUnitario * float64(in.Quantidade),
		StatusPedido:  statusPedido,
		DataEntrega:   in.DataEntrega,
	}
	order, err := o.OrderDB.CreateOrder(orderDTO)
	if err != nil {
		return nil, err
	}

	orderResponse := &pb.Order{
		PedidoId:      order.PedidoID,
		Cliente:       order.Cliente,
		Vendedor:      order.Vendedor,
		Produto:       order.Produto,
		Quantidade:    int32(order.Quantidade),
		PrecoUnitario: order.PrecoUnitario,
		PrecoTotal:    order.PrecoTotal,
		StatusPedido:  order.StatusPedido.String(),
		DataEntrega:   order.DataEntrega,
	}

	return &pb.OrderResponse{Order: orderResponse}, nil
}

func (o *OrderService) ListOrders(ctx context.Context, input *pb.Blank) (*pb.OrderList, error) {
	orders, err := o.OrderDB.FindAll()
	if err != nil {
		return nil, err
	}

	orderList := &pb.OrderList{}
	for _, order := range orders {
		orderList.Orders = append(orderList.Orders, &pb.Order{
			PedidoId:      order.PedidoID,
			Cliente:       order.Cliente,
			Vendedor:      order.Vendedor,
			Produto:       order.Produto,
			Quantidade:    int32(order.Quantidade),
			PrecoUnitario: order.PrecoUnitario,
			PrecoTotal:    order.PrecoTotal,
			StatusPedido:  order.StatusPedido.String(),
			DataEntrega:   order.DataEntrega,
		})
	}

	return orderList, nil
}
