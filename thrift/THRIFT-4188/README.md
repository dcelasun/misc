# Bug test for THRIFT-4188

Clone the repo, `cd` to this directory and run:

```sh
$ go build -o bug
```

Then start the server:

```sh
$ ./bug -server
```

And run the client from another terminal:

```sh
$ ./bug
```
