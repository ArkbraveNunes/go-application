## Go Basic

Olá, bom dia, aqui será apresentado uma documentação relativa a estrutura implementada nessa aplicação go, o objetivo dessa documentação é expôr de maneira clara e objetiva o funcionamento da aplicação bem como sua estrutura.
O padrão implementado aqui segue conceitos como Clean Code, Clean Architecture, SOLID, DRY, KISS e Design Patterns. Essa documentação também tem correlação ao [Layout Padrão de Projetos em Go](https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md#visÃ£o-geral) seguindo os modelos estipulados de diretórios.

## Estrutura

O Projeto segue o seguinte modelo de orgranização de arquivos e pastas:

```
  /api
  /cmd
  /internal
  /pkg
  /vendor

```

- /src - Aqui se localiza a raíz da aplicação em questão;

- /common - Aqui é armazendo pastas contendo arquivos cuja finalidade é utilizá-los na aplicação como um todo independente do seu local;

  - /config - Local os se localiza o mapeamento das variáveis de ambiente(env) da aplicação
  - /enum - Essa pasta compreende os enums(identificadores) utilizados na aplicação como um todo

- /module -> Aqui são armazenados os módulos da aplicação;

- /module/moduleName/application:

  - /controller - Essa pasta compreende os controllers da aplicação;
  - /dto - Aqui se localiza as definições de entrada(IN) e saída(OUT) dos controllers;

- /module/moduleName/domain:

  - /service: Compreende principalmente os services da aplicação;
  - /entity - Entidade de conversão de dados, atuando como um de/para para os dados fornecidos pelo Banco de Dados;
  - /contract - Local onde fica definido os métodos(queries) enviadas ao banco de dados;

- /module/moduleName/infra:
  - /repository: Aqui se localiza as queries executadas no banco de dados;
  - /adapter: Diretório que compreende conexões com provedores de terceiros ou outras aplicações em diferentes contextos;
  - /schema: Local onde são armazenadas as Tabelas e/ou Schemas/Models relativas ao Banco de Dados;
