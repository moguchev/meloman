# meloman
Курсовая работа по БСБД

## Зависимости
Для установки и запуска приложения потребуется:
1) [docker](https://www.docker.com)
2) docker-compose
3) [go 1.16](https://golang.org)
4) make
5) [buf](https://github.com/bufbuild/buf) (optional)
6) [protobuf](https://developers.google.com/protocol-buffers/docs/downloads) (optional)
7) возможно что-то еще понадобится (optional)

## Установка
```
$ git clone https://github.com/moguchev/meloman.git && cd meloman
$ git clone https://github.com/googleapis/googleapis.git # optional
$ make deps && make build # optional
```
## Запуск
```
$ make up
```

## Работа
Документацию API проекта можно посмотреть в [браузере](http://localhost:8080/swaggerui/)  (__localhost:8080/swaggerui/__), там же можно повыполнять запросы к сервису
Сервис слушает два порта __:8080__ - HTTP и __:8090__ - gRPC

## Остановка
После того как вы поигрались и вам больше не нужен этот сервис выполните:
```
$ make down
```