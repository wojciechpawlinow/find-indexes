# find-indexes

The aim of this project is to create working HTTP REST API service with just one endpoint.

## Description

For a given `value` search for all indexes in a given `slice` where `slice[index] == value`.  

Load the _input.txt_ file (containing sorted numbers from 0 to 1000000) into a slice once service starts.  

In case service is not able to find the index for a given value, it can return an index for any other existing values,
assuming that conformation is at 10% level.

## Built with

- [gin](https://github.com/gin-gonic/gin) - The web framework used
- [viper](https://github.com/spf13/viper) - Configuration management
- [di](https://github.com/sarulabs/di) - Dependency injection framework
- [zap](https://pkg.go.dev/go.uber.org/zap) - Logging

## Usage

```bash
git clone https://github.com/wojciechpawlinow/find-indexes.git
cd find-indexes
go run cmd/server/main.go # or `make run`
```

To find the index, i.e:
```curl
curl localhost:8080/index/1150
```

### Sample output
```
❯ curl localhost:8080/index/0
{"index":0,"match":"exact","value":0}

❯ curl localhost:8080/index/100                                                                                                                                                                                                                              
{"index":1,"match":"exact","value":100}

❯ curl localhost:8080/index/150
{"error":"value not found"}

❯ curl localhost:8080/index/1150
{"index":12,"match":"nearest","value":1200}

❯ curl localhost:8080/index/1200
{"index":12,"match":"exact","value":1200}

❯ curl localhost:8080/index/-1
{"error":"value must be greater than zero"}

❯ curl localhost:8080/index/"asda"
{"error":"value must be an integer"}
```

## Unit tests

```bash
make test
```

## Lint

```bash
make format
```

## Configuration

Configuration is loaded from _config.yaml_ file.

```yaml
port: 8080
log_level: debug
```