# full-cycle-sonarcloud

Exercício do curso **Full Cycle** sobre análise de qualidade de código com **SonarCloud** e cobertura de testes em Go.

## O que tem aqui

- `soma.go`: funções simples de soma e subtração.
- `sum_test.go`: testes com o pacote `testing`.
- [`ci.yaml`](.github/workflows/ci.yaml): workflow que roda em pull requests para `develop` e gera o relatório de cobertura (`coverage.out`).
- `sonar-project.properties`: configuração do SonarCloud, que separa código de testes e lê o relatório de cobertura do Go.

## Como rodar localmente

```bash
go mod init mysum
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Tecnologias

Go · GitHub Actions · SonarCloud
