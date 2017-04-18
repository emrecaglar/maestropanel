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

func (m *SSL) GetSSLCert(domainName string) (result SSLResult, err error) {
	result = SSLResult{}

	extra := struct {
		Name string `json:"name"`
	}{
		domainName,
	}

	response, err := m.mp.writeData(getSSLAction.Method, m.mp.getURL(getSSLAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *SSL) DeleteSSLRequest(domainName string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct {
		Name string `json:"name"`
	}{
		domainName,
	}

	response, err := m.mp.writeData(deleteSSLRequestAction.Method, m.mp.getURL(deleteSSLRequestAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *SSL) CompleteSSLRequest(domainName string, responseCertificate string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct {
		Name                string `json:"name"`
		ResponseCertificate string `json:"responseCertificate"`
	}{
		domainName,
		responseCertificate,
	}

	response, err := m.mp.writeData(completeSSLRequestAction.Method, m.mp.getURL(completeSSLRequestAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *SSL) DeleteSSLCert(domainName string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct {
		Name string `json:"name"`
	}{
		domainName,
	}

	response, err := m.mp.writeData(deleteSSLCertAction.Method, m.mp.getURL(deleteSSLCertAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *SSL) ImportPfxFile(domainName string, pfxPassword string, pfxFilePath string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		FilePath string `json:"pfx" type:"filepath"`
	}{
		domainName,
		pfxPassword,
		pfxFilePath,
	}

	response, err := m.mp.writeData(importPfxFileAction.Method, m.mp.getURL(importPfxFileAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}
