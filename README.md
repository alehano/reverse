# reverse
Go (golang) URL reverse

Simple URL reverse package. It's useful for templates. You can get a URL by a name and params and not depend on URL structure.

It fits to any router. All it does is just stores urls by a name and replace params when you retrieve a URL.
To use it you have to add a URL with a name, raw URL with placeholders (params) and a list of these params.

```go
// To set a URL and return raw URL use:
reverse.Add("UrlName", "/url_path/:param1/:param2", ":param1", ":param2")
// OUT: "/url_path/:param1/:param2"

// To set a URL with group (subrouter) prefix and return URL without prefix use:
reverse.AddGr("UrlName", "/url_path", "/:param1/:param2", ":param1", ":param2")
// OUT: "/:param1/:param2"

// Note, that these funcs panic if errors. Instead you can use Urls.Add() and Urls.Reverse() 
// that return errors. Or you can make your own wrapper for them.

// To retrieve a URL by name with given params use:
reverse.Rev("UrlName", "value1", "value2")
// OUT: "/url_path/value1/value2"

// Get raw url by name
reverse.Get("UrlName")
// OUT: "/url_path/:param1/:param2"

// Get all url as map[string]string
reverse.GetAllUrls()

// Get all urls params as map[string][]string
reverse.GetAllParams()

```

Example for Gin router (https://github.com/gin-gonic/gin):

```go
func main() {
    router := gin.Default()
    
    // URL: "/"
    // To fetch the url use: reverse.Rev("home")
    router.GET(reverse.Add("home", "/"), indexEndpoint)
    
    // URL: "/get/123"
    // With param: c.Param("id")
    // To fetch the URL use: reverse.Rev("get_url", "123")
    router.GET(reverse.Add("get_url", "/get/:id"), getUrlEndpoint)
    

    // Simple group: v1 (each URL starts with /v1 prefix)
    groupName := "/v1"
    v1 := router.Group(groupName)
    {
        // URL: "/v1"
        // To fetch the URL use: reverse.Rev("v1_root")
        v1.GET(reverse.AddGr("v1_root", groupName, ""), v1RootEndpoint)
        
        // URL: "v1/read/cat123/id456"
        // With params (c.Param): catId, articleId
        // To fetch the URL use: reverse.Rev("v1_read", "123", "456")
        v1.GET(reverse.AddGr("v1_read", groupName, "/read/cat:catId/id:articleId", ":catId", ":articleId"), readEndpoint)

        // URL: /v1/login
        // To fetch the URL use: reverse.Rev("v1_login")
        v1.GET(reverse.AddGr("v1_login", groupName, "/login"), loginGetEndpoint)
        
    }

    router.Run(":8080")
}

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


Example subrouters for [Chi](https://github.com/go-chi/chi) router:

```go
// Original code
r.Route("/articles", func(r chi.Router) {
	r.Get("/", listArticles)
	r.Route("/{articleID}", func(r chi.Router) {
		r.Get("/", getArticle)
	})
})

// With reverse package
r.Route("/articles", func(r chi.Router) {
	r.Get(reverse.AddGr("list_articles", "/articles", "/"), listArticles)
	r.Route("/{articleID}", func(r chi.Router) {
		r.Get(reverse.AddGr("get_article", "/articles/{articleID}", "/", "{articleID}"), getArticle)
	})
})

// Get a reverse URL:
reverse.Rev("get_article", "123")
// Output: /articles/123/

// One more example (without tailing slashes)
r.Route(reverse.Add("admin.index", "/admin"), func(r chi.Router) {
	r.Get("/", index)

	r.Route(reverse.AddGr("admin.login", "/admin", "/login"), func(r chi.Router) {
		r.Get("/", login)
		r.Post("/", loginPost)
	})
})

```
