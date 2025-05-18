# Entidade: Categoria (Category)

**Contexto de Domínio:** Usuário

## Descrição

Representa uma classificação definida pelo usuário para organizar suas transações financeiras (ex: Alimentação, Moradia, Transporte).

## Atributos

| Atributo         | Tipo (Go/BD)        | Obrigatório | Descrição                                                                     | Observações / ADRs Relevantes                                     |
| :--------------- | :------------------ | :---------- | :---------------------------------------------------------------------------- | :---------------------------------------------------------------- |
| `id`             | `uuid.UUID` / `UUID`| Sim         | Identificador único universal da categoria.                                   | ADR-003 (UUID v7)                                                 |
| `user_id`        | `uuid.UUID` / `UUID`| Sim         | **ID do usuário proprietário desta categoria.** | Chave estrangeira para `user.id`.                                 |
| `name`           | `string` / `VARCHAR`| Sim         | Nome da categoria (ex: "Alimentação").                                        | Deve ser único por `user_id` e `parent_id`. Máx: 255.             |
| `parent_id`      | `*uuid.UUID` / `UUID NULL` | Não    | ID da categoria pai (se for uma subcategoria). Nulo para categorias raiz.   | Permite estrutura hierárquica.                                  |
| `created_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando a categoria foi criada.                                 | ADR-002                                                           |
| `updated_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização da categoria.                               | ADR-002                                                           |
| `version`        | `int` / `INTEGER`   | Sim         | Número da versão do registro.                                                 | ADR-006.                                                          |
| `archived_at`    | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a categoria foi arquivada.                            | ADR-004.                                                          |
| `deleted_at`     | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a categoria foi marcada como excluída (soft delete).   | ADR-001.                                                          |

## Relacionamentos

* **Uma Categoria pertence a um:** Usuário ([`user.md`](./user.md)) (via `user_id`)
* **Uma Categoria pode ter uma:** Categoria Pai (auto-relacionamento via `parent_id`)
* **Uma Categoria pode ter muitas:** Subcategorias (auto-relacionamento)
* **Muitas Categorias podem estar associadas a muitas:** Transações ([`../wallet/transaction.md`](../wallet/transaction.md))

## Regras de Negócio e Validações (Exemplos)

* O nome da categoria deve ser único para um mesmo usuário dentro do mesmo nível hierárquico (mesmo `parent_id`).
