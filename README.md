# Clean Architecture - Desafio

## Descrição
```
Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.
```

## Passos:

1. Subir a instância do MySQL

    ```
    docker-compose up -d
    ```

2. Iniciar os Servicos

   ### Endpoint REST (GET /order)
    ```
    go run cmd/api/main.go
    ```

   ### Query ListOrders GraphQL
    ```
    go run cmd/graphql/main.go
    ```

   ### Service ListOrders com GRPC
    ```
    go run cmd/grpc/main.go
    ```

   ### Evans
    ```
    evans --path proto/order.proto
    ```

## Chamadas HTTP:

Para Criação e Listagem de Orders
Utilizar o arquivo clean-architecture.http

## Auxilio ao desenvolvimento:

### Subir Docker
    ```
    colima stop && colima start --mount-type 9p
    ```

### Acessar MySQL
    ```
    /usr/local/opt/mysql-client/bin/mysql -uroot -p mysql -h127.0.0.1
    ```

### Para Docker
    ```
    docker-compose stop
    ```

### Alterações de proto
    ```
    evans --path /path/to --path . proto/order.proto 
    ```