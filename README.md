## Introdução
Esse projeto foi feito pra um server de minecraft de um amigo (@flipbit03)   
A ideia é um bot do discord que response o status atual do servidor usando comandos.

## Como rodar o projeto
Versão do Golang usada no projeto: 1.22.3

- Tenha o Go instalado na sua máquina
- Crie um arquivo `.env` na raiz do projeto e popule ele com as variáveis baseado no arquivo pré-existente chamado `.env.template`. Você também pode usar as variáveis diretamente no seu ambiente, sem usar o arquivo `.env`
- Execute o projeto utilizando `go run main.go` ou buildando e executando o projeto.

## Env vars
- GUILD_ID -> Id do server do discord que o bot irá registrar os comandos.   
- BOT_TOKEN -> Token do seu bot   
- HOST -> IP do servidor de minecraft com a porta (caso necessário), ex: "localhost:25565"