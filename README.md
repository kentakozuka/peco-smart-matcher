# peco-smart-matcher

### build

```sh
go build -o peco-smart-matcher
```

### config.json

```json
  "CustomFilter": {
      "InvertMatch": {
          "Cmd": "/path/peco-smart-matcher",
        "Args": [ "$QUERY" ]
    }
  }
```

### How to use

if want,

`input == "hoge" && input != "foo"`

type,

`hoge !foo`
