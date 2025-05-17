# ADR-012: Adoção do PostgreSQL como Banco de Dados Principal

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A escolha do sistema de gerenciamento de banco de dados (SGBD) é crucial para o RedToGreen, um SaaS de gestão financeira pessoal. O SGBD deve garantir a integridade dos dados financeiros, suportar relacionamentos complexos (usuários, wallets, transações, categorias, compartilhamentos), oferecer boa performance, escalabilidade e ser compatível com a linguagem Go (conforme ADR-011). Os dados financeiros exigem alta consistência (ACID). A arquitetura definida é um monolito modular (conforme ADR-010).

Adicionalmente, a aplicação RedToGreen será hospedada em uma plataforma de nuvem (considerando GCP ou AWS). A escolha do SGBD deve, portanto, levar em conta a disponibilidade e maturidade dos serviços gerenciados, custos, facilidade de operação e escalabilidade oferecidos por esses provedores.

A pergunta que estamos tentando responder é: Qual SGBD devemos adotar para o RedToGreen que melhor atenda aos requisitos de integridade, performance, escalabilidade, modelagem de dados e suporte em ambientes de nuvem gerenciados?

O escopo desta decisão abrange o SGBD principal a ser utilizado para persistência de todos os dados de negócio da aplicação RedToGreen.

**Critérios de Decisão Chave:**
* Integridade e Consistência dos Dados (suporte ACID robusto).
* Capacidade de modelar relacionamentos complexos.
* Performance para cargas de trabalho mistas (leitura e escrita).
* Escalabilidade (vertical e horizontal, facilitada por serviços de nuvem).
* Extensibilidade (tipos de dados, funções).
* Maturidade, Confiabilidade e Suporte como Serviço Gerenciado em Nuvem.
* Suporte da comunidade e ecossistema (especialmente para Go).
* Facilidade de uso e gerenciamento (incluindo a operação via serviços de nuvem).
* Custo (incluindo o TCO dos serviços gerenciados na nuvem).

## Decisão

Adotaremos o **PostgreSQL** como o sistema de gerenciamento de banco de dados principal para o RedToGreen. A intenção é utilizar um serviço gerenciado de PostgreSQL em uma plataforma de nuvem (como AWS RDS for PostgreSQL, GCP Cloud SQL for PostgreSQL, ou suas variantes de alta performance como Amazon Aurora PostgreSQL-Compatible Edition ou GCP AlloyDB for PostgreSQL), com a escolha específica do serviço e provedor a ser definida com base em uma análise detalhada de custo-benefício no momento da implementação da infraestrutura.

**Justificativas:**
1.  **Conformidade ACID e Integridade de Dados:** O PostgreSQL é renomado por sua robustez e forte conformidade com os princípios ACID, o que é fundamental para a segurança e confiabilidade dos dados financeiros.
2.  **Modelagem Relacional Robusta:** Permite a modelagem eficaz dos relacionamentos complexos inerentes ao sistema RedToGreen, como usuários, carteiras, transações, categorias e as regras de compartilhamento (conforme RBAC definido no ADR-008).
3.  **Tipos de Dados Avançados e Extensibilidade:** O suporte nativo a JSONB é particularmente valioso, permitindo flexibilidade para campos com dados semiestruturados (ex: metadados de transações) sem sacrificar a capacidade de indexação e consulta eficiente. Outros tipos como arrays, hstore, e a capacidade de definir tipos customizados aumentam a flexibilidade.
4.  **Funcionalidades SQL Avançadas:** Recursos como Common Table Expressions (CTEs), window functions, e um rico conjunto de operadores e funções facilitam a escrita de consultas complexas, o que será útil para a geração de relatórios financeiros e análises de dados.
5.  **Excelente Suporte como Serviço Gerenciado em Nuvem:** Tanto AWS quanto GCP oferecem serviços gerenciados maduros, de alta qualidade e com opções de alta performance para PostgreSQL. Isso simplifica significativamente a operação, manutenção (patching, upgrades), backups, alta disponibilidade (HA) e escalabilidade.
6.  **Performance e Escalabilidade na Nuvem:** As plataformas de nuvem fornecem ferramentas eficazes para escalar instâncias PostgreSQL verticalmente e configurar réplicas de leitura para escalabilidade horizontal de consultas, atendendo às necessidades de crescimento do RedToGreen.
7.  **Comunidade Ativa e Ecossistema Rico:** PostgreSQL possui uma comunidade global ativa e um vasto ecossistema de ferramentas. Para Go, existem drivers de alta qualidade como o `pgx`, conhecido por sua performance e suporte a funcionalidades específicas do PostgreSQL.
8.  **Recursos de Segurança:** Oferece funcionalidades de segurança robustas, incluindo Row-Level Security (RLS), que pode ser benéfico para implementar lógicas de acesso refinadas, como no compartilhamento de carteiras.
9.  **Alinhamento com Decisões Arquiteturais Anteriores:**
    * Suporta bem o uso de UUID v7 (ADR-003) como chaves primárias (com tipo de dado `UUID` nativo).
    * Facilita a implementação de colunas de auditoria `created_at`/`updated_at` (ADR-002) e a coluna de versionamento `version` (ADR-006).

## Alternativas Consideradas

* **MySQL (utilizando serviço gerenciado em nuvem):**
    * *Prós:* Grande popularidade, vasta documentação, ecossistema maduro e excelentes serviços gerenciados em AWS (RDS for MySQL, Aurora MySQL) e GCP (Cloud SQL for MySQL). Conhecido por sua boa performance em aplicações web e facilidade de uso percebida.
    * *Motivo da Rejeição:* Embora seja uma alternativa viável e bem suportada na nuvem, o PostgreSQL geralmente oferece maior conformidade com os padrões SQL, superior extensibilidade (especialmente com JSONB, tipos de dados geoespaciais e customizados), e um conjunto de funcionalidades SQL avançadas mais rico. Para um sistema financeiro onde a integridade de dados, a capacidade de modelar relacionamentos complexos e o potencial para análises sofisticadas são críticos, o PostgreSQL é considerado mais vantajoso a longo prazo.

* **MongoDB (utilizando MongoDB Atlas ou serviço nativo da nuvem):**
    * *Prós:* Alta flexibilidade de esquema, escalabilidade horizontal nativa facilitada por serviços como o Atlas, e pode acelerar o desenvolvimento inicial para certos modelos de dados.
    * *Motivo da Rejeição:* A natureza fundamentalmente relacional dos dados financeiros do RedToGreen (transações, saldos, categorias, usuários, compartilhamentos) e a necessidade crítica de transações ACID consistentes através de múltiplos "documentos" ou "coleções", juntamente com a integridade referencial, tornam um RDBMS como o PostgreSQL uma escolha mais segura e apropriada. Gerenciar esses aspectos em um banco NoSQL orientado a documentos adicionaria complexidade desnecessária à camada de aplicação para este domínio, apesar da disponibilidade de bons serviços gerenciados.

## Consequências

**Positivas:**
* **Alta Integridade e Consistência dos Dados:** Proporciona uma base confiável para as operações financeiras do RedToGreen.
* **Flexibilidade na Modelagem de Dados:** A combinação de modelagem relacional com o tipo JSONB oferece o melhor dos dois mundos para diferentes tipos de dados.
* **Capacidade para Consultas Analíticas Complexas:** Suporta a geração de relatórios e insights financeiros detalhados.
* **Operação Simplificada e Escalabilidade:** Através de serviços gerenciados em nuvem, tarefas como backups, patching, configuração de alta disponibilidade e escalonamento são facilitadas.
* **Boa Integração com o Ecossistema Go:** Drivers eficientes e bem mantidos.
* **Menor Risco de "Vendor Lock-in" no Nível do SGBD:** Sendo open-source, oferece maior flexibilidade em caso de mudança de provedor de nuvem ou estratégia de hospedagem, embora a dependência de funcionalidades específicas de serviços premium da nuvem (ex: Aurora, AlloyDB) deva ser gerenciada.

**Negativas / Trade-offs:**
* **Custo dos Serviços Gerenciados:** Os serviços gerenciados em nuvem, especialmente instâncias de alta performance, com múltiplas zonas de disponibilidade (Multi-AZ) ou réplicas de leitura, implicam custos que precisam ser planejados e otimizados.
* **Complexidade de Gerenciamento (Percebida para Auto-Hospedagem):** Embora mitigada por serviços gerenciados, a administração fina e o tuning avançado do PostgreSQL podem ser percebidos como mais complexos do que o MySQL para equipes com menos experiência específica, caso se opte por uma abordagem menos gerenciada no futuro.
* **Lock-in em Nível de Serviço de Nuvem Premium:** Se o projeto optar por utilizar intensivamente funcionalidades exclusivas de variantes como Amazon Aurora PostgreSQL-Compatible ou GCP AlloyDB for PostgreSQL, a portabilidade para outros serviços PostgreSQL "vanilla" ou outros provedores pode exigir adaptações.

## (Opcional) Notas Adicionais
* A escolha específica do serviço gerenciado (ex: AWS RDS for PostgreSQL vs. GCP Cloud SQL for PostgreSQL, ou suas variantes premium) será feita com base em uma análise detalhada de custos, funcionalidades específicas disponíveis, latência para os usuários alvo, e preferências/experiência da equipe no momento da configuração da infraestrutura.
* Recomenda-se a utilização do driver `pgx` para Go devido à sua performance e suporte a funcionalidades específicas do PostgreSQL.
* A versão do PostgreSQL a ser utilizada será a mais recente estável que seja bem suportada pelo serviço de nuvem escolhido e pelas ferramentas do projeto.
* A estratégia de backup, retenção de dados e recuperação de desastres será primariamente definida e gerenciada pelas funcionalidades oferecidas pelo serviço de banco de dados em nuvem escolhido.
* Considerar o uso de schemas PostgreSQL para organizar logicamente as tabelas, potencialmente alinhando-os com os módulos definidos na arquitetura monolítica modular (ADR-010), para melhorar a organização e o gerenciamento do banco de dados à medida que o sistema cresce.
