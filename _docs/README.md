# Documentação Central do Projeto RedToGreen

Bem-vindo ao coração da documentação do projeto RedToGreen! Este diretório é o repositório central para todos os artefatos de documentação que definem, guiam e registram a evolução do nosso produto, desde a sua concepção até as decisões técnicas de arquitetura e as especificações funcionais.

O objetivo desta documentação é servir como uma fonte única da verdade (Single Source of Truth - SSOT), promovendo um entendimento compartilhado entre todos os membros da equipe e stakeholders, facilitando o onboarding de novos colaboradores e apoiando a tomada de decisões futuras.

## Estrutura da Documentação

A documentação está organizada nos seguintes diretórios principais:

* **[`./GLOSSARIO.md`](./GLOSSARIO.md)**
    * Contém as definições dos termos e conceitos chave utilizados em todo o projeto RedToGreen. Essencial para garantir uma linguagem e entendimento comuns.

* **[`./produto/`](./produto/)**
    * **Propósito:** Documentos relacionados à visão, estratégia e definição do produto RedToGreen.
    * **Conteúdo Chave:**
        * `visao_geral_v1.md`: Descreve a proposta de valor, público-alvo, objetivos e funcionalidades centrais do MVP (Produto Mínimo Viável).

* **[`./adr/`](./adr/) - Architecture Decision Records**
    * **Propósito:** Registra as decisões arquiteturais significativas tomadas durante o desenvolvimento do projeto. Cada ADR detalha o contexto da decisão, a solução escolhida, as alternativas consideradas e as consequências.
    * **Conteúdo Chave:**
        * `000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md`: O manifesto que guia nossa filosofia de desenvolvimento e arquitetura.
        * Demais ADRs numerados (ex: `001-*.md` a `020-*.md`) cobrindo desde a escolha da linguagem até estratégias de teste.
        * `template.md`: Modelo para criação de novos ADRs.

* **[`./dominio/`](./dominio/) - Modelagem do Domínio**
    * **Propósito:** Descreve o coração do negócio do RedToGreen – as entidades, seus atributos, relacionamentos e os contextos de domínio em que operam.
    * **Conteúdo Chave:**
        * `README.md`: Um índice e introdução às entidades e contextos do domínio.
        * `CONTEXTO_USUARIO.md` e `CONTEXTO_CARTEIRA.md`: Descrições detalhadas dos principais contextos de domínio.
        * Subdiretórios `user/` e `wallet/` contendo a documentação de cada entidade específica (ex: [`user/user.md`](./dominio/user/user.md), [`wallet/transaction.md`](./dominio/wallet/transaction.md)).

* **[`./epicos/`](./epicos/) - Épicos do Produto**
    * **Propósito:** Documenta os Épicos, que são grandes funcionalidades ou temas de alto nível do produto. Cada épico é geralmente decomposto em User Stories menores.
    * *(Este diretório será populado conforme o planejamento do produto avança).*

* **[`./epicos/**/us_**.md`](./epicos/) - Histórias de Usuário (US)**
    * **Propósito:** Captura os requisitos sob a perspectiva do usuário final, descrevendo uma funcionalidade e o valor que ela entrega. Incluirá Critérios de Aceite (ACs).
    * **Exemplos:**
        * `us_001_cadastro_de_usuarios.md`/`: Descreve a US de cadastro de usuários e os ACs relacionados.

* **[`./epicos/**/rf_**.md` e `./epicos/**/nrf_**.md`/`](./epicos/) - Requisitos Detalhados**
    * **Propósito:** Especifica de forma mais granular os requisitos funcionais e não funcionais do sistema.
    * **Exemplos:**
        * `rf_001_gerenciamento_de_usuarios.md`/`: Documenta os Requisitos Funcionais (RFs) relacionados ao gerenciamento de usuários.
        * `nrf_001_gerenciamento_de_usuarios.md`/`: Documenta os Requisitos Não Funcionais (NRFs) relacionados ao gerenciamento de usuários.

* **[`./politicas/`](./politicas/) - Políticas e Diretrizes do Projeto**
    * **Propósito:** Armazena documentos que estabelecem políticas, padrões e guias para o desenvolvimento e arquitetura, complementando os ADRs.
    * **Conteúdo Chave:**
        * `README.md`: Explica o propósito deste diretório e lista as políticas disponíveis.
        * Subdiretórios como `desenvolvimento/`, `seguranca/`, `arquitetura/` conterão os documentos de política específicos (ex: Guia de Estilo Go, Processo de Code Review).

## Como Navegar

* Para uma introdução rápida ao produto, comece pela [Visão Geral do Produto](./produto/visao_geral_v1.md).
* Para entender nossa filosofia de trabalho, leia os [Princípios Fundamentais (ADR-000)](./adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md).
* Para termos técnicos e de negócio, consulte o [Glossário](./GLOSSARIO.md).
* Para decisões técnicas específicas, navegue pelos [ADRs](./adr/).
* Para entender a estrutura de dados e regras de negócio, explore a [Modelagem do Domínio](./dominio/).
* Para guias de desenvolvimento e padrões do projeto, consulte as [Políticas](./politicas/).

Esta documentação é um artefato vivo e será continuamente atualizada para refletir o estado atual e a evolução do projeto RedToGreen.
