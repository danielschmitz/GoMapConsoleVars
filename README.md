# mapconsolevars

Read Console Vars and save to a map

## Example:

When you run: `go run *.go key1=value1 key2=value2`

And need the value of a key:

```go
value1, found := mapconsolevars.GetArgsValue("key1")
if found {
    fmt.Println("Value of key1 is", value1)
}
```

## Install


```bash
$  go get github.com/danielschmitz/mapconsolevars@latest
```

```go
import "github.com/danielschmitz/mapconsolevars"
```

## Bonus: Add params to a VSCODE Run&Debug

- Press F5, choice Go, and add this .code/launch.json file

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Run",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceRoot}",
            "args": ["key1=value1", "key2=value2"]
        }
    ]
}

