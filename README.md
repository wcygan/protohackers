# Protohackers

Solutions to the [Protohackers Challenges](https://protohackers.com/).

## Deployment

Run `main.go`  in any of the challenges directories.

```bash
go run main.go
```

## Testing

Use `telnet` to connect to the server.

```bash
telnet hostname 8080
```

## Server Hosting

https://protohackers.com/faq

> Do I have to host the server myself?

I use a [Virtual Machine on Azure](https://azure.microsoft.com/en-us/products/virtual-machines).

To host the server for this project, I've [installed Go](https://go.dev/dl/) on the VM, I've cloned this
repository, and I've run `go run main.go` in the challenge directory.

Once the server is running, I've exposed the port `8080` to the internet.

After this, I'm able to use `telnet` to connect to the server & run the solution checker against my server.