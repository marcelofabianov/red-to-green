# ADR-003: Adoção de UUID v7 como Estratégia Principal de Identificadores

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A escolha de uma estratégia de geração de identificadores (IDs) é uma decisão fundamental na arquitetura de qualquer sistema. Para o RedToGreen, precisamos de uma abordagem que garanta unicidade, seja performática, segura. Um estudo prévio analisou diversas opções, incluindo inteiros auto-incrementais, UUID v4, ULID, NanoID e abordagens híbridas.

Os principais requisitos considerados foram:

- **Unicidade:** Garantir que os IDs sejam únicos globalmente ou, no mínimo, dentro do escopo da aplicação, mesmo considerando futuras expansões ou integrações.
- **Performance:** Impacto no desempenho do banco de dados, especialmente em operações de escrita (inserção) e leitura (indexação, joins).
- **Segurança:** Evitar que IDs sejam facilmente adivinháveis ou enumeráveis, caso expostos.
- **Escalabilidade e Sistemas Distribuídos:** Capacidade de gerar IDs de forma descentralizada.
- **Ordenação:** Desejável que os IDs tenham alguma propriedade de ordenação temporal para otimizar a localidade dos dados e performance de índices.

A pergunta que estamos tentando responder é: Qual estratégia de geração de IDs deve ser adotada como padrão para as chaves primárias e identificadores de recursos no sistema RedToGreen?

O escopo desta decisão abrange as chaves primárias de todas as entidades persistidas no banco de dados e, potencialmente, os identificadores de recursos expostos via APIs.

## Decisão

Adotaremos **UUID versão 7 (UUID v7)** como a estratégia padrão para a geração de identificadores únicos para as entidades do sistema RedToGreen.

Principais aspectos da decisão:

1.  **Chaves Primárias:** UUIDs v7 serão utilizados como chaves primárias para as tabelas no banco de dados.
2.  **Geração:** Os UUIDs v7 serão gerados pela camada de aplicação antes da persistência dos dados, utilizando bibliotecas confiáveis e compatíveis com a especificação IETF RFC 9562 (anteriormente draft-ietf-uuidrev-rfc4122bis).
3.  **Armazenamento:** O formato de armazenamento no banco de dados será otimizado para performance e espaço (ex: `BINARY(16)` ou tipo `UUID` nativo, se disponível no SGBD), em vez de representações textuais longas como `CHAR(36)`, sempre que possível.

## Alternativas Consideradas (Opcional)

- **Inteiros Auto-incrementais:**

  - Descartados devido à previsibilidade (risco de segurança se expostos), dificuldades de geração em ambientes distribuídos ou offline, e potenciais conflitos em cenários de merge de bancos de dados.

- **UUID v4 (Aleatório):**

  - Considerado pela sua ampla adoção e garantia de unicidade.
  - Descartado como estratégia principal devido à sua natureza puramente aleatória, que pode levar à fragmentação de índices B-tree e, consequentemente, a uma performance de escrita e leitura inferior em comparação com UUIDs ordenáveis, especialmente sob alta carga.

- **ULID (Universally Unique Lexicographically Sortable Identifier):**

  - Uma alternativa forte, oferecendo unicidade, ordenação lexicográfica e representação mais compacta e URL-friendly que UUIDs tradicionais.
  - UUID v7 foi preferido por ser um padrão emergente da IETF com potencial para maior adoção e suporte nativo em ferramentas e SGBDs no futuro, oferecendo benefícios de ordenação temporal e performance semelhantes.

- **NanoID / ShortUUID:**

  - Descartados para chaves primárias devido à maior probabilidade de colisão com o crescimento do volume de dados (para tamanhos curtos) e por não serem inerentemente ordenáveis por tempo de criação, o que é um benefício do UUID v7.

- **Abordagem Híbrida (PK Inteiro Auto-incremental + `public_id` UUID/ULID):**
  - Reconhecida como uma prática robusta que combina a performance interna de inteiros com a segurança de IDs não sequenciais para exposição externa.
  - Para este projeto, optou-se por UUID v7 como chave primária única para simplificar o modelo de dados (evitando a gestão de dois IDs por entidade) e aproveitar diretamente os benefícios de performance de escrita e unicidade que o UUID v7 oferece, sendo também adequado para exposição externa.

## Consequências

**Positivas:**

- **Unicidade Global:** Praticamente elimina o risco de colisões de ID, mesmo em cenários distribuídos ou federados, permitindo a geração de IDs de forma descentralizada.
- **Excelente Performance de Indexação e Escrita:** A natureza monotonicamente crescente (baseada em timestamp) dos UUIDs v7 resulta em melhor localidade de dados nos índices B-tree do banco de dados, similar a inteiros sequenciais. Isso reduz a fragmentação do índice e melhora significativamente a performance de inserção e de consultas que utilizam range scans.
- **Não Adivinháveis:** Como UUIDs, não são sequenciais no sentido tradicional, dificultando a enumeração de recursos ou a estimativa do volume de dados, o que é positivo para a segurança.
- **Ordenáveis (Aproximadamente por Tempo):** A componente de timestamp UNIX de 48 bits no início do UUID v7 permite que os IDs sejam, em geral, ordenados cronologicamente, o que pode ser útil para queries e análise de dados.
- **Padrão IETF:** Sendo baseado em um padrão proposto pela IETF (RFC 9562), há uma expectativa de crescente adoção e suporte em bibliotecas, frameworks e SGBDs.
- **Geração na Aplicação:** Permite que os IDs sejam gerados pela aplicação antes mesmo da comunicação com o banco de dados, o que pode ser útil em certos padrões arquiteturais (ex: Domain-Driven Design).

**Negativas / Trade-offs:**

- **Tamanho de Armazenamento:** UUIDs (128 bits / 16 bytes) são maiores que inteiros (ex: 64 bits / 8 bytes para `BIGINT`), resultando em chaves primárias e índices maiores, o que consome mais espaço em disco e pode impactar ligeiramente a performance de leitura de índices muito grandes.
- **Legibilidade Humana:** São significativamente menos legíveis e mais difíceis de manusear por humanos (em logs, depuração manual, queries ad-hoc) em comparação com inteiros.
- **Suporte e Implementação:** Embora o padrão esteja definido, é crucial usar bibliotecas que implementem corretamente a especificação UUID v7. O suporte nativo em todos os SGBDs ainda pode estar em evolução.
- **Complexidade Relativa de Geração:** A lógica de geração é mais complexa que a de um simples contador auto-incremental.
- **Não Estritamente Sequencial/Monotônico em Todos os Cenários:** A monotonicidade é garantida dentro de um mesmo gerador e clock, mas colisões de timestamp entre geradores distintos podem ocorrer (embora a parte aleatória vise resolver isso). A ordenação é primariamente por tempo, mas não é tão simples quanto um inteiro sequencial.

**(Opcional) Notas Adicionais:**

- É fundamental selecionar e utilizar bibliotecas robustas e bem testadas para a geração de UUID v7 nas linguagens de programação do projeto (Go, PHP, Node).
- A escolha do tipo de dado para armazenar UUIDs no SGBD (ex: `UUID` nativo, `BINARY(16)`) deve ser feita considerando a otimização de armazenamento e performance de consulta. Evitar o armazenamento como strings (`CHAR(36)`) se alternativas mais eficientes estiverem disponíveis.
- Considerar a criação de funções no banco de dados para converter entre a representação binária e textual do UUID para facilitar a depuração, se necessário.
