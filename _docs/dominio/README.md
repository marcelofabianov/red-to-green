# Modelagem do Domínio RedToGreen

Este diretório contém a documentação detalhada das entidades centrais que compõem o domínio da aplicação RedToGreen. A modelagem do domínio é fundamental para o entendimento da lógica de negócios e da estrutura de dados do sistema.

A organização das entidades segue os principais contextos de domínio identificados:

* **Contexto do Usuário:** Entidades e configurações gerenciadas diretamente pelo usuário para personalizar sua experiência e organizar seus dados.
* **Contexto da Carteira:** Entidades relacionadas ao agrupamento e registro das movimentações financeiras do usuário.

## Entidades do Domínio

A seguir, uma lista das entidades de domínio documentadas, com uma breve descrição e links para seus detalhes:

### Contexto do Usuário

Este contexto gerencia as informações e configurações pessoais do usuário.

* **[Usuário (`user.md`)](./user/user.md)**
    * Representa um indivíduo registrado e autenticado no sistema RedToGreen.
* **[Categoria (`category.md`)](./user/category.md)**
    * Representa uma classificação definida pelo usuário para organizar suas transações financeiras.
* **[Tag (`tag.md`)](./user/tag.md)**
    * Representa uma etiqueta ou palavra-chave flexível definida pelo usuário para marcar suas transações.

### Contexto da Carteira

Este contexto lida com o agrupamento das transações financeiras e as contas de origem/destino dos fundos.

* **[Carteira (`wallet.md`)](./wallet/wallet.md)**
    * Representa um agrupador financeiro principal ou "centro de custo" pertencente a um usuário.
* **[Conta Bancária (`bank_account.md`)](./wallet/bank_account.md)**
    * Representa uma conta bancária real ou virtual do usuário, usada como referência em transações.
* **[Transação (`transaction.md`)](./wallet/transaction.md)**
    * Representa uma movimentação financeira individual (receita ou despesa).

## Documentos de Apoio ao Domínio

Para um entendimento mais aprofundado do domínio e seus termos, consulte também:

* **Contextos de Domínio Detalhados:**
    * [`CONTEXTO_USUARIO.md`](./CONTEXTO_USUARIO.md)
    * [`CONTEXTO_CARTEIRA.md`](./CONTEXTO_CARTEIRA.md)
* **Glossário de Termos:**
    * [`GLOSSARIO.md`](../GLOSSARIO.md
---
