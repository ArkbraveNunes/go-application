## Go Basic

Olá, bom dia, aqui será apresentado uma documentação relativa a estrutura implementada nessa aplicação go, o objetivo dessa documentação é expôr de maneira clara e objetiva o funcionamento da aplicação bem como sua estrutura.
O padrão implementado aqui segue conceitos como Clean Code, Clean Architecture, SOLID, DRY, KISS e Design Patterns. Essa documentação também tem correlação ao [Layout Padrão de Projetos em Go](https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md) seguindo os modelos estipulados de diretórios.

## IMPORTANTE

Caso tenha algum problema em rodar qualquer um dos projetos rode o comando

```shell
make install
```

assim todos os pacotes go serão atualizados ou instalados, caso não use o make pode instalar de maneira manual usando o go

```shell
go mod tidy
```

## Swagger

Rota do swagger de todas as aplicações é sempre o virtual host definido no docker-compose dentro da pasta .docker do [docker-development](https://github.com/ArkbraveNunes/docker-development.git) com o sufixo "/docs/index.html" no final!

## Estrutura

O Projeto segue o seguinte modelo de orgranização de arquivos e pastas:

```
../
  /api
  /cmd
  /config
  /internal
  /pkg
  /test
../
```

- /api -> Especificações OpenAPI/Swagger, arquivos de esquema JSON, arquivos de definição de protocolo;

- /cmd -> Principais aplicações para este projeto; O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter (ex. /cmd/myapp); Não coloque muitos códigos no diretório da aplicação. Se você acha que o código pode ser importado e usado em outros projetos, ele deve estar no diretório /pkg. Se o código não for reutilizável ou se você não quiser que outros o reutilizem, coloque esse código no diretório /internal. O padro ter uma pequena função main que importa e invoca o código dos diretórios /internal e /pkg e nada mais;

- /internal -> Aplicação privada e código de bibliotecas. Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas. Observe que esse padrão de layout é imposto pelo próprio compilador Go; Você pode ter mais de um diretório internal em qualquer nível da árvore do seu projeto; Opcionalmente, você pode adicionar um pouco de estrutura extra aos seus pacotes internos para separar o seu código interno compartilhado e não compartilhado. Não é obrigatório (especialmente para projetos menores), mas é bom ter dicas visuais que mostram o uso pretendido do pacote. Seu atual código da aplicação pode ir para o diretório /internal/app (ex. /internal/app/myapp) e o código compartilhado por essas aplicações no diretório /internal/pkg (ex. /internal/pkg/myprivlib);

- /pkg - Código de bibliotecas que podem ser usados por aplicativos externos (ex. /pkg/mypubliclib). Outros projetos irão importar essas bibliotecas esperando que funcionem, então pense duas vezes antes de colocar algo aqui; Observe que o diretório internal é a melhor maneira de garantir que seus pacotes privados não sejam importáveis porque é imposto pelo Go. O diretório /pkg contudo é uma boa maneira de comunicar explicitamente que o código naquele diretório é seguro para uso.

- /test -> Aplicações de testes externos adicionais e dados de teste. Sinta-se à vontade para estruturar o diretório /test da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter /test/data ou /test/testdata se precisar que o Go ignore o que está naquele diretório; Observe que o Go também irá ignorar diretórios ou arquivos que começam com "." ou "\_", para que você tenha mais flexibilidade em termos de como nomear seu diretório de dados de teste.
