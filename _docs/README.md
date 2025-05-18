# Documentação do Projeto RedToGreen

Bem-vindo ao diretório central de documentação do projeto RedToGreen! Este espaço organiza todos os artefatos importantes que definem, guiam e registram a evolução do nosso produto e de sua arquitetura técnica.

## Navegação da Documentação

A documentação está estruturada nos seguintes diretórios principais:

* **[`./adr/`](./adr/) - Architecture Decision Records (Registros de Decisão de Arquitetura)**
    * Contém todos os ADRs que documentam as decisões arquiteturais significativas tomadas para o projeto. Cada ADR detalha o contexto, a decisão, as alternativas consideradas e as consequências.
    * Comece pelo [ADR-000: Princípios Fundamentais da Cultura de Arquitetura e Desenvolvimento](./adr/000-principios_fundamentais_cultura_arquitetura_desenvolvimento.md) para entender nossa filosofia de desenvolvimento.
    * Um `template.md` está disponível para a criação de novos ADRs.

* **[`./dominio/`](./dominio/) - Modelagem do Domínio**
    * Descreve as entidades de negócio centrais, seus atributos, relacionamentos e os contextos de domínio do RedToGreen.
    * O [`README.md`](./dominio/README.md) dentro deste diretório serve como um índice para as entidades e contextos.
    * Inclui descrições detalhadas para:
        * Contextos: [`CONTEXTO_USUARIO.md`](./dominio/CONTEXTO_USUARIO.md), [`CONTEXTO_CARTEIRA.md`](./dominio/CONTEXTO_CARTEIRA.md)
        * Entidades do Usuário: [`user/user.md`](./dominio/user/user.md), [`user/category.md`](./dominio/user/category.md), [`user/tag.md`](./dominio/user/tag.md)
        * Entidades da Carteira: [`wallet/wallet.md`](./dominio/wallet/wallet.md), [`wallet/bank_account.md`](./dominio/wallet/bank_account.md), [`wallet/transaction.md`](./dominio/wallet/transaction.md)

* **[`./epicos/`](./epicos/) - Épicos do Produto**
    * Este diretório conterá a documentação dos Épicos, que são grandes funcionalidades ou temas de alto nível do produto. Cada épico será decomposto em User Stories.
    * *(Este diretório será populado conforme o planejamento do produto avança).*

* **[`./produto/`](./produto/) - Documentação do Produto**
    * Contém documentos que descrevem o produto RedToGreen sob uma perspectiva de negócio e do usuário.
    * Principal documento: [`visao_geral_v1.md`](./produto/visao_geral_v1.md) - Descreve a proposta de valor, público-alvo e funcionalidades centrais do MVP.

* **[`GLOSSARIO.md`](./GLOSSARIO.md) - Glossário de Termos**
    * Define os termos chave e conceitos utilizados em todo o projeto RedToGreen para garantir um entendimento comum e consistente. É uma leitura recomendada para todos os envolvidos.

## Propósito

O objetivo desta documentação é servir como uma fonte única da verdade (Single Source of Truth - SSOT) para o design, arquitetura e entendimento funcional do RedToGreen. Ela deve ser mantida atualizada e ser uma referência constante para a equipe de desenvolvimento e outros stakeholders.

---
