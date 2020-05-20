# WASM Test




# Run
```
GOOS=js GOARCH=wasm go build -o main.wasm && cp main.wasm ./dist  
```
then open your site in dist path.


# Trouble shooting

## can't load package: named files must be .go files: -o
因为没有 go.mod
```
go mod init github.com/lumixraku/wasmtest
```

## "could not import syscall/js" in vscode
https://github.com/microsoft/vscode-go/issues/1874

Add this in vscode settings.json

```
{
    "go.toolsEnvVars": {
        "GOARCH":"wasm",
        "GOOS":"js",
    },
    "go.testEnvVars": {
        "GOARCH":"wasm",
        "GOOS":"js",
    },
    "go.installDependenciesWhenBuilding": false,
}

```

## Read More
https://www.youtube.com/watch?v=4kBvvk2Bzis
https://tutorialedge.net/golang/go-webassembly-tutorial/  这个教程已经旧了

https://www.aaron-powell.com/posts/2019-02-06-golang-wasm-3-interacting-with-js-from-go/  比较新的教程

https://github.com/mattn/golang-wasm-example/   比较新的例子
