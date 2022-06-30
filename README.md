# Borzoi

A basic customer management system.

## Resources

- Registration with username/password, social networks, mobile number, etc.
- Manage your profile by adding photo, description, social networks, etc.
- Add your customers with only basic information
- Add notes to your customers, like a diary
- Import and export information in CSV, PDF, etc.
- Create QRCode and share

## Installation

There are three ways to run Borzoi:

1. Direct download of [releases](https://github.com/elga-io/borzoi/releases).

2. Use **make** to create your own binary:

    ```bash
    # You need Go installed in your system.
    # More info: https://go.dev/dl/
    make build
    ```

3. Or create a Docker image and run it:

    ```bash
    make docker
    make docker-run
    ```
    And go to http://localhost to use the system.

## Usage

Once the binary is created, just run the command **borzoi** and access port 8080 (default):

```bash
./borzoi
```

Go to http://localhost:8080 and be happy.

## Development

### Requirements

- Go 1.18+
- npm 8.5.4

### Run

- Server: run `make run` to start server API.
- Web UI: run `make run-web` to start frontend web UI.

## License

[MIT](LICENSE)
