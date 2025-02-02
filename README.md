starting parser:
flag --config is optional
go run cmd\parser\main.go --config path/to/config.json

Example of config for parser
{
    "cmc_apikey":"apikey",
    "database": {
      "name": "coinder",
      "user":"postgres",
      "password": "password",
      "host": "127.0.0.1",
      "port":"5432"
    }
}
  