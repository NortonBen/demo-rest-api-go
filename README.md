#Demo API REST GOLANG - POSTGRESQL

## init
### install swagger gen api of fiber

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```
or
```sh
$ make init
```

## generate docs API
```sh
$ swag init
```
or
```sh
$ make gen
```

### Run app
```sh
$ go run apm
```

### Run API When ENV_GO not is product
```editorconfig
http://localhost:5010/swagger/index.html
```

### Run app run go-micro v3
#### run micro local
```sh
$ micro server
```

#### run app as dev
```sh
$ go run apm/micro/dev
```

#### run app as cloud
```sh
$ go run apm/micro
```
