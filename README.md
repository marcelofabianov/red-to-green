# RedToGreen - Seu Dinheiro, Suas Regras, Sem Complicação!

Bem-vindo ao repositório do projeto RedToGreen! Esta aplicação tem como objetivo ajudar usuários a organizar suas finanças pessoais de forma simples, intuitiva e eficaz.

## 1. Por Onde Começar?

Para entender o RedToGreen, recomendamos os seguintes pontos de partida na nossa documentação:

* **Visão Geral do Produto:** Entenda o "quê" e o "porquê" do RedToGreen.
    * [`_docs/produto/visao_geral_v1.md`](_docs/produto/visao_geral_v1.md)
* **Guia da Documentação:** Um mapa para navegar por toda a nossa documentação.
    * [`_docs/README.md`](_docs/README.md)
* **Princípios Fundamentais:** Conheça a cultura e a filosofia que guiam nosso desenvolvimento.
    * [`_docs/adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md`](_docs/adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md)
* **Glossário de Termos:** Para um entendimento comum dos termos chave do projeto.
    * [`_docs/GLOSSARIO.md`](_docs/GLOSSARIO.md)
* **Modelagem do Domínio:** Explore as entidades centrais e seus relacionamentos.
    * [`_docs/dominio/README.md`](_docs/dominio/README.md)

## 2. Decisões Arquiteturais (ADRs)

Todas as decisões arquiteturais significativas são documentadas utilizando o padrão Architecture Decision Record (ADR). A leitura dos ADRs é essencial para compreender a fundo a estrutura técnica do projeto.

* **Índice Completo de ADRs:** [`_docs/adr/`](_docs/adr/)

**ADRs Chave (Leitura Recomendada):**

* **Tecnologias Core:**
    * [ADR-011: Adoção da Linguagem Go para o Backend](_docs/adr/011-adocao_linguagem_go_para_backend.md)
    * [ADR-012: Adoção do PostgreSQL como Banco de Dados Principal](_docs/adr/012-adocao_postgresql_como_banco_dados.md)
* **Estrutura e Comunicação:**
    * [ADR-010: Adoção de Arquitetura Monolítica Modular](_docs/adr/010-arquitetura_backend_monolitica_modular.md)
    * [ADR-020: Adoção da Arquitetura Hexagonal (Ports and Adapters) para a Implementação dos Módulos](_docs/adr/020-adocao_arquitetura_hexagonal_nos_modulos.md)
    * [ADR-015: Adoção de REST como Estilo Principal para Comunicação Backend-Frontend](_docs/adr/015-adocao_rest_como_estilo_principal_comunicacao_backend_frontend.md)
    * [ADR-016: Adoção do Roteador `go-chi/chi` para Serviços HTTP](_docs/adr/016-adocao_router_chi_para_servicos_http.md)
* **Dados e Persistência:**
    * [ADR-013: Adoção do `database/sql` com Driver `pgx` para Acesso a Dados](_docs/adr/013-adocao_database_sql_e_pgx_para_acesso_dados.md)
    * [ADR-014: Adoção da Ferramenta `goose` para Migrações de Esquema](_docs/adr/014-adocao_ferramenta_goose_para_migracoes_schema.md)
* **Configuração e Observabilidade:**
    * [ADR-018: Adoção de `envconfig` e Arquivos `.env` para Gerenciamento de Configurações](_docs/adr/018-adocao_envconfig_para_gerenciamento_configuracoes.md)
    * [ADR-017: Adoção do Pacote `slog` para Logging Estruturado](_docs/adr/017-adocao_pacote_slog_para_logging_estruturado.md)
* **Qualidade e Testes:**
    * [ADR-019: Estratégia de Testes Unitários em Go com a Biblioteca Testify](_docs/adr/019-testes_unitarios_go_com_testify.md)

*(Consulte o diretório `_docs/adr/` para a lista completa e atualizada.)*

## 3. Primeiros Passos para Desenvolvedores

### 3.1. Pré-requisitos

* Go (versão 1.22 ou superior - *Verificar ADR-011 e ADR-017 para a versão exata*)
* Docker e Docker Compose (para ambiente de desenvolvimento com PostgreSQL e outros serviços)
* `pressly/goose` CLI (para migrações de banco de dados - conforme ADR-014).
    * Instalação: `go install github.com/pressly/goose/v3/cmd/goose@latest`
* Um editor de código com bom suporte a Go (ex: VS Code com a extensão oficial Go).

### 3.2. Configuração do Ambiente de Desenvolvimento Local

1.  **Clone o Repositório:**
    ```bash
    git clone [URL_DO_SEU_REPOSITORIO_GIT_AQUI] redtogreen
    cd redtogreen
    ```

2.  **Configure as Variáveis de Ambiente:**
    * Copie o arquivo de exemplo `.env.example` (localizado na raiz do projeto) para um arquivo `.env` na raiz do projeto:
        ```bash
        cp .env.example .env
        ```
    * Edite o arquivo `.env` com as configurações específicas para o seu ambiente local (especialmente `APP_DATABASE_URL` e `APP_JWT_SECRET_KEY`).
    * **Importante:** O arquivo `.env` está (ou deveria estar) no `.gitignore` e não deve ser comitado no repositório.

3.  **Inicie os Serviços de Apoio (ex: Banco de Dados PostgreSQL):**
    * Assumindo que um arquivo `docker-compose.yml` está configurado na raiz do projeto.
        ```bash
        docker compose up -d
        ```
    * Aguarde alguns instantes para que o banco de dados esteja pronto para aceitar conexões.

4.  **Execute as Migrações do Banco de Dados:**
    * As migrações SQL estão localizadas no diretório `./db/migrations/`.
    * Certifique-se de que a string de conexão no seu `.env` (ex: `APP_DATABASE_URL`) está correta.
    * Execute o comando `goose`:
        ```bash
        # Exemplo de comando (certifique-se que APP_DATABASE_URL está corretamente definida no seu .env):
        goose -dir ./db/migrations postgres "${APP_DATABASE_URL}" up
        ```
        *(Consulte a documentação do `goose` e o ADR-014 para detalhes).*

5.  **Instale as Dependências do Projeto Go:**
    ```bash
    go mod tidy
    go mod download
    ```

6.  **Execute a Aplicação (Backend):**
    * O ponto de entrada principal está em `./cmd/api/main.go`.
    * Para executar:
        ```bash
        go run ./cmd/api/main.go
        ```
    * A aplicação deverá logar que está escutando na porta configurada (padrão 8080, conforme ADR-018).

### 3.3. Executando Testes

* Para rodar todos os testes unitários:
    ```bash
    go test ./...
    ```
* Para visualizar a cobertura de testes:
    ```bash
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    ```
* Consulte [ADR-019](_docs/adr/019-testes_unitarios_go_com_testify.md) para nossa estratégia de testes.

## 4. Estrutura do Projeto (Visão de Alto Nível)

A estrutura do projeto visa seguir as boas práticas da comunidade Go, promovendo modularidade e clareza.

```bash
redtogreen/
├── _docs/                # Documentação geral (ADRs, Visão, Domínio, etc.)
├── api/                  # Definições/especificações de API (ex: OpenAPI)
├── cmd/                  # Aplicações executáveis
│   └── api/              # Ponto de entrada para o servidor da API REST
├── db/
│   └── migrations/       # Arquivos de migração SQL (goose)
├── internal/             # Código privado da aplicação
│   ├── config/
│   ├── core/             # Contextos, Domínios...
│   ├── database/
│   └── transport/
│       └── http/
├── pkg/                  # Bibliotecas reutilizáveis (se houver)
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── LICENSE
└── README.md             # Este documento
```

## 5. Como Contribuir

1.  Leia e compreenda o manifesto os **Princípios Fundamentais Cultura de Arquitetura e Desenvolvimento** ([ADR-000](_docs/adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md)) e os ADRs relevantes.
2.  Siga as **Convenções de Código** (a serem definidas).
3.  Escreva **Testes Unitários** (conforme [ADR-019](_docs/adr/019-testes_unitarios_go_com_testify.md)).
4.  Proponha **Novos ADRs** para decisões arquiteturais significativas.
5.  Utilize **Pull Requests** para todas as alterações de código.
