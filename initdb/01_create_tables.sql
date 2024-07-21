CREATE TABLE orders (
   `pedido_id` VARCHAR(255) NOT NULL,
   `cliente` VARCHAR(255) NOT NULL,
   `vendedor` VARCHAR(255) NOT NULL,
   `produto` VARCHAR(255) NOT NULL,
   `quantidade` INT NOT NULL,
   `preco_unitario` DECIMAL(10,2) NOT NULL,
   `preco_total` DECIMAL(10,2) NOT NULL,
   `status_pedido` ENUM('Aguardando Pagamento', 'Pago', 'Em Processamento', 'Enviado', 'Finalizado') NOT NULL,
   `data_entrega` DATE NULL,
   PRIMARY KEY (`pedido_id`)
);