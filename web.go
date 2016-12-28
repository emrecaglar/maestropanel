package maestropanel

import (
    "encoding/json"
)

type Web struct {
    mp MaestroPanel
}

func (m *Web) GetList() (result GetListResult, err error) {
	result = GetListResult{}

	response, err := m.mp.writeData(getListAction.Method, m.mp.getURL(getListAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}
	
	return
}

func (m *Web) CreateDomain(domain Domain) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(createDomainAction.Method, m.mp.getURL(createDomainAction), domain)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) DeleteDomain(domainName string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct{
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(deleteDomainAction.Method, m.mp.getURL(deleteDomainAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) StopDomain(domainName string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct{
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(stopDomainAction.Method, m.mp.getURL(stopDomainAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) Password(password DomainPassword) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(passwordChangeDomainAction.Method, m.mp.getURL(passwordChangeDomainAction), password)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) AddDomainAlias(alias DomainAlias) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(addDominAliasAction.Method, m.mp.getURL(addDominAliasAction), alias)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) DeleteDomainAlias(alias DomainAlias) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(deleteDominAliasAction.Method, m.mp.getURL(deleteDominAliasAction), alias)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) GetDomainAliases(domainName string) (result GetDomainAliasResult, err error) {
	result = GetDomainAliasResult{}

	extra := struct {
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(getDomainAliasesAction.Method, m.mp.getURL(getDomainAliasesAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) AddSubDomain(subdomain SubDomain) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(addSubdomainAction.Method, m.mp.getURL(addSubdomainAction), subdomain)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) DeleteSubDomain(domainName string, subDomain string) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	extra := struct {
		Name string `json:"name"`
		SubDomain string `json:"subdomain"`
	} {
		domainName,
		subDomain,
	}

	response, err := m.mp.writeData(deleteSubdomainAction.Method, m.mp.getURL(deleteSubdomainAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) GetSubDomains(domainName string) (result GetSubDomainsResult, err error) {
	result = GetSubDomainsResult{}

	extra := struct {
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(subDomainsAction.Method, m.mp.getURL(subDomainsAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) SetSubDomainFTPAccount(ftpAccount SetSubDomainFTPAccount) (result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(setSubDomainFTPAccountAction.Method, m.mp.getURL(setSubDomainFTPAccountAction), ftpAccount)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}