# ADR-018: Adoção de `envconfig` e Arquivos `.env` para Gerenciamento de Configurações da Aplicação

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A aplicação RedToGreen necessita de uma forma robusta, segura e de fácil utilização para gerenciar suas configurações em diferentes ambientes (desenvolvimento local, staging, produção). As configurações incluem informações como strings de conexão com o banco de dados (ADR-012), níveis de log (ADR-017), portas de servidor (ADR-016), segredos de API, e parâmetros de feature flags (ADR-009).

Os critérios principais para a escolha da estratégia de configuração são:
* **Facilidade de Uso para Desenvolvedores:** Simplicidade na configuração do ambiente de desenvolvimento local.
* **Carregamento Estruturado e Tipado:** As configurações devem ser carregadas em structs Go com tipos de dados corretos.
* **Valores Padrão e Checagem de Obrigatoriedade:** Suporte para definir valores padrão e garantir que configurações essenciais sejam fornecidas.
* **Consistência com Princípios do Twelve-Factor App:** Preferência por configuração via variáveis de ambiente, especialmente para produção.
* **Mínimo de Dependências Externas:** Alinhamento com a filosofia do projeto (ADR-000) de utilizar soluções leves e, quando possível, da biblioteca padrão ou com baixo acoplamento.
* **Segurança:** Tratamento adequado de informações sensíveis (não versionar segredos).
* **Clareza e Baixa Verbosidade de Código:** O código para carregar e acessar as configurações deve ser o mais conciso e claro possível.

A pergunta que estamos tentando responder é: Qual abordagem devemos adotar para carregar, validar (obrigatoriedade e defaults) e acessar as configurações da aplicação RedToGreen de forma eficiente, segura e amigável para os desenvolvedores, priorizando o uso de arquivos `.env` para desenvolvimento local?

O escopo desta decisão abrange o mecanismo pelo qual a aplicação Go lê e processa suas configurações de tempo de execução.

## Decisão

Adotaremos a seguinte estratégia para o gerenciamento de configurações da aplicação RedToGreen:

1.  **Estrutura de Configuração em Go:** As configurações da aplicação serão definidas em uma ou mais `structs` Go. Estas structs utilizarão tags para metadados de configuração.
2.  **Biblioteca `kelseyhightower/envconfig`:** Utilizaremos a biblioteca `github.com/kelseyhightower/envconfig` para popular as structs de configuração a partir de variáveis de ambiente.
    * As tags `envconfig:"NOME_DA_VARIAVEL"` serão usadas para mapear variáveis de ambiente para os campos da struct.
    * As tags `default:"valor_padrao"` serão usadas para especificar valores padrão para os campos.
    * As tags `required:"true"` serão usadas para marcar campos cuja variável de ambiente correspondente deve obrigatoriamente estar definida.
    * `envconfig` também cuidará da conversão automática de strings de variáveis de ambiente para os tipos corretos dos campos da struct (int, bool, string, `time.Duration`, slices, etc.).
3.  **Arquivos `.env` para Desenvolvimento Local:** Para facilitar a configuração do ambiente de desenvolvimento local, utilizaremos arquivos `.env` (ex: `.env.local`, `.env.development`, ou um `.env` padrão).
    * A biblioteca `github.com/joho/godotenv` será usada no início da aplicação (apenas em ambientes de desenvolvimento ou quando explicitamente habilitado) para carregar as variáveis definidas nesses arquivos `.env` para o ambiente da aplicação, de onde o `envconfig` poderá lê-las.
    * Os arquivos `.env` contendo segredos ou configurações específicas de máquina não deverão ser versionados no Git (serão adicionados ao `.gitignore`). Um arquivo `.env.example` será versionado como template.
4.  **Variáveis de Ambiente em Produção:** Em ambientes de staging e produção, as configurações serão fornecidas diretamente como variáveis de ambiente, injetadas pela plataforma de hospedagem (ex: Kubernetes, Docker Compose, configurações de PaaS). `godotenv` não será utilizado ativamente nesses ambientes, ou sua falha em encontrar um arquivo `.env` não será tratada como um erro fatal.
5.  **Prefixo para Variáveis de Ambiente:** Um prefixo (ex: `APP_`) será usado para todas as variáveis de ambiente lidas pelo `envconfig` para evitar colisões com variáveis de ambiente do sistema.
6.  **Validações Adicionais:** Validações de configuração que vão além de "obrigatório" ou "padrão" (ex: checar se um valor está dentro de um conjunto permitido, validar ranges numéricos específicos) serão implementadas em código Go após a configuração ser carregada pela `envconfig`.

## Alternativas Consideradas

* **Implementação Manual Pura em Go (com `os.Getenv` e `strconv`):**
    * *Prós:* Controle total, sem dependências externas.
    * *Motivo da Rejeição:* Leva a um código de carregamento muito verboso e propenso a erros para conversão de tipos, aplicação de defaults e checagem de obrigatoriedade para múltiplas configurações. A verbosidade não se alinha com o desejo de clareza e concisão para esta tarefa.

* **`spf13/viper`:**
    * *Prós:* Biblioteca de configuração muito poderosa e flexível, com suporte a múltiplas fontes (arquivos JSON/YAML/TOML, variáveis de ambiente, flags, configuração remota via etcd/Consul), hierarquia de precedência, e unmarshalling para structs.
    * *Motivo da Rejeição (para a abordagem primária atual):* Embora muito capaz e com excelente suporte para evolução futura (configurações remotas), o setup inicial do Viper é mais verboso do que `envconfig` para o caso de uso principal de carregar configurações de variáveis de ambiente (populadas por `.env` em dev). A preferência por uma solução mais leve e focada, que ainda atenda aos requisitos de tipagem, estrutura e obrigatoriedade via tags, tornou `envconfig` mais atraente para o momento, dado o foco na simplicidade dos arquivos `.env` para os desenvolvedores. A integração com sistemas remotos, se necessária no futuro, pode ser alcançada por mecanismos que injetem configurações como variáveis de ambiente, mantendo a lógica da aplicação Go simples.

* **Outras Bibliotecas de Configuração Menores ou Mais Opinativas:**
    * *Motivo da Rejeição:* `envconfig` se destacou por seu foco específico em variáveis de ambiente, sua leveza, e a clareza de uso com tags, que foi um requisito apreciado.

## Consequências

**Positivas:**
* **Simplicidade e Baixa Verbosidade no Código:** O carregamento da configuração e a definição de defaults/obrigatoriedade são feitos de forma concisa utilizando `envconfig` e suas tags.
* **Facilidade para Desenvolvedores:** O uso de arquivos `.env` com `godotenv` torna a configuração do ambiente de desenvolvimento local extremamente simples e familiar.
* **Tipagem Clara e Estrutura:** As configurações são carregadas diretamente em structs Go, com conversão de tipos automática e segura.
* **Checagem de Obrigatoriedade e Defaults:** As tags `required` e `default` do `envconfig` lidam com esses aspectos de forma declarativa.
* **Alinhamento com Twelve-Factor App:** Prioriza a configuração via variáveis de ambiente, especialmente para produção.
* **Mínimo de Dependências:** Adiciona apenas duas dependências leves e bem focadas (`joho/godotenv`, `kelseyhightower/envconfig`).
* **Controle sobre Validações Complexas:** Validações de negócio específicas para as configurações ainda podem ser implementadas em Go, de forma explícita.

**Negativas / Trade-offs:**
* **Foco em Variáveis de Ambiente:** `envconfig` é primariamente desenhado para ler de variáveis de ambiente. Se houver uma necessidade futura forte de ler configurações de arquivos com estruturas complexas (JSON/YAML aninhados que não mapeiam bem para variáveis de ambiente) diretamente pela biblioteca de configuração, esta abordagem exigiria workarounds ou uma reconsideração.
* **Integração com Configuração Remota Indireta:** A integração com sistemas como Consul ou Vault dependeria de um mecanismo externo para popular as variáveis de ambiente da aplicação, em vez de a biblioteca de configuração se conectar diretamente a esses sistemas. Isso mantém a aplicação Go mais simples, mas move a complexidade da integração para a infraestrutura/deploy.
* **Validações Avançadas Manuais:** Regras de validação além de "obrigatório" e "padrão" precisam ser codificadas manualmente (embora isso seja comum e mantenha o controle explícito).

## (Opcional) Notas Adicionais
* Um arquivo `.env.example` será mantido no repositório para documentar todas as variáveis de ambiente esperadas pela aplicação e seus formatos.
* Arquivos `.env` reais (contendo segredos ou configurações específicas de desenvolvedor) serão listados no `.gitignore`.
* A função de carregamento de configuração (`config.Load()`) será chamada no início da aplicação (`main.go`), e a falha em carregar configurações obrigatórias resultará em um erro fatal, impedindo a inicialização da aplicação.
* Uma função ou método `GetSanitized()` na struct de configuração será implementada para permitir o log da configuração carregada sem expor segredos.
