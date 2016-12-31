# MaestroPanelGoApi
Golang MaestroPanel Api

[![wercker status](https://app.wercker.com/status/b25da712119d17c7ef50d8e918f1413c/s/master "wercker status")](https://app.wercker.com/project/byKey/b25da712119d17c7ef50d8e918f1413c)

## Examples

Create Domain Example

```go
    m := maestropanel.MaestroPanel{
            "maestropanel url (http://domain.com:9715)", 
            "authorization key",
            "api version (v1)"
        }

    domain := maestropanel.Domain{}
    domain.Name = "domain.com"
    domain.UserName = "admin"
    domain.Password = "111111"
    domain.PlanAlias = "default"

    web := m.Web()

    result, _ := web.CreateDomain(domain)
```