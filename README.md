# Hooktail

<img src="https://www.mariowiki.com/images/thumb/d/dd/Hooktail_Artwork_-_Paper_Mario_The_Thousand-Year_Door.png/1254px-Hooktail_Artwork_-_Paper_Mario_The_Thousand-Year_Door.png" width="150">

Hooktail is a server written in go that can be used for github webhook deployments.

## REQUIREMENTS

- [Go](https://golang.org/) >= **1.13**

## CONFIGURATION

First generate an SSL certificate:

```bash
openssl req -newkey rsa:4096 \
  -nodes -keyout key.pem \
  -x509 -days 365 \
  -out cert.pem
```

Copy the sample configuration file and edit as per your liking:

```bash
cp config.example.yml config.yml
```

The available configuration is listed in the **config.example.yml** file.

## BUILD & RUN

You can build and run the server using the following commands:

```bash
go build && sudo ./hooktail
```

**NOTE:** **sudo** is required in order to be able to run a deployment as a
specific user.

In order to specify a custom configuration file (other than **config.yml**)
you can use the following flag:

```bash
./hooktail -config <path-to-config.yml>
```