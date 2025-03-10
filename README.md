starting parser:
flag --config is optional
go run cmd\parser\main.go --config path/to/config.json

Example of config for parser
```
  {
      "cmc_apikey":"apikey",
      "database": {
        "name": "coinder",
        "user":"postgres",
        "password": "password",
        "host": "database",
        "port":"5432"
      }
  }

```
  
Example of config for app:
env :
LOCAL            = "local" starts without ngrok and auth middleware
LOCAL_WITH_NGROK = "local_with_ngrok"
DEV              = "dev"
PROD             = "prod"

```
  {
      "env" : "local_with_ngrok",
      "database": {
        "name": "coinder",
        "user":"postgres",
        "password": "password",
        "host": "database",
        "port":"5432"
      },
      "parser": {
          "cmc_apikey":"apikey",
          "timestamp": "24h",
          "timeout_for_req": "10s"
      }
  }
```
  