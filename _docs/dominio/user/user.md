# Entidade: Usuário (User)

**Contexto de Domínio:** Usuário

## Descrição

Representa um indivíduo registrado e autenticado no sistema RedToGreen. É a entidade central à qual todas as outras informações financeiras e de configuração pessoal são vinculadas.

## Atributos

| Atributo         | Tipo (Go/BD)        | Obrigatório | Descrição                                                                 | Observações / ADRs Relevantes                                     |
| :--------------- | :------------------ | :---------- | :------------------------------------------------------------------------ | :---------------------------------------------------------------- |
| `id`             | `uuid.UUID` / `UUID`| Sim         | Identificador único universal do usuário.                                 | ADR-003 (UUID v7)                                                 |
| `name`           | `string` / `VARCHAR`| Sim         | Nome completo ou apelido do usuário.                                      |                                                                   |
| `email`          | `string` / `VARCHAR`| Sim         | Endereço de e-mail do usuário, usado para login e comunicação.            | Deve ser único no sistema.                                        |
| `password_hash`  | `string` / `VARCHAR`| Sim         | Hash da senha do usuário.                                                 | ADR-007 (Argon2). Nunca armazenar a senha em texto plano.         |
| `created_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando o registro do usuário foi criado.                   | ADR-002                                                           |
| `updated_at`     | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização do registro do usuário.                 | ADR-002                                                           |
| `version`        | `int` / `INTEGER`   | Sim         | Número da versão do registro, para controle de concorrência.              | ADR-006. Inicializa com 1.                                        |
| `archived_at`    | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando o usuário foi arquivado (se aplicável).         | ADR-004.                                                          |
| `deleted_at`     | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando o usuário foi marcado como excluído (soft delete). | ADR-001.                                                          |

## Relacionamentos

* **Um Usuário possui muitas:**
    * Categorias ([`category.md`](./category.md))
    * Tags ([`tag.md`](./tag.md))
    * Carteiras ([`../wallet/wallet.md`](../wallet/wallet.md))
    * Contas Bancárias ([`../wallet/bank_account.md`](../wallet/bank_account.md))

## Regras de Negócio e Validações (Exemplos)

* O e-mail deve ser único em todo o sistema.
* A senha deve atender a critérios mínimos de complexidade e ser armazenada em um formato seguro.
