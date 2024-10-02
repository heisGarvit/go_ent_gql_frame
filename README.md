# go_ent_gql_frame

## Install dependencies
```curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.24.0 sh```

## Generate DB migrations (Local only)
```./migration.sh```
```./upgrade.sh```

## Apply DB migrations (Prod)
```./upgrade.sh```

## Run the project
```go generate . && go run main.go```