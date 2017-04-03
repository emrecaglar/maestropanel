package maestropanel

import "encoding/json"

type SSL struct {
	mp MaestroPanel
}

func (m *SSL) CreateSSLRequest(createSSL CreateSSL) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(createSSLRequestAction.Method, m.mp.getURL(createSSLRequestAction), createSSL)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}
