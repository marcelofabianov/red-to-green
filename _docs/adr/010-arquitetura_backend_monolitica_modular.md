# ADR-010: Adoção de Arquitetura Monolítica Modular para o Backend

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Para o produto RedToGreen, é fundamental definir a arquitetura macro do componente backend. A arquitetura escolhida deve suportar os requisitos iniciais do MVP, priorizando a simplicidade de desenvolvimento e deploy, ao mesmo tempo em que oferece um caminho claro para evolução, manutenibilidade e escalabilidade futuras. É crucial evitar a criação de um "monolito big ball of mud" que se torne difícil de gerenciar com o tempo.

O RedToGreen visa ser um produto evolutivo, começando com um escopo simples e crescendo com novas features. A arquitetura do backend deve facilitar essa evolução incremental.

A pergunta que estamos tentando responder é: Qual arquitetura de backend devemos adotar para o RedToGreen que ofereça um bom equilíbrio entre agilidade no desenvolvimento inicial, simplicidade operacional e capacidade de adaptação e crescimento sustentável?

O escopo desta decisão abrange a estrutura de alto nível e os princípios de design para o backend da aplicação RedToGreen.

## Decisão

O backend da aplicação RedToGreen será desenvolvido inicialmente como uma **aplicação monolítica** com um forte foco em **modularidade interna**.

1.  **Estrutura Monolítica Inicial:** A aplicação backend será um único processo/serviço deployável. Isso simplifica o desenvolvimento, o teste, o deploy e o monitoramento nas fases iniciais do projeto.

2.  **Design Modular Interno (Monolito Modular):**
    * A base de código do monolito será organizada em **módulos ou contextos de domínio** bem definidos e com baixo acoplamento entre si. Exemplos de módulos podem incluir: "Gerenciamento de Usuários", "Gerenciamento de Carteiras", "Gerenciamento de Transações", etc.
    * Cada módulo deve ter alta coesão interna, encapsulando sua lógica de negócios específica, entidades de domínio e, idealmente, suas abstrações de acesso a dados.
    * A comunicação entre os módulos dentro do monolito deve ocorrer através de **interfaces claramente definidas** (ex: chamadas de método via interfaces Go, canais para eventos síncronos/assíncronos internos). Deve-se evitar o acoplamento direto a implementações concretas de outros módulos.

3.  **Preparação para Evolução Futura:**
    * Esta arquitetura modular visa facilitar a manutenção contínua e a evolução do monolito, permitindo que diferentes partes do sistema sejam desenvolvidas, testadas e compreendidas de forma mais independente.
    * A clara separação de módulos e suas interfaces bem definidas manterá a opção de, se e quando for estrategicamente vantajoso (devido a requisitos de escalabilidade, tamanho da equipe, ou complexidade de um módulo específico), **extrair um ou mais módulos para se tornarem microsserviços separados**. Essa extração seria facilitada pelo baixo acoplamento já existente, minimizando a necessidade de refatorações extensas. No entanto, a migração para microsserviços não é um objetivo primário inicial, mas uma possibilidade que a arquitetura escolhida não impede, e sim facilita.

## Alternativas Consideradas (Opcional)

* **Arquitetura de Microsserviços desde o Início:**
    * Descrição: Desenvolver o backend como um conjunto de múltiplos serviços pequenos, independentes e deployáveis separadamente desde o começo.
    * Motivo da Rejeição: Introduz uma alta complexidade operacional (deploy, monitoramento distribuído, resiliência da comunicação inter-serviços, consistência de dados eventual, gerenciamento de infraestrutura) e de desenvolvimento que é geralmente desproporcional para um produto em estágio inicial (MVP) e para equipes que podem ser pequenas. O "imposto arquitetural" dos microsserviços é significativo.

* **Monolito Tradicional (sem foco explícito em modularidade):**
    * Descrição: Desenvolver uma aplicação monolítica sem uma preocupação rigorosa com a separação interna de responsabilidades em módulos com baixo acoplamento.
    * Motivo da Rejeição: Esta abordagem frequentemente leva a um "big ball of mud", onde o código se torna altamente acoplado, difícil de entender, manter, testar e evoluir. Dificulta a identificação de limites para futuras extrações de serviços e pode levar a gargalos de desenvolvimento.

* **Arquitetura Serverless (Functions as a Service - FaaS) como base principal:**
    * Descrição: Construir a totalidade ou a maior parte do backend utilizando funções serverless.
    * Motivo da Rejeição (como arquitetura principal): Embora o FaaS seja excelente para certas tarefas (ex: processamento de eventos, APIs simples), construir um sistema de gestão completo com lógica de domínio potencialmente complexa e estado pode levar a desafios de orquestração, "cold starts", e gerenciamento de estado. Um monolito modular oferece uma base mais coesa para o domínio principal nesta fase.

## Consequências

**Positivas:**

* **Simplicidade de Desenvolvimento Inicial:** Reduz a carga cognitiva e a complexidade de setup, build, e desenvolvimento em comparação com arquiteturas distribuídas.
* **Simplicidade de Deploy e Operação Inicial:** Um único artefato a ser deployado e monitorado simplifica a infraestrutura e as operações no início.
* **Performance:** A comunicação intra-processo entre módulos bem definidos dentro de um monolito é geralmente muito mais rápida e menos sujeita a falhas de rede do que a comunicação inter-serviços.
* **Consistência de Dados:** É mais simples garantir transações ACID e consistência forte dos dados quando se opera majoritariamente sobre um único banco de dados.
* **Facilidade de Refatoração Interna:** Reorganizar código e responsabilidades entre módulos dentro de um mesmo monolito é significativamente mais fácil do que em um ambiente distribuído.
* **Caminho de Evolução Claro:** A modularidade interna e o baixo acoplamento preparam o sistema para uma evolução sustentável, seja como um "monolito modular evolutivo" ou através da extração de microsserviços no futuro, se necessário.
* **Menor Custo Operacional Inicial:** Menos componentes distintos para gerenciar, escalar e monitorar.
* **Iteração Rápida:** Permite que a equipe foque na entrega de valor de negócio rapidamente durante as fases iniciais do produto.

**Negativas / Trade-offs:**

* **Escalabilidade Granular Limitada:** O monolito como um todo precisa ser escalado horizontalmente, mesmo que apenas um de seus módulos esteja sob alta carga. (A modularidade pode, no entanto, permitir otimizações de performance mais direcionadas dentro do monolito).
* **Adoção de Stack Tecnológica Única (Predominante):** Embora não seja uma regra absoluta, os monolitos tendem a ser construídos predominantemente com uma única linguagem de programação e stack tecnológica principal, limitando a flexibilidade de usar tecnologias diferentes para módulos distintos.
* **Impacto de Falhas Potencialmente Maior:** Uma falha crítica não contida em um módulo mal isolado tem o potencial de afetar toda a aplicação monolítica. Um design modular robusto com bom tratamento de erros pode mitigar isso.
* **Tempo de Build e Testes Completos:** À medida que o monolito cresce, os tempos de build da aplicação completa e a execução de todos os testes podem aumentar. A modularidade permite, no entanto, testar módulos de forma mais isolada.
* **Disciplina de Modularidade Requerida:** Manter o baixo acoplamento e a alta coesão entre os módulos exige disciplina contínua da equipe de desenvolvimento e uma arquitetura interna clara (ex: uso de Portas e Adaptadores, Injeção de Dependência). O risco de "vazamentos" e aumento do acoplamento entre módulos sempre existe.

**(Opcional) Notas Adicionais:**

* A adoção de padrões arquiteturais como a Arquitetura Hexagonal (Portas e Adaptadores) ou Clean Architecture *dentro* de cada módulo pode ser fortemente considerada para reforçar o desacoplamento, a testabilidade e a separação de preocupações.
* O uso de ferramentas de análise estática de código pode ajudar a monitorar e impor os limites entre os módulos, prevenindo acoplamentos indesejados.
* A estratégia de comunicação entre módulos deve ser bem definida, priorizando interfaces e, possivelmente, um broker de eventos interno para cenários que se beneficiem de maior desacoplamento assíncrono (mesmo dentro do monolito).
