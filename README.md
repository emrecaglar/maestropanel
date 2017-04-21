# MaestroPanel


[![wercker status](https://app.wercker.com/status/b25da712119d17c7ef50d8e918f1413c/s/master "wercker status")](https://app.wercker.com/project/byKey/b25da712119d17c7ef50d8e918f1413c)

## MaestroPanel Go API Client Package

This package allows you to use the functions on the MaestroPanel API service with golang.


## Install Package

```batch
go get github.com/emrecaglar/maestroPanel
```

## Usage


```go
    m := maestroPanel.MaestroPanel{
            Url:        "maestropanel url (http://domain.com:9715)", 
            Key:        "api key",
            Version:    "api version (v1)"
    }
```

[Follow for create api key](https://wiki.maestropanel.com/maestropanelde-api-anahtari-olusturma/)

## Basic Example
```go
package main

import (
    "fmt"
    "github.com/emrecaglar/maestroPanel"
)

func main(){
    m := maestroPanel.MaestroPanel{
            Url:        "http://domain.com:9715", 
            Key:        "xxxxxx",
            Version:    "v1"
    }

    domain := maestroPanel.Domain{
        Name:       "domain.com",
        UserName:   "admin",
        Password:   "111111",
        PlanAlias:  "default",
    }

    web := m.Web()

    result, _ := web.CreateDomain(domain)

    fmt.Println(result)
}
```

## Resources 
[Other API Client Libraries and API License](https://wiki.maestropanel.com/api-dokumantasyonu-ve-ornek-kodlar/)

[API Documentation](https://docs.google.com/document/d/1rmXwq6gx6E6LbCkhRuzXk_6v998R018cN72oAw9_vYs/edit)
