# Entidade: Tag

**Contexto de Domínio:** Usuário

## Descrição

Representa uma etiqueta ou palavra-chave flexível definida pelo usuário para marcar suas transações.

## Atributos

| Atributo         | Tipo (Go/BD)        | Obrigatório | Descrição                                                                     | Observações / ADRs Relevantes                                     |
| :--------------- | :------------------ | :---------- | :---------------------------------------------------------------------------- | :---------------------------------------------------------------- |
| `id`             | `uuid.UUID` / `UUID`| Sim         | Identificador único universal da tag.                                         | ADR-003 (UUID v7)                                                 |
| `user_id`        | `uuid.UUID` / `UUID`| Sim         | **ID do usuário proprietário desta tag.** | Chave estrangeira para `user.id`.                                 |
| `name`           | `string` / `VARCHAR`| Sim         | Nome da tag (ex: "viagem", "reembolsavel").                                   | Deve ser único por `user_id`. Máx: 100.                           |
| `created_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando a tag foi criada.                                       | ADR-002                                                           |
| `updated_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização da tag.                                     | ADR-002                                                           |
| `version`        | `int` / `INTEGER`   | Sim         | Número da versão do registro.                                                 | ADR-006.                                                          |
| `archived_at`    | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a tag foi arquivada.                                  | ADR-004.                                                          |
| `deleted_at`     | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a tag foi marcada como excluída (soft delete).         | ADR-001.                                                          |

## Relacionamentos

* **Uma Tag pertence a um:** Usuário ([`user.md`](./user.md)) (via `user_id`)
* **Muitas Tags podem estar associadas a muitas:** Transações ([`../wallet/transaction.md`](../wallet/transaction.md)) (via tabela de junção `transaction_tags`)

## Regras de Negócio e Validações (Exemplos)

* O nome da tag deve ser único para um mesmo usuário.
