package maestropanel

import (
    "encoding/json"
)

type Mail struct {
    mp MaestroPanel
}

func (m *Mail) AddMailBox(mailBox MailBox)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(addMailBoxAction.Method,m.mp.getURL(addMailBoxAction), mailBox)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *Mail) DeleteMailBox(domainName string, account string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    extra := struct {
        Name string `json:"name"`
        Account string `json:"account"`
    } {
        domainName,
        account,
    }

    response, err := m.mp.writeData(deleteMailBoxAction.Method,m.mp.getURL(deleteMailBoxAction),extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

