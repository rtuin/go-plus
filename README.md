# go-plus

Provides a collection of HTTP helper functions to use with the Go language.

## Usage

```go
import (
    "net/http"
    "github.com/rtuin/go-plus"
)

func main() {
    response, _ := http.Get("http://www.google.com/")

    // IsHTTPServerError example
    fmt.Printf("IsHTTPServerError: %v", IsHTTPServerError(response)) // 
}
```

## Credits

- [Richard Tuin](https://github.com/rtuin)
- [All Contributors](https://github.com/rtuin/go-plus/contributors)

## License

The MIT License (MIT). Please see the [license file](https://github.com/rtuin/go-plus/blob/master/LICENSE) for more information.
