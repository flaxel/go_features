# Vendoring

Vendoring is the act of making your own copy of the third party packages your project is using. Those copies are traditionally placed inside each project and then saved in the project repository. 

## Initialize module

Create a new `go.mod` file for the new module. 

```
go mod init
```

## Create Copies

Create `vendor` folder with the copies of the dependencies and the `go.sum` file with the expected cryptographic checksums.

```
go mod vendor
```

## Cleaning

It is possible to add missing and remove unused modules.

```
go mod tidy
```

At the end you can check the dependencies of the current module.

```
go mod verify
```
