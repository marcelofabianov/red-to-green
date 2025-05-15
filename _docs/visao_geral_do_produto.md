# Documentação Inicial do Produto: RedToGreen (Revisão 1)

## 1. Visão Geral do Produto

O **RedToGreen** é uma aplicação de controle financeiro pessoal, concebida com foco na simplicidade e facilidade de uso. O objetivo principal é auxiliar usuários, especialmente aqueles com pouca experiência em gestão financeira, a organizar suas finanças, saindo de um cenário de descontrole ("vermelho") para uma situação de clareza e planejamento ("verde").

O desenvolvimento inicial seguirá o modelo de MVP (Minimum Viable Product), entregando as funcionalidades essenciais para o registro e acompanhamento de transações financeiras. O produto evoluirá com funcionalidades incrementais em versões futuras, baseadas no feedback e nas necessidades dos usuários.

## 2. Nome do Produto

RedToGreen

## 3. Funcionalidades Centrais do MVP

O MVP do RedToGreen se concentrará no gerenciamento de transações financeiras, que serão consolidadas dentro de "Carteiras". Todas as entidades de organização (Carteiras, Contas Bancárias, Tags, Categorias) serão vinculadas ao usuário logado, permitindo que cada usuário gerencie seus próprios dados.

### 3.1. Carteiras (Centros de Custo)

* As "Carteiras" funcionarão como centros de custo principais, vinculadas ao usuário, onde todas as suas transações financeiras serão agrupadas (ex: "Minhas Despesas Pessoais", "Contas da Casa").
* O usuário poderá cadastrar, visualizar, editar e excluir suas próprias Carteiras.

### 3.2. Contas Bancárias

* O usuário poderá cadastrar, visualizar, editar e excluir suas próprias Contas Bancárias (ex: "Conta Corrente Banco A", "Conta Poupança Banco B").
* Para o MVP, cada conta bancária será identificada principalmente por um nome/apelido fornecido pelo usuário. Detalhes adicionais como instituição, agência e número da conta podem ser considerados para futuras evoluções.

### 3.3. Tags

* O usuário poderá cadastrar, visualizar, editar e excluir seu próprio conjunto de Tags.
* Tags são palavras-chave (ex: #supermercado, #saude, #educacao) que o usuário pode criar e associar às suas transações para facilitar a classificação livre e a filtragem.

### 3.4. Categorias e Subcategorias

* O usuário poderá cadastrar, visualizar, editar e excluir sua própria estrutura hierárquica de Categorias e Subcategorias (ex: Categoria: "Alimentação", Subcategoria: "Restaurante"; Categoria: "Moradia", Subcategoria: "Aluguel").
* Esta estrutura ajudará o usuário a organizar suas transações de forma mais padronizada para análises futuras.

### 3.5. Transações Financeiras

O registro de transações é a funcionalidade chave do sistema. Cada transação conterá os seguintes campos:

* **Título:** Uma breve descrição para identificar a transação (obrigatório, máximo de 255 caracteres).
* **Descrição Adicional:** Um campo para detalhes complementares sobre a transação (opcional, máximo de 255 caracteres).
* **Valor:** O montante financeiro da transação (obrigatório).
* **Tipo:** Indicador se a transação é uma "Receita" ou uma "Despesa" (obrigatório).
* **Data da Transação:** Data em que a transação ocorreu ou foi lançada no sistema (obrigatório).
* **Data de Vencimento:** Data limite para pagamento (para despesas) ou recebimento (para receitas) (opcional).
* **Data de Pagamento/Recebimento:** Data em que a transação foi efetivamente liquidada (opcional).
* **Status (Posição):** Indica a situação atual da transação (ex: "Pendente", "Pago", "Vencido"). *Nota: Este status pode ser inicialmente derivado da presença da "Data de Pagamento" e da comparação da "Data de Vencimento" com a data atual.*
* **Conta Bancária Vinculada:** Especifica a qual **Conta Bancária (previamente cadastrada pelo usuário)** a transação está associada (obrigatório). *Nota: Nesta fase inicial (MVP), as transações não serão vinculadas diretamente a cartões de crédito individuais ou dinheiro em espécie; o foco é exclusivamente em contas bancárias.*
* **Tags:** Um conjunto de **Tags (previamente cadastradas pelo usuário)** para classificar a transação (obrigatório, permitindo múltiplas tags por transação).
* **Categoria/Subcategoria:** Uma estrutura hierárquica de **Categorias/Subcategorias (previamente cadastradas pelo usuário)** para classificar a transação (obrigatório).
* **Vinculação à Carteira:** Toda transação deverá ser obrigatoriamente associada a uma **Carteira (previamente cadastrada pelo usuário)** (obrigatório).

## 4. Evolução Contínua

Após a validação do MVP, o RedToGreen será enriquecido com novas funcionalidades, como:
* Relatórios financeiros (fluxo de caixa, despesas por categoria, etc.).
* Planejamento orçamentário.
* Definição de metas financeiras.
* Alertas e notificações.
* Consideração de outras formas de pagamento/recebimento como cartões de crédito.
