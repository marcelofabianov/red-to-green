# ADR-014: Adoção da Ferramenta `pressly/goose` para Migrações de Esquema de Banco de Dados

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Com o PostgreSQL definido como SGBD (ADR-012) e a abordagem de acesso a dados via `database/sql` + `pgx` (ADR-013), é fundamental estabelecer uma estratégia para gerenciar e versionar as alterações no esquema do banco de dados (Data Definition Language - DDL) de forma controlada, repetível e rastreável. As migrações de esquema são essenciais para evoluir a estrutura do banco de dados à medida que a aplicação RedToGreen se desenvolve.

Critérios importantes para a escolha da ferramenta de migração incluem:
* **Suporte a SQL Puro:** Capacidade de escrever DDL diretamente em arquivos `.sql`.
* **Simplicidade de Uso:** Facilidade para criar, aplicar e reverter migrações.
* **Versionamento:** Manutenção de um histórico claro das migrações aplicadas.
* **Flexibilidade:** Suporte para operações de `up` (aplicar) e `down` (reverter).
* **Formato dos Arquivos de Migração:** Preferência por um formato que seja fácil de gerenciar e entender.
* **Independência de ORM:** Não estar acoplado a um ORM específico.
* **Comunidade e Maturidade:** Ferramenta estabelecida e confiável.

A pergunta que estamos tentando responder é: Qual ferramenta de migração de esquema de banco de dados devemos adotar para gerenciar as alterações de DDL no PostgreSQL para o projeto RedToGreen, alinhada com nossa preferência por SQL puro e simplicidade operacional?

O escopo desta decisão abrange o processo e a ferramenta para criar, aplicar, versionar e reverter alterações no esquema do banco de dados PostgreSQL do RedToGreen.

## Decisão

Adotaremos a ferramenta **`pressly/goose`** para o gerenciamento das migrações de esquema do banco de dados PostgreSQL no projeto RedToGreen.

As migrações serão escritas predominantemente como arquivos SQL puros.

**Justificativas para a escolha do `goose`:**
1.  **Trabalho com SQL Puro Nativo:** `goose` permite que as migrações sejam escritas diretamente em arquivos `.sql`, utilizando a sintaxe nativa do PostgreSQL. Isso garante controle total sobre o DDL, permite otimizações específicas do SGBD e evita a necessidade de aprender uma DSL (Domain Specific Language) ou usar código Go para definir alterações de esquema simples.
2.  **Facilidade de Executar Migrações e Rollbacks:** O CLI do `goose` é intuitivo e oferece comandos claros para aplicar migrações (`goose up`), reverter a última migração (`goose down`), reverter para uma versão específica (`goose down-to VERSION`), verificar o status (`goose status`), entre outros.
3.  **Trabalho com Arquivos `.sql`:** As migrações são organizadas em arquivos `.sql` individuais, versionados cronologicamente (via timestamp no nome do arquivo), o que facilita o rastreamento e a revisão das alterações no sistema de controle de versão (Git).
4.  **Unificação de `Up` e `Down` em Único Arquivo:** `goose` utiliza um único arquivo `.sql` por migração, contendo seções `_-- +goose Up_` e `_-- +goose Down_`. Isso reduz o número total de arquivos a serem mantidos em comparação com ferramentas que usam arquivos separados para `up` e `down`, simplificando a organização.
5.  **Simplicidade e Leveza:** `goose` é uma ferramenta focada e relativamente simples, sem um grande overhead de configuração ou dependências complexas.
6.  **Independência de ORM:** Funciona perfeitamente com a abordagem de `database/sql` + `pgx` (ADR-013), não exigindo o uso de um ORM.
7.  **Maturidade e Comunidade:** É uma ferramenta bem estabelecida e utilizada na comunidade Go.
8.  **Opção de Uso como Biblioteca:** Além do CLI, `goose` pode ser importado como uma biblioteca Go, permitindo integrar a lógica de migração dentro da própria aplicação, se necessário (por exemplo, para aplicar migrações durante o deploy ou inicialização da aplicação).

## Alternativas Consideradas

* **`golang-migrate/migrate`:**
    * *Prós:* Ferramenta robusta, amplamente utilizada, suporta SQL puro, CLI e uso como biblioteca. Muito flexível em termos de fontes de migração e SGBDs suportados.
    * *Motivo da Rejeição:* Embora muito capaz, a preferência pelo formato de arquivo único do `goose` para as seções `Up` e `Down` foi um fator. As funcionalidades principais são comparáveis, mas o fluxo de trabalho com `goose` pareceu ligeiramente mais alinhado com as preferências da equipe.

* **`Atlas` (atlasgo.io):**
    * *Prós:* Capacidade de gerar DDL automaticamente através de "schema diffing" a partir de um estado desejado (declarativo), visualização de esquema, e pode exportar migrações para o formato do `goose`.
    * *Motivo da Rejeição (como ferramenta primária de execução de migrações):* A preferência principal é pela escrita manual e controle total do SQL de migração. A geração automática, embora poderosa, pode introduzir uma camada de "mágica" e requer uma curva de aprendizado maior para a definição do esquema desejado. No entanto, `Atlas` pode ser considerado no futuro como uma ferramenta auxiliar para *gerar* o SQL inicial para os arquivos de migração do `goose`, especialmente para alterações de esquema mais complexas, sujeito à revisão cuidadosa do DDL gerado.

* **Migrações Embutidas em ORMs (ex: `GORM AutoMigrate`, `Ent Migrations`):**
    * *Prós:* Podem simplificar as migrações se já estiver utilizando o ORM, derivando as alterações dos modelos Go.
    * *Motivo da Rejeição:* Incompatível com a decisão de não utilizar um ORM completo para acesso a dados (ADR-013) e não atende à preferência por SQL puro e controle explícito do DDL.

## Consequências

**Positivas:**
* **Controle Total sobre o DDL:** Os desenvolvedores escrevem SQL nativo do PostgreSQL, permitindo otimizações e uso de todas as funcionalidades do SGBD.
* **Clareza e Rastreabilidade:** As migrações são arquivos SQL versionáveis, fáceis de ler, revisar e entender o histórico de alterações do esquema.
* **Simplicidade Operacional:** O CLI do `goose` é simples para aplicar, reverter e verificar o status das migrações.
* **Consistência com a Stack Tecnológica:** Alinha-se bem com a escolha de Go, PostgreSQL, e a abordagem de acesso a dados com `database/sql`.
* **Redução de Complexidade:** Evita a necessidade de aprender DSLs complexas ou depender de geração de código "mágico" para o DDL.
* **Manutenção Simplificada de Arquivos:** O formato de arquivo único para `Up` e `Down` é preferido pela equipe.

**Negativas / Trade-offs:**
* **Escrita Manual do DDL:** Requer que os desenvolvedores escrevam manualmente todo o SQL para `Up` e `Down`, o que pode ser propenso a erros se não for feito com cuidado, especialmente para as seções `Down`.
* **Responsabilidade pelas Migrações `Down`:** A corretude e a testabilidade das migrações de `Down` são cruciais e inteiramente responsabilidade do desenvolvedor. Rollbacks complexos podem ser difíceis de implementar e testar. (Muitas equipes adotam uma estratégia de "roll-forward" para corrigir problemas em produção, em vez de depender de rollbacks complexos).
* **Coordenação em Equipe:** Em equipes maiores, a criação concorrente de arquivos de migração pode levar a conflitos de timestamp/ordenação se não houver um processo de desenvolvimento coordenado.
* **Sem Verificação Estática Avançada do DDL:** A correção do SQL escrito nos arquivos de migração só é totalmente verificada no momento da aplicação contra o banco de dados (embora linters de SQL possam ajudar).

## (Opcional) Notas Adicionais
* Os arquivos de migração SQL gerados pelo `goose` serão armazenados em um diretório dedicado dentro do repositório do projeto (ex: `db/migrations/`).
* A tabela de versionamento do `goose` (padrão: `goose_db_version`) será criada automaticamente no banco de dados na primeira execução.
* Recomenda-se testar as migrações (tanto `Up` quanto `Down`) em ambientes de desenvolvimento e staging antes de aplicá-las em produção.
* A política para lidar com migrações problemáticas em produção deve ser definida (ex: priorizar "roll-forward" com uma nova migração corretiva em vez de "roll-back", especialmente para alterações que envolvem perda de dados ou transformações complexas).
