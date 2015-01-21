# reverse
Go (golang) URL reverse

It's useful for templates. You can get a URL by a name and params and not depend on URL structure.

Simple URL reverse package. It fits to any router. All it does is just stores urls by a name and replace params when you retrieve a URL.
To use it you have to add a URL with a name, raw URL with placeholders (params) and a list of these params.

```go
// To set a URL and return raw URL use:
reverse.Urls.MustAdd("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")

// To retrieve a URL by name with given params use:
reverse.Urls.MustReverse("UrlName", "value1", "value2")

// OUT: "/url_path/value1/value2"
```

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
        // We can get reversed URL by it's name and a list of params:
        // reverse.Urls.MustReverse("UrlName", "value1", "value2")

        fmt.Fprintf(w, "Hello, %s", reverse.Urls.MustReverse("HelloUrl", c.URLParams["name"]))
}

func main() {
        // Set a URL and Params and return raw URL to a router
        // reverse.Urls.MustAdd("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")

        goji.Get(reverse.Urls.MustAdd("HelloUrl", "/hello/:name", ":name"), hello)
        
        // For regexp you can save a url separately and then get it as usual:
        // reverse.Urls.MustReverse("DeleteCommentUrl", "123")
        reverse.Urls.MustAdd("DeleteCommentUrl", "/comment/:id", ":id")
        re := regexp.MustCompile("^/comment/(?P<id>\d+)$")
        goji.Delete(re, deleteComment)
        
        goji.Serve()
}
```

Example for Gorilla Mux

```go

// Original set: r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
r.HandleFunc(reverse.Urls.MustAdd("ArticleCatUrl", "/articles/{category}/{id:[0-9]+}", "{category}", "{id:[0-9]+}"), ArticleHandler)

// So, if we want to retrieve URL "/articles/news/123", we call:
fmt.Println( reverse.Urls.MustReverse("ArticleCatUrl", "news", "123") )

```
