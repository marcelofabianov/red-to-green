# ADR-017: Adoção do Pacote `log/slog` para Logging Estruturado

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Um sistema de logging eficaz é crucial para a observabilidade, depuração, monitoramento e segurança da aplicação RedToGreen. Os logs precisam ser estruturados para facilitar a análise por máquinas e ferramentas de agregação, devem suportar níveis de severidade e permitir a inclusão de contexto rico para rastreabilidade.

Considerando a linguagem Go (ADR-011), a necessidade de integração futura com OpenTelemetry (OTel) para tracing distribuído (com ferramentas como Jaeger) e a preferência por utilizar soluções da biblioteca padrão sempre que possível (ADR-000), a escolha da biblioteca de logging deve refletir esses requisitos. A saída em formato JSON é um requisito chave para a integração com sistemas modernos de gerenciamento de logs.

A pergunta que estamos tentando responder é: Qual biblioteca e abordagem de logging devemos adotar para o RedToGreen, garantindo logs estruturados em JSON, performáticos, contextuais e preparados para integração com um ecossistema de observabilidade moderno, priorizando soluções nativas do Go?

O escopo desta decisão abrange a principal biblioteca e as práticas de logging a serem implementadas no backend Go do RedToGreen.

## Decisão

Adotaremos o pacote **`log/slog`** (disponível a partir do Go 1.21) como a biblioteca padrão e principal para logging estruturado no sistema RedToGreen.

1.  **Formato de Saída Principal:** Os logs serão formatados como **JSON** utilizando o `slog.JSONHandler` embutido. Esta saída será direcionada para `os.Stdout` (ou `os.Stderr` para logs de erro ou de nível mais alto, conforme a necessidade). Este formato facilita a coleta por sistemas de gerenciamento de logs em ambientes de contêiner (Docker, Kubernetes) e a integração com plataformas de observabilidade (ex: Elastic Stack, Splunk, Datadog, Grafana Loki, Google Cloud Logging, AWS CloudWatch Logs).
2.  **Níveis de Log:** Serão utilizados os níveis de log padrão definidos pelo `slog` (DEBUG, INFO, WARN, ERROR). O nível de log ativo será configurável (ex: via variável de ambiente), com `INFO` como padrão para ambientes de produção e `DEBUG` para ambientes de desenvolvimento e teste.
3.  **Contexto Estruturado e Atributos:** Todas as mensagens de log deverão incluir campos contextuais relevantes como pares chave-valor no objeto JSON.
    * **Atributos Padrão por Requisição:** Para logs dentro do contexto de uma requisição HTTP, atributos como `request_id` (para rastrear uma requisição específica), e, quando aplicável, `user_id`.
    * **Atributos Específicos da Operação:** Outros atributos relevantes para a operação sendo logada (ex: `wallet_id`, `transaction_id`, parâmetros chave) serão adicionados conforme necessário.
    * **Informações de Erro:** Ao logar erros, o log deve incluir a mensagem de erro (via `err.Error()`, idealmente utilizando a estrutura definida no ADR-005) e, para erros inesperados ou panics recuperados, um stack trace.
4.  **Integração com `context.Context`:** O `context.Context` da requisição será utilizado para propagar um logger enriquecido com atributos contextuais (como `request_id` e, futuramente, `trace_id` e `span_id` do OpenTelemetry) através das camadas da aplicação.
5.  **Preparação para OpenTelemetry (OTel):** A escolha do `slog` e a prática de incluir IDs de rastreamento (como `request_id` inicialmente, e depois `trace_id`, `span_id` do OTel) no contexto dos logs são passos preparatórios para uma futura integração robusta com o OpenTelemetry. Isso permitirá a correlação eficaz entre logs e traces distribuídos (visualizados em ferramentas como Jaeger).

## Alternativas Consideradas

* **Pacote `log` da Biblioteca Padrão (anterior ao `slog`):**
    * *Prós:* Sem dependências externas, simples.
    * *Motivo da Rejeição:* Não oferece logging estruturado nativo (especialmente JSON), nem níveis de log ou fácil adição de contexto, tornando-o inadequado para as necessidades de observabilidade de uma aplicação moderna e para integração com sistemas de gerenciamento de logs.

* **`uber-go/zap`:**
    * *Prós:* Performance excepcional, logging estruturado robusto, maturidade e ampla adoção na comunidade.
    * *Motivo da Rejeição:* Embora seja uma excelente biblioteca, a disponibilidade do `slog` na biblioteca padrão (Go 1.21+) oferece uma solução nativa, performática e alinhada com a preferência por minimizar dependências externas (Princípio Fundamental do ADR-000), sem sacrificar as funcionalidades essenciais de logging estruturado em JSON. A API do `slog` também é projetada para ser idiomática e de fácil adoção. A diferença de performance, embora `zap` possa ser marginalmente superior em benchmarks, não é considerada crítica para o RedToGreen a ponto de justificar uma dependência externa sobre a solução da stdlib.

* **`rs/zerolog`:**
    * *Prós:* Alta performance, API fluida para logging JSON, popular.
    * *Motivo da Rejeição:* Semelhante ao `zap`, é uma biblioteca de terceiros de alta qualidade. No entanto, com a introdução do `slog` na biblioteca padrão, a necessidade de uma dependência externa para logging estruturado em JSON de alta qualidade diminuiu, e `slog` se torna a escolha preferencial por ser nativo e atender aos requisitos.

* **Uso de `fmt.Print*` para `os.Stdout`/`os.Stderr`:**
    * *Prós:* Simplicidade extrema para depuração local rápida.
    * *Motivo da Rejeição:* Completamente inadequado para logging de produção devido à falta de estrutura, níveis, contexto, controle de saída e capacidade de integração com sistemas de gerenciamento de logs.

## Consequências

**Positivas:**
* **Logging Estruturado JSON Nativo:** Utiliza a solução padrão do Go para logs estruturados, garantindo boa performance e integração com o ecossistema da linguagem, com saída JSON ideal para sistemas de análise.
* **Melhoria na Depuração e Rastreabilidade:** Logs contextuais e estruturados permitem identificar problemas e rastrear fluxos de requisição de forma mais eficiente.
* **Redução de Dependências:** Sendo parte da biblioteca padrão (Go 1.21+), reduz a necessidade de gerenciar uma dependência externa para uma funcionalidade core como logging.
* **Preparação para Observabilidade Avançada:** A estrutura e o conteúdo dos logs facilitam a futura integração com OpenTelemetry para correlação com traces (visualizados em ferramentas como Jaeger) e métricas.
* **Consistência e Padronização:** Promove um padrão de logging consistente em toda a aplicação, alinhado com as ferramentas e práticas modernas do Go.

**Negativas / Trade-offs:**
* **Requisito de Versão do Go:** Exige o uso do Go 1.21 ou superior (o que é uma premissa para este novo projeto, RedToGreen).
* **Curva de Aprendizagem (Menor):** A equipe precisará se familiarizar com a API e os conceitos do `slog` (ex: `slog.Attr`, `slog.Handler`), embora seja projetado para ser idiomático e de fácil adoção.
* **Ecossistema de Handlers de Terceiros:** Embora o `slog` seja extensível através da interface `slog.Handler`, o ecossistema de handlers de terceiros altamente especializados pode ainda estar em crescimento em comparação com bibliotecas mais antigas como `zap` ou `zerolog`. No entanto, os handlers JSON e Text embutidos cobrem as necessidades primárias da maioria das aplicações.

## (Opcional) Notas Adicionais
* A configuração do nível de log (DEBUG, INFO, WARN, ERROR) será gerenciada através de variáveis de ambiente ou um sistema de configuração centralizado (a ser definido em um ADR futuro, ex: ADR sobre Configuração da Aplicação).
* **Implementação de Logger por Requisição:** Será implementado um middleware HTTP (utilizando `chi`, conforme ADR-016) responsável por:
    * Gerar ou obter um `request_id` único para cada requisição HTTP.
    * Futuramente, quando OpenTelemetry for integrado, extrair `trace_id` e `span_id` do contexto OTel.
    * Criar uma instância de `slog.Logger` (ou obter do pool) enriquecida com esses IDs e outros dados da requisição (ex: método HTTP, path, IP do cliente).
    * Injetar este logger específico da requisição no `context.Context` da requisição, para que possa ser acessado por todos os handlers de rota e camadas de serviço subsequentes.
* Recomenda-se a definição de campos de log padronizados (convenções de nomenclatura) para atributos comuns (ex: `error_message`, `stack_trace`, `duration_ms`) para facilitar as consultas, a criação de dashboards e alertas.
* A política de rotação e retenção de logs, caso os logs sejam temporariamente armazenados em arquivos localmente antes da coleta por um agente, será definida como parte da estratégia de deploy e infraestrutura. Em ambientes de nuvem, a coleta e retenção são geralmente gerenciadas pelo serviço de logging da plataforma (ex: CloudWatch Logs, Google Cloud Logging).
