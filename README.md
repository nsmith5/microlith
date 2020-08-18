# Microlith

Golang architecture that can be monolithic or microservice deployed. Check
out the article [here](https://www.nfsmith.ca/articles/monolith_by_default/)
to learn why you might want to do that.

To compile the code run

```shell
$ go build
```

To run in monolith mode run

```shell
$ ./microlith
```

and test out the service with curl as follows:

```shell
$ curl localhost:8080/walk
Oh yeahh you can really talk
$ curl localhost:8080/dance
Uh oh, looks like you can never dance
```

To run the permissions part in a microservice run

```shell
$ ./microlith permissions & // Binds to localhost:8081
$ ./microlith -perms-url http://locahost:8081
```