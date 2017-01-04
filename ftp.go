package maestropanel

import (
    "encoding/json"
)

type FTP struct {
    mp MaestroPanel
}

func (m *FTP) AddFtpAccount(ftpAccount FTPAccount)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(addFtpAccountAction.Method, m.mp.getURL(addFtpAccountAction), ftpAccount)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *FTP) DeleteFtpAccount(domainName string, account string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    extra := struct {
        Name string `json:"name"`
        Account string `json:"account"`
    }

    response, err := m.mp.writeData(addFtpAccountAction.Method, m.mp.getURL(addFtpAccountAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

