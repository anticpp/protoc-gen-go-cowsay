Demonstrate creating a protoc plugin(`go-cowsay`) to generate Go code with `protogen`. Every message has a method named `CowSay()`, which prints a `cowsay` message.

```
./bin/cowsay
_____________
< Hello World! >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## install

Install `protoc-gen-go-cowsay` with `go install`.

## gen

Generate code in `api/example/`.

## build

Build all packages output to `bin/`.

## TODO

- How to rename package name to `cowsay` from `v2`?
