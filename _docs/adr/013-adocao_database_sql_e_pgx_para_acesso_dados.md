# ADR-013: Adoção do Pacote `database/sql` com Driver `pgx` para Acesso a Dados PostgreSQL

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Com a linguagem Go (ADR-011) e o PostgreSQL (ADR-012) definidos como tecnologias chave para o backend do RedToGreen, é necessário decidir a estratégia e as ferramentas para interagir com o banco de dados a partir da aplicação Go. As opções variam desde o uso direto da biblioteca padrão, passando por query builders, geradores de código a partir de SQL, até Object-Relational Mappers (ORMs) completos.

Os principais critérios para esta decisão incluem:
* **Controle sobre o SQL:** Capacidade de escrever e otimizar queries SQL diretamente.
* **Performance:** Minimizar o overhead entre a aplicação e o banco de dados.
* **Produtividade do Desenvolvedor:** Equilibrar o controle com a eficiência no desenvolvimento de operações CRUD e queries complexas.
* **Manutenibilidade:** Clareza do código de acesso a dados e facilidade de refatoração.
* **Segurança:** Prevenção de SQL injection.
* **Alinhamento com a Filosofia Go:** Preferência por ferramentas explícitas, simples e que utilizem bem os recursos da linguagem e sua biblioteca padrão.
* **Suporte às Funcionalidades do PostgreSQL:** Capacidade de utilizar tipos de dados e funcionalidades específicas do PostgreSQL.

A pergunta que estamos tentando responder é: Qual abordagem e conjunto de ferramentas devemos utilizar para o acesso ao banco de dados PostgreSQL em nossa aplicação Go, visando controle, performance e clareza?

O escopo desta decisão abrange a camada de acesso a dados (Data Access Layer - DAL) da aplicação RedToGreen, definindo como as queries SQL serão construídas, executadas e como os resultados serão mapeados para estruturas de dados Go.

## Decisão

Adotaremos o uso do pacote padrão **`database/sql`** da biblioteca Go em conjunto com o driver **`jackc/pgx`** (na sua modalidade compatível com `database/sql`, ou seja, `pgx/v5/stdlib`) como a principal forma de interação com o banco de dados PostgreSQL.

1.  **Interface Padrão `database/sql`:** Utilizaremos as interfaces e tipos fornecidos pelo pacote `database/sql` (`sql.DB`, `sql.Tx`, `sql.Rows`, `sql.Stmt`, etc.) para executar queries, gerenciar transações e escanear resultados. Isso garante portabilidade conceitual e familiaridade para desenvolvedores Go.
2.  **Driver `pgx`:** O `jackc/pgx/v5/stdlib` será o driver específico para PostgreSQL. Ele é escolhido por sua alta performance, robustez e excelente suporte a funcionalidades específicas do PostgreSQL, que podem ser acessadas quando necessário.
3.  **SQL Explícito:** As queries SQL serão escritas manualmente como strings no código Go. Isso proporciona controle total sobre a otimização e a lógica das queries.
4.  **Mapeamento Manual (com Auxílio Potencial):** O mapeamento dos resultados das queries (`sql.Rows`) para as structs Go será feito explicitamente usando o método `rows.Scan()`.
5.  **Prevenção de SQL Injection:** Todas as queries que envolvem dados de entrada do usuário serão escritas utilizando queries parametrizadas (placeholders), conforme o padrão suportado por `database/sql` e `pgx`.

## Alternativas Consideradas

* **Query Builders (ex: `Masterminds/squirrel`):**
    * *Prós:* Ajudam a construir queries SQL de forma programática e fluida, facilitando queries dinâmicas e reduzindo o risco de erros de sintaxe SQL.
    * *Motivo da Rejeição (como abordagem primária):* Embora úteis, introduzem uma camada de abstração adicional e uma nova API a ser aprendida. A preferência é por manter o SQL explícito e visível. Podem ser considerados para casos de uso muito específicos de construção de queries dinâmicas complexas, se a escrita manual se provar excessivamente verbosa.

* **`sqlc` (Geração de Código a partir de SQL):**
    * *Prós:* Gera código Go type-safe a partir de queries SQL escritas pelo desenvolvedor, combinando controle sobre o SQL com produtividade e segurança de tipos. Reduz o boilerplate de mapeamento.
    * *Motivo da Rejeição (para esta fase inicial):* Embora seja uma alternativa muito forte e alinhada com muitos dos objetivos, a equipe opta por começar com a abordagem mais fundamental de `database/sql` para garantir um entendimento profundo da interação com o banco. `sqlc` pode ser reavaliado no futuro se o boilerplate de mapeamento manual se tornar um gargalo significativo de produtividade, sendo uma evolução natural da abordagem escolhida.

* **ORMs (Object-Relational Mappers - ex: `GORM`, `Ent`):**
    * *Prós:* Alta produtividade para operações CRUD, gerenciamento de relacionamentos, e funcionalidades como migrations e hooks. `Ent` oferece forte type-safety através de geração de código.
    * *Motivo da Rejeição:* Introduzem uma camada de abstração significativa sobre o SQL, o que pode levar a "mágica", dificultar a depuração de performance e limitar o uso de funcionalidades específicas do PostgreSQL. Afasta-se da preferência por controle explícito e simplicidade da biblioteca padrão. A complexidade adicional de um ORM completo não é considerada necessária para os requisitos atuais, dado o desejo de manter o controle direto sobre o SQL.

* **Uso direto da API nativa do `pgx` (sem `database/sql`):**
    * *Prós:* Potencialmente a maior performance em alguns cenários, acesso direto a todas as funcionalidades do `pgx`.
    * *Motivo da Rejeição (como padrão):* Embora poderoso, o uso da API `database/sql` oferece uma interface mais padronizada e familiar para a maioria dos desenvolvedores Go e facilita a portabilidade teórica para outros drivers SQL, se isso algum dia fosse necessário (embora o foco seja no PostgreSQL). A versão `stdlib` do `pgx` já oferece excelente performance e acesso à maioria das funcionalidades do `pgx`.

## Consequências

**Positivas:**
* **Controle Máximo sobre o SQL:** Permite otimizações finas e o uso de todas as funcionalidades do PostgreSQL sem as limitações de uma camada de abstração.
* **Performance Potencialmente Mais Alta:** A interação direta (via `database/sql` e `pgx`) minimiza o overhead.
* **Simplicidade Conceitual:** Utiliza componentes da biblioteca padrão do Go, que são bem compreendidos e documentados.
* **Mínimo de Dependências Externas:** Além do driver `pgx`, não introduz grandes dependências de terceiros para a lógica de acesso a dados core.
* **Transparência e Depuração Facilitada:** O SQL executado é exatamente o que está escrito no código, facilitando a depuração e o profiling.
* **Segurança Explícita:** A necessidade de usar queries parametrizadas para evitar SQL injection é clara e direta.
* **Flexibilidade:** Adapta-se bem a queries simples ou complexas.

**Negativas / Trade-offs:**
* **Boilerplate para CRUD e Mapeamento:** Escrever SQL manualmente para todas as operações CRUD e mapear os resultados com `rows.Scan()` pode ser verboso e repetitivo.
* **Risco de Erros no Mapeamento Manual:** Erros de digitação em nomes de colunas ou na ordem dos campos no `rows.Scan()` podem levar a bugs em tempo de execução. (Bibliotecas como `scany` podem mitigar isso parcialmente).
* **Refatoração Manual do SQL:** Alterações no esquema do banco de dados (nomes de tabelas/colunas) exigirão a atualização manual de todas as strings SQL correspondentes no código, o que é propenso a erros.
* **Construção de Queries Dinâmicas:** Pode ser mais complexo e verboso construir queries SQL dinâmicas de forma segura em comparação com query builders ou ORMs.
* **Menor Produtividade para Tarefas Rotineiras:** Comparado a ORMs ou `sqlc`, o desenvolvimento de operações CRUD padrão pode levar mais tempo.

## (Opcional) Notas Adicionais
* A equipe deve ser diligente na escrita de queries parametrizadas para garantir a segurança contra SQL injection.
* Recomenda-se a criação de uma camada de repositório bem definida que encapsule a lógica de acesso a dados, mantendo as queries SQL e o uso de `database/sql` isolados do restante da lógica de negócios.
* Testes de integração que interajam com um banco de dados real (ou um banco de testes) serão cruciais para garantir a correção das queries SQL e do mapeamento de dados.
* A biblioteca `georgysavva/scany` será avaliada como um possível auxiliar para reduzir o boilerplate do `rows.Scan()` em listagens, sem comprometer o controle sobre o SQL.
* Esta decisão não impede o uso futuro de `sqlc` como uma otimização de produtividade, caso o boilerplate se torne um problema, pois `sqlc` também se baseia na escrita de SQL e usa `database/sql`.
