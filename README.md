#<img src="https://user-images.githubusercontent.com/<id>.svg" width="240"/>

[![Latest GitHub release](https://img.shields.io/github/v/release/elga-io/borzoi?color=25ae8f)](https://github.com/dstotijn/borzoi/releases/latest)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fdstotijn%2Fborzoi%2Fbadge%3Fref%3Dmain&label=build&color=24ae8f)](https://github.com/dstotijn/borzoi/actions/workflows/build-test.yml)
![GitHub download count](https://img.shields.io/github/downloads/dstotijn/borzoi/total?color=25ae8f)
[![GitHub](https://img.shields.io/github/license/dstotijn/borzoi?color=25ae8f)](https://github.com/dstotijn/borzoi/blob/master/LICENSE)
[![Documentation](https://img.shields.io/badge/borzoi-docs-25ae8f)](https://borzoi.xyz/)

* [Portuguese](README_pt-br.md)

**Borzoi** is a shortener app. It aims to become an open
source alternative to commercial software like Bitly, with simple features.

<img src="https://borzoi.xyz/img/hero.png" width="907" alt="Borzoi (screenshot)" />

## Features

- Machine-in-the-middle (MITM) HTTP proxy, with logs and advanced search
- HTTP client for manually creating/editing requests, and replay proxied requests
- Scope support, to help keep work organized
- Easy-to-use web based admin interface
- Project based database storage, to help keep work organized

👷‍♂️ Borzoi is under active development. Check the <a
href="https://github.com/dstotijn/borzoi/projects/1">backlog</a> for the current
status.

📣 Are you pen testing professionaly in a team? I would love to hear your
thoughts on tooling via [this 5 minute
survey](https://forms.gle/36jtgNc3TJ2imi5A8). Thank you!

## Getting started

💡 The [Getting started](https://borzoi.xyz/docs/getting-started) doc has more
detailed install and usage instructions.

### Installation

The quickest way to install and update Borzoi is via a package manager:

#### macOS

```sh
brew install borzoisoft/tap/borzoi
```

#### Linux

```sh
sudo snap install borzoi
```

#### Windows

```sh
scoop bucket add borzoisoft https://github.com/borzoisoft/scoop-bucket.git
scoop install borzoisoft/borzoi
```

#### Other

Alternatively, you can [download the latest release from
GitHub](https://github.com/dstotijn/borzoi/releases/latest) for your OS and
architecture, and move the binary to a directory in your `$PATH`. If your OS is
not available for one of the package managers or not listed in the GitHub
releases, you can compile from source _(link coming soon)_ or use a Docker image
_(link coming soon)_.

### Usage

Once installed, start Borzoi via:

```sh
borzoi
```

💡 Read the [Getting started](https://borzoi.xyz/docs/getting-started) doc for
more details.

To list all available options, run: `borzoi --help`:

```
$ borzoi --help

Usage:
    borzoi [flags] [subcommand] [flags]

Runs an HTTP server with (MITM) proxy, GraphQL service, and a web based admin interface.

Options:
    --cert         Path to root CA certificate. Creates file if it doesn't exist. (Default: "~/.borzoi/borzoi_cert.pem")
    --key          Path to root CA private key. Creates file if it doesn't exist. (Default: "~/.borzoi/borzoi_key.pem")
    --db           Database directory path. (Default: "~/.borzoi/db")
    --addr         TCP address for HTTP server to listen on, in the form \"host:port\". (Default: ":8080")
    --chrome       Launch Chrome with proxy settings applied and certificate errors ignored. (Default: false)
    --verbose      Enable verbose logging.
    --json         Encode logs as JSON, instead of pretty/human readable output.
    --version, -v  Output version.
    --help, -h     Output this usage text.

Subcommands:
    - cert  Certificate management

Run `borzoi <subcommand> --help` for subcommand specific usage instructions.

Visit https://borzoi.xyz to learn more about Borzoi.
```

## Documentation

📖 [Read the docs](https://borzoi.xyz/docs)

## Support

Use [issues](https://github.com/dstotijn/borzoi/issues) for bug reports and
feature requests, and
[discussions](https://github.com/dstotijn/borzoi/discussions) for questions and
troubleshooting.

## Community

💬 [Join the Borzoi Discord server](https://discord.gg/3HVsj5pTFP)

## Contributing

Want to contribute? Great! Please check the [Contribution
Guidelines](CONTRIBUTING.md) for details.

## Acknowledgements

- Thanks to the [Hacker101 community on Discord](https://www.hacker101.com/discord)
  for the encouragement and early feedback.
- The font used in the logo and admin interface is [JetBrains
  Mono](https://www.jetbrains.com/lp/mono/).

## Sponsors

<a href="https://www.tines.com/?utm_source=oss&utm_medium=sponsorship&utm_campaign=borzoi">
<img src="https://borzoi.xyz/img/tines-sponsorship-badge.png" width="140" alt="Sponsored by Tines">
</a>

## License

[MIT](LICENSE)

© 2022 Borzoi Software
