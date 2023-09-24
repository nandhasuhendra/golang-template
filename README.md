# Golang Template

This is my personal template for golang projects. Feel Free for anyone want to use this template

# Run docker compose for development

This command will using existing network. You may can change the network base on what you need
`docker-compose -f docker-compose.dev.yml up`

# Run migration

Create new migration
`bin/migrate -c users`

Up migration
`bin/migrate -u`

Down migration
`bin/migrate -d MIGRATION_VERSION`

Force run migration
`bin/migrate -f MIGRATION_VERSION`

# Run project

`go run main.go`
