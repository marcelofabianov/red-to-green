# Contexto de Domínio: Usuário

## 1. Visão Geral

O **Contexto do Usuário** é fundamental para o RedToGreen, pois representa o indivíduo para o qual todas as funcionalidades de gestão financeira são direcionadas. Este contexto é responsável por gerenciar a identidade do usuário, suas configurações pessoais e as entidades de apoio que ele cria para personalizar sua experiência de organização financeira.

A principal responsabilidade deste contexto é garantir que cada usuário tenha um espaço seguro e isolado para gerenciar suas informações financeiras, com ferramentas que se adaptem às suas preferências de categorização e marcação.

## 2. Entidades Principais do Contexto

As seguintes entidades são centrais para o Contexto do Usuário:

* **Usuário (`user.md`):**
    * Representa a pessoa física que utiliza o sistema.
    * Contém informações de identificação (nome, e-mail) e credenciais de acesso (hash da senha).
    * É o "proprietário" de todas as outras entidades personalizáveis neste contexto.
* **Categoria (`category.md`):**
    * Permite ao usuário criar uma estrutura hierárquica (categorias e subcategorias) para classificar suas receitas e despesas de forma padronizada.
    * Exemplos: "Alimentação" (com subcategorias "Supermercado", "Restaurante"), "Moradia", "Transporte".
    * Cada categoria pertence a um usuário específico.
* **Tag (`tag.md`):**
    * Permite ao usuário criar marcadores flexíveis (palavras-chave) para adicionar um nível adicional de classificação ou contexto às suas transações.
    * Exemplos: `#viagem2025`, `#reembolsável`, `#trabalho`, `#pessoal`.
    * Cada tag pertence a um usuário específico.
* **Conta Bancária (`bank_account.md`):**
    * Permite ao usuário registrar as contas financeiras que ele utiliza (conta corrente, poupança, conta digital, etc.).
    * Embora as transações ocorram dentro das "Carteiras", elas são originadas ou destinadas a estas Contas Bancárias.
    * Cada conta bancária pertence a um usuário específico.

## 3. Responsabilidades e Funcionalidades Chave

* **Gerenciamento de Identidade e Perfil:**
    * Registro de novos usuários.
    * Autenticação de usuários (verificação de credenciais).
    * Atualização de informações do perfil do usuário (nome, e-mail, senha).
* **Gerenciamento de Entidades de Apoio Personalizadas:**
    * Permitir que o usuário crie, visualize, edite, arquive/desarquive e exclua (soft delete) suas próprias Categorias e Subcategorias.
    * Permitir que o usuário crie, visualize, edite, arquive/desarquive e exclua (soft delete) suas próprias Tags.
    * Permitir que o usuário crie, visualize, edite, arquive/desarquive e exclua (soft delete) suas próprias Contas Bancárias.
* **Fornecimento de Dados para Outros Contextos:**
    * Disponibilizar as Categorias, Tags e Contas Bancárias ativas do usuário para serem selecionadas durante o registro de transações no Contexto da Carteira.

## 4. Interações com Outros Contextos

* **Contexto da Carteira:** O Usuário é o proprietário das Carteiras. As Categorias, Tags e Contas Bancárias definidas neste contexto são utilizadas para detalhar as Transações dentro de uma Carteira.
* **Segurança e Autenticação (Contexto Implícito/Transversal):** Este contexto lida com a entidade Usuário, que é a base para os processos de autenticação e autorização (RBAC, conforme ADR-008, focado nas permissões da Carteira).

## 5. Considerações de Design

* **Isolamento de Dados do Usuário:** Todas as entidades personalizáveis (Categorias, Tags, Contas Bancárias) devem ser estritamente isoladas por `user_id` para garantir a privacidade e a organização individual.
* **Usabilidade:** A interface para gerenciar essas entidades deve ser intuitiva, refletindo a simplicidade proposta pelo RedToGreen.
* **Flexibilidade:** O usuário deve ter liberdade para criar as estruturas de organização que melhor se adaptem às suas necessidades.
