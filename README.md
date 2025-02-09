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
  
Example of config for app:
env :
LOCAL            = "local" starts without ngrok and auth middleware
LOCAL_WITH_NGROK = "local_with_ngrok"
DEV              = "dev"
PROD             = "prod"


{
    "env" : "local_with_ngrok",
    "database": {
      "name": "coinder",
      "user":"postgres",
      "password": "GK986vBBdOOkjh08Bvzzz",
      "host": "127.0.0.1",
      "port":"5432"
    },
    "parser": {
        "cmc_apikey":"adb5310b-ece6-40c1-9904-caa8f3cc704e",
        "timestamp": "24h",
        "timeout_for_req": "10s"
    }
}
  