/*
Package maestroPanel is a MaestroPanel API Client library

For Example:
package main

import (
	"fmt"
	"github.com/emrecaglar/maestropanel"
)

func main(){
	m := maestropanel.MaestroPanel{
            Url:        "http://domain.com:9715",
            Key:        "xxxxxx",
            Version:    "v1"
    }

    domain := maestropanel.Domain{
        Name:       "domain.com",
        UserName:   "admin",
        Password:   "111111",
        PlanAlias:  "default",
    }

    web := m.Web()

    result, _ := web.CreateDomain(domain)

    fmt.Println(result)
}

MastroPanel'in aldığı parametreler
Url: MaestroPanel API servisinin adres ve port bilgisi
Key: MaestroPanel üzerinde oluşturulan API anahtarı
Version: API versiyonu. Eğer versiyon belirtilmezse varsayılan olarak v1 kullanılacaktır
*/

package maestropanel
