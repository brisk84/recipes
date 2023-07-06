# Recipes [Backend Application] ![GO][go-badge]

[go-badge]: https://img.shields.io/github/go-mod/go-version/p12s/furniture-store?style=plastic
[go-url]: https://github.com/p12s/furniture-store/blob/master/go.mod

## Build & Run (Locally)
### Prerequisites
- go 1.20
- docker & docker-compose
- [golangci-lint](https://github.com/golangci/golangci-lint) (<i>optional</i>, used to run code checks)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) (<i>optional</i>, used to re-generate server)

Use `make up` to build and run project, `make lint` to check code with linter.