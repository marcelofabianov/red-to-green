# ADR-004: Estratégia de Arquivamento de Registros para Indisponibilidade em Novos Usos

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Em sistemas de gestão como o RedToGreen, os usuários frequentemente configuram entidades de apoio (como Categorias, Tags, Contas Bancárias) que são usadas para classificar ou vincular transações. Com o tempo, algumas dessas entidades podem se tornar obsoletas ou não mais desejadas para _novos_ registros, mas ainda precisam estar presentes e visíveis em registros históricos e para fins de gerenciamento.

Por exemplo, um usuário pode querer parar de usar uma determinada Categoria para novas despesas, mas todas as despesas antigas já categorizadas com ela devem manter essa informação. Simplesmente "excluir" (mesmo com soft-delete, conforme ADR-001) não é ideal, pois o soft-delete geralmente implica que o registro não é mais visível na maioria das consultas. Precisamos de uma forma de marcar um registro como "inativo para novos usos" mas "ativo para visualização e histórico".

A pergunta que estamos tentando responder é: Como podemos permitir que o usuário "aposente" certos registros (ex: Categorias), impedindo seu uso em novas transações, mas mantendo sua visibilidade para consulta, gerenciamento e em dados históricos?

O escopo desta decisão abrange entidades configuráveis pelo usuário que são frequentemente vinculadas a outros registros (ex: Categorias, Subcategorias, Tags, e potencialmente Contas Bancárias ou Carteiras no futuro), onde o ciclo de vida do item pode incluir um estado de "arquivado" ou "inativo para novos usos".

## Decisão

Adotaremos uma estratégia de **arquivamento lógico** para entidades selecionadas, que permitirá que sejam marcadas como indisponíveis para novos relacionamentos, mas permaneçam visíveis para gerenciamento e em vínculos existentes.

Esta estratégia será implementada da seguinte forma:

1.  As tabelas das entidades relevantes (ex: Categorias, Tags) incluirão uma coluna chamada `archived_at`.
2.  A coluna `archived_at` será do tipo `TIMESTAMP` (ou `DATETIME` equivalente) e permitirá valores nulos.
3.  Um valor `NULL` na coluna `archived_at` indica que o registro está "disponível" (não arquivado) e pode ser selecionado/vinculado em novas operações.
4.  Um valor de data/hora (timestamp) na coluna `archived_at` indica que o registro foi "arquivado" no momento especificado.
5.  **Comportamento de Registros Arquivados:**
    - **Visibilidade em Gerenciamento:** Registros arquivados (`archived_at IS NOT NULL`) continuarão visíveis nas telas de listagem e gerenciamento dessas entidades (ex: lista de todas as categorias). A interface do usuário deverá indicar visualmente o status de arquivamento (ex: ícone, cor de fundo, texto informativo).
    - **Visibilidade em Vínculos Existentes:** Transações ou outros registros que já foram vinculados a um item antes de seu arquivamento continuarão a exibir esse vínculo normalmente.
    - **Disponibilidade para Novos Vínculos:** Em formulários de criação ou edição onde o usuário precisa selecionar um item de uma dessas entidades (ex: selecionar uma Categoria ao criar uma nova Transação), apenas os itens com `archived_at IS NULL` (não arquivados) serão apresentados como opções selecionáveis.

## Alternativas Consideradas (Opcional)

- **Utilizar a coluna `deleted_at` (soft-delete) com semântica diferente:**

  - Descrição: Sobrecarga da flag de soft-delete (ADR-001) para significar "arquivado" para certas entidades.
  - Motivo da Rejeição: O soft-delete, conforme definido no ADR-001, tem a intenção de tornar o registro indisponível e invisível para a maioria das operações, simulando uma exclusão. Misturar essa semântica com a de "arquivamento" (onde o item ainda é visível e relevante para o histórico) levaria a uma lógica confusa e a uma violação do significado estabelecido para `deleted_at`.

- **Flag booleana `is_archived`:**

  - Descrição: Usar uma coluna do tipo `BOOLEAN` (ex: `is_archived`) para indicar o status de arquivamento.
  - Motivo da Rejeição: Embora mais simples, uma coluna `TIMESTAMP` como `archived_at` oferece a vantagem adicional de registrar _quando_ o item foi arquivado. Essa informação pode ser útil para auditoria, para entender o contexto do arquivamento ou para futuras funcionalidades (ex: relatórios de itens arquivados em um determinado período). O custo adicional de um timestamp versus um booleano é considerado marginal em face desse benefício.

- **Sem marcador no banco, apenas lógica de UI:**
  - Descrição: Gerenciar o "arquivamento" apenas como um filtro visual na interface do usuário, sem um status persistido no banco.
  - Motivo da Rejeição: Insuficiente para garantir a integridade da regra de "não usar para novos relacionamentos" a nível de backend e dificultaria a aplicação consistente da regra em diferentes partes do sistema ou APIs.

## Consequências

**Positivas:**

- **Clareza Semântica:** Distingue claramente o estado de "arquivado" (ainda visível, utilizável em histórico, mas não para novos vínculos) do estado de "deletado" (efetivamente removido da visão do usuário).
- **Manutenção da Integridade Histórica:** Permite que os usuários "aposentem" itens de configuração sem afetar a integridade ou a visualização de dados históricos que dependem desses itens.
- **Melhoria da Experiência do Usuário:** Reduz a desordem em listas de seleção para novas transações, apresentando apenas opções relevantes e ativas, ao mesmo tempo que preserva o acesso completo aos dados para gerenciamento.
- **Rastreabilidade do Arquivamento:** O uso de um timestamp para `archived_at` permite saber quando o status de arquivamento foi aplicado.
- **Controle Granular:** Oferece ao usuário um controle mais fino sobre o ciclo de vida de suas entidades de configuração.

**Negativas / Trade-offs:**

- **Complexidade Adicional na Lógica de Consulta:** A aplicação precisará implementar lógicas de consulta diferenciadas:
  - Para telas de gerenciamento: buscar todos os registros, incluindo os arquivados.
  - Para seletores em formulários de novos vínculos: buscar apenas registros com `archived_at IS NULL`.
- **Coluna Adicional:** Introduz uma nova coluna (`archived_at`) nas tabelas das entidades relevantes, com um pequeno impacto no armazenamento.
- **Considerações na Interface do Usuário (UI):** Requer que a UI diferencie visualmente os registros arquivados dos não arquivados nas telas de gerenciamento e forneça uma ação clara para o usuário arquivar/desarquivar itens.
- **Potencial Confusão para o Usuário:** A diferença entre "arquivar" e "excluir" deve ser claramente comunicada ao usuário através da UI e, possivelmente, de mensagens de ajuda ou documentação, para evitar mal-entendidos.

**(Opcional) Notas Adicionais:**

- A interface do usuário deve permitir que o usuário possa "arquivar" um item e também "desarquivar" um item (ou seja, setar `archived_at` de volta para `NULL`).
- A criação de índices na coluna `archived_at` deve ser considerada se esta for frequentemente utilizada em cláusulas `WHERE` para filtrar registros ativos.
