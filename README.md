# Hooktail

[![CodeFactor](https://www.codefactor.io/repository/github/klipitkas/hooktail/badge?s=afe43e6de344ceb5e154110b38cbecf8dcbf3568)](https://www.codefactor.io/repository/github/klipitkas/hooktail)

---

<img src="https://www.nicepng.com/png/full/133-1332501_502px-hooktail-artwork-paper-mario-hooktail.png" width="150">

Hooktail is an HTTP server written in Go that can be used for github webhook
deployments.

## REQUIREMENTS

- [Go](https://golang.org/) >= **1.13**

## CONFIGURATION

Copy the sample configuration file and edit as per your liking:

```
cp config.example.yml config.yml
```

The available configuration is listed in the **config.example.yml** file.

## RUN THE TESTS

You can run all the available unit tests using the command below:

```
go test -v ./...
```

## BUILD & RUN

You can build and run the server using the following commands:

```
go build && sudo ./hooktail
```

**NOTE:** **sudo** is required in order to be able to run a deployment as a
specific user.

In order to specify a custom configuration file (other than **config.yml**)
you can use the following flag:

```
sudo ./hooktail -config <path-to-config.yml>
```

## TLS / SSL SUPPORT

Since **Hooktail** only supports HTTP, it cannot handle SSL termination. In
order to handle SSL termination you need a reverse proxy such as:
- [Nginx](https://www.nginx.com)
- [Haproxy](https://www.haproxy.org)