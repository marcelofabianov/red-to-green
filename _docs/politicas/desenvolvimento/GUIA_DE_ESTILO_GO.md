# Guia de Estilo Go para o Projeto RedToGreen

**Última Atualização:** 2025-05-18

## 1. Introdução

Este documento define as convenções de estilo de código Go a serem seguidas no desenvolvimento do RedToGreen. O objetivo é promover código limpo, legível, consistente e manutenível, alinhado com os [Princípios Fundamentais (ADR-000)](../../adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md), especialmente os de "Qualidade e Robustez" e "Utilização da Biblioteca Padrão e Ferramental Nativo".

## 2. Formatação Automática

* **Ferramenta:** Todo o código Go **DEVE** ser formatado utilizando `goimports` (que inclui `gofmt` e também organiza os imports).
* **Aplicação:** Recomenda-se configurar o editor/IDE para aplicar `goimports` ao salvar. Esta verificação também **DEVE** fazer parte do pipeline de CI e, idealmente, de hooks de pre-commit.
    * _Relacionado ao ADR-000: Utilização da Biblioteca Padrão e Ferramental Nativo._

## 3. Nomenclatura

Seguir as convenções idiomáticas do Go:
* **Pacotes:** Nomes curtos, concisos, em minúsculas, uma única palavra (ex: `config`, `models`, `services`).
* **Variáveis, Funções, Tipos, Métodos:** `camelCase`.
* **Constantes:** `CamelCase` (preferencial) ou `ALL_CAPS` para constantes verdadeiramente imutáveis e globais.
* **Interfaces:** Sufixo `-er` quando apropriado (ex: `Reader`, `UserCreator`) ou nomeadas pelo que representam.
* **Exportado vs. Não Exportado:** Utilizar letra inicial maiúscula para membros exportados e minúscula para não exportados.
* **Nomes de Arquivo:** `snake_case.go`.
* **Nomes de Testes:** Seguir o padrão `TestNomeDaFuncao_ContextoOuCenario` (conforme ADR-019).
* **DTOs vs. Modelos de Domínio vs. Modelos de Banco:** Usar sufixos ou prefixos claros se necessário para distinguir (ex: `UserDTO`, `UserModel`, `UserDBEntity`), mas priorizar nomes que expressem o propósito no contexto em que são usados.

## 4. Comentários

* **Comentar o Porquê, Não o Quê:** O código deve ser autoexplicativo sobre "o quê" faz. Comentários devem explicar a intenção, decisões de design complexas ou o "porquê" de uma abordagem específica.
* **Documentação `godoc`:** Todos os pacotes, funções, tipos, constantes e variáveis exportadas **DEVEM** ter comentários `godoc` claros e informativos.
* **`// TODO:` e `// FIXME:`:** Usar para marcar trabalho pendente ou problemas conhecidos, idealmente com uma referência a uma issue ou nome/data.

## 5. Tratamento de Erros

* **Retorno Explícito:** Seguir o padrão Go de retornar erros como o último valor.
* **Embrulho de Erros (Wrapping):** Utilizar `fmt.Errorf("contexto da falha: %w", errOriginal)` para adicionar contexto ao propagar erros, permitindo o uso de `errors.Is` e `errors.As`.
* **Estrutura Padronizada de Erros:** Utilizar a estrutura `MessageError` definida no [ADR-005](./../../adr/005-estrutura_padronizada_erros_com_contexto.md) para erros que precisam de códigos específicos da aplicação ou contexto adicional para APIs ou logging.
* **Não Ignorar Erros:** Erros retornados por funções **DEVEM** ser checados, a menos que haja uma justificativa explícita e documentada para ignorá-los (raro).

## 6. Linters e Análise Estática

* **Ferramenta Principal:** `golangci-lint` **DEVE** ser utilizado.
* **Configuração:** Um arquivo `.golangci.yml` (ou `.golangci.toml`) será mantido na raiz do projeto, definindo os linters ativos e suas configurações. (A definição inicial destes linters pode ser um ADR ou uma decisão documentada aqui).
* **Linters Recomendados (Conjunto Inicial):** `errcheck`, `govet`, `staticcheck`, `unused`, `ineffassign`, `gofmt`, `goimports`, `misspell`, `stylecheck`, `unconvert`, `prealloc`, `typecheck`.
* **Integração:** `golangci-lint` **DEVE** ser integrado ao pipeline de CI/CD e, idealmente, a hooks de pre-commit. Falhas no linter devem impedir o merge/deploy.
    * _Relacionado ao ADR-000: Qualidade e Robustez._

## 7. Organização de Código

* **Funções Curtas:** Manter funções e métodos com responsabilidade única e o menor tamanho possível.
* **Evitar Aninhamento Excessivo:** Usar "guard clauses" e "early returns" para reduzir o aninhamento de condicionais e loops.
* **Imports Agrupados:** `goimports` cuida disso (padrão, stdlib, terceiros, locais).
* **Variáveis Globais:** Evitar ao máximo. Se necessárias, devem ser justificadas e seu escopo minimizado.

## 8. Concorrência

* Seguir as práticas idiomáticas de Go para concorrência (goroutines e channels).
* Utilizar `context.Context` para cancelamento e timeouts em operações concorrentes e I/O.
* Proteger acesso a dados compartilhados com mutexes ou outras primitivas de sincronização quando necessário.
* Cuidado com deadlocks e race conditions (usar o race detector do Go: `go test -race`).

Este guia será revisado e atualizado conforme necessário pela equipe.
