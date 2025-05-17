# ADR-007: Adoção do Algoritmo Argon2 para Hashing de Senhas

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A proteção das senhas dos usuários é um pilar fundamental da segurança de qualquer aplicação, especialmente uma que lida com dados financeiros como o RedToGreen. Senhas armazenadas incorretamente (ex: em texto plano ou com algoritmos de hashing fracos/rápidos) são alvos fáceis para diversos tipos de ataques, como força bruta, rainbow tables e vazamentos de dados. É imperativo utilizar um algoritmo de hashing de senhas que seja forte, moderno e resistente a ataques que utilizam hardware especializado (GPUs, ASICs).

A pergunta que estamos tentando responder é: Qual algoritmo de hashing de senhas devemos implementar no RedToGreen para garantir o mais alto nível de segurança para as credenciais dos nossos usuários?

O escopo desta decisão abrange o mecanismo de hashing utilizado durante o registro de novos usuários e a verificação de senhas durante o processo de login.

## Decisão

Adotaremos o algoritmo **Argon2** como o padrão para o hashing de senhas no sistema RedToGreen.

1.  **Variante:** Utilizaremos a variante **Argon2id**. Esta variante é um híbrido que combina a resistência do Argon2d a ataques de cracking por GPU (com dependência de dados no acesso à memória) e a resistência do Argon2i a ataques de side-channel e time-memory tradeoff.
2.  **Parâmetros de Configuração:** Os parâmetros de configuração do Argon2id (custo de memória - `m`, número de iterações/passes - `t`, e grau de paralelismo - `p`) serão definidos com base nas recomendações atuais de segurança (ex: OWASP Cheat Sheet Series) e ajustados para o ambiente de produção do RedToGreen, buscando um equilíbrio entre a robustez do hash e o tempo de processamento aceitável para operações de login e registro. Esses parâmetros serão configuráveis na aplicação.
3.  **Salt:** Para cada senha de usuário, um salt criptograficamente seguro e único será gerado aleatoriamente. Este salt será armazenado junto com o hash da senha resultante.
4.  **Armazenamento do Hash:** O formato de armazenamento do hash incluirá idealmente o identificador do algoritmo, a versão, os parâmetros utilizados, o salt e o hash em si (ex: no formato modular crypt format).

## Alternativas Consideradas (Opcional)

- **bcrypt:**

  - Descrição: Um algoritmo de hashing de senhas amplamente adotado, baseado na cifra Blowfish, conhecido por sua resistência devido ao fator de trabalho (custo) configurável que o torna lento.
  - Motivo da Rejeição: Embora seja uma alternativa forte e comprovada, Argon2 foi o vencedor da Password Hashing Competition (2015) e é projetado para oferecer maior resistência contra ataques que utilizam hardware customizado (GPUs, FPGAs, ASICs) devido ao seu design que faz uso intensivo de memória e oferece maior configurabilidade de resistência.

- **scrypt:**

  - Descrição: Um algoritmo de hashing de senhas que também é projetado para ser caro em termos de memória, dificultando ataques com hardware especializado.
  - Motivo da Rejeição: Semelhante ao bcrypt, é uma boa alternativa. No entanto, Argon2 oferece mais dimensões de ajuste de resistência (memória, tempo, paralelismo) e é geralmente considerado o estado da arte.

- **PBKDF2 (Password-Based Key Derivation Function 2):**

  - Descrição: Um padrão de derivação de chave que pode ser usado para hashing de senhas, aplicando uma função pseudoaleatória (como HMAC-SHA256) repetidamente.
  - Motivo da Rejeição: Embora ainda seja considerado seguro se configurado com um número muito alto de iterações, é geralmente menos resistente a ataques com hardware especializado em comparação com algoritmos mais modernos como bcrypt, scrypt e, especialmente, Argon2.

- **Algoritmos de Hashing de Propósito Geral (ex: SHA-256, SHA-512) usados diretamente:**
  - Descrição: Utilizar funções de hash rápidas, mesmo com salt.
  - Motivo da Rejeição: Estes algoritmos são projetados para serem rápidos, o que os torna inadequados para hashing de senhas, pois são vulneráveis a ataques de força bruta de alta velocidade, mesmo com a adição de um salt. Eles não possuem o conceito de "fator de trabalho" ou "custo de memória" ajustável.

## Consequências

**Positivas:**

- **Segurança de Ponta:** Argon2 (especialmente Argon2id) é atualmente considerado um dos algoritmos de hashing de senhas mais seguros disponíveis, oferecendo forte proteção contra ataques de força bruta, ataques baseados em dicionário, rainbow tables e hardware especializado.
- **Resistência a Diversos Vetores de Ataque:** Argon2id é projetado para ser resistente a ataques de side-channel e ataques de time-memory tradeoff.
- **Configurabilidade e Escalabilidade da Segurança:** Os parâmetros de memória, tempo (iterações) e paralelismo podem ser ajustados para aumentar a força do hash à medida que o poder computacional evolui, permitindo manter um nível de segurança adequado ao longo do tempo.
- **Conformidade com Recomendações Atuais:** A escolha está alinhada com as recomendações de especialistas e organizações de segurança (como OWASP).
- **Proteção contra Rainbow Tables:** O uso obrigatório de um salt único por usuário impede a eficácia de ataques de rainbow table precomputados.

**Negativas / Trade-offs:**

- **Consumo de Recursos:** Sendo um algoritmo deliberadamente intensivo em recursos (especialmente memória e CPU), o processo de hashing de senhas com Argon2 consumirá mais recursos do servidor durante o registro e login em comparação com algoritmos mais antigos ou menos seguros. Isso precisa ser dimensionado adequadamente.
- **Tempo de Processamento:** O tempo para calcular o hash pode ser na ordem de centenas de milissegundos a segundos, dependendo da configuração dos parâmetros. Embora isso seja desejável para a segurança (dificultando ataques de força bruta), deve ser balanceado para não impactar negativamente a experiência do usuário.
- **Complexidade de Implementação e Configuração:** Requer o uso de bibliotecas criptográficas que suportem Argon2. A escolha e configuração correta dos parâmetros (memória, iterações, paralelismo) são cruciais e exigem entendimento.
- **Portabilidade e Disponibilidade de Bibliotecas:** É necessário garantir o uso de bibliotecas Argon2 seguras, bem mantidas e em conformidade com a especificação para as linguagens do projeto (Go, PHP, Node).

**(Opcional) Notas Adicionais:**

- Os parâmetros do Argon2id (`m`, `t`, `p`) devem ser escolhidos com base em benchmarks no ambiente de produção e nas recomendações de segurança atuais. Esses parâmetros devem ser configuráveis e, idealmente, armazenados junto com o hash para permitir a atualização futura dos parâmetros sem invalidar os hashes existentes (implementando uma estratégia de re-hashing gradual para os usuários no próximo login quando os parâmetros forem atualizados).
- É fundamental utilizar implementações de bibliotecas Argon2 de fontes confiáveis e que sejam mantidas ativamente.
