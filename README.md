# wi18n

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wi18n?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wi18n)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wi18n)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wi18n/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wi18n)](https://goreportcard.com/report/github.com/wyy-go/wi18n)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wi18n?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wi18n/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wi18n)

## Usage

Download and install it:

```sh
go get github.com/wyy-go/wi18n
```

Import it in your code:

```go
import ginI18n "github.com/wyy-go/wi18n"
```

Canonical example:

```go
package main

import (
  "log"
  "net/http"

  "github.com/wyy-go/wi18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(wi18n.Localize(wi18n.WithBundle(&wi18n.Config{
     RootPath:         "./localize",
  })))

  router.GET("/", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage("welcome"))
  })

  router.GET("/:name", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage(&i18n.LocalizeConfig{
      MessageID: "welcomeWithName",
      TemplateData: map[string]string{
        "name": context.Param("name"),
      },
    }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Bundle

```go
package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/wyy-go/wi18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
  "golang.org/x/text/language"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(wi18n.Localize(wi18n.WithBundle(&wi18n.Config{
    RootPath:         "./_example/localizeJSON",
    AcceptLanguage:   []language.Tag{language.German, language.English},
    DefaultLanguage:  language.English,
    UnmarshalFunc:    json.Unmarshal,
    FormatBundleFile: "json",
  })))

  router.GET("/", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage("welcome"))
  })

  router.GET("/:name", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage(&i18n.LocalizeConfig{
      MessageID: "welcomeWithName",
      TemplateData: map[string]string{
        "name": context.Param("name"),
      },
    }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Get Language Handler

```go
package main

import (
  "log"
  "net/http"

  "github.com/wyy-go/wi18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(wi18n.Localize(
	  wi18n.WithGetLngHandle(
      func(context *gin.Context, defaultLng string) string {
        lng := context.Query("lng")
        if lng == "" {
          return defaultLng
        }
        return lng
      },
    ),
  ))

  router.GET("/", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage("welcome"))
  })

  router.GET("/:name", func(context *gin.Context) {
    context.String(http.StatusOK, wi18n.MustGetMessage(&i18n.LocalizeConfig{
      MessageID: "welcomeWithName",
      TemplateData: map[string]string{
        "name": context.Param("name"),
      },
    }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
