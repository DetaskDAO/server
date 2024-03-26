## Detask-Go Backend Project

```shell
├── cmd
├── config
├── internal
│   └── app
│       ├── api
│       ├── config
│       ├── core
│       ├── global
│       ├── initialize
│       ├── middleware
│       ├── model
│       ├── resource
│       ├── router
│       ├── service
│       ├── task
│       └── utils
```

| Directory    | instruction      | description                               |
|--------------|------------------|----------------------------------|
| `config`     | Config directory | Config directory                 |
| `internal`   | Internal         |                                  |
| `--app`      |                  |                                  |
| `api`        | HTTP API         | HTTP API                         |
| `config`     | Config struct    | Config struct              |
| `core`       | Core Files       | Initialization (zap, viper, server) |
| `global`     | Global objects   | Global objects                   |
| `initialize` | initialize       | router,redis,gorm,validator,timer |
| `middleware` | Middleware layer | `gin` middleware                 |
| `model`      | Model layer      | Model layer                      |
| `router`     | Router layer     | Router layer                     |
| `service`    | Service layer    | Service logic                    |
| `utils`      | Utilities        | Utilities                           |
| `timer`      | Timer            | timer                          |


## Cross-compilation
```shell
# Compile to Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
# Compile to macOS ARM64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build
# Compile to Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 
```
## Deployment

#### 1、Configuration file

```
cp ./config/config.demo.yaml ./config/config.yaml
vi ./config/config.yaml
```

```yaml
# system configuration
system:
  env: public  # Running Environment(public/develop)
  addr: 8888 # Running port

# pgsql configuration
pgsql:
  path: "127.0.0.1"   # Database ip
  port: "5432"        # Database port
  config: ""          # Advanced config
  db-name: "itom-admin" # Database name
  username: "postgres"  # Database user
  password: "123456"    # Database password
  max-idle-conns: 10    # Database idle connections
  max-open-conns: 100   # Database max connections
  log-mode: "info"      # Log mode (info/warn/error/slice)
  log-zap: false        # Use zap

# blockchain configuration
blockchain:
  - name: "mumbai"      # Name of the chain
    provider: ["https://rpc.ankr.com/polygon_mumbai"]  # Rpc provider

# contract configuration
contract:
  default-net: "mumbai"            # Default chain name
  signature: "Welcome to Detask!"  # Signature content
 
# jwt configuration
jwt:
  signing-key: "Detask"  # Signing key
  expires-time: 86400    # Expires in seconds
  issuer: "Detask"       # Issuer

# local configuration
local:
  path: 'uploads/file'   # Upload directory

# ipfs configuration
ipfs:
  - api: "https://ipfs.io/ipfs"                 # IPFS API
    upload-api: "http://192.168.1.10:3022/v1"   # Upload API
```

#### 2、Configuration Contract Address
```
vi ./abi/mumbai/DeTask.json     # enter the DeTask contract address
vi ./abi/mumbai/DeOrder.json    # enter the DeOrder contract address
```

#### 3、Run Project

```shell
./code-market-admin
```

#### 4、Import database

```
psql -p 5432 -h 127.0.0.1 -U web3 -d detask -f ./db/db.sql;
```