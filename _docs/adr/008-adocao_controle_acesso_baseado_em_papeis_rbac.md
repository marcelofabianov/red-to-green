# ADR-008: Adoção do Controle de Acesso Baseado em Papéis (RBAC)

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Para a aplicação RedToGreen, é crucial implementar um sistema de autorização que defina o que cada usuário pode fazer, especialmente considerando a funcionalidade de compartilhamento de "Wallets" (Carteiras) com diferentes níveis de permissão. O produto visa a simplicidade para usuários leigos em finanças. Uma análise prévia de diversos modelos de controle de acesso foi realizada, incluindo RBAC, ABAC, PBAC, ReBAC, entre outros.

A principal funcionalidade que necessita de um modelo de autorização claro é o compartilhamento de uma Wallet por seu proprietário com outros usuários, concedendo-lhes capacidades específicas (ex: apenas visualizar, visualizar e adicionar/editar transações).

A pergunta que estamos tentando responder é: Qual modelo de controle de acesso devemos adotar para gerenciar as permissões dos usuários no RedToGreen, de forma que seja seguro, funcional para os casos de uso de compartilhamento e, ao mesmo tempo, mantenha a simplicidade para o usuário final e para a implementação inicial?

O escopo desta decisão abrange a autorização de todas as ações significativas que um usuário pode realizar sobre os recursos da aplicação, com foco principal nas Wallets e nas Transações financeiras contidas nelas.

## Decisão

Adotaremos o modelo de **Controle de Acesso Baseado em Papéis (RBAC)** como a estratégia principal para autorização no sistema RedToGreen.

1.  **Definição de Papéis:** Serão definidos papéis claros e limitados, associados principalmente ao contexto de uma Wallet específica. Exemplos de papéis iniciais incluem:
    * `DONO_DA_WALLET`: O criador e proprietário da Wallet. Possui todas as permissões sobre ela, incluindo visualizar, adicionar/editar/excluir transações, gerenciar o compartilhamento (convidar outros usuários, modificar seus papéis, remover acesso) e excluir a própria Wallet.
    * `VISUALIZADOR_DA_WALLET`: Um usuário convidado que pode apenas visualizar os dados da Wallet e suas transações, sem permissão para realizar modificações.
    * `COLABORADOR_DA_WALLET`: Um usuário convidado que pode visualizar os dados da Wallet e também adicionar e editar transações dentro dela. A capacidade de excluir transações por este papel será avaliada (pode ser uma permissão separada ou um papel distinto, como "Editor"). Este papel não permite excluir a Wallet ou gerenciar seus compartilhamentos.

2.  **Atribuição de Papéis:** Um usuário terá um papel específico *em relação a uma Wallet específica*.
    * O criador de uma Wallet é automaticamente atribuído o papel `DONO_DA_WALLET` para essa Wallet.
    * Ao compartilhar uma Wallet, o `DONO_DA_WALLET` atribui um dos papéis de convidado (ex: `VISUALIZADOR_DA_WALLET`, `COLABORADOR_DA_WALLET`) a outro usuário para *aquela Wallet*.

3.  **Verificação de Permissões:** A lógica de autorização do sistema verificará se o papel do usuário autenticado em relação à Wallet alvo lhe concede a permissão para realizar a ação solicitada sobre a Wallet ou seus recursos (ex: transações).

## Alternativas Consideradas (Opcional)

* **ReBAC (Relationship-Based Access Control):**
    * Descrição: Modelo que baseia permissões em relacionamentos diretos entre entidades (ex: "Usuário A" é "proprietário de" "Wallet X"; "Usuário B" "tem permissão de visualização para" "Wallet X" concedida por "Usuário A").
    * Motivo da Rejeição (como estratégia principal única): Embora muito adequado para modelar o compartilhamento, a implementação e o gerenciamento da lógica de grafo de relacionamentos podem introduzir uma complexidade maior do que o RBAC para os requisitos iniciais. O RBAC, quando aplicado contextualmente a um recurso (Wallet), pode abstrair esses relacionamentos de forma mais simples para o MVP.

* **ABAC (Attribute-Based Access Control) / PBAC (Policy-Based Access Control):**
    * Descrição: Modelos que permitem regras de autorização extremamente flexíveis e granulares baseadas em atributos do usuário, recurso, ação e ambiente, definidas através de políticas.
    * Motivo da Rejeição: Considerados excessivamente complexos para a configuração inicial, gerenciamento e depuração, dado o foco do RedToGreen na simplicidade para o usuário leigo e para o desenvolvimento do MVP. A flexibilidade oferecida não é um requisito primordial nesta fase.

* **Abordagem Híbrida (ex: ReBAC com interface simplificada via RBAC):**
    * Descrição: Combinar a precisão do ReBAC na lógica de backend com uma interface de atribuição de papéis simplificada para o usuário.
    * Motivo da Rejeição (para a fase inicial): A decisão por RBAC puro (aplicado contextualmente à Wallet) visa priorizar a simplicidade de implementação e entendimento para o MVP. Uma evolução para um modelo mais híbrido pode ser considerada no futuro se a complexidade das regras de compartilhamento e permissões aumentar significativamente.

## Consequências

**Positivas:**

* **Simplicidade e Intuitividade:** O modelo RBAC é conceitualmente simples e fácil de ser entendido tanto pelos desenvolvedores quanto pelos usuários finais (ao atribuir papéis como "Visualizador" ou "Colaborador" no contexto de um compartilhamento).
* **Gerenciamento Facilitado:** As permissões são agrupadas em papéis, o que pode simplificar o gerenciamento, especialmente com um número limitado e bem definido de papéis.
* **Implementação Direta para o MVP:** A implementação do RBAC para os cenários de compartilhamento de Wallets com níveis de acesso distintos (visualizar, colaborar) é relativamente direta.
* **Boa Base para Começar:** Fornece uma estrutura de autorização funcional e segura para o MVP, que pode ser estendida ou adaptada no futuro.
* **Clareza nas Responsabilidades:** Os papéis definem claramente as capacidades de cada usuário em relação a uma Wallet compartilhada.

**Negativas / Trade-offs:**

* **Granularidade Limitada:** Se houver necessidade de permissões muito específicas que não se encaixam bem nos papéis predefinidos (ex: "Colaborador que só pode adicionar transações, mas não editar" ou "Colaborador que só pode editar transações abaixo de um certo valor"), pode levar à proliferação de papéis ou à necessidade de complementar o RBAC.
* **Menos Flexível para Regras Dinâmicas:** O RBAC é menos adequado para cenários que exigem decisões de autorização baseadas em múltiplos atributos dinâmicos (onde ABAC/PBAC seriam mais apropriados).
* **Escalabilidade dos Papéis:** Se o número de papéis distintos e complexos crescer muito, o gerenciamento pode se tornar um desafio. É crucial manter o conjunto de papéis conciso.
* **Necessidade de Contextualização ao Recurso:** A principal característica da implementação aqui é que o papel é atribuído a um usuário *para uma Wallet específica*. A lógica de aplicação deve garantir essa contextualização corretamente.

**(Opcional) Notas Adicionais:**

* A implementação deve garantir que um usuário possa ter papéis diferentes para Wallets diferentes.
* As permissões exatas associadas a cada papel (`DONO_DA_WALLET`, `VISUALIZADOR_DA_WALLET`, `COLABORADOR_DA_WALLET`) devem ser claramente definidas e documentadas (ex: em Requisitos Funcionais ou na documentação do Domínio).
* O uso de bibliotecas ou frameworks que auxiliem na implementação do RBAC (ex: Casbin, ou funcionalidades nativas de frameworks web) deve ser considerado para agilizar o desenvolvimento e garantir a robustez.
