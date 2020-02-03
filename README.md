# go-mocker
A simple http mock server implemented by Golang.

## How to Use
### install

```shell script
go get github.com/amooly/go-mocker
```

### Start a Http Server
Firstly you need to create a json file to define the server config, just like this(You can also see the sample file [here](https://github.com/amooly/go-mocker/blob/master/sample/mock.json)):
```json
{
  "server": {
    "address": "0.0.0.0",
    "port": 8080
  },
  "routes": [
    {
      "request": {
        "method": "GET",
        "url": "/hello"
      },
      "response": {
        "status": 200,
        "body": "{\"message\": \"World!\"}"
      }
    }
  ]
}
```


After Then, run the command. The format of go-mocker command is :
```shell script
go-mocker [{path_to_mock_file}]
```
> You can define your own file using absolute path or relative path.
> Or it will use the [sample file](https://github.com/amooly/go-mocker/blob/master/sample/mock.json)


> Inspired by [go-http-mock-server](https://github.com/jaswdr/go-http-mock-server)
