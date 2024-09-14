# go_ent_gql_frame

## Install dependencies
```curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.24.0 sh```

## Generate DB migrations (Local only)
```./migration.sh```

## Apply DB migrations (Local & Prod)
```./migrate.sh```
```./upgrade.sh```

## Run the project
```go generate . && go run main.go```