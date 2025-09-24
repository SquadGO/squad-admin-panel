# Squad Admin Panel

Create `.env` file and see `.env.example`

## Production

DockerCompose: [Install](https://docs.docker.com/compose/install)

```bash
# Use make OR command
docker-compose -f docker-compose.prod.yml up --build -d
# OR
make prod
```

## Development

golangci-lint: [Install](https://golangci-lint.run/welcome/install/) [IDE Integration](https://golangci-lint.run/welcome/integrations/)\
sqlc: [Install](https://docs.sqlc.dev/en/latest/overview/install.html)\
dbmate: [Install](https://github.com/amacneil/dbmate)\
make: [MacOS](https://formulae.brew.sh/formula/make) [Windows](https://gnuwin32.sourceforge.net/packages/make.htm)

```bash
# Setup git hooks to run linters before commit
git config core.hooksPath .githooks
# for generate db/queries/*.sql
make sqlc-gen
# run
make dev-compose
make dev-apps # need wait for start compose, mb 30sec for first start!?
# see makefile for all commands
```

```bash
# create new migration
dbmate new create_users_table
```
