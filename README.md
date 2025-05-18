# RedToGreen - Seu Dinheiro, Suas Regras, Sem Complicação!

Bem-vindo ao repositório do projeto RedToGreen! Esta aplicação tem como objetivo ajudar usuários a organizar suas finanças pessoais de forma simples, intuitiva e eficaz.

## Visão Geral do Produto

O RedToGreen é uma ferramenta de controle financeiro pessoal projetada com foco na simplicidade e facilidade de uso. Nossa missão é capacitar os usuários, especialmente aqueles com pouca experiência em gestão financeira, a transformar sua relação com o dinheiro, saindo de um cenário de descontrole ("vermelho") para uma situação de clareza, planejamento e tranquilidade ("verde").

Para uma descrição completa da proposta de valor, público-alvo e funcionalidades planejadas (incluindo o MVP), por favor, consulte o documento de Visão Geral do Produto:

* **Documento Principal:** [`_docs/produto/visao_geral_v1.md`](_docs/produto/visao_geral_v1.md)

## Nossa Cultura de Arquitetura e Desenvolvimento

No RedToGreen, pautamos nosso desenvolvimento por uma cultura que valoriza a simplicidade, o pragmatismo, a robustez e a capacidade de evoluir de forma sustentável. Nossos princípios fundamentais estão codificados e servem como guia para todas as nossas decisões técnicas.

Convidamos você a ler nosso manifesto de princípios:

* **Princípios Fundamentais:** [`_docs/adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md`](_docs/adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md)

## Decisões Arquiteturais (ADRs)

Todas as decisões arquiteturais significativas que moldam o RedToGreen são documentadas utilizando o padrão Architecture Decision Record (ADR). Esses registros fornecem o contexto da decisão, a solução escolhida, as alternativas que foram consideradas e as consequências esperadas.

A leitura dos ADRs é essencial para compreender a fundo a estrutura técnica e as razões por trás das escolhas feitas. Você pode encontrar todos os ADRs no diretório:

* **Índice de ADRs:** [`_docs/adr/`](_docs/adr/)

Alguns ADRs chave para iniciar seu entendimento do projeto incluem:

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

## Primeiros Passos para Desenvolvedores

### Pré-requisitos

* Go (versão 1.22 ou superior - *Verificar ADR-011 e ADR-017 para a versão exata*)
* Docker e Docker Compose (para ambiente de desenvolvimento com PostgreSQL e outros serviços)
* `pressly/goose` CLI (para migrações de banco de dados - conforme ADR-014).
    * Instalação: `go install github.com/pressly/goose/v3/cmd/goose@latest`
* Um editor de código com bom suporte a Go (ex: VS Code com a extensão oficial Go).

### Configuração do Ambiente de Desenvolvimento Local

1.  **Clone o Repositório:**
    ```bash
    git clone [URL_DO_SEU_REPOSITORIO_GIT_AQUI] redtogreen
    cd redtogreen
    ```

2.  **Configure as Variáveis de Ambiente:**
    * Copie o arquivo de exemplo `.env.example` (localizado na raiz do projeto ou em `_docs/exemplos/`) para um arquivo `.env` na raiz do projeto:
        ```bash
        cp .env.example .env
        ```
    * Edite o arquivo `.env` com as configurações específicas para o seu ambiente local (especialmente `APP_DATABASE_URL` e `APP_JWT_SECRET_KEY`).
    * **Importante:** O arquivo `.env` está (ou deveria estar) no `.gitignore` e não deve ser comitado no repositório.

3.  **Inicie os Serviços de Apoio (ex: Banco de Dados PostgreSQL):**
    * Assumindo que um arquivo `docker-compose.yml` está configurado na raiz do projeto para gerenciar serviços como o PostgreSQL.
        ```bash
        docker compose up -d
        ```
    * Aguarde alguns instantes para que o banco de dados esteja pronto para aceitar conexões.

4.  **Execute as Migrações do Banco de Dados:**
    * As migrações SQL estão localizadas no diretório `./db/migrations/`.
    * Certifique-se de que a string de conexão no seu `.env` (ex: `APP_DATABASE_URL`) está correta.
    * Execute o comando `goose` para aplicar as migrações:
        ```bash
        # Exemplo de comando (certifique-se que APP_DATABASE_URL está corretamente definida no seu .env):
        goose -dir ./db/migrations postgres "${APP_DATABASE_URL}" up
        ```
        *(Consulte a documentação do `goose` e o ADR-014 para detalhes sobre a configuração e execução).*

5.  **Instale as Dependências do Projeto Go:**
    ```bash
    go mod tidy
    go mod download
    ```

6.  **Execute a Aplicação (Backend):**
    * O ponto de entrada principal da aplicação backend provavelmente estará em um subdiretório de `cmd/` (ex: `cmd/api/main.go` ou `cmd/server/main.go`).
    * Para executar:
        ```bash
        go run ./cmd/api/main.go
        ```
    * A aplicação deverá iniciar e logar que está escutando na porta configurada (padrão 8080, conforme ADR-018).

### Executando Testes

* Para rodar todos os testes unitários do projeto:
    ```bash
    go test ./...
    ```
* Para visualizar a cobertura de testes:
    ```bash
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    ```
* Consulte [ADR-019: Estratégia de Testes Unitários em Go com a Biblioteca Testify](_docs/adr/019-testes_unitarios_go_com_testify.md) para mais detalhes sobre nossa abordagem de testes.
