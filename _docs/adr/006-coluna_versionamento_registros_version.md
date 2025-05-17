# ADR-006: Adoção de Coluna de Versionamento de Registros (`version`)

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Além das colunas de auditoria temporal `created_at` (ADR-002) e `updated_at` (ADR-002), que registram _quando_ um registro foi criado e sua última modificação, respectivamente, há a necessidade de rastrear _quantas vezes_ um registro foi alterado. Esta informação é valiosa para:

- Auditoria básica, indicando a frequência de modificações.
- Implementação de mecanismos de controle de concorrência, como o Bloqueio Otimista (Optimistic Locking).
- Entendimento da volatilidade e do ciclo de vida dos dados.

A pergunta que estamos tentando responder é: Como podemos manter uma contagem simples e incremental do número de modificações para cada registro em nossas entidades?

O escopo desta decisão abrange a maioria das entidades persistidas no banco de dados da aplicação RedToGreen que são sujeitas a operações de atualização e onde o rastreamento do número de versões é considerado útil.

## Decisão

Adotaremos a inclusão de uma coluna chamada `version` em todas as tabelas relevantes do banco de dados para rastrear o número de modificações de um registro.

1.  **Definição da Coluna**:

    - Nome da Coluna: `version`
    - Tipo de Dado: `INTEGER` (ou `BIGINT` para cenários com um número extremamente alto de atualizações por registro), não nulo.

2.  **Comportamento**:

    - **Inicialização:** Quando um novo registro é criado, o valor da coluna `version` será definido como `1`.
    - **Incrementação:** A cada vez que um registro existente é modificado (operação de `UPDATE`), o valor da coluna `version` será incrementado em 1.

3.  **Responsabilidade pela Atualização**:
    - A lógica de inicialização (para `1` na criação) e incrementação (em cada atualização) da coluna `version` será gerenciada pela camada de aplicação (ex: dentro dos repositórios ou serviços, antes da persistência da alteração). Esta abordagem oferece flexibilidade e controle explícito sobre o versionamento.

## Alternativas Consideradas (Opcional)

- **Não ter uma coluna de versionamento explícita:**

  - Descrição: Confiar apenas nas colunas `created_at` e `updated_at` ou em logs para inferir atividade de modificação.
  - Motivo da Rejeição: Não fornece uma contagem direta e simples do número de alterações, que é o requisito principal aqui. `updated_at` apenas indica o tempo da última mudança, não a frequência.

- **Versionamento completo de dados (ex: tabelas de histórico/auditoria dedicadas):**

  - Descrição: Manter um histórico completo de cada estado anterior do registro.
  - Motivo da Rejeição (para esta decisão específica): Esta é uma solução significativamente mais complexa, com maior overhead de armazenamento e processamento. A coluna `version` é uma forma leve de versionamento, focada unicamente na contagem de modificações e como base para otimistic locking, não no conteúdo de cada versão. Uma solução de auditoria completa é complementar e pode ser abordada separadamente se necessário.

- **Uso de Triggers no Banco de Dados para incrementação:**
  - Descrição: Delegar a lógica de incrementação da versão para triggers no SGBD.
  - Motivo da Rejeição (como decisão primária): Embora possível, gerenciar a lógica na aplicação oferece mais portabilidade entre diferentes SGBDs e mantém a lógica de negócio (mesmo que simples como esta) mais centralizada e visível no código da aplicação. No entanto, o uso de triggers pode ser reconsiderado se a consistência na aplicação se mostrar um desafio.

## Consequências

**Positivas:**

- **Auditoria Simplificada de Modificações:** Fornece uma maneira direta de determinar quantas vezes um registro específico foi alterado desde sua criação.
- **Base para Bloqueio Otimista (Optimistic Locking):** A coluna `version` é um mecanismo comum e eficaz para implementar o bloqueio otimista, ajudando a prevenir atualizações perdidas em ambientes concorrentes (ex: `UPDATE ... WHERE id = ? AND version = ?;`).
- **Indicador de Atividade do Registro:** Pode servir como um indicador da "volatilidade" ou frequência de alteração de determinados registros.
- **Implementação Relativamente Simples:** A lógica de inicialização e incrementação na camada de aplicação é direta.

**Negativas / Trade-offs:**

- **Overhead de Escrita Adicional:** Cada operação de `UPDATE` em um registro também exigirá a atualização da coluna `version`, adicionando uma pequena sobrecarga à escrita.
- **Gerenciamento pela Aplicação:** A consistência do valor da `version` depende da correta implementação da lógica de incrementação em todas as rotinas de atualização da aplicação. Falhas em aplicar essa lógica podem levar a valores de versão incorretos.
- **Não Fornece Histórico de "O Quê Mudou":** A coluna `version` apenas informa _quantas_ vezes o registro foi alterado, mas não fornece detalhes sobre _quais_ campos foram modificados em cada versão ou os valores anteriores. Para isso, seria necessária uma solução de auditoria mais completa.
- **Potencial para Valores Grandes:** Em registros que são atualizados com extrema frequência, o valor da versão pode se tornar muito grande, embora para a maioria dos casos de uso um `INTEGER` padrão seja suficiente.

**(Opcional) Notas Adicionais:**

- A atualização da coluna `version` deve ser realizada como parte da mesma transação da atualização do registro principal para garantir atomicidade e consistência.
- Se a estratégia de Bloqueio Otimista for formalmente adotada utilizando este campo, isso deverá ser documentado (possivelmente em um ADR futuro ou como parte das convenções de desenvolvimento).
- É importante que todas as partes do código da aplicação que modificam entidades versionadas sigam consistentemente a regra de incrementação da versão.
