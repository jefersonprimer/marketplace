
# Backend - Marketplace API

Este é um servidor GraphQL baseado em [gqlgen](https://gqlgen.com/) escrito em Go, implementando uma API completa para um marketplace.

## Arquitetura

O projeto segue a arquitetura Clean Architecture com as seguintes camadas:

- **Domain**: Entidades e regras de negócio
- **Application**: Casos de uso (usecases)
- **Infrastructure**: Implementações de repositórios e configurações
- **Graph**: Resolvers GraphQL

## Estrutura do Banco de Dados

O projeto inclui um schema completo de banco de dados PostgreSQL com as seguintes principais entidades:

- **Users/Profiles**: Usuários do sistema
- **Stores**: Lojas do marketplace
- **Categories**: Categorias de produtos
- **Products**: Produtos das lojas
- **Orders**: Pedidos dos clientes
- **Payments**: Pagamentos
- **Reviews**: Avaliações de produtos e lojas
- **Coupons**: Cupons de desconto
- **Promotions**: Promoções

## Pré-requisitos

- Go 1.25.0 ou superior
- PostgreSQL 12 ou superior

## Instalação

1. Clone o repositório
2. Instale as dependências:
   ```bash
   go mod tidy
   ```

## Configuração

### Configuração do Supabase (Recomendado)

1. Crie um projeto no [Supabase](https://supabase.com)
2. Vá para Settings > Database para obter as credenciais
3. Copie o arquivo de configuração de exemplo:
   ```bash
   cp config.example .env
   ```

4. Atualize o arquivo `.env` com suas credenciais do Supabase:
   ```
   SUPABASE_URL=postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   SUPABASE_ANON_KEY=your_supabase_anon_key_here
   ```

### Configuração PostgreSQL Local (Alternativa)

Se preferir usar PostgreSQL local, configure as variáveis:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=marketplace
   ```

## Configuração do Banco de Dados

1. Crie um banco PostgreSQL chamado `marketplace`
2. Execute a migração:
   ```bash
   go run cmd/migrate/main.go
   ```

## Executando o Servidor

```bash
go run server.go
```

O GraphQL playground estará disponível em `http://localhost:8080`

## Desenvolvimento

Para regenerar o código GraphQL após mudanças no schema:

```bash
go run github.com/99designs/gqlgen generate
```

## Principais Funcionalidades Implementadas

### Queries
- `categories`: Lista todas as categorias
- `category(id)`: Busca categoria por ID
- `products(categoryID)`: Lista produtos (opcionalmente filtrados por categoria)
- `product(id)`: Busca produto por ID
- `stores`: Lista todas as lojas
- `store(id)`: Busca loja por ID

### Mutations
- `createCategory(input)`: Cria uma nova categoria
- `createProduct(input)`: Cria um novo produto
- `createStore(input)`: Cria uma nova loja

## Próximos Passos

- Implementar autenticação e autorização
- Adicionar mais operações CRUD
- Implementar sistema de pagamentos
- Adicionar upload de imagens
- Implementar sistema de notificações
- Adicionar testes unitários e de integração
