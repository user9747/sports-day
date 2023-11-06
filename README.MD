# Sports Day

## Getting started

### Start dev server

```bash
docker-compose up
```

### Create new table

```bash
make create-table <table-name>
```

### Run migration

```make
make migrate
```

### Run tests

```bash
go test -race ./...
```

## API doc

https://documenter.getpostman.com/view/14156563/2s9YXe8jFx

## Folder structure

- **router** - All API routes
- **db** - Db init and migrations
- **cache** - redis utility
- **conf** - configs for db, cache, server...etc
- **internal/entity** - All entity types and creation
- **internal/repository** - DB layer
- **internal/usecases** - All business logic on entities
- **internal/requesthandler** - Handlers for API requests
- **internal/utils** - utility functions
