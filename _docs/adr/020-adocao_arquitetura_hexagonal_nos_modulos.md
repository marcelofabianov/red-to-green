# ADR-020: Adoção da Arquitetura Hexagonal (Ports and Adapters) para a Implementação dos Módulos

**Status:** Aceito

**Data:** 2025-05-18

## Contexto

Após a decisão de implementar uma Arquitetura Monolítica Modular para o backend (ADR-010) e a escolha da linguagem Go (ADR-011), torna-se necessário definir o padrão arquitetural interno para cada módulo do sistema RedToGreen. Este padrão deve orientar como estruturar o código dentro de cada módulo, garantindo consistência, manutenibilidade e alinhamento com os princípios fundamentais estabelecidos (ADR-000).

A escolha do padrão arquitetural interno tem impacto direto na:
- Qualidade e legibilidade do código
- Facilidade de manutenção e evolução do sistema
- Capacidade de testar componentes isoladamente
- Possibilidade de substituir implementações específicas sem alterar a lógica de negócio
- Preparação para possível extração de módulos para microsserviços no futuro

A pergunta que estamos tentando responder é: Qual padrão arquitetural interno deve ser adotado para cada módulo do sistema RedToGreen que melhor se alinha com os princípios de simplicidade, modularidade, testabilidade e evolução sustentável já estabelecidos?

O escopo desta decisão abrange a estrutura interna de cada módulo do monolito modular, definindo como as responsabilidades serão organizadas e como os componentes se comunicarão entre si dentro de cada contexto de domínio.

## Decisão

Adotaremos a **Arquitetura Hexagonal (Ports and Adapters)** como o padrão arquitetural para implementação interna de cada módulo do sistema RedToGreen.

A implementação seguirá estas diretrizes:

1. **Estrutura funcional alinhada com convenções Go:**
   - **Domínio:** Entidades e regras de negócio (`type.go`, `[module].go`)
   - **Portas:** Interfaces de entrada e saída (`port.go`)
   - **Aplicação:** Casos de uso que orquestram operações de negócio (`/usecase/`)
   - **Adaptadores:** Implementações concretas para interfaces de entrada (`/handler/`) e saída (`/repository/`)

2. **Portas bem definidas:**
   - **Portas primárias (driving/entrada):** Interfaces que definem como o mundo externo pode interagir com o módulo
   - **Portas secundárias (driven/saída):** Interfaces que definem como o módulo interage com serviços externos (banco de dados, outros módulos, etc.)

3. **Adaptadores desacoplados:**
   - **Adaptadores primários (handlers):** Implementam as portas de entrada (ex: API REST, CLI, consumidores de eventos)
   - **Adaptadores secundários (repositories):** Implementam as portas de saída (ex: repositórios de banco de dados, clientes HTTP, publishers de eventos)

4. **Dependências apontando para dentro:**
   - O domínio (definido em `type.go` e `[module].go`) não depende de nenhuma camada externa
   - As camadas externas dependem das interfaces (portas) definidas em `port.go`
   - Aplicação de injeção de dependência para resolver as dependências em tempo de execução

5. **Estrutura de pastas sugerida por módulo, alinhada com as convenções da comunidade Go:**
```
/internal/[nome-do-modulo]/       # Ex: internal/wallet
  /handler/                       # Adaptadores primários (HTTP handlers)
    create_[entidade]_handler.go    # Ex: create_wallet_handler.go
    update_[entidade]_handler.go    # Ex: update_wallet_handler.go
    list_[entidade]_handler.go      # Ex: list_wallet_handler.go
  /repository/                    # Adaptadores secundários (implementações de repositório)
    postgres_[entidade]_repository.go # Ex: postgres_wallet_repository.go
    memory_[entidade]_repository.go # Ex: memory_wallet_repository.go (para testes)
  /usecase/                       # Casos de uso da aplicação
    create_[entidade].go            # Ex: create_wallet.go
    update_[entidade].go            # Ex: update_wallet.go
    list_[entidade].go              # Ex: list_wallet.go
  port.go                         # Interfaces (portas) de entrada e saída
  type.go                         # Definições de tipos e estruturas do domínio
  msg.go                          # Mensagens, erros e constantes específicas do módulo
  [entidade].go                     # Ex: wallet.go - Lógica de domínio principal e regras de negócio
```

## Alternativas Consideradas

* **Clean Architecture:**
    * Descrição: Arquitetura em camadas concêntricas (Entities, Use Cases, Interface Adapters, Frameworks) com regra de dependência apontando para o centro.
    * Motivo da Rejeição: Embora bastante similar à Arquitetura Hexagonal, apresenta uma complexidade ligeiramente maior em termos de número de camadas. A Arquitetura Hexagonal oferece os mesmos benefícios de proteção do domínio com uma estrutura um pouco mais simples, alinhando-se melhor ao princípio de simplicidade e pragmatismo (KISS, YAGNI) do ADR-000.

* **Arquitetura em Camadas Tradicional:**
    * Descrição: Divisão em camadas horizontais (Apresentação, Aplicação, Domínio, Infraestrutura) com dependências diretas entre camadas adjacentes.
    * Motivo da Rejeição: Não isolaria o domínio de forma tão eficaz quanto a Arquitetura Hexagonal, pois as dependências tradicionais fluem de cima para baixo, o que pode levar a vazamentos de conceitos entre camadas. A inversão de dependência aplicada na Arquitetura Hexagonal protege melhor o domínio.

* **Domain-Driven Design Tático Puro:**
    * Descrição: Foco apenas nos padrões táticos de DDD (Aggregates, Entities, Value Objects, Repositories, etc.) sem uma estrutura arquitetural específica.
    * Motivo da Rejeição: Embora os padrões táticos de DDD sejam valiosos (e serão incorporados dentro da Arquitetura Hexagonal), eles não definem por si só uma estrutura completa de organização de código e gerenciamento de dependências.

## Consequências

**Positivas:**

* **Isolamento da Lógica de Negócio:** O domínio fica protegido de mudanças em tecnologias externas, tornando o sistema mais manutenível e evoluível.
* **Testabilidade Aprimorada:** As portas bem definidas facilitam a criação de mocks e testes unitários isolados, permitindo testar o domínio sem dependências externas.
* **Flexibilidade de Adaptadores:** Possibilidade de trocar implementações de infraestrutura (banco de dados, serviços externos) sem afetar a lógica de negócio.
* **Alinhamento com o Monolito Modular:** Reforça a modularidade interna estabelecida no ADR-010, com interfaces claras entre os componentes.
* **Evolução Facilitada:** Prepara o caminho para possível extração de microsserviços no futuro, caso necessário.
* **Clareza de Responsabilidades:** A separação clara entre domínio, aplicação e adaptadores ajuda novos desenvolvedores a entenderem onde implementar diferentes tipos de funcionalidade.
* **Comunicação Clara entre Módulos:** As interfaces bem definidas facilitam a integração entre diferentes módulos do monolito.

**Negativas / Trade-offs:**

* **Complexidade Inicial:** A implementação inicial pode parecer mais complexa do que abordagens mais simples, exigindo disciplina na estruturação do código.
* **Mais código Boilerplate:** A necessidade de definir interfaces e suas implementações gera mais código em comparação com abordagens mais diretas.
* **Curva de Aprendizado:** Desenvolvedores não familiarizados com o padrão podem precisar de tempo para se adaptar à separação entre portas e adaptadores.
* **Overhead para Funcionalidades Simples:** Para operações CRUD simples, a arquitetura pode parecer excessiva inicialmente, embora isso seja compensado pela consistência e escalabilidade a longo prazo.
* **Necessidade de Disciplina:** Manter a arquitetura requer disciplina da equipe para não criar atalhos que violem as regras de dependência.

**(Opcional) Notas Adicionais:**

* Esta decisão promove uma estrutura de código mais idiomática para Go, seguindo o princípio de "achatamento" de hierarquias e organização funcional, em vez de uma organização estritamente baseada em camadas. Isso mantém os benefícios da arquitetura hexagonal, enquanto adere às convenções da comunidade Go.
* Os princípios da Arquitetura Hexagonal são mantidos mesmo com esta estrutura de diretórios mais simplificada:
  - `port.go` define as interfaces (portas)
  - `type.go` e `[module].go` contêm o domínio
  - `/usecase/` contém a camada de aplicação
  - `/handler/` e `/repository/` são os adaptadores primários e secundários, respectivamente
* A estrutura proposta reduz a complexidade de navegação no código e favorece a coesão funcional, alinhando-se ao princípio de simplicidade mencionado no ADR-000.
* Esta decisão não impede a aplicação de princípios e padrões de Domain-Driven Design (DDD) dentro da estrutura Hexagonal. Na verdade, DDD complementa bem esta arquitetura, especialmente para módulos com domínios mais complexos.
* A implementação específica pode variar ligeiramente entre módulos dependendo da complexidade de cada domínio, mas o padrão geral de Portas e Adaptadores deve ser mantido para consistência.
* Ferramentas como injeção de dependência serão importantes para manter o desacoplamento entre os componentes. Em Go, isso pode ser implementado através de interfaces bem definidas e composição de estruturas.
* Para módulos mais complexos, pode-se adicionar subdivisões adicionais como `service/` para serviços de domínio ou `event/` para manipuladores de eventos.
* Recomenda-se criar exemplos e guias internos de como implementar esta arquitetura na linguagem Go, aproveitando seus recursos específicos como interfaces implícitas e composição de structs.
