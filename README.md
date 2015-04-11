# reverse
Go (golang) URL reverse

Simple URL reverse package. It's useful for templates. You can get a URL by a name and params and not depend on URL structure.

It fits to any router. All it does is just stores urls by a name and replace params when you retrieve a URL.
To use it you have to add a URL with a name, raw URL with placeholders (params) and a list of these params.

```go
// To set a URL and return raw URL use:
reverse.Add("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")
// OUT: "/url_path/:param1/:param2"

// To retrieve a URL by name with given params use:
reverse.Rev("UrlName", "value1", "value2")
// OUT: "/url_path/value1/value2"

// Note, that these funcs panic if errors. Instead you can use Urls.Add() and Urls.Reverse() 
// that return errors. Or you can make your own wrapper for them.
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
        // reverse.Rev("UrlName", "value1", "value2")

        fmt.Fprintf(w, "Hello, %s", reverse.Rev("HelloUrl", c.URLParams["name"]))
}

func main() {
        // Set a URL and Params and return raw URL to a router
        // reverse.Add("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")

        goji.Get(reverse.Add("HelloUrl", "/hello/:name", ":name"), hello)
        
        // In regexp instead of: re := regexp.MustCompile("^/comment/(?P<id>\\d+)$")
        re := regexp.MustCompile(reverse.Add("DeleteCommentUrl", "^/comment/(?P<id>\\d+)$", "(?P<id>\\d+)$"))
        goji.Delete(re, deleteComment)
        
        goji.Serve()
}
```

Example for Gorilla Mux

```go

// Original set: r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
r.HandleFunc(reverse.Add("ArticleCatUrl", "/articles/{category}/{id:[0-9]+}", "{category}", "{id:[0-9]+}"), ArticleHandler)

// So, if we want to retrieve URL "/articles/news/123", we call:
fmt.Println( reverse.Rev("ArticleCatUrl", "news", "123") )

```