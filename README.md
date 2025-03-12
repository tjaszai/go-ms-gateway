# Go MS Gateway

The MSG provides tools for managing microservices.

## Requirements

* `Go` ( v1.24.0 )
* `Google wire`
  ```Bash
  go install github.com/google/wire/cmd/wire@latest
  ```
* `Swaggo`
  ```Bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

## In case of local development

* `Docker` ( v27.5.1+ )
* `Docker Compose` ( v2.32.4+ )

## Tech stack

* `Go` ( v1.24.0 )
* `Fiber` ( v2.52.6 )
* `Gorm` ( v1.25.12 )
* `PostgreSQL`

## Commands

* Run Wire ( compile-time Dependency Injection ): `cd ./cmd && wire`
* Generate API docs: `swag init --dir ./cmd,./internal --output ./docs`
