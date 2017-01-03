package maestropanel

import (
    "encoding/json"
)

type WebStats struct {
    mp MaestroPanel
}

func (m *WebStats) ProtectStatsArea(model ProtectStatsArea)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(protectStatsAreaAction.Method, m.mp.getURL(protectStatsAreaAction), model)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *WebStats) UnProtectStatsArea(domainName string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}
    
    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(unProtectStatsArea.Method,m.mp.getURL(unProtectStatsArea), extra)

    if err == nil {
        err = json.Unmarshal(response, &result)
    }

    return
}

func (m *WebStats) EnableStatsProtection(domainName string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}
    
    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(enableStatsProtection.Method,m.mp.getURL(enableStatsProtection), extra)

    if err == nil {
        err = json.Unmarshal(response, &result)
    }

    return
}

func (m *WebStats) DisableStatsProtection(domainName string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}
    
    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(disableStatsProtection.Method,m.mp.getURL(disableStatsProtection), extra)

    if err == nil {
        err = json.Unmarshal(response, &result)
    }

    return
}

func (m *WebStats) SetFtpUserStatsArea(domainName string, ftpUser string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}
    test = ""
    extra := struct {
        Name string `json:"name"`
        FTPUser string `json:"ftpuser"`
    } {
        domainName,
        ftpUser,
    }

    response, err := m.mp.writeData(setFtpUserStatsArea.Method,m.mp.getURL(setFtpUserStatsArea), extra)

    if err == nil {
        err = json.Unmarshal(response, &result)
    }

    return
}


