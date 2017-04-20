# MaestroPanelGoApi
Golang MaestroPanel Api

[![wercker status](https://app.wercker.com/status/b25da712119d17c7ef50d8e918f1413c/s/master "wercker status")](https://app.wercker.com/project/byKey/b25da712119d17c7ef50d8e918f1413c)

## MaestroPanel REST Web Service Arayüzü
MaestroPanel REST* API, Domain, Bayi (Reseller) ve Sunucu özelliklerine HTTP üzerinden

belirli kurallar çerçevesinden erişebileceğiniz bir programlama arayüzüdür.

Api, REST (Representational State Transfer​) olarak çalıştığından herhangi bir programlama

diline ihtiyacınız olmadan herhangi bir HTTP istemcisi ile (örneğin browser'ınız) rahatlıkla

komutlar gönderebilir ve kendi geliştirdiğiniz yazılımlarla erişebilirsiniz.

MaestroPanel API kendi iş akışınıza MaestroPanel'i entegre etmenizi kolayca sağlamaktadır.

    REST kavramı ile ilgili daha detaylı bilgi almak için:

    http://en.wikipedia.org/wiki/Representational_state_transfer


## Examples

Create Domain Example

```go
    m := maestropanel.MaestroPanel{
            Url:        "maestropanel url (http://domain.com:9715)", 
            Key:        "authorization key",
            Version:    "api version (v1)"
        }

    domain := maestropanel.Domain{
        Name:       "domain.com",
        UserName:   "admin",
        Password:   "111111",
        PlanAlias:  "default",
    }

    web := m.Web()

    result, _ := web.CreateDomain(domain)
```