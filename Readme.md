# Go API

## Descrição

Esta é uma API em desenvolvimento, criada para aplicar e aprofundar meus conhecimentos na linguagem Go (Golang). O objetivo principal deste projeto é explorar conceitos fundamentais de desenvolvimento de APIs, incluindo estruturação de código, manipulação de rotas, conexão com banco de dados e autenticação.

## Funcionalidades

- **Gerenciamento de Usuários**: Permite a criação, leitura, atualização e exclusão de usuários.
- **Rotas**: Utiliza o pacote `gorilla/mux` para o roteamento de requisições HTTP.
- **Estrutura Modular**: O código é organizado em pacotes para facilitar a manutenção e a escalabilidade.
- **Persistência de Dados**: Futuramente, o projeto incluirá integração com bancos de dados como PostgreSQL.
- **Autenticação**: Planejamento para implementação de autenticação com JWT.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação principal.
- **Gorilla Mux**: Pacote de roteamento para gerenciamento de rotas HTTP.
- **GORM**: ORM para Go (será integrado em versões futuras).
- **PostgreSQL**: Banco de dados (planejado para versões futuras).

## Como Executar

1. **Clone o repositório**:

   ```bash

   ```

2. **Inicialize o módulo Go**:

```bash
    go mod tidy
```

3. **Iniciar o banco de dados**:

```bash
    docker-compose up -d
```

4. **Inicie o servidor:**

```bash
    make run
```

5. **Acesse a API**:

- A API estará disponível em http://localhost:4000.
