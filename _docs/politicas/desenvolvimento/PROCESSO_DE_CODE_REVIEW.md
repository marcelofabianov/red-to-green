# Política de Revisão de Código (Code Review) do RedToGreen

**Última Atualização:** 2025-05-18

## 1. Propósito

O processo de revisão de código é uma etapa fundamental no ciclo de desenvolvimento do RedToGreen para garantir a qualidade, manutenibilidade, segurança e consistência do código, além de promover o compartilhamento de conhecimento entre a equipe. Esta política define as diretrizes para a realização de revisões de código.

* _Relacionado ao ADR-000: Qualidade e Robustez, Documentação Contínua._

## 2. Quando Realizar Code Review?

* **Todas as Alterações de Código:** Qualquer código que será integrado à branch principal de desenvolvimento (ex: `develop` ou `main`, dependendo do Gitflow - ADR a ser definido) **DEVE** passar por revisão de código. Isso inclui novas features, correções de bugs, refatorações e atualizações de dependências.
* **Pull Requests (PRs) / Merge Requests (MRs):** O mecanismo padrão para submeter código para revisão será através de PRs/MRs na plataforma de controle de versão (ex: GitHub, GitLab).

## 3. Responsabilidades

### 3.1. Autor do Código (Desenvolvedor)

* **Preparação para Revisão:**
    * Garantir que o código compila sem erros.
    * Garantir que todos os testes unitários (ADR-019) e de integração relevantes passam.
    * Garantir que o código está formatado com `goimports` e passa no linter (`golangci-lint`) configurado para o projeto.
    * Escrever uma descrição clara no PR, explicando o "quê" e o "porquê" da mudança, e como testá-la. Referenciar User Stories, Tasks ou Issues relevantes.
    * Realizar uma auto-revisão antes de solicitar a revisão de outros.
* **Durante a Revisão:**
    * Estar aberto a feedback construtivo.
    * Responder prontamente às dúvidas e comentários dos revisores.
    * Realizar as alterações solicitadas ou justificar por que não foram feitas.

### 3.2. Revisor(es)

* **Abordagem:** Revisar com foco construtivo, visando melhorar a qualidade do código e do produto, não criticar o autor.
* **Escopo da Revisão (O que verificar):**
    * **Corretude:** A lógica implementada atende aos requisitos e resolve o problema proposto?
    * **Clareza e Legibilidade:** O código é fácil de entender? A nomenclatura é clara?
    * **Simplicidade (KISS):** A solução é a mais simples possível para o problema? Há complexidade desnecessária? (ADR-000)
    * **Aderência aos Princípios e ADRs:** O código segue os [Princípios Fundamentais (ADR-000)](../../adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md) e os ADRs relevantes?
    * **Aderência ao Guia de Estilo:** O código segue o [Guia de Estilo Go do RedToGreen](./GUIA_DE_ESTILO_GO.md)?
    * **Testes:** Os testes unitários são adequados, cobrem os cenários importantes e seguem as diretrizes (ADR-019)?
    * **Performance:** Há gargalos óbvios de performance ou uso ineficiente de recursos?
    * **Segurança:** Há vulnerabilidades de segurança aparentes? (Ex: SQL injection, XSS, tratamento inadequado de dados sensíveis).
    * **Documentação:** Comentários `godoc` para código exportado estão presentes e claros? A lógica complexa está bem comentada?
    * **Impacto:** Quais são os possíveis impactos desta mudança em outras partes do sistema?
* **Feedback:**
    * Ser específico, educado e fornecer sugestões claras.
    * Distinguir entre sugestões obrigatórias (bloqueiam o merge) e sugestões opcionais/nitpicking.
* **Aprovação:** Aprovar o PR apenas quando estiver satisfeito com a qualidade e a correção do código.

## 4. Processo

1.  O autor cria uma branch de feature (seguindo o Gitflow - ADR a ser definido), implementa a funcionalidade e commita as alterações.
2.  O autor abre um Pull Request (PR) para a branch de destino (ex: `develop`).
3.  **No mínimo 1 (um) outro desenvolvedor da equipe** deve revisar o PR. Para features críticas ou complexas, 2 (dois) revisores são recomendados.
4.  Os revisores adicionam comentários e sugestões.
5.  O autor responde aos comentários e faz as alterações necessárias.
6.  O ciclo de revisão/alteração continua até que os revisores aprovem o PR.
7.  Após a aprovação e a passagem de todos os checks de CI, o PR é mergeado.

## 5. Ferramentas

* **Plataforma de Controle de Versão:** (ex: GitHub, GitLab) para gerenciar PRs e discussões.
* **Comunicação:** Ferramenta de chat da equipe para discussões rápidas, se necessário, mas o feedback principal deve ficar registrado no PR.
