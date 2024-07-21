package model

type Order struct {
	PedidoID      string       `json:"pedido_id"`
	Cliente       string       `json:"cliente"`
	Vendedor      string       `json:"vendedor"`
	Produto       string       `json:"produto"`
	Quantidade    int          `json:"quantidade"`
	PrecoUnitario float64      `json:"preco_unitario"`
	PrecoTotal    float64      `json:"preco_total"`
	StatusPedido  StatusPedido `json:"status_pedido"`
	DataEntrega   string       `json:"data_entrega"`
}

type StatusPedido string

const (
	AGUARDANDOPAGAMENTO StatusPedido = "Aguardando Pagamento"
	PAGO                StatusPedido = "Pago"
	EMPROCESSAMENTO     StatusPedido = "Em Processamento"
	ENVIADO             StatusPedido = "Enviado"
	FINALIZADO          StatusPedido = "Finalizado"
	INVALIDO            StatusPedido = "Invalido"
)

type OrderDTO struct {
	Cliente       string       `json:"cliente"`
	Vendedor      string       `json:"vendedor"`
	Produto       string       `json:"produto"`
	Quantidade    int          `json:"quantidade"`
	PrecoUnitario float64      `json:"preco_unitario"`
	PrecoTotal    float64      `json:"preco_total"`
	StatusPedido  StatusPedido `json:"status_pedido"`
	DataEntrega   string       `json:"data_entrega"`
}

func (s StatusPedido) String() string {
	return string(s)
}
