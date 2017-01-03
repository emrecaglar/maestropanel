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