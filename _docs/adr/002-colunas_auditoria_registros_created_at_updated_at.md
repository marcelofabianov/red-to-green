# ADR-002: Adoção de Colunas de Auditoria de Registros (`created_at`, `updated_at`)

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Para a aplicação RedToGreen, é essencial ter a capacidade de rastrear quando cada registro foi criado e quando foi a última vez que ele sofreu alguma alteração. Essas informações são cruciais para diversos fins, incluindo:

- Depuração de problemas e análise de causa raiz.
- Auditoria básica do ciclo de vida dos dados.
- Ordenação de registros por data de criação ou modificação.
- Implementação de lógicas de negócio que dependem do tempo (ex: processamento de registros recentes).
- Entendimento geral da atividade e evolução dos dados no sistema.

A pergunta que estamos tentando responder é: Qual é a maneira padronizada e eficiente de registrar os momentos de criação e última atualização dos nossos dados?

O escopo desta decisão abrange a maioria das entidades persistidas no banco de dados da aplicação RedToGreen, especialmente aquelas que representam dados de negócio ou cujo ciclo de vida é relevante para o sistema ou para os usuários.

## Decisão

Adotaremos a inclusão de duas colunas de auditoria temporal em todas as tabelas relevantes do banco de dados:

1.  **`created_at`**:

    - Tipo de Dado: `TIMESTAMP` (ou `DATETIME` equivalente, dependendo do SGBD), não nulo.
    - Propósito: Armazenar a data e hora exatas em que o registro foi inserido pela primeira vez no banco de dados.
    - Comportamento: Este valor deve ser definido no momento da criação do registro e não deve ser alterado subsequentemente.

2.  **`updated_at`**:
    - Tipo de Dado: `TIMESTAMP` (ou `DATETIME` equivalente, dependendo do SGBD), não nulo.
    - Propósito: Armazenar a data e hora da última modificação efetuada no registro.
    - Comportamento: Este valor deve ser definido no momento da criação do registro (geralmente com o mesmo valor de `created_at`) e deve ser atualizado automaticamente para o momento atual sempre que qualquer campo do registro for modificado.

A responsabilidade pelo preenchimento e atualização automática dessas colunas deve ser, preferencialmente, delegada ao Sistema de Gerenciamento de Banco de Dados (SGBD) através de mecanismos como valores padrão (`DEFAULT CURRENT_TIMESTAMP`) e triggers (`ON UPDATE CURRENT_TIMESTAMP`), se suportados e adequados para o SGBD escolhido. Alternativamente, essa lógica será garantida pela camada de aplicação ou pelo Object-Relational Mapper (ORM) de forma consistente.

## Alternativas Consideradas (Opcional)

- **Não utilizar colunas de auditoria dedicadas:**

  - Descrição: Confiar exclusivamente em logs da aplicação ou do SGBD para inferir informações de criação/atualização.
  - Motivo da Rejeição: Menos direto para consultas, dificulta a ordenação e a lógica de negócio baseada nessas datas, e depende da configuração, disponibilidade e política de retenção de logs externos, que podem não ser consistentes ou facilmente acessíveis.

- **Solução de auditoria completa e versionamento de registros (ex: tabelas de log/histórico dedicadas):**
  - Descrição: Implementar um sistema que rastreie cada alteração em cada campo, os valores anteriores e quem fez a alteração.
  - Motivo da Rejeição (para esta decisão específica): Embora seja uma capacidade valiosa para auditoria detalhada, é significativamente mais complexa de implementar e gerenciar do que as simples colunas `created_at`/`updated_at`. As colunas de auditoria temporal oferecem um primeiro nível de rastreabilidade essencial com baixo custo de implementação. Uma solução de auditoria completa pode ser considerada como um ADR futuro e complementar, se necessário.

## Consequências

**Positivas:**

- **Rastreabilidade Melhorada:** Facilidade em identificar quando um registro foi criado e quando ocorreu sua última modificação.
- **Suporte à Depuração:** Auxilia na investigação de problemas ao fornecer um contexto temporal para os dados.
- **Capacidade de Ordenação Temporal:** Permite que os dados sejam facilmente ordenados por data de criação ou última atualização.
- **Padronização:** Garante uma forma consistente e uniforme de acessar essas informações de auditoria em todas as entidades relevantes.
- **Base para Funcionalidades Futuras:** Pode ser utilizada como base para otimizações de cache, processos de sincronização de dados, ou interfaces que mostram atividade recente.
- **Baixo Custo de Implementação:** Relativamente simples de adicionar e gerenciar, especialmente se o SGBD suportar o preenchimento automático.

**Negativas / Trade-offs:**

- **Overhead de Armazenamento:** Adiciona duas colunas a cada tabela, resultando em um pequeno aumento no uso de espaço em disco.
- **Overhead de Escrita (I/O):** Há uma pequena sobrecarga de escrita para definir/atualizar a coluna `updated_at` em cada operação de inserção e atualização de registro.
- **Necessidade de Consistência na Implementação:** Se não for totalmente gerenciado pelo SGBD, requer disciplina e mecanismos na camada de aplicação (ex: ORM, interceptadores) para garantir que as colunas sejam sempre preenchidas e atualizadas corretamente.
- **Auditoria Limitada:** Estas colunas não fornecem um histórico completo de todas as alterações (quem alterou, quais campos, valores anteriores), apenas a data/hora da criação e da última modificação.

**(Opcional) Notas Adicionais:**

- Recomenda-se a criação de índices nas colunas `created_at` e `updated_at` se elas forem frequentemente usadas em cláusulas `WHERE` ou `ORDER BY`.
- A precisão do timestamp (segundos, milissegundos) dependerá do SGBD e da necessidade da aplicação.
