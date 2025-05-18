# Entidade: Carteira (Wallet)

**Contexto de Domínio:** Carteira

## Descrição

Representa um agrupador financeiro principal ou um "centro de custo" pertencente a um usuário. Todas as transações financeiras do usuário são organizadas dentro de uma ou mais Carteiras, permitindo separar diferentes contextos financeiros (ex: "Pessoal", "Casa", "Projeto X"). A Carteira em si não gerencia saldos; os valores são derivados da soma das suas transações.

## Atributos

| Atributo         | Tipo (Go/BD)        | Obrigatório | Descrição                                                                     | Observações / ADRs Relevantes                                     |
| :--------------- | :------------------ | :---------- | :---------------------------------------------------------------------------- | :---------------------------------------------------------------- |
| `id`             | `uuid.UUID` / `UUID`| Sim         | Identificador único universal da carteira.                                    | ADR-003 (UUID v7)                                                 |
| `user_id`        | `uuid.UUID` / `UUID`| Sim         | ID do usuário proprietário desta carteira.                                    | Chave estrangeira para `user.id`.                                 |
| `name`           | `string` / `VARCHAR`| Sim         | Nome da carteira (ex: "Despesas Pessoais").                                   | Deve ser único por `user_id`. Máx: 255.                           |
| `description`    | `string` / `TEXT NULL`| Não       | Descrição opcional da carteira.                                               |                                                                   |
| `created_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando a carteira foi criada.                                  | ADR-002                                                           |
| `updated_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização da carteira.                                | ADR-002                                                           |
| `version`        | `int` / `INTEGER`   | Sim         | Número da versão do registro.                                                 | ADR-006.                                                          |
| `archived_at`    | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a carteira foi arquivada (não pode ser usada para novas transações). | ADR-004.                                                          |
| `deleted_at`     | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a carteira foi marcada como excluída (soft delete).    | ADR-001.                                                          |

## Relacionamentos

* **Uma Carteira pertence a um:** Usuário ([`../user/user_entity.md`](../user/user.md)) (via `user_id`)
* **Uma Carteira possui muitas:** Transações ([`transaction_entity.md`](./transaction.md))
* **Uma Carteira pode ser compartilhada com muitos:** Usuários (via tabela de junção `user_wallet_permissions` ou similar, implementando ADR-008 - RBAC)

## Regras de Negócio e Validações (Exemplos)

* O nome da carteira deve ser único para o usuário.
* Uma carteira não pode ser excluída (hard delete) se contiver transações; deve ser arquivada.
* **Saldos:** Os saldos relacionados a uma carteira (total de receitas, total de despesas, saldo final) são calculados dinamicamente a partir da soma de suas transações em um determinado período, quando necessário para exibição ou relatórios. A entidade `Wallet` não armazena saldos.

## Considerações MVP

* MVP inclui CRUD de carteiras (nome, descrição).
* A lógica de cálculo de saldos para exibição será feita sob demanda.
* Compartilhamento de carteiras (ADR-008) é uma funcionalidade chave, mas sua implementação completa pode ser faseada.
