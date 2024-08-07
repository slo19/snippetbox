## Run

In root folder:

```
$ go run ./cmd/web
```

## DB

Using docker image of ~MYSQL~ MariaDB;
 - [x] [installation](https://plus.diolinux.com.br/t/dica-como-instalar-o-mysql-no-docker-e-conectar-no-dbeaver-atualizado-08-2022/47274)
 - [x] [connect to docker by terminal](https://baumannalexj.medium.com/connect-your-db-tool-to-a-dockerized-mysql-server-container-bc18853524ed)
 - ~[ ] [find and run a container in host]()~
 - ~[ ] [conectar ao banco no docker]()~
 - [x] connect to a local db
 - [x] import and use [alice package](https://github.com/justinas/alice) for middleware chaining
 - [X] use http router [](https://github.com/julienschmidt/httprouter")
 - [X] use form validator [go-playground](https://github.com/go-playground/form/v4)
 - [X] use session manager [scs](https://github.com/alexedwards/scs)
 - [X] use https (generating a TLS certificate)




### *Generating a TLS certificate with crypto/tls*

```$ go run <path_to_go_library>/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost```
