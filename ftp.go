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
    } {
        domainName,
        account,
    }

    response, err := m.mp.writeData(deleteFtpAccountAction.Method, m.mp.getURL(deleteFtpAccountAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *FTP) ChangeFTPAccountPassword(model ChangeFTPAccountPassword)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(changeFtpAccountPasswordAction.Method,m.mp.getURL(changeFtpAccountPasswordAction),model)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *FTP) GetFTPAccounts(domainName string)(result GetFTPAccountsResult, err error)  {
    result = GetFTPAccountsResult{}

    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(getFtpAccountsAction.Method, m.mp.getURL(getFtpAccountsAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}