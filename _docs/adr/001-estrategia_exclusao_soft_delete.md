# ADR-001: Estratégia de Exclusão de Registros com Soft-Delete

**Status:** Aceito

**Data:** 2025-05-15

## Contexto

Para a aplicação RedToGreen, é necessário definir uma estratégia consistente para a exclusão de registros no banco de dados. A exclusão física de dados (hard delete) pode levar à perda permanente de informações, dificultar auditorias, impedir a recuperação de dados excluídos acidentalmente e quebrar a integridade referencial se não tratada com extremo cuidado. Buscamos uma abordagem que minimize esses riscos, especialmente considerando a natureza dos dados financeiros onde o histórico pode ser importante.

A pergunta que estamos tentando responder é: Como devemos tratar a "exclusão" de registros de forma segura e que permita rastreabilidade e potencial recuperação?

O escopo desta decisão abrange todas as entidades persistidas no banco de dados da aplicação RedToGreen que são sujeitas à operações de exclusão por parte do usuário ou do sistema.

## Decisão

Adotaremos a estratégia de **soft-delete** para todos os registros passíveis de exclusão no sistema.

Esta estratégia será implementada da seguinte forma:
1.  Todas as tabelas cujos registros podem ser "excluídos" incluirão uma coluna chamada `deleted_at`.
2.  A coluna `deleted_at` será do tipo `TIMESTAMP` (ou `DATETIME` equivalente, dependendo do SGBD) e permitirá valores nulos.
3.  Um valor `NULL` na coluna `deleted_at` indica que o registro está disponível e não foi excluído.
4.  Um valor de data/hora (timestamp) na coluna `deleted_at` indica que o registro foi logicamente excluído. O valor representa o momento exato da exclusão.
5.  Todas as consultas padrão da aplicação (leituras, listagens, relatórios) devem, por padrão, retornar apenas registros onde `deleted_at IS NULL`.
6.  O acesso a registros "excluídos" (onde `deleted_at IS NOT NULL`) será restrito e só deve ocorrer em contextos específicos, como funcionalidades de lixeira, restauração de dados ou auditoria por administradores.

## Alternativas Consideradas (Opcional)

* **Hard Delete (Exclusão Física):**
    * Descrição: Remover permanentemente os registros do banco de dados usando o comando `DELETE FROM table WHERE ...`.
    * Motivo da Rejeição: Impede a recuperação de dados, dificulta a auditoria, pode levar à perda de histórico importante e aumenta o risco de problemas de integridade referencial.

* **Tabelas de Histórico/Arquivamento Separadas:**
    * Descrição: Mover os registros excluídos da tabela principal para uma tabela de "arquivo morto" ou "histórico".
    * Motivo da Rejeição (para o momento): Embora seja uma estratégia válida para arquivamento de longo prazo, ela adiciona complexidade na implementação de consultas que podem precisar de dados ativos e arquivados, bem como na lógica de restauração. A abordagem de soft-delete na mesma tabela é considerada mais simples para o escopo inicial do MVP e atende aos requisitos de recuperação e auditoria básica.

## Consequências

Liste as consequências resultantes da decisão tomada. É importante ser honesto sobre os trade-offs.

**Positivas:**

* **Recuperação de Dados:** Registros "excluídos" podem ser facilmente restaurados simplesmente atualizando a coluna `deleted_at` para `NULL`.
* **Auditoria e Rastreabilidade:** Mantém um histórico completo dos dados, incluindo quando um registro foi "desativado", o que é crucial para dados financeiros.
* **Integridade Referencial:** Reduz o risco de quebrar chaves estrangeiras e a necessidade de exclusões em cascata complexas, pois o registro fisicamente ainda existe.
* **Simplicidade na Implementação da Lógica de Exclusão:** A "exclusão" se torna uma operação de `UPDATE`.
* **Consistência:** Padroniza a forma como a exclusão é tratada em toda a aplicação.

**Negativas / Trade-offs:**

* **Aumento do Tamanho do Banco de Dados:** Como os registros não são fisicamente removidos, as tabelas podem crescer mais rapidamente ao longo do tempo.
* **Complexidade nas Consultas:** Todas as consultas de leitura de dados ativos precisam incluir a condição `WHERE deleted_at IS NULL`. Isso pode ser mitigado com abstrações na camada de acesso a dados (ex: scopes globais em ORMs).
* **Potencial Impacto na Performance:** Tabelas maiores e a necessidade do filtro adicional podem, com o tempo, impactar a performance das consultas se não houver indexação adequada na coluna `deleted_at`.
* **Necessidade de Gerenciamento de Dados "Excluídos":** A longo prazo, pode ser necessário implementar políticas de arquivamento ou expurgo (hard delete) para registros em `deleted_at` muito antigos e que não precisam mais ser mantidos online (fora do escopo do MVP).
* **Risco de Erro Humano:** Desenvolvedores podem esquecer de adicionar o filtro `deleted_at IS NULL` em novas consultas manuais ou em partes do sistema não cobertas pela abstração de acesso a dados, levando à exibição de dados "excluídos".

**(Opcional) Notas Adicionais:**

* Recomenda-se a criação de índices na coluna `deleted_at` em tabelas onde ela é frequentemente usada em cláusulas `WHERE`.
* A lógica de filtrar automaticamente registros com `deleted_at IS NOT NULL` deve ser, preferencialmente, implementada de forma centralizada na camada de acesso a dados do sistema (ex: através de views, scopes globais em ORMs, ou repositórios base) para garantir consistência e reduzir a repetição de código.
