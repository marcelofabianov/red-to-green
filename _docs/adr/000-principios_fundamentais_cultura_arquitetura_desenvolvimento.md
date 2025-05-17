# ADR-000: Princípios Fundamentais da Cultura de Arquitetura e Desenvolvimento do RedToGreen

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Este documento estabelece os princípios arquiteturais e a filosofia de desenvolvimento que guiarão a construção e evolução do sistema RedToGreen. Ele serve como uma referência fundamental para todas as decisões técnicas, visando garantir consistência, qualidade e alinhamento com os objetivos de longo prazo do produto. Este ADR não introduz uma nova decisão, mas sim codifica os valores e aprendizados extraídos das decisões arquiteturais tomadas até o momento (refletindo as discussões e os ADRs de 001 a 017, e os que virão).

## Princípios Fundamentais

Os seguintes princípios são a base da nossa abordagem para o desenvolvimento do RedToGreen:

1.  **Simplicidade e Pragmatismo (KISS, YAGNI):**
    * **Motivação:** Priorizar soluções claras, diretas e eficientes, evitando complexidade desnecessária e a implementação de funcionalidades ou adoção de tecnologias antes que sejam genuinamente necessárias e justificadas. Foco em entregar valor de forma incremental.
    * **Refletido em:**
        * Adoção de uma arquitetura Monolítica Modular inicial (ADR-010) em vez de microsserviços prematuros.
        * Escolha do RBAC para controle de acesso (ADR-008) pela sua simplicidade para os requisitos iniciais.
        * Adoção do `log/slog` (ADR-017) por ser uma solução nativa e eficaz para logging estruturado.
        * Adoção do `database/sql` + `pgx` (ADR-013) para acesso a dados, valorizando controle e clareza sobre abstrações pesadas de ORMs.
        * Preferência pelo `go-chi/chi` (ADR-016) como um roteador HTTP leve e compatível com `net/http`.
        * Início com REST/OpenAPI para comunicação backend-frontend (ADR-015) devido à sua simplicidade, vasta adoção e ecossistema maduro.
        * Uso do `pressly/goose` (ADR-014) para migrações de banco de dados com SQL puro.

2.  **Utilização da Biblioteca Padrão e Ferramental Nativo do Go (Quando Adequado):**
    * **Motivação:** Reduzir o número de dependências externas, minimizar a sobrecarga de configuração ("fadiga de ferramentas"), e aproveitar a estabilidade, performance e os padrões idiomáticos oferecidos pela linguagem Go e seu ecossistema padrão.
    * **Refletido em:**
        * Escolha da linguagem Go (ADR-011), valorizando sua biblioteca padrão (`stdlib`) robusta e ferramentas integradas (`go fmt`, `go test`, `go mod`, etc.).
        * Adoção do `log/slog` (ADR-017) para logging.
        * Utilização do `database/sql` (ADR-013) como interface primária de acesso ao banco.
        * Escolha do `go-chi/chi` (ADR-016) pela sua forte compatibilidade e integração com `net/http`.

3.  **Controle Explícito e Transparência sobre Abstrações "Mágicas":**
    * **Motivação:** Manter um entendimento claro do que acontece "por baixo dos panos" para facilitar a depuração, otimizar a performance de forma precisa, e evitar comportamentos inesperados ou limitações impostas por bibliotecas ou frameworks excessivamente opinativos ou que ocultam a complexidade de forma prejudicial.
    * **Refletido em:**
        * Preferência pelo `database/sql` + `pgx` (ADR-013) sobre ORMs completos para a fase inicial.
        * Adoção do `pressly/goose` (ADR-014) para migrações, utilizando SQL puro e explícito.
        * A linguagem Go (ADR-011) em si incentiva código explícito e menos "mágico".

4.  **Qualidade, Robustez e Manutenibilidade do Código (Clean Code, SOLID, DRY, TDA, LD):**
    * **Motivação:** Construir um sistema que seja fácil de entender, manter, testar e evoluir, reduzindo a probabilidade de bugs e facilitando a colaboração da equipe a longo prazo.
    * **Refletido em:**
        * Adoção de uma estrutura padronizada para erros com contexto adicional (ADR-005).
        * Implementação de colunas de auditoria (`created_at`, `updated_at` - ADR-002) e versionamento de registros (`version` - ADR-006).
        * Escolha de algoritmos de hashing de senhas seguros e modernos (Argon2 - ADR-007).
        * Adoção de UUID v7 para identificadores únicos (ADR-003), visando unicidade e otimização de performance de escrita no banco.
        * O princípio da modularidade interna na arquitetura monolítica (ADR-010).

5.  **Segurança como Prioridade (Security by Design):**
    * **Motivação:** Proteger os dados financeiros sensíveis e a privacidade dos usuários do RedToGreen, incorporando considerações de segurança em todas as fases do desenvolvimento.
    * **Refletido em:**
        * Adoção do Argon2 para hashing de senhas (ADR-007).
        * Implementação do RBAC para controle de acesso (ADR-008).
        * Atenção à prevenção de SQL injection ao usar `database/sql` (ADR-013).
        * Planejamento para observabilidade (ADR-017 e discussão sobre OTel/Jaeger) que também auxilia na detecção de anomalias de segurança.

6.  **Design Evolutivo e Preparação para o Futuro (Arquitetura Adaptável):**
    * **Motivação:** Iniciar com uma base simples e sólida, mas projetar o sistema de forma que possa crescer e se adaptar a novas funcionalidades, aumento de carga e mudanças nos requisitos de negócio sem a necessidade de grandes refatorações disruptivas.
    * **Refletido em:**
        * Arquitetura Monolítica Modular (ADR-010), que permite a extração futura de microsserviços se necessário.
        * Adoção de Feature Flags (ADR-009) para lançamentos graduais e gerenciamento flexível de funcionalidades.
        * Escolha do PostgreSQL (ADR-012), um SGBD robusto, extensível e escalável.
        * Abertura para adoção de outros protocolos de comunicação além do REST quando apropriado (ADR-015).
        * Adoção de uma estratégia de arquivamento de registros (ADR-004) que considera o ciclo de vida dos dados.

7.  **Documentação Contínua e Compartilhamento de Conhecimento (ADRs como Ferramenta):**
    * **Motivação:** Manter um registro claro e acessível das decisões arquiteturais, suas justificativas, alternativas consideradas e consequências. Isso facilita o entendimento do sistema ao longo do tempo, o onboarding de novos membros da equipe e a tomada de decisões futuras consistentes.
    * **Refletido em:** A própria prática de criar, revisar e manter estes ADRs como parte central do processo de design e desenvolvimento.

8.  **Observabilidade como Pilar Fundamental:**
    * **Motivação:** Garantir a capacidade de entender o comportamento do sistema em produção, diagnosticar problemas rapidamente, monitorar a saúde da aplicação e obter insights operacionais.
    * **Refletido em:**
        * Adoção de logging estruturado e contextualizado com `slog` (ADR-017).
        * Planejamento explícito para a futura implementação de OpenTelemetry (OTel) e Jaeger para tracing distribuído e métricas.
        * A estrutura padronizada para erros (ADR-005) que facilita o monitoramento e a criação de alertas.

## Consequências da Adoção destes Princípios

* **Positivas:**
    * Maior clareza e foco no desenvolvimento, com menos tempo gasto em debates sobre ferramentas e abordagens que divergem da cultura estabelecida.
    * Código-fonte mais consistente, legível, testável e manutenível.
    * Maior controle sobre a performance, segurança e comportamento geral do sistema.
    * Uma base arquitetural sólida que suporta escalabilidade e evolução futura de forma sustentável.
    * Equipe alinhada com uma filosofia de desenvolvimento compartilhada, promovendo melhor colaboração.
    * Redução de dívida técnica não intencional e retrabalho.
* **Negativas / Trade-offs:**
    * Pode exigir maior disciplina e esforço inicial para aderir aos princípios, especialmente ao resistir à tentação de soluções "rápidas e fáceis" que os contradigam.
    * A preferência por soluções da biblioteca padrão ou ferramentas mais leves e focadas pode, em cenários específicos, exigir a escrita de mais código boilerplate para funcionalidades que frameworks "tudo incluso" poderiam fornecer prontas (embora, geralmente, com o custo de maior complexidade, "mágica" e menor controle).
    * A responsabilidade pela implementação correta de aspectos como segurança, otimização de performance e lógica de negócios complexa recai mais diretamente sobre a equipe de desenvolvimento ao se evitar abstrações muito pesadas.

## (Opcional) Notas Adicionais

Este documento é um artefato vivo e deve ser revisitado e refinado periodicamente à medida que o projeto RedToGreen evolui, novos desafios surgem e novos aprendizados são adquiridos pela equipe. No entanto, os princípios fundamentais aqui descritos devem servir como um guia duradouro e consistente para a tomada de decisões.
