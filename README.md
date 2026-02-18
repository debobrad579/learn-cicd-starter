![Tests](https://github.com/debobrad579/notely/actions/workflows/ci.yml/badge.svg)

# Notely

A note-taking REST API built in Go, created as part of the boot.dev CI/CD course. The main focus of this project was setting up a full CI/CD pipeline with GitHub Actions, automated testing, security scanning, and deployment to Google Cloud Platform.

## CI/CD Pipeline

**CI** runs on every pull request to `main` and includes:
- Unit tests with coverage (`go test -cover ./...`)
- Security scanning via `gosec`
- Format checking via `go fmt`
- Linting via `staticcheck`

**CD** runs on every push to `main` and:
1. Builds the binary via `./scripts/buildprod.sh`
2. Builds and pushes a Docker image to Google Artifact Registry
3. Runs database migrations via `./scripts/migrateup.sh`
4. Deploys to Google Cloud Run

## License

This project is part of the boot.dev curriculum.
