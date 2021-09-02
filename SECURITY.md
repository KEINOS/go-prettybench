# Security Policy

## Supported Versions

We test with the following Go versions.

| Version | Supported          | Docker Image |
| ------- | ------------------ | :--: |
| 1.14  | :white_check_mark: | golang:1.14-alpine |
| 1.15  | :white_check_mark: | golang:1.15-alpine |
| 1.16  | :white_check_mark: | golang:1.16-alpine |
| 1.17  | :white_check_mark: | golang:1.17-alpine |

## Security Analysis

As a basic manner, we run the below automated checks on any push to the repo and weekly-run as well.

- Scurity analysis: [CodeQL for Go](https://codeql.github.com/docs/codeql-language-guides/codeql-for-go/)
- Vulnerability check: [GitHub Dependabot](https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/about-dependabot-version-updates)

## Testing Policy

Besides to try keep high coverage persentage of the unit tests, we use [GolangCI-Lint](https://golangci-lint.run/) for static analysis and lint checking.

- See configuration: [./.golangci.yml](https://github.com/KEINOS/go-prettybench/.golangci.yml)

## Reporting a Vulnerability

Please report any vulnerabilities to the [issues](https://github.com/KEINOS/go-prettybench/issues) page. Issues with reproducible examples helps us a lot.
