# ADR-011: Adoção da Linguagem Go (Golang) para o Desenvolvimento do Backend

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

A seleção da linguagem de programação principal para o backend é uma decisão tecnológica fundamental que impacta a produtividade do desenvolvimento, performance da aplicação, escalabilidade, manutenibilidade e a capacidade de atrair e reter talentos. Para o RedToGreen, um produto SaaS de gestão financeira pessoal com arquitetura inicial de monolito modular (conforme ADR-010), a linguagem escolhida deve suportar esses aspectos de forma eficaz.

Um estudo comparativo considerou Go, Node.js e PHP, avaliando características como ecossistema, performance, facilidade de desenvolvimento, ferramentas, cultura da comunidade e adequação aos requisitos do produto.

A pergunta que estamos tentando responder é: Qual linguagem de programação principal devemos adotar para o desenvolvimento do backend do RedToGreen, que melhor se alinhe com nossos objetivos de simplicidade de ferramental, performance, concorrência eficiente, e uma cultura de desenvolvimento que valoriza a robustez da biblioteca padrão?

O escopo desta decisão abrange a linguagem primária a ser utilizada para a implementação de toda a lógica de negócios, APIs, e demais funcionalidades do componente backend do RedToGreen.

## Decisão

Adotaremos a linguagem **Go (Golang)** como a principal linguagem de programação para o desenvolvimento do backend do sistema RedToGreen.

A escolha é fundamentada nos seguintes pontos fortes identificados para Go, que se alinham com os objetivos do projeto:
1.  **Simplicidade e Clareza da Linguagem:** Sintaxe concisa e um conjunto pequeno de funcionalidades ortogonais que promovem código legível e de fácil manutenção.
2.  **Performance Elevada:** Go é compilada para código de máquina nativo, resultando em executáveis rápidos e com baixo overhead, além de um garbage collector otimizado para baixa latência.
3.  **Concorrência Nativa Eficiente:** Goroutines e channels oferecem um modelo poderoso e simplificado para construir aplicações concorrentes e escaláveis, essencial para serviços web modernos.
4.  **Biblioteca Padrão Robusta e Abrangente (stdlib):** A stdlib de Go é extensa, cobrindo muitas necessidades comuns (HTTP/2, JSON, criptografia, I/O, SQL, testes, etc.), o que reduz a dependência de bibliotecas de terceiros e a "fadiga de dependências".
5.  **Ferramental Integrado e Padronizado:** Ferramentas como `go fmt`, `go test`, `go mod`, `go vet`, e `pprof` são parte da distribuição padrão, promovendo consistência e produtividade sem a necessidade de configurar um ecossistema complexo de ferramentas externas.
6.  **Compilação Rápida e Binário Único:** A compilação rápida melhora o ciclo de desenvolvimento, e a geração de um binário único (na maioria dos casos) simplifica drasticamente o deploy, especialmente em ambientes de contêineres (Docker).
7.  **Baixo Consumo de Memória:** Go tende a ser mais eficiente no uso de memória em comparação com linguagens interpretadas ou aquelas com VMs mais pesadas.
8.  **Tipagem Estática Forte:** Contribui para a robustez do código, detecção de erros em tempo de compilação e refatorações mais seguras.

## Alternativas Consideradas (Opcional)

* **Node.js (JavaScript/TypeScript no Backend):**
    * *Prós:* Vasto ecossistema NPM, possibilidade de usar JavaScript/TypeScript full-stack, I/O não bloqueante ideal para aplicações I/O-bound, comunidade grande.
    * *Motivo da Rejeição:* Principalmente devido à "fadiga de configuração" (necessidade de escolher e configurar múltiplas ferramentas para linting, formatação, transpilação, bundling, testes), à complexidade e ao tamanho potencial da pasta `node_modules` e ao gerenciamento de um grande número de dependências. A natureza single-threaded para tarefas CPU-bound também foi uma consideração.

* **PHP:**
    * *Prós:* Maturidade, grande base instalada, facilidade de hospedagem tradicional, frameworks poderosos e completos (ex: Laravel, Symfony), gerenciador de dependências moderno (Composer), e melhorias significativas de performance e funcionalidades nas versões recentes (PHP 8+).
    * *Motivo da Rejeição:* Percepção de que, para certos cenários de serviços web de alta performance ou algoritmos específicos, Go poderia oferecer vantagens. Além disso, a filosofia de Go em relação à biblioteca padrão e ao ferramental integrado pareceu mais alinhada com o desejo de simplicidade e menor dependência de ecossistemas de frameworks muito opinativos para as funcionalidades core. O modelo de execução tradicional do PHP (ex: PHP-FPM) também foi considerado menos ideal para certos tipos de serviços persistentes sem o uso de ferramentas adicionais.

## Consequências

**Positivas:**

* **Alta Performance e Eficiência de Recursos:** Aplicações Go são conhecidas por seu baixo consumo de memória e alta performance, adequadas para serviços que precisam lidar com muitas requisições concorrentes.
* **Desenvolvimento Concorrente Simplificado:** O modelo de concorrência de Go é um dos seus maiores trunfos, permitindo construir sistemas responsivos e escaláveis de forma mais direta.
* **Produtividade do Desenvolvedor:** A simplicidade da linguagem, compilação rápida e o ferramental integrado contribuem para um ciclo de desenvolvimento eficiente.
* **Manutenibilidade:** Código Go tende a ser claro e legível, facilitado pela obrigatoriedade do `go fmt` e pela simplicidade da sintaxe.
* **Deploy Simplificado:** A capacidade de gerar binários únicos e estaticamente vinculados (na maioria das vezes) torna o deploy em servidores e contêineres (Docker) muito mais simples e as imagens de contêiner menores.
* **Redução da "Fadiga de Ferramentas":** A cultura de Go e seu ferramental built-in minimizam o tempo gasto na escolha e configuração de ferramentas externas.
* **Forte Suporte para Aplicações de Rede:** A biblioteca padrão oferece excelente suporte para a construção de serviços de rede, incluindo servidores HTTP/2.
* **Comunidade Crescente e Ativa:** A comunidade Go é vibrante e focada em construir software robusto e de alta qualidade.

**Negativas / Trade-offs:**

* **Gerenciamento de Erros Explícito:** O padrão `if err != nil` é uma característica distintiva de Go. Embora promova um tratamento de erro cuidadoso, pode ser percebido como verboso por desenvolvedores acostumados com exceções (try-catch).
* **Frameworks Web Menos Opinativos (em geral):** Comparado a ecossistemas como Ruby on Rails, Django (Python) ou Laravel (PHP), os frameworks web em Go (ex: Gin, Echo, Chi) tendem a ser mais modulares e menos "tudo-incluso". Isso oferece flexibilidade, mas pode exigir mais trabalho de configuração inicial para montar uma stack completa de aplicação.
* **Genéricos São Relativamente Novos:** Embora os genéricos tenham sido adicionados no Go 1.18 e sejam uma adição poderosa, o ecossistema e os padrões de uso ainda estão amadurecendo em torno deles.
* **Curva de Aprendizagem para Paradigmas Específicos de Go:** Embora a sintaxe seja simples, dominar os aspectos idiomáticos de Go, como a concorrência com goroutines e channels e o tratamento de erros, requer aprendizado e prática.
* **Tamanho do Ecossistema de Bibliotecas de Terceiros:** Embora a biblioteca padrão seja forte, para nichos muito específicos, o número de bibliotecas de terceiros maduras pode ser menor do que em ecossistemas como Node.js (NPM) ou Python (PyPI). No entanto, para desenvolvimento web e de sistemas, Go é muito bem servido.

**(Opcional) Notas Adicionais:**

* A escolha de Go é consistente com a decisão de uma arquitetura monolítica modular (ADR-010), pois a performance e a eficiência de Go beneficiam um monolito, e sua tipagem estática e estrutura facilitam a criação de módulos bem definidos e interfaces claras.
* A ênfase da comunidade Go em utilizar a biblioteca padrão e evitar o excesso de dependências alinha-se com o desejo expresso de evitar a complexidade de gerenciamento de pacotes observada em outros ecossistemas.
