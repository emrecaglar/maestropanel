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

func (m *Mail) ChangeMailBoxPassword(domainName string, account string, password string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    extra := struct {
        Name string `json:"name"`
        Account string `json:"account"`
        Password string `json:"newpassword"`
    } {
        domainName,
        account,
        password,
    }

    response, err := m.mp.writeData(changeMailBoxPasswordAction.Method, m.mp.getURL(changeMailBoxPasswordAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *Mail) GetMailList(domainName string)(result GetMailListResult, err error)  {
    result = GetMailListResult{}

    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(getMailListAction.Method,m.mp.getURL(getMailListAction),extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

