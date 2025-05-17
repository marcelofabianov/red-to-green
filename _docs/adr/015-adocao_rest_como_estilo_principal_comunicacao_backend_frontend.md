# ADR-015: Adoção de REST como Estilo Principal para Comunicação Backend-Frontend com OpenAPI

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A definição do protocolo e estilo de comunicação entre o backend (Go - ADR-011) e os clientes (principalmente o frontend web) é fundamental para o desenvolvimento, a usabilidade, a performance e a manutenibilidade da API do RedToGreen. Uma interface bem definida é crucial para o desacoplamento entre frontend e backend e para facilitar o desenvolvimento paralelo.

Foram consideradas alternativas como REST, GraphQL e gRPC. A escolha deve priorizar a simplicidade para o desenvolvimento inicial (MVP), ampla adoção, facilidade de consumo por clientes web, boa documentação e a capacidade de evoluir.

A pergunta que estamos tentando responder é: Qual estilo arquitetural e protocolo devemos adotar como principal para a comunicação entre o backend e os clientes do RedToGreen, e como garantiremos que essa API seja bem documentada e compreensível?

O escopo desta decisão abrange a principal interface de comunicação exposta pelo backend do RedToGreen para ser consumida por seus clientes primários (ex: aplicação web frontend).

## Decisão

Adotaremos **REST (Representational State Transfer)** como o estilo arquitetural principal para a API de comunicação entre o backend e o frontend do RedToGreen. A intenção é evoluir o design da API para seguir os princípios **RESTful** à medida que o sistema amadurece.

1.  **Protocolo e Formato de Dados:** A comunicação será realizada sobre **HTTP/S**, utilizando **JSON** como o formato padrão para os corpos de requisição e resposta.
2.  **Princípios RESTful:** Serão empregados os princípios RESTful, incluindo:
    * Uso semântico de métodos HTTP (GET, POST, PUT, DELETE, PATCH).
    * Identificação de recursos através de URLs claras e consistentes.
    * Uso de códigos de status HTTP para indicar o resultado das operações.
    * Statelessness das requisições.
3.  **Documentação com OpenAPI:** A API será documentada utilizando a especificação **OpenAPI**. O arquivo de definição OpenAPI será a fonte da verdade para o contrato da API, facilitando a compreensão, o desenvolvimento de clientes e os testes.
4.  **Flexibilidade para Protocolos Adicionais:** Embora REST seja a principal forma de comunicação, o sistema manterá a flexibilidade para adotar outros protocolos de comunicação (ex: gRPC para comunicação interna entre futuros microsserviços, WebSockets ou Server-Sent Events para funcionalidades de tempo real, WebHooks para integrações com terceiros) para casos de uso específicos onde estes ofereçam vantagens claras.

## Alternativas Consideradas

* **GraphQL:**
    * *Prós:* Permite que o cliente solicite exatamente os dados necessários, evitando over-fetching/under-fetching; endpoint único; fortemente tipado; bom para UIs complexas e evolutivas.
    * *Motivo da Rejeição (como principal forma inicial):* Introduz maior complexidade inicial de setup e desenvolvimento no backend (resolvers, schema) e no frontend (bibliotecas cliente). O caching e o monitoramento podem ser mais complexos. Para as necessidades iniciais do RedToGreen, os benefícios podem não superar a simplicidade e familiaridade do REST. Pode ser considerado para futuras APIs específicas ou para agregar dados de múltiplos serviços, se a arquitetura evoluir para microsserviços.

* **gRPC (Google Remote Procedure Call):**
    * *Prós:* Alta performance devido ao uso de HTTP/2 e Protocol Buffers; contrato forte e tipado; ideal para comunicação interna entre microsserviços; suporte a streaming bidirecional.
    * *Motivo da Rejeição (como principal forma de comunicação com o frontend):* Menos amigável para consumo direto por navegadores web sem uma camada de proxy (como gRPC-Web), o que adiciona complexidade. A serialização binária dificulta a depuração manual sem ferramentas específicas. Embora excelente para comunicação inter-serviços, REST é mais direto para a comunicação cliente-servidor web inicial.

* **Outras Abordagens RPC (ex: JSON-RPC, XML-RPC):**
    * *Prós:* Podem ser simples para chamadas de procedimento.
    * *Motivo da Rejeição:* Menos padronizadas e com ecossistema de ferramentas (especialmente para documentação e gateways de API) menos rico que REST com OpenAPI. Não oferecem os mesmos benefícios de cacheabilidade e uniformidade de interface do REST.

## Consequências

**Positivas:**
* **Simplicidade e Familiaridade:** REST é um padrão amplamente conhecido e compreendido por desenvolvedores, facilitando a curva de aprendizado e a integração de novos membros na equipe.
* **Vasto Ecossistema:** Grande quantidade de ferramentas, bibliotecas e frameworks disponíveis para todas as linguagens (incluindo Go para o backend e JavaScript/TypeScript para o frontend) que suportam REST e JSON.
* **Facilidade de Consumo por Clientes Web:** Navegadores e frameworks frontend modernos consomem APIs REST/JSON de forma nativa e eficiente.
* **Documentação Clara com OpenAPI:** A adoção de OpenAPI garante um contrato de API bem definido, facilitando a colaboração entre equipes de backend e frontend, a geração de SDKs de cliente e a automação de testes.
* **Statelessness e Cacheability:** Facilita a escalabilidade horizontal do backend e permite o uso de mecanismos de caching HTTP para melhorar a performance.
* **Boa Performance para Casos de Uso Comuns:** Para operações CRUD e interações cliente-servidor típicas, REST oferece boa performance.
* **Flexibilidade para Evolução:** A decisão não impede a adoção de outros protocolos para necessidades específicas, permitindo uma arquitetura de comunicação híbrida e otimizada no futuro.
* **Interoperabilidade:** APIs REST são facilmente consumidas por uma ampla gama de clientes e serviços.

**Negativas / Trade-offs:**
* **Potencial para Over-fetching/Under-fetching:** Com APIs REST tradicionais, o cliente pode receber mais dados do que precisa ou precisar fazer múltiplas chamadas para obter todos os dados necessários. Um bom design de API (com parâmetros de consulta para seleção de campos, endpoints específicos) pode mitigar isso parcialmente.
* **Gerenciamento de Múltiplos Endpoints:** À medida que a aplicação cresce, o número de endpoints pode aumentar, exigindo uma boa organização e documentação (que o OpenAPI ajuda a manter).
* **Versionamento da API:** Requer uma estratégia clara para versionamento (ex: na URL, headers) para gerenciar alterações que quebram a compatibilidade, o que pode adicionar complexidade.
* **Menos Eficiente para Dados Altamente Interconectados ou em Tempo Real (comparado a GraphQL/WebSockets):** Para buscar grafos de dados complexos em uma única requisição ou para comunicação bidirecional em tempo real, REST pode ser menos eficiente ou exigir soluções complementares.

## (Opcional) Notas Adicionais
* A especificação OpenAPI será mantida como parte do código-fonte do backend e idealmente integrada ao processo de CI/CD para validação e publicação da documentação.
* Serão exploradas ferramentas para gerar código (ex: stubs de servidor, modelos de dados) a partir da especificação OpenAPI para Go, visando aumentar a consistência e reduzir boilerplate (ex: `oapi-codegen`).
