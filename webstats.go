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