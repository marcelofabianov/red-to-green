# ADR-019: Estratégia de Testes Unitários em Go com a Biblioteca Testify

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Testes unitários são fundamentais para garantir a qualidade, robustez e manutenibilidade do código da aplicação RedToGreen. É necessário definir uma estratégia clara e consistente para a escrita de testes unitários, abrangendo a organização dos arquivos de teste, a criação de mocks para dependências, a forma de realizar asserções e a estruturação dos casos de teste.

A estratégia deve promover testes que sejam fáceis de escrever, ler, manter e que forneçam feedback claro sobre o comportamento do código. Deve também se alinhar com os princípios de desenvolvimento do projeto (ADR-000), como simplicidade, pragmatismo e utilização de ferramentas que melhorem a produtividade sem adicionar complexidade desnecessária.

A pergunta que estamos tentando responder é: Qual conjunto de ferramentas e práticas devemos adotar para a escrita de testes unitários em Go no projeto RedToGreen, visando clareza, produtividade e eficácia?

O escopo desta decisão abrange as convenções e bibliotecas a serem utilizadas para a escrita de testes unitários para os componentes do backend Go do RedToGreen.

## Decisão

Adotaremos as seguintes práticas e ferramentas para a escrita de testes unitários no RedToGreen, com a biblioteca `stretchr/testify` como principal ferramenta de apoio:

1.  **Pacote de Teste e Organização de Arquivos:**
    * Os arquivos de teste seguirão a convenção Go, nomeados com o sufixo `_test.go` (ex: `meu_arquivo.go` terá testes em `meu_arquivo_test.go`).
    * **Os testes unitários serão predominantemente escritos dentro do mesmo pacote do código que está sendo testado** (ex: `package meupacote`). Esta abordagem (teste de caixa branca) permite o teste de funções e métodos não exportados, facilitando o teste granular da lógica interna.
    * Testes para a API pública de um pacote (teste de caixa preta) podem, opcionalmente, residir em um pacote com sufixo `_test` (ex: `package meupacote_test`) se essa distinção for benéfica para testar o contrato do pacote como um cliente externo.
    * Dados de teste (fixtures, arquivos de entrada/saída) serão armazenados em subdiretórios `testdata` dentro do pacote correspondente.

2.  **Biblioteca de Asserções: `stretchr/testify/assert` e `stretchr/testify/require`:**
    * Utilizaremos os pacotes `assert` e `require` da biblioteca `github.com/stretchr/testify` para realizar asserções nos testes.
    * O pacote `assert` será usado para verificações onde múltiplas falhas podem ser reportadas em um único teste sem interromper sua execução imediatamente.
    * O pacote `require` será usado para verificações críticas onde o teste deve parar imediatamente em caso de falha (ex: erro em setup, condição prévia não atendida, asserções que tornam o restante do teste inválido).
    * Esta escolha visa aumentar a legibilidade dos testes, reduzir o boilerplate de código de verificação (`if/t.Errorf`) e fornecer mensagens de erro mais claras e informativas por padrão.

3.  **Criação de Mocks: `stretchr/testify/mock`:**
    * Utilizaremos o pacote `mock` da biblioteca `github.com/stretchr/testify` para criar mocks de dependências (que devem ser definidas como interfaces no código da aplicação).
    * Este pacote permite definir o comportamento esperado dos mocks (ex: argumentos esperados, valores de retorno, número de chamadas) e verificar se essas expectativas foram atendidas durante a execução do teste.
    * Ferramentas de geração de código para mocks baseadas em interfaces (ex: `vektra/mockery`, que pode gerar mocks compatíveis com o formato do `testify/mock`) podem ser consideradas para interfaces mais complexas, a fim de reduzir o boilerplate da escrita manual de mocks, mantendo a consistência com `testify/mock`.

4.  **Estrutura de Casos de Teste: Table-Driven Tests com `t.Run()`:**
    * Adotaremos amplamente o padrão de "table-driven tests" (testes orientados a tabela) para testar funções e métodos com múltiplos cenários de entrada/saída e diferentes condições de forma organizada.
    * Uma slice de structs (a "tabela") será definida, onde cada struct representa um caso de teste individual, contendo um nome descritivo (para o sub-teste), as entradas necessárias e os resultados esperados (valores e/ou erros).
    * Um loop `for _, tc := range testCases { ... }` iterará sobre os casos de teste. Cada caso será executado como um sub-teste nomeado utilizando `t.Run(tc.name, func(t *testing.T) { ... })`.
    * O método `t.Parallel()` será utilizado dentro dos sub-testes `t.Run()` sempre que os casos de teste forem independentes entre si e não compartilharem estado mutável, para permitir a execução paralela e acelerar a suíte de testes.

## Alternativas Consideradas

* **Uso Exclusivo do Pacote `testing` Nativo (sem a biblioteca `testify`):**
    * *Prós:* Sem dependências externas, total controle.
    * *Motivo da Rejeição:* Leva a código de asserção mais verboso e propenso a erros (`if/t.Errorf`), com mensagens de erro que precisam ser construídas manualmente para serem informativas. A escrita manual de mocks para todas as interfaces pode ser excessivamente trabalhosa. O ganho de produtividade, clareza e expressividade com `testify` justifica a adição desta dependência leve e amplamente adotada pela comunidade Go.

* **Outras Bibliotecas de Mocking (ex: `golang/mock`):**
    * *Prós:* `golang/mock` é uma ferramenta originada no Google e bastante poderosa.
    * *Motivo da Rejeição:* O pacote `stretchr/testify/mock` é frequentemente considerado mais simples de usar, com uma API mais intuitiva e alinhada com o estilo programático de `testify/assert`. A curva de aprendizado e a verbosidade do `golang/mock` (com seu `gomock.Controller` e a sintaxe `EXPECT()`) foram consideradas menos ideais em comparação com a abordagem mais direta do `testify/mock` para as necessidades do projeto.

* **Não Utilizar Table-Driven Tests Consistentemente:**
    * *Prós:* Pode parecer mais rápido escrever um teste para um único caso de forma isolada.
    * *Motivo da Rejeição:* Dificulta a adição e o gerenciamento de múltiplos cenários de teste, leva a código de teste mais repetitivo, menos organizado e potencialmente com menor cobertura de casos de borda. Table-driven tests são um padrão idiomático e eficaz em Go para melhorar a qualidade e a abrangência dos testes.

## Consequências

**Positivas:**
* **Testes Mais Legíveis, Concisos e Expressivos:** O uso de `testify/assert` e `testify/require` tornará as verificações mais claras, fáceis de entender e menos propensas a erros.
* **Mocks Gerenciáveis e Poderosos:** `testify/mock` oferece uma forma estruturada, testável e flexível de mockar dependências, facilitando o isolamento das unidades sob teste.
* **Cobertura Abrangente e Organizada de Cenários:** Table-driven tests com `t.Run()` facilitam a criação, manutenção e compreensão de um conjunto diversificado de casos de teste.
* **Melhor Feedback de Testes:** Sub-testes nomeados (`t.Run`) e mensagens de erro detalhadas fornecidas por `testify` ajudam a identificar rapidamente a causa e o contexto das falhas.
* **Aumento da Produtividade do Desenvolvedor:** Menos boilerplate na escrita de asserções e na definição de múltiplos casos de teste.
* **Alinhamento com Práticas Comuns e Recomendadas da Comunidade Go:** A biblioteca `testify` é amplamente utilizada e reconhecida como um padrão de fato para testes em Go.
* **Testes de Caixa Branca Facilitados:** A decisão de colocar os testes no mesmo pacote permite testar componentes internos de forma eficaz quando necessário.

**Negativas / Trade-offs:**
* **Adição de Dependência Externa (`stretchr/testify`):** Embora seja uma biblioteca popular, leve e bem mantida, representa uma dependência externa adicional ao projeto. Este trade-off é considerado aceitável devido aos significativos ganhos de produtividade e clareza.
* **Curva de Aprendizagem para `testify`:** Desenvolvedores novos à biblioteca `testify` precisarão se familiarizar com sua API e funcionalidades (embora seja geralmente considerada intuitiva e bem documentada).
* **Potencial para Mocks Excessivamente Complexos:** Se não usado com disciplina, o poder do `testify/mock` pode levar à criação de mocks muito intrincados que tornam os testes frágeis ou difíceis de entender. A preferência ainda é por definir interfaces simples e focadas para as dependências, resultando em mocks mais simples.

## (Opcional) Notas Adicionais
* A cobertura de código dos testes unitários será monitorada regularmente utilizando as ferramentas padrão do Go (`go test -cover` e `go tool cover`).
* Todos os novos desenvolvimentos de lógica de negócios e funcionalidades críticas devem ser acompanhados de testes unitários abrangentes que validem tanto os caminhos felizes quanto os casos de erro e de borda.
* Os testes devem ser mantidos limpos, focados em uma única responsabilidade (quando possível) e projetados para serem rápidos. Testes unitários lentos podem desencorajar sua execução frequente.
* O uso de `t.Helper()` será consistentemente aplicado em funções auxiliares de teste para garantir que as mensagens de falha dos asserts sejam reportadas corretamente na linha do código de teste que invocou a função auxiliar.
