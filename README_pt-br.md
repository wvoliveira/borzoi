# Borzoi

[English](README.md)

Um sistema básico de gerenciamento de clientes.

## Recursos

- Cadastro com usuário/senha, redes sociais, número de celular, etc
- Gerencie o seu perfil, adicionando foto, descrição, redes sociais, etc
- Adicione seus clientes apenas com informações básicas
- Adicione notas aos seus clientes, como um diário
- Importe e exporte informações em formatos CSV, PDF, etc
- Crie QRCode e compartilhe

## Instalação

Há dois exemplos de instalação:

1. Download direto das [releases](https://github.com/elga-io/borzoi/releases).

2. Use o **make** criar o seu próprio binário:

    ```bash
    make build
    ```

    Ou caso não tenha o Go instalado, poderá utilizar o docker:

    ```bash
    make build-docker-for-windows
    # make build-docker-for-linux
    ```

3. Ou crie uma imagem docker e rode no próprio docker:

    ```bash
    make docker
    make docker-run
    ```
    E acesse http://localhost para utilizar o sistema.

## Utilização

Uma vez criado o binário, basta rodar o comando **borzoi** e acessar a porta 8080 (padrão):

```bash
./borzoi
```

Vá para http://localhost:8080 e seja feliz.

Para listar todas as opções disponíveis, rode: borzoi --help:

```bash
$ borzoi --help

Borzoi is a basic client management. Easy to install and use.

Usage:
  borzoi [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  server      Options for Spitz server/API.

Flags:
      --config string   config file (default is $HOME/.spitz/spitz.yaml)
  -h, --help            help for spitz
  -v, --version         version for spitz

Use "spitz [command] --help" for more information about a command.
```

## License

[MIT](LICENSE)