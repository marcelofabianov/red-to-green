# Entidade: Transação (Transaction)

**Contexto de Domínio:** Carteira

## Descrição

Representa uma movimentação financeira individual, seja uma receita (entrada de fundos) ou uma despesa (saída de fundos). É a entidade central para o controle financeiro no RedToGreen e está sempre associada a uma Carteira e a uma Conta Bancária do usuário, além de ser classificada por Categorias e Tags.

## Atributos

| Atributo                | Tipo (Go/BD)        | Obrigatório | Descrição                                                                     | Observações / ADRs Relevantes                                     |
| :---------------------- | :------------------ | :---------- | :---------------------------------------------------------------------------- | :---------------------------------------------------------------- |
| `id`                    | `uuid.UUID` / `UUID`| Sim         | Identificador único universal da transação.                                   | ADR-003 (UUID v7)                                                 |
| `wallet_id`             | `uuid.UUID` / `UUID`| Sim         | ID da Carteira à qual esta transação pertence.                                | Chave estrangeira para `wallet.id`.                               |
| `user_id`               | `uuid.UUID` / `UUID`| Sim         | ID do usuário proprietário (identifica o dono da `wallet_id`).                | Chave estrangeira para `user.id`. Desnormalização para facilitar queries diretas de transações por usuário; a autorização primária deve vir pela `wallet_id`. |
| `bank_account_id`       | `uuid.UUID` / `UUID`| Sim         | ID da Conta Bancária do usuário afetada por esta transação.                   | Chave estrangeira para `bank_account.id`.                         |
| `category_id`           | `uuid.UUID` / `UUID`| Sim         | ID da Categoria (ou Subcategoria) do usuário à qual esta transação está classificada. | Chave estrangeira para `category.id`.                             |
| `title`                 | `string` / `VARCHAR(255)`| Sim     | Descrição curta/título da transação (ex: "Almoço", "Salário").                | Conforme `visao_geral_do_produto.md`.                             |
| `description`           | `string` / `VARCHAR(255) NULL`| Não | Detalhes adicionais opcionais sobre a transação.                            | Conforme `visao_geral_do_produto.md`.                             |
| `amount`                | `decimal` / `DECIMAL`| Sim         | Valor monetário da transação (deve ser armazenado como um valor absoluto positivo). | Precisão e escala a definir (ex: 15,2).                           |
| `type`                  | `string` / `VARCHAR(10)`| Sim     | Tipo da transação: "INCOME" (Receita) ou "EXPENSE" (Despesa).               | Pode ser um ENUM no banco (`transaction_type`).                   |
| `transaction_date`      | `time.Time` / `DATE`  | Sim         | Data em que a transação ocorreu ou foi lançada pelo usuário.                  |                                                                   |
| `due_date`              | `*time.Time` / `DATE NULL`| Não   | Data de vencimento (para despesas/contas a pagar ou receitas a receber).      | Opcional, conforme `visao_geral_do_produto.md`.                   |
| `payment_date`          | `*time.Time` / `DATE NULL`| Não   | Data em que a transação foi efetivamente liquidada (paga/recebida).         | Opcional, conforme `visao_geral_do_produto.md`.                   |
| `status`                | `string` / `VARCHAR(10)`| Sim     | Posição/Situação da transação (ex: "PENDING", "PAID", "OVERDUE").             | Default "PENDING". Pode ser derivado das datas ou gerenciado explicitamente. |
| `created_at`            | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora de quando a transação foi criada no sistema.                      | ADR-002                                                           |
| `updated_at`            | `time.Time` / `TIMESTAMPTZ` | Sim   | Data e hora da última atualização da transação.                               | ADR-002                                                           |
| `version`               | `int` / `INTEGER`   | Sim         | Número da versão do registro, para controle de concorrência.                  | ADR-006.                                                          |
| `deleted_at`            | `*time.Time` / `TIMESTAMPTZ NULL` | Não | Data e hora de quando a transação foi marcada como excluída (soft delete).   | ADR-001.                                                          |

## Relacionamentos

* **Uma Transação pertence a uma:** Carteira ([`wallet.md`](./wallet.md)) (via `wallet_id`)
* **Uma Transação está associada a uma:** Conta Bancária ([`bank_account.md`](./bank_account.md)) (via `bank_account_id`)
* **Uma Transação está associada a uma:** Categoria ([`../user/category.md`](../user/category.md)) (via `category_id`)
* **Uma Transação pertence a um:** Usuário ([`../user/user.md`](../user/user.md)) (via `user_id`, que é o proprietário da `wallet_id`)
* **Muitas Transações podem estar associadas a muitas:** Tags ([`../user/tag.md`](../user/tag.md)) (Esta relação é Muitos-para-Muitos e será implementada através de uma tabela de junção, por exemplo, `transaction_tags` contendo `transaction_id` e `tag_id`).

## Regras de Negócio e Validações (Exemplos)

* `amount` deve ser sempre um valor positivo. O campo `type` ("INCOME" ou "EXPENSE") determina como este valor afeta os cálculos de saldo.
* Se `payment_date` estiver preenchido, o `status` da transação deve ser "PAID".
* Se `due_date` for anterior à data atual e `payment_date` estiver nulo, o `status` da transação pode ser automaticamente considerado "OVERDUE" (ou atualizado por um processo).
* A `transaction_date` não pode ser uma data futura distante (validação de sanidade).
* A `bank_account_id`, `category_id` e `wallet_id` devem referenciar entidades existentes e pertencentes (ou acessíveis) ao `user_id` associado.
* Ao menos uma Tag deve ser associada (conforme `visao_geral_do_produto.md`).

## Considerações MVP

* Todos os campos marcados como "Obrigatório" na `visao_geral_do_produto.md` estão refletidos.
* A lógica de derivação do campo `status` (Pendente, Pago, Vencido) com base nas datas (`transaction_date`, `due_date`, `payment_date`) será implementada.
* Não há campo `currency_code` na transação, assumindo uma moeda única para o sistema/usuário no MVP (BRL).
