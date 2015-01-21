# reverse
Go (golang) url reverse

Simple url reverse package. It fits to any router. All it does is just stores urls by a name and replace params when you retrieve a url.
To use it you have to add a url with a name, raw url with placeholders (params) and a list of these params.

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
        // We can get reversed url by it's name and a list of params:
        // reverse.Urls.MustReverse("UrlName", "param1", "param2")

        fmt.Fprintf(w, "Hello, %s", reverse.Urls.MustReverse("HelloUrl", c.URLParams["name"]))
}

func main() {
        // Set a Url and Params and return raw url to a router
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

// So, if we want to retrieve url "/articles/news/123", we call:
fmt.Println( reverse.Urls.MustReverse("ArticleCatUrl", "news", "123") )

```
