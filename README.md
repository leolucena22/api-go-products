# API de Produtos em Go

## Sobre o Projeto

Esta é uma API RESTful desenvolvida em Go utilizando o framework Gin, com arquitetura baseada em camadas (Controller, Usecase, Repository e Model). A API permite a manipulação de produtos em um banco de dados PostgreSQL, sendo possível criar, listar e buscar produtos por ID.

## Tecnologias Utilizadas

- **Go** (Golang)
- **Gin** (framework web)
- **PostgreSQL** (Banco de Dados)
- **Docker** e **Docker Compose** (para conteinerização)
- **GORM** (ORM para Go)
- **Gin-Gonic** (framework web para APIs)

## Arquitetura do Projeto

```
api/
│── controller/
│   ├── product_controller.go
│── db/
│   ├── conn.go
│── model/
│   ├── product.go
│   ├── response.go
│── repository/
│   ├── product_repository.go
│── usecase/
│   ├── product_usecase.go
│── main.go
│── docker-compose.yml
│── Dockerfile
```

## Configuração e Execução

### **1. Clonar o repositório**
```sh
git clone https://github.com/leolucena22/api-go-products.git
cd api-go-products
```

### **2. Executar com Docker**

Para subir a aplicação e o banco de dados PostgreSQL, utilize:
```sh
docker-compose up --build
```
A API estará acessível em `http://localhost:8000`.

### **3. Executar sem Docker**

Caso prefira rodar a API localmente sem Docker, siga os passos abaixo:

1. Instale as dependências:
```sh
go mod tidy
```
2. Exporte as variáveis de ambiente para conexão com o PostgreSQL:
```sh
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=1234
export POSTGRES_DB=postgres
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
```
3. Execute a API:
```sh
go run main.go
```
A API será iniciada em `http://localhost:8000`.

## Endpoints

### **1. Listar todos os produtos**
```sh
GET /products
```
**Resposta:**
```json
[
  {
    "id_product": 1,
    "name": "Produto Exemplo",
    "price": 99.99
  }
]
```

### **2. Criar um novo produto**
```sh
POST /product
```
**Corpo da Requisição:**
```json
{
  "name": "Produto Teste",
  "price": 49.99
}
```
**Resposta:**
```json
{
  "id_product": 2,
  "name": "Produto Teste",
  "price": 49.99
}
```

### **3. Buscar um produto por ID**
```sh
GET /product/{id}
```
**Exemplo:**
```sh
GET /product/1
```
**Resposta:**
```json
{
  "id_product": 1,
  "name": "Produto Exemplo",
  "price": 99.99
}
```

## Contribuição

Contribuições são bem-vindas! Para sugerir melhorias, abra um Pull Request ou uma Issue.

## Licença

Este projeto está sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.

