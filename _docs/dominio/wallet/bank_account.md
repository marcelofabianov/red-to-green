# Entidade: Conta Bancária (BankAccount)

**Contexto de Domínio:** Usuário (gerenciada pelo usuário) / Carteira (usada como referência em transações)

## Descrição

Representa uma conta bancária pertencente ao usuário. As transações são vinculadas a uma conta bancária para indicar a origem ou destino dos fundos. A Conta Bancária em si não gerencia saldos; os valores são derivados da soma das transações a ela associadas.

## Atributos

| Atributo            | Tipo (Go/BD)        | Obrigatório | Descrição                                                                          | Observações / ADRs Relevantes                                     |
| :------------------ | :------------------ | :---------- | :--------------------------------------------------------------------------------- | :---------------------------------------------------------------- |
| `id`                | `uuid.UUID` / `UUID`| Sim         | Identificador único universal da conta bancária.                                   | ADR-003 (UUID v7)                                                 |
| `user_id`           | `uuid.UUID` / `UUID`| Sim         | ID do usuário proprietário desta conta bancária.                                   | Chave estrangeira para `user.id`.                                 |
| `name`              | `string` / `VARCHAR`| Sim         | Nome/apelido da conta bancária (ex: "Conta Corrente BB", "Poupança Nubank").         | Deve ser único por `user_id`. Máx: 255.                           |
| `institution_name`  | `string` / `VARCHAR NULL`| Não    | Nome da instituição financeira (ex: "Banco do Brasil", "Nubank").                   | MVP: opcional.                                                    |
| `account_type`      | `string` / `VARCHAR NULL`| Não    | Tipo de conta (ex: "Corrente", "Poupança", "Cartão de Crédito", "Investimento").   | MVP: opcional. Pode ser um enum/tipo definido.                   |
| `created_at`        | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando a conta bancária foi criada.                                 | ADR-002                                                           |
| `updated_at`        | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização da conta bancária.                               | ADR-002                                                           |
| `version`           | `int` / `INTEGER`   | Sim         | Número da versão do registro.                                                      | ADR-006.                                                          |
| `archived_at`       | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a conta foi arquivada (não pode ser usada para novas transações). | ADR-004.                                                          |
| `deleted_at`        | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a conta foi marcada como excluída (soft delete).            | ADR-001.                                                          |

## Relacionamentos

* **Uma Conta Bancária pertence a um:** Usuário ([`../user/user_entity.md`](../user/user.md)) (via `user_id`)
* **Uma Conta Bancária pode estar associada a muitas:** Transações ([`transaction_entity.md`](./transaction.md))

## Regras de Negócio e Validações (Exemplos)

* O nome da conta bancária deve ser único para o usuário.
* Uma conta bancária não pode ser excluída (hard delete) se tiver transações associadas; deve ser arquivada.
* **Saldos:** Os saldos relacionados a uma conta bancária são calculados dinamicamente a partir da soma das suas transações em um determinado período, quando necessário para exibição ou relatórios. A entidade `BankAccount` não armazena saldos.

## Considerações MVP

* MVP inclui CRUD de contas bancárias (foco no nome/apelido).
* `institution_name` e `account_type` são opcionais no MVP.
* No MVP, as transações são vinculadas exclusivamente a estas "Contas Bancárias". A lógica de cartões de crédito com faturas é uma evolução futura.
