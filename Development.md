## Run Go server

```
go run ./cmd/ace/*.go install
```

## Build UI

```
make build-ui
```

## API Docs

- `options.json`: http://localhost:4000/apis/options.json
- `schems.json`: http://localhost:4000/apis/schema.json

- POST options.json to http://localhost:4000/apis/install

## README Example:

https://github.com/tamalsaha/bb-installer-demo/blob/master/ace-installer/README.md

## Test Server

```
curl -X POST -d @ace-installer/options.yaml http://localhost:4000/apis/install
```
