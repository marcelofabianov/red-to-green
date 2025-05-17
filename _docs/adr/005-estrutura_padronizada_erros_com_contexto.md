# ADR-005: Adoção de Estrutura Padronizada para Erros com Contexto Adicional

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

Um tratamento de erros robusto e informativo é crucial para a manutenibilidade, depuração e monitoramento de qualquer aplicação. Os mecanismos de erro padrão do Go (`errors.New`, `fmt.Errorf`) são flexíveis, mas por si só não impõem uma estrutura para informações adicionais como códigos de erro específicos da aplicação, mensagens amigáveis ou contexto dinâmico sobre a falha.

Para o RedToGreen, que envolve lógica de negócios e interações com dados sensíveis, é importante que os erros:

- Sejam consistentes em toda a aplicação.
- Forneçam informações suficientes para que os desenvolvedores identifiquem rapidamente a causa raiz.
- Permitam a comunicação de erros de forma clara para APIs ou outras camadas do sistema (ex: frontend).
- Facilitem a categorização e o rastreamento em sistemas de logging e monitoramento.

A pergunta que estamos tentando responder é: Como podemos padronizar a criação e o tratamento de erros na aplicação para fornecer mais contexto, facilitar a depuração e permitir uma melhor comunicação de falhas?

O escopo desta decisão abrange a forma como os erros são criados, propagados e inspecionados em toda a base de código Go do backend do RedToGreen.

## Decisão

Adotaremos uma estrutura customizada e padronizada para erros, encapsulada em uma struct chamada `MessageError`. Esta estrutura permitirá adicionar um código de erro específico da aplicação, uma mensagem descritiva e um contexto dinâmico a qualquer erro.

1.  **Estrutura `MessageError`**:
    Será definida uma struct `MessageError` (ex: no pacote `pkg/msg` ou similar) com os seguintes campos:

    ```go
    // pkg/msg/err.go (ou similar)
    package msg

    type MessageError struct {
        Err     error            // Erro base original (para errors.Is/As)
        Message string           // Mensagem descritiva do erro
        Code    string           // Código de erro específico da aplicação (ex: "USER_NOT_FOUND")
        Context map[string]any   // Dados contextuais adicionais (ex: IDs, parâmetros)
    }
    ```

2.  **Implementação da Interface `error`**:
    A struct `MessageError` implementará a interface `error` padrão do Go:

    ```go
    func (e *MessageError) Error() string {
        if e.Message != "" {
            return e.Message
        }
        if e.Err != nil {
            return e.Err.Error()
        }
        return "Unknown error" // Fallback, though should ideally always have a message or base error
    }

    func (e *MessageError) Unwrap() error {
        return e.Err
    }
    ```

3.  **Função Construtora e Métodos Auxiliares**:
    Uma função construtora `NewMessageError` e um método `WithContext` serão fornecidos:

    ```go
    func NewMessageError(err error, message string, code string, context map[string]any) *MessageError {
        return &MessageError{
            Err:     err,
            Message: message,
            Code:    code,
            Context: context,
        }
    }

    func (e *MessageError) WithContext(key string, value any) *MessageError {
        if e.Context == nil {
            e.Context = make(map[string]any)
        }
        e.Context[key] = value
        return e
    }
    ```

4.  **Criação de Erros Específicos por Módulo**:
    Cada módulo/pacote da aplicação terá um arquivo `msg.go` (ou similar) para centralizar a definição de suas mensagens e erros específicos, utilizando a estrutura `MessageError`.
    Exemplo para um erro "usuário não encontrado":

    ```go
    // pkg/user/msg.go (ou similar)
    // import "errors"
    // import "fmt"
    // import "path/to/pkg/msg"
    // import "path/to/pkg/vo" // Value Object package

    var errUserNotFoundBase = errors.New("user not found") // Sentinel error

    func UserNotFoundError(userID vo.ID) *msg.MessageError {
        return msg.NewMessageError(
            errUserNotFoundBase,
            fmt.Sprintf("user with ID %s not found", userID.String()),
            "USER_NOT_FOUND",
            map[string]any{"user_id": userID.String()},
        )
    }
    ```

    A variável `errUserNotFoundBase` serve como um "sentinel error" que pode ser usado com `errors.Is` para verificar o tipo de erro base.

## Alternativas Consideradas (Opcional)

- **Uso exclusivo de `errors.New` e `fmt.Errorf` com `errors.Wrap`:**

  - Descrição: Utilizar as funcionalidades padrão do Go para criar e aninhar erros.
  - Motivo da Rejeição: Embora permita o aninhamento (wrapping) e a verificação com `errors.Is/As`, não fornece uma estrutura padronizada para códigos de erro específicos da aplicação ou para carregar contexto adicional de forma consistente sem boilerplate manual em cada local de criação de erro.

- **Bibliotecas de terceiros para tratamento de erros (ex: `pkg/errors`, `cockroachdb/errors`):**

  - Descrição: Adotar uma biblioteca externa que já oferece funcionalidades avançadas de erro.
  - Motivo da Rejeição: A estrutura `MessageError` proposta é relativamente simples, atende diretamente aos requisitos identificados (código, mensagem, contexto, wrapping) e evita adicionar uma dependência externa para uma funcionalidade tão central, mantendo o controle sobre a evolução da estrutura. Se necessidades mais complexas (como stack traces formatados) surgirem, essa decisão pode ser reavaliada.

- **Definir apenas sentinel errors e verificar com `errors.Is`:**
  - Descrição: Definir variáveis globais para cada tipo de erro (ex: `var ErrUserNotFound = errors.New("user not found")`).
  - Motivo da Rejeição: Esta abordagem é boa para identificar tipos de erro, mas não resolve nativamente a necessidade de adicionar contexto dinâmico (como o ID do usuário que não foi encontrado) ou códigos de erro padronizados para APIs de forma estruturada. A proposta atual de `MessageError` pode e deve ser usada em conjunto com sentinel errors como base (conforme demonstrado no exemplo `errUserNotFoundBase`).

## Consequências

**Positivas:**

- **Padronização e Consistência:** Todos os erros gerados pela aplicação seguirão uma estrutura uniforme, facilitando seu entendimento e tratamento.
- **Contexto Rico para Depuração:** A inclusão de um mapa de `Context` permite anexar dados específicos da ocorrência do erro (ex: IDs de entidades, parâmetros de entrada), o que é invaluable para a depuração.
- **Códigos de Erro Explícitos:** O campo `Code` permite que as APIs retornem códigos de erro bem definidos, facilitando a integração com clientes (frontends, outros serviços) e o tratamento específico de erros por parte deles.
- **Melhor Rastreabilidade e Monitoramento:** Erros estruturados são mais fáceis de serem processados, categorizados e analisados por sistemas de logging e monitoramento.
- **Clareza na Origem do Erro:** A combinação do erro base (`Err`), mensagem e código ajuda a identificar rapidamente a natureza e a origem do problema.
- **Modularidade:** A centralização da definição de erros específicos em arquivos `msg.go` por módulo/pacote organiza o código e facilita a manutenção.
- **Compatibilidade com `errors.Is` e `errors.As`:** A implementação do método `Unwrap()` garante a interoperabilidade com as funcionalidades padrão de inspeção de erros do Go.

**Negativas / Trade-offs:**

- **Overhead de Implementação Inicial:** Requer a criação da struct `MessageError` e das funções construtoras para os diversos tipos de erro da aplicação.
- **Pequeno Boilerplate:** Pode haver um leve aumento de código boilerplate ao definir cada erro com sua respectiva função construtora e erro base.
- **Curva de Aprendizagem:** A equipe de desenvolvimento precisará se familiarizar e aderir ao novo padrão de criação e tratamento de erros.
- **Impacto Mínimo na Performance:** A alocação da struct `MessageError` e do mapa de contexto para cada erro pode introduzir um overhead de performance mínimo. Em cenários com uma frequência extremamente alta de erros, isso poderia ser uma consideração, mas para a maioria das aplicações web, o benefício da informação contextual supera esse impacto.

**(Opcional) Notas Adicionais:**

- O campo `Message` na `MessageError` pode ser inicialmente usado tanto para logs detalhados quanto para respostas de API. Se houver necessidade de mensagens distintas para o usuário final (mais genéricas ou traduzidas) e para logs (mais técnicas), a estrutura `MessageError` pode ser estendida ou complementada no futuro.
- A correta implementação e uso do método `Unwrap()` é fundamental para garantir que a cadeia de erros seja preservada e que `errors.Is` e `errors.As` funcionem como esperado.
