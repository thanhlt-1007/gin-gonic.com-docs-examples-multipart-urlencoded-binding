# gin-gonic.com-docs-examples-multipart-urlencoded-binding

- Multipart/Urlencoded binding

- Reference: https://gin-gonic.com/docs/examples/multipart-urlencoded-binding/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

- login

```sh
curl --location 'localhost:8080/login' \
--form 'user="admin"' \
--form 'password="Aa@123456"'
```
