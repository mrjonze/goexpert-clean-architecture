# goexpert-clean-architecture
Resposta para o desafio de Clean Architecture da pós Go Expert


1 - Para execução do banco de dados e do RabbitMQ, utilizar o docker-compose.yml que está na raiz do projeto.

```docker-compose up -d```

2 - Para execução do projeto, que irá subir os três servidores (HTTP, GraphQL e gRPC), utilizar o comando abaixo:

```go run main.go wire_gen.go```

3 - Para testar o HTTP Server é possível utilizar o arquivo localizado em ```api/list_orders.http```

4 - Para testar o GraphQL Server é possível utilizar o playground no browser, acessando ```http://localhost:8080``` e executando uma query com a seguinte sintaxe:

```graphql
 
 query {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

5 - Para testar o gRPC Server é possível utilizar o evans com o comando ```evans -r repl``` e selecionar o package ```pb```, o service ```OrderService``` e a chamada ```ListOrders```. 

6 - Para testar o RabbitMQ, é possível utilizar o RabbitMQ Management, acessando ```http://localhost:15672``` com o usuário ```guest``` e senha ```guest```.

Após a criação da fila ```orders``` e o bind da mesma com a exchange ```amq.direct```, é possível enviar mensagens para a fila através do HTTP Server, GraphQL Server e gRPC Server para os eventos de listagem dos pedidos.