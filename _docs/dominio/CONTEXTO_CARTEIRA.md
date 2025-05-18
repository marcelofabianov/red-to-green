# Contexto de Domínio: Carteira

## 1. Visão Geral

O **Contexto da Carteira** no RedToGreen é onde a gestão financeira ativa do usuário acontece. Ele se concentra no registro, acompanhamento e organização das movimentações financeiras (transações) dentro de diferentes agrupadores que o usuário define como "Carteiras".

Uma "Carteira" funciona como um centro de custo ou um envelope financeiro que permite ao usuário separar e analisar diferentes aspectos de suas finanças. Este contexto depende fortemente das entidades de apoio definidas no Contexto do Usuário (como Categorias, Tags e Contas Bancárias) para classificar e detalhar as transações.

## 2. Entidades Principais do Contexto

* **Carteira (Wallet - `wallet.md`):**
    * É o principal agrupador de transações. Um usuário pode ter múltiplas carteiras (ex: "Pessoal", "Despesas da Casa", "Projeto Viagem").
    * Pertence a um usuário e pode ser compartilhada com outros usuários (conforme ADR-008).
    * Não armazena saldo diretamente; os valores são derivados das transações contidas nela.
* **Transação (Transaction - `transaction.md`):**
    * Representa qualquer receita (entrada) ou despesa (saída) de dinheiro.
    * Está sempre associada a uma única Carteira.
    * É vinculada a uma Conta Bancária (para indicar origem/destino dos fundos) e classificada por Categorias e Tags (definidas no Contexto do Usuário).

## 3. Responsabilidades e Funcionalidades Chave

* **Gerenciamento de Carteiras:**
    * Permitir que o usuário (proprietário) crie, visualize, edite, arquive/desarquive e exclua (soft delete) suas Carteiras.
* **Gerenciamento de Transações:**
    * Permitir o registro de novas transações (receitas e despesas) dentro de uma Carteira específica.
    * Permitir a visualização de transações (com filtros por período, tipo, categoria, tag, conta bancária, status).
    * Permitir a edição e exclusão (soft delete) de transações.
    * Gerenciar o status das transações (Pendente, Paga, Vencida).
* **Cálculo e Apresentação de Saldos e Resumos:**
    * Calcular dinamicamente o saldo de uma Carteira (total de receitas - total de despesas) para um período específico.
    * Fornecer resumos de gastos/ganhos por categoria/tag dentro de uma Carteira.
* **Compartilhamento de Carteiras (ADR-008):**
    * Permitir que o proprietário de uma Carteira a compartilhe com outros usuários, atribuindo papéis específicos (ex: Visualizador, Colaborador).
    * Gerenciar as permissões de acesso às transações de uma Carteira com base nesses papéis.

## 4. Interações com Outros Contextos

* **Contexto do Usuário:**
    * Uma Carteira sempre pertence a um Usuário.
    * Utiliza as Categorias, Tags e Contas Bancárias definidas pelo Usuário para classificar e detalhar as Transações.
    * A lógica de compartilhamento de Carteiras envolve múltiplos Usuários e seus papéis (RBAC).
* **Relatórios e Análises (Contexto Futuro):** Os dados de Transações e Carteiras serão a principal fonte para a geração de relatórios financeiros e análises.

## 5. Considerações de Design

* **Foco no Fluxo de Transações:** O coração deste contexto é o ciclo de vida das transações.
* **Derivação de Saldos:** Conforme decisão de simplificação, os saldos não são armazenados diretamente nas entidades `Wallet` ou `BankAccount`, mas calculados sob demanda a partir das transações para garantir consistência e simplicidade no MVP.
* **Clareza na Vinculação:** As transações devem estar claramente vinculadas à sua Carteira, Conta Bancária, Categoria e Tags para permitir uma organização eficaz.
