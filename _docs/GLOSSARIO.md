# Glossário de Termos do Domínio RedToGreen

Este documento define os termos chave utilizados no contexto do produto RedToGreen para garantir um entendimento comum entre todos os envolvidos no projeto.

---

**AC (Critério de Aceite / Acceptance Criteria)**
* **Definição:** Um conjunto de condições que uma funcionalidade de software (geralmente uma User Story) deve satisfazer para ser considerada completa e aceita pelo stakeholder ou product owner. São testáveis e definem o escopo e a qualidade esperada da entrega.

**ADR (Architecture Decision Record)**
* **Definição:** Um documento que captura uma decisão arquitetural importante, o contexto por trás dela, as alternativas consideradas e suas consequências. Utilizado para registrar o histórico e a justificativa das escolhas técnicas do projeto.
* **Referência:** ADR-000

**Arquivamento (Archiving)**
* **Definição:** Processo de marcar um registro (ex: Categoria, Tag, Conta Bancária, Carteira) como inativo para novos usos, mas mantendo-o visível para consulta em dados históricos e para gerenciamento. Diferente de `soft delete`.
* **Referência:** ADR-004

**Atributo (Attribute)**
* **Definição:** Uma característica ou propriedade de uma entidade. Por exemplo, `nome` e `email` são atributos da entidade `Usuário`.

**Autenticação (Authentication)**
* **Definição:** O processo de verificar a identidade de um usuário, geralmente através de credenciais como e-mail e senha.

**Autorização (Authorization)**
* **Definição:** O processo de determinar quais ações um usuário autenticado tem permissão para realizar sobre determinados recursos.
* **Referência:** ADR-008 (RBAC)

**Backend**
* **Definição:** A parte do sistema RedToGreen que lida com a lógica de negócios, processamento de dados, persistência e a API que serve os clientes (como o frontend web).
* **Referência:** ADR-010, ADR-011

**Carteira (Wallet)**
* **Definição:** Principal centro de custo ou agrupador financeiro de um usuário no RedToGreen. Todas as transações são vinculadas a uma Carteira. Permite ao usuário separar e organizar suas finanças em diferentes contextos (ex: "Despesas Pessoais", "Contas da Casa"). Não armazena saldo diretamente; os valores são derivados de suas transações.
* **Entidade:** [`./dominio/wallet/wallet.md`](./dominio/wallet/wallet.md)

**Categoria (Category)**
* **Definição:** Um rótulo estruturado, definido e gerenciado pelo usuário, para classificar suas transações financeiras (ex: "Alimentação", "Transporte"). Pode incluir Subcategorias para maior detalhamento.
* **Entidade:** [`./dominio/user/category.md`](./dominio/user/category.md)

**Conta Bancária (BankAccount)**
* **Definição:** Representa uma conta financeira real ou virtual do usuário (ex: conta corrente, poupança, conta digital) registrada no sistema. As transações são vinculadas a uma Conta Bancária para indicar a origem ou destino dos fundos. Não armazena saldo diretamente.
* **Entidade:** [`./dominio/wallet/bank_account.md`](./dominio/wallet/bank_account.md) (Nota: Gerenciada pelo Usuário, usada no contexto de Transações da Carteira)

**Contexto de Domínio (Domain Context)**
* **Definição:** Uma delimitação lógica dentro do domínio da aplicação que agrupa entidades e funcionalidades com alta coesão e responsabilidades relacionadas. Ajuda a organizar a complexidade do sistema.
* **Exemplos:** Contexto do Usuário, Contexto da Carteira.
* **Referência:** [`./dominio/CONTEXTO_USUARIO.md`](./dominio/CONTEXTO_USUARIO.md), [`./dominio/CONTEXTO_CARTEIRA.md`](./dominio/CONTEXTO_CARTEIRA.md)

**CRUD (Create, Read, Update, Delete)**
* **Definição:** Acrônimo para as quatro operações básicas de persistência de dados: Criar, Ler, Atualizar e Excluir.

**Despesa (Expense)**
* **Definição:** Um tipo de transação que representa uma saída de dinheiro. Parte da entidade Transação.

**Entidade (Entity)**
* **Definição:** Um objeto fundamental no domínio do problema, geralmente com uma identidade única e um ciclo de vida. Exemplos: Usuário, Transação, Carteira.

**Épico (Epic)**
* **Definição:** Um grande corpo de trabalho ou uma feature de alto nível que pode ser dividida em várias User Stories menores. Ajuda a organizar o backlog e a comunicar o progresso em um nível estratégico.
* **Exemplo:** [`./epicos/EPIC_001_Gerenciamento_Configuracoes_Usuario.md`](./epicos/EPIC_001_Gerenciamento_Configuracoes_Usuario.md)

**MVP (Minimum Viable Product / Produto Mínimo Viável)**
* **Definição:** A versão inicial do RedToGreen contendo apenas o conjunto essencial de funcionalidades necessárias para entregar valor aos primeiros usuários e validar as principais hipóteses do produto.
* **Referência:** [`./produto/visao_geral_v1.md`](./produto/visao_geral_v1.md)

**RBAC (Role-Based Access Control / Controle de Acesso Baseado em Papéis)**
* **Definição:** Modelo de autorização onde as permissões são atribuídas a papéis, e os usuários recebem papéis, determinando o que podem acessar e fazer.
* **Referência:** ADR-008

**Receita (Income)**
* **Definição:** Um tipo de transação que representa uma entrada de dinheiro. Parte da entidade Transação.

**Requisito Funcional (RF)**
* **Definição:** Uma descrição específica de *o que* o sistema deve fazer para atender a uma necessidade do usuário ou do negócio, geralmente derivado de uma User Story.

**Requisito Não Funcional (NRF)**
* **Definição:** Uma descrição de *como* o sistema deve operar ou uma qualidade que ele deve possuir (ex: performance, segurança, usabilidade, confiabilidade). Pode ser global ou específico.

**Soft Delete**
* **Definição:** Estratégia de marcar um registro como "excluído" sem removê-lo fisicamente do banco de dados, geralmente preenchendo uma coluna como `deleted_at`.
* **Referência:** ADR-001

**Tag**
* **Definição:** Uma palavra-chave ou marcador livre, definido e gerenciado pelo usuário, para adicionar contexto ou facilitar a filtragem de transações (ex: `#viagem`, `#urgente`, `#reembolsavel`). Uma transação pode ter múltiplas tags.
* **Entidade:** [`./dominio/user/tag.md`](./dominio/user/tag.md)

**Transação (Transaction)**
* **Definição:** Qualquer movimentação financeira registrada pelo usuário no RedToGreen, podendo ser uma Receita ou uma Despesa. É a unidade fundamental de registro financeiro.
* **Entidade:** [`./dominio/wallet/transaction.md`](./dominio/wallet/transaction.md)

**User Story (US / História de Usuário)**
* **Definição:** Uma descrição informal e concisa de uma funcionalidade sob a perspectiva do usuário final, focando no valor que ela entrega. Geralmente no formato: "Como um `<tipo de usuário>`, eu quero `<fazer algo>` para que `<obtenha um benefício>`."
* **Exemplo:** [`./user_stories/US_001_Gerenciar_Categorias_Pessoais.md`](./user_stories/US_001_Gerenciar_Categorias_Pessoais.md)

**Usuário (User)**
* **Definição:** A pessoa que se registra e utiliza o sistema RedToGreen para gerenciar suas finanças pessoais.
* **Entidade:** [`./dominio/user/user.md`](./dominio/user/user.md)

*(Este glossário é um documento vivo e deve ser atualizado conforme novos termos e conceitos são introduzidos ou refinados no projeto RedToGreen.)*
