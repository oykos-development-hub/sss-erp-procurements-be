# Procurements API

## Description

This project is a RESTful API developed in Go that provides public procurements operations.

## Installation

### Prerequisites

- Go version 1.20 or higher
- Docker

### Clone the repository

```bash
git clone git@gitlab.sudovi.me:erp/procurements-api.git
```

### Environment Setup

Copy the .env-example file to create a .env that will contain your environment-specific configurations.

```bash
cp .env-example .env
```

After copying the example environment file, you'll need to open .env and fill in the necessary details such as your database connection information.

## Usage

Procurements API is RESTful API built using Go.

### Running the Service

From the root directory, run the following commands to start the service:

```bash
docker-compose up -d
make start
```

The service will be available on port that is specified in .env file.

### Interacting with the RESTful API

Access the REST API endpoint at /api. For testing and exploring the API, we recommend using Postman.

### Postman collection

You can easily set up Postman to work with our API by clicking the "Run in Postman" button below:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/27680596-8e47dce6-074e-4ad0-aca2-38c818ad30b5?action=collection%2Ffork&collection-url=entityId%3D27680596-8e47dce6-074e-4ad0-aca2-38c818ad30b5%26entityType%3Dcollection%26workspaceId%3De27156de-19d4-4c8a-ae2d-9b845d9ea484)

- Go version 1.20 or higher
- Docker

## Contributing

Please follow the guidelines below for branch naming and commits.

### Commit and branch naming

#### Branch naming convention

1. state the type of change you are making: `build, fix, refactor, feat`
2. add forward slash `/`
3. state the task ID (if applicable) - TSK-123
4. add forward slash `/`
5. add 2-3 words separated by `-` that describe changes you are making
6. Example: `fix/TSK-123/fixing-border-radius`

#### Commit & Push

We use the same convention as for Branch naming.

Only difference is that we use `:` instead of `/` in the commit message. And we describe in the message what we did without `-` between words.

Example: `fix: changed border radius from 4px to 2px`

## License

This project is licensed under the [Internal License](LICENCE).

## Credits

- [Unidoc v3](github.com/unidoc/unipdf/v3)
