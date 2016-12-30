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

func (m *Web) ChangeIpAddr(ipaddres ChangeIpAddres)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(changeIPAddressAction.Method, m.mp.getURL(changeIPAddressAction), ipaddres)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) GetListItem(domainName string)(result DomainListItemResult, err error) {
	result = DomainListItemResult{}

	extra := struct {
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(getDomainListItemAction.Method, m.mp.getURL(getDomainListItemAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) GetLimits(domainName string)(result LimitResult, err error) {
	result = LimitResult{}

	extra := struct {
		Name string `json:"name"`
	} {
		domainName,
	}

	response, err := m.mp.writeData(getLimitsAction.Method, m.mp.getURL(getLimitsAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) Forwarding(forward Forward)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(forwardingAction.Method, m.mp.getURL(forwardingAction), forward)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) ChangeReseller(reseller Reseller)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(changeResellerAction.Method, m.mp.getURL(changeResellerAction), reseller)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Web) SetDomainPlan(domainPlan DomainPlan)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(setDomainPlanAction.Method, m.mp.getURL(setDomainPlanAction), domainPlan)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}