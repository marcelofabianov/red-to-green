# ADR-016: Adoção do Roteador `go-chi/chi` para Serviços HTTP

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Com a decisão de utilizar REST como o estilo principal para a API (ADR-015) e Go como linguagem de backend (ADR-011), é necessário selecionar um roteador HTTP (multiplexer) para gerenciar as requisições HTTP, aplicar middlewares e despachar para os handlers apropriados. A biblioteca padrão `net/http` do Go oferece funcionalidades básicas de roteamento, mas para APIs mais complexas, um roteador mais completo pode aumentar a produtividade e a organização do código.

Os critérios chave para esta decisão incluem:
* **Compatibilidade com `net/http`:** Utilização das interfaces padrão do Go (`http.Handler`, `http.HandlerFunc`) para garantir interoperabilidade com o ecossistema existente.
* **Simplicidade e Design Idiomático:** Facilidade de uso e alinhamento com as práticas comuns em Go.
* **Leveza e Performance:** Mínimo overhead e bom desempenho.
* **Funcionalidades de Roteamento:** Suporte a parâmetros de rota, roteamento por método HTTP, agrupamento de rotas.
* **Sistema de Middlewares:** Capacidade de aplicar lógica transversal (logging, autenticação, CORS, etc.) de forma flexível.
* **Manutenibilidade e Ecossistema:** Projeto ativo, bem documentado e com boa adoção pela comunidade.

A pergunta que estamos tentando responder é: Qual roteador HTTP Go devemos adotar para construir nossa API REST, que equilibre funcionalidades, performance, simplicidade e compatibilidade com o ecossistema `net/http`?

O escopo desta decisão abrange a biblioteca principal a ser utilizada para roteamento de requisições HTTP e aplicação de middlewares no backend do RedToGreen.

## Decisão

Adotaremos o roteador **`go-chi/chi`** (especificamente `github.com/go-chi/chi/v5`) como a biblioteca principal para o roteamento de requisições HTTP e gerenciamento de middlewares no backend do RedToGreen.

**Justificativas para a escolha do `go-chi/chi`:**
1.  **Total Compatibilidade com `net/http`:** `chi` é construído sobre as interfaces padrão do Go (`http.Handler`, `http.HandlerFunc`). Isso permite o uso transparente de qualquer middleware ou handler compatível com `net/http`, garantindo acesso a um vasto ecossistema de bibliotecas e promovendo a interoperabilidade.
2.  **Simplicidade e Design Idiomático:** A API do `chi` é considerada elegante, concisa e segue as melhores práticas do Go. É fácil de aprender e usar, sem introduzir complexidade desnecessária.
3.  **Leveza e Performance:** `chi` é um roteador leve que adiciona um overhead mínimo e é conhecido por sua excelente performance, adequada para aplicações que exigem baixa latência.
4.  **Sistema de Middlewares Flexível e Poderoso:** `chi` oferece um sistema robusto para encadear e agrupar middlewares. Middlewares podem ser aplicados globalmente, a grupos de rotas (sub-roteadores) ou a rotas individuais. O pacote `chi/middleware` já fornece um conjunto útil de middlewares prontos para uso (ex: Logger, Recoverer, CORS, RealIP, RequestID, Timeout, Compress).
5.  **Funcionalidades de Roteamento Completas:** Suporta parâmetros de URL nomeados, wildcards, roteamento baseado em método HTTP, e agrupamento de rotas, facilitando a organização de APIs complexas e RESTful.
6.  **Integração com `context.Context`:** Utiliza o `context.Context` do Go de forma eficaz para passar informações específicas da requisição (como parâmetros de rota) através da cadeia de handlers e middlewares.
7.  **Comunidade Ativa e Boa Documentação:** `chi` é um projeto bem mantido, com uma comunidade ativa e documentação clara.
8.  **Controle e Não Opinativo:** `chi` foca em ser um excelente roteador e sistema de middlewares, sem se tornar um framework web pesado e opinativo. Isso dá ao desenvolvedor o controle sobre a arquitetura da aplicação, alinhando-se com a preferência por simplicidade e controle.

## Alternativas Consideradas

* **`net/http` (Biblioteca Padrão):**
    * *Prós:* Sem dependências externas, estável, performático.
    * *Motivo da Rejeição:* Funcionalidades de roteamento muito básicas (sem suporte nativo a parâmetros de rota ou middlewares componíveis de forma elegante), levando a boilerplate significativo para APIs RESTful mais complexas.

* **`gin-gonic/gin`:**
    * *Prós:* Alta performance, popular, rico em funcionalidades (binding, rendering).
    * *Motivo da Rejeição:* Embora performático, o uso de um `gin.Context` próprio, que encapsula `http.Request` e `http.ResponseWriter`, pode tornar a interoperabilidade com middlewares `net/http` puros menos direta em comparação com `chi`. A preferência é por uma integração mais nativa com as interfaces padrão.

* **`gorilla/mux`:**
    * *Prós:* Maduro, estável, roteamento poderoso.
    * *Motivo da Rejeição:* Preocupações sobre o ritmo de manutenção ativa do projeto em alguns períodos e uma API que pode ser percebida como menos fluida ou mais verbosa em comparação com alternativas mais modernas como `chi`.

* **`go-fiber/fiber`:**
    * *Prós:* Inspirado no Express.js, alta performance (usa `fasthttp`).
    * *Motivo da Rejeição:* Construído sobre `fasthttp`, o que o torna **incompatível** com o ecossistema `net/http` (`http.Handler`, etc.). Essa incompatibilidade é um fator decisivo contra, dada a importância da interoperabilidade com o ecossistema padrão do Go.

* **`julienschmidt/httprouter`:**
    * *Prós:* Performance extrema, baixo overhead.
    * *Motivo da Rejeição:* A assinatura de sua função de handler não é diretamente compatível com `http.HandlerFunc`, exigindo wrappers para usar middlewares `net/http` padrão, o que reduz a conveniência da compatibilidade direta buscada.

## Consequências

**Positivas:**
* **Desenvolvimento Produtivo de APIs RESTful:** Facilita a criação de APIs bem estruturadas, com bom suporte a middlewares e roteamento avançado.
* **Manutenção da Compatibilidade com o Ecossistema Go:** Permite o uso irrestrito de middlewares e handlers `net/http` existentes.
* **Código Limpo e Organizado:** A API do `chi` incentiva a escrita de código claro e modular.
* **Performance Confiável:** Adequado para aplicações que necessitam de boa performance e baixa latência.
* **Curva de Aprendizagem Suave:** Fácil de adotar para desenvolvedores Go, especialmente aqueles já familiarizados com `net/http`.
* **Controle sobre a Stack:** `chi` não impõe uma estrutura de aplicação rígida, permitindo que a equipe defina a arquitetura.

**Negativas / Trade-offs:**
* **Adição de uma Dependência Externa:** Embora `chi` seja leve e estável, é uma dependência externa a ser gerenciada.
* **Foco em Roteamento e Middlewares:** `chi` não é um framework "full-stack" com todas as baterias inclusas (como ORM, sistema de templates complexo, etc.). Para funcionalidades além de roteamento e middlewares HTTP, outras bibliotecas precisarão ser integradas (o que é intencional e alinhado com a filosofia de compor software em Go).

## (Opcional) Notas Adicionais
* A versão de `chi` a ser utilizada será a `v5` ou a mais recente estável disponível.
* A estrutura de organização dos handlers e middlewares seguirá as boas práticas recomendadas pela comunidade Go e `chi` para manter o código modular e testável.
* O conjunto de middlewares padrão do `chi/middleware` será avaliado e utilizado conforme a necessidade (ex: `middleware.Logger`, `middleware.Recoverer`, `middleware.RequestID`, `middleware.RealIP`, `middleware.Compress`, `middleware.Timeout`, `middleware.StripSlashes`, `middleware.Recoverer`).
