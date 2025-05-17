# ADR-009: Adoção de Feature Flags para Gerenciamento de Lançamentos e Funcionalidades

**Status:** Aceito

**Data:** 2025-05-17

## Contexto

No desenvolvimento moderno de software, especialmente para produtos SaaS como o RedToGreen, é crucial ter flexibilidade no lançamento de novas funcionalidades, mitigar riscos associados a deploys e segmentar o acesso a features para diferentes grupos de usuários ou planos contratuais. Práticas tradicionais de lançamento, como branches de feature de longa duração ou deploys monolíticos de grandes conjuntos de funcionalidades, podem introduzir atrasos, complexidade em merges e aumentar o risco de instabilidade.

O RedToGreen necessita de mecanismos para:
* Desacoplar o deploy de código do lançamento efetivo de features.
* Permitir lançamentos graduais (canary releases) e testes A/B.
* Ter a capacidade de desabilitar rapidamente uma feature problemática (kill switch).
* Oferecer funcionalidades específicas para diferentes segmentos de usuários ou planos de assinatura ("entregas granulares").
* Facilitar o Desenvolvimento Baseado em Tronco (Trunk-Based Development) e a integração contínua.

A pergunta que estamos tentando responder é: Como podemos gerenciar o ciclo de vida das funcionalidades do RedToGreen de forma flexível, segura e controlada, permitindo lançamentos graduais, testes e segmentação de usuários?

O escopo desta decisão abrange a estratégia de desenvolvimento, deploy e lançamento de novas funcionalidades, bem como a modificação de funcionalidades existentes na aplicação RedToGreen.

## Decisão

Adotaremos o uso de **Feature Flags (FFs)** como uma prática padrão e integral no ciclo de desenvolvimento e operação do RedToGreen.

1.  **Tipos de Feature Flags:** Serão considerados e utilizados diferentes tipos de flags conforme a necessidade, incluindo:
    * **Release Toggles:** Para permitir que código novo ou incompleto seja integrado à branch principal e deployado em produção sem estar ativo, facilitando o Trunk-Based Development.
    * **Experiment Toggles (Testes A/B):** Para testar diferentes versões de uma feature ou UI/UX com subconjuntos de usuários e tomar decisões baseadas em dados.
    * **Ops Toggles:** Para controlar aspectos operacionais do sistema, permitindo habilitar/desabilitar funcionalidades que possam impactar a performance, estabilidade ou para gerenciar a degradação graciosa de serviços.
    * **Permission Toggles (Contrato/Plano):** Para controlar o acesso a features específicas com base no plano de assinatura do usuário, tipo de contrato, ou outros critérios de segmentação de usuários. Este é um caso de uso chave para o RedToGreen.

2.  **Backend como Fonte da Verdade:** A lógica primária para determinar o estado de uma feature flag (ativa/inativa para um dado contexto de usuário/ambiente) residirá no backend (desenvolvido em Go). Isso garante a aplicação consistente das regras de ativação, mesmo que a interface do usuário seja contornada.

3.  **Frontend para Experiência do Usuário:** O frontend consumirá o estado das flags (seja diretamente de um serviço de FF ou via uma API do backend) para adaptar dinamicamente a interface e a experiência do usuário – mostrando ou ocultando elementos, habilitando/desabilitando funcionalidades, etc.

4.  **Ciclo de Vida das Flags:** Será estabelecida uma política para o gerenciamento do ciclo de vida das FFs, incluindo a revisão periódica e a remoção de flags obsoletas (features 100% lançadas, universalmente ativas/inativas, ou abandonadas) para mitigar a dívida técnica no código.

## Alternativas Consideradas (Opcional)

* **Branches de Feature de Longa Duração:**
    * Descrição: Manter o desenvolvimento de novas features em branches separadas por longos períodos, fazendo o merge apenas quando a feature está "completa".
    * Motivo da Rejeição: Aumenta o risco e a complexidade dos merges ("merge hell"), atrasa a integração e o feedback, e dificulta o deploy contínuo. As FFs permitem integrar código frequentemente na branch principal (Trunk-Based Development).

* **Deploy de Múltiplas Versões da Aplicação:**
    * Descrição: Manter e operar múltiplas versões da aplicação em paralelo para diferentes conjuntos de features ou usuários.
    * Motivo da Rejeição: Extremamente complexo em termos operacionais, de infraestrutura e de manutenção. As FFs permitem uma única base de código e um único artefato de deploy com múltiplas "personalidades" ativadas dinamicamente.

* **Ausência de Gerenciamento Explícito de Features (Deploy Direto):**
    * Descrição: Todas as features mergeadas na branch principal são automaticamente lançadas para todos os usuários no próximo deploy.
    * Motivo da Rejeição: Não oferece flexibilidade para lançamentos graduais, testes A/B, segmentação de usuários ou desativação rápida de features problemáticas (kill switch). Aumenta significativamente o risco associado a cada deploy.

## Consequências

**Positivas:**

* **Desacoplamento entre Deploy e Lançamento:** Novas funcionalidades podem ser deployadas em produção de forma segura, permanecendo inativas até que sejam explicitamente habilitadas.
* **Redução Significativa de Risco:** Facilita lançamentos graduais (canary releases), testes em produção com um subconjunto de usuários e a implementação de "kill switches" para desativar instantaneamente features problemáticas.
* **Segmentação Avançada de Usuários:** Permite direcionar features específicas para diferentes grupos de usuários com base em diversos critérios (plano, contrato, comportamento, localização, etc.), crucial para a estratégia de monetização e personalização do RedToGreen.
* **Testes A/B e Experimentação:** Viabiliza a execução de testes A/B para comparar diferentes versões de uma feature e tomar decisões de produto baseadas em dados.
* **Suporte ao Desenvolvimento Contínuo e Trunk-Based Development:** Desenvolvedores podem integrar seu código na branch principal com frequência, mantendo features incompletas ou instáveis desativadas por FFs, o que simplifica o processo de merge e acelera o ciclo de desenvolvimento.
* **Feedback Mais Rápido:** Permite liberar features para grupos beta de usuários para coletar feedback antecipado.

**Negativas / Trade-offs:**

* **Aumento da Complexidade no Código-Fonte:** A introdução de lógica condicional (`if feature_enabled... else...`) para controlar as FFs pode aumentar a complexidade do código e o número de caminhos a serem testados.
* **Dívida Técnica ("Flag Debt"):** Se flags obsoletas não forem removidas ativamente, elas podem se acumular e tornar o código mais difícil de entender, manter e testar. É essencial ter uma política de gerenciamento do ciclo de vida das flags.
* **Overhead de Gerenciamento das Flags:** Requer um sistema ou processo para gerenciar o estado das flags, suas configurações de targeting, quem pode modificá-las e para auditar alterações.
* **Complexidade nos Testes:** A matriz de testes aumenta, pois é necessário validar os diferentes comportamentos da aplicação com as flags ativadas e desativadas, e para diferentes segmentos de usuários.
* **Potencial Impacto na Performance:** A lógica para determinar o estado de uma flag (consulta a um banco de dados, arquivo de configuração ou serviço externo) pode introduzir uma pequena latência. Ferramentas de FF dedicadas geralmente mitigam isso com SDKs eficientes que utilizam caching local e/ou atualizações em tempo real.
* **Consistência:** Garantir que o estado da flag seja verificado consistentemente em todas as camadas relevantes da aplicação (backend, frontend) é importante.

**(Opcional) Notas Adicionais:**

* Para um produto financeiro como o RedToGreen, a capacidade de auditar quem alterou o estado de uma feature flag e quando é um aspecto de segurança e conformidade importante. Ferramentas de FF dedicadas geralmente oferecem essa funcionalidade.
* É crucial garantir a consistência na avaliação do estado das flags, especialmente para funcionalidades que manipulam dados ou têm implicações de segurança.
* O desempenho do sistema de avaliação das flags deve ser monitorado para garantir que não se torne um gargalo.
