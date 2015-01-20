# reverse
Go (golang) url reverse


Example using Goji router:

```go
package main

import (
        "fmt"
        "net/http"
        "github.com/alehano/reverse"
        "github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
        // We can get reversed url by name and a list of params:
        // reverse.Urls.MustReverse("UrlName", "param1", "param2")

        fmt.Fprintf(w, "Hello, %s", reverse.Urls.MustReverse("helloUrl", c.URLParams["name"]))
}

func main() {
        // Set an Url and Params and return raw url to router
        // reverse.Urls.MustAdd("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")

        goji.Get(reverse.Urls.MustAdd("helloUrl", "/hello/:name", ":name"), hello)
        goji.Serve()
}
```
