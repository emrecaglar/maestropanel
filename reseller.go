package maestropanel

import (
	"encoding/json"
)

type ResellerExecutionResult struct {
	Result
	Details ResellerOperationResult
}

type ResellerOperationResult struct {
	OperationResult
	ClientId   int32
	ClientName string
}

type ResellerModel struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	PlanAlias    string `json:"planAlias"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Country      string `json:"country"`
	Organization string `json:"organization"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postalCode"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
}

type ResellerDomain struct {
	UserName         string `json:"username"`
	Password         string `json:"password"`
	PlanAlias        string `json:"planAlias"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	DomainUserName   string `json:"domainUserName"`
	DomainPassword   string `json:"domainPassword"`
	ActiveDomainUser string `json:"activeDomainUser"`
	Email            string `json:"email"`
	Expiration       string `json:"expiration"`
	IPAddr           string `json:"ipaddr"`
}

type ResellerListItemResult struct {
	Result
	Details []ResellerListItem
}

type ResellerListItem struct {
	ApiAccess      bool
	Email          string
	ExpirationDate string
	FirstName      string
	Id             int32
	LastName       string
	LoginType      int32
	Organization   string
	Status         int32
	Username       string
}

type Reseller struct {
	mp MaestroPanel
}

func (m *Reseller) Create(reseller ResellerModel) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	response, err := m.mp.writeData(createResellerAction.Method, m.mp.getURL(createResellerAction), reseller)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) Delete(username string) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	extra := struct {
		UserName string `json:"username"`
	}{
		username,
	}

	response, err := m.mp.writeData(deleteResellerAction.Method, m.mp.getURL(deleteResellerAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) Stop(username string) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	extra := struct {
		UserName string `json:"username"`
	}{
		username,
	}

	response, err := m.mp.writeData(stopResellerAction.Method, m.mp.getURL(stopResellerAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) Start(username string) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	extra := struct {
		UserName string `json:"username"`
	}{
		username,
	}

	response, err := m.mp.writeData(startResellerAction.Method, m.mp.getURL(startResellerAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) ChangePassword(username string, newpassword string) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	extra := struct {
		UserName    string `json:"username"`
		NewPassword string `json:"newpassword"`
	}{
		username,
		newpassword,
	}

	response, err := m.mp.writeData(changePasswordResellerAction.Method, m.mp.getURL(changePasswordResellerAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) AddDomain(domain ResellerDomain) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	response, err := m.mp.writeData(addResellerDomainAction.Method, m.mp.getURL(addResellerDomainAction), domain)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) DeleteDomain(domainName string, username string) (result ResellerExecutionResult, err error) {
	result = ResellerExecutionResult{}

	extra := struct {
		DomainName string `json:"domainName"`
		UserName   string `json:"username"`
	}{
		domainName,
		username,
	}

	response, err := m.mp.writeData(deleteResellerDomainAction.Method, m.mp.getURL(deleteResellerDomainAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) GetDomains(username string) (result DomainListItemResult, err error) {
	result = DomainListItemResult{}

	extra := struct {
		UserName string `json:"username"`
	}{
		username,
	}

	response, err := m.mp.writeData(deleteResellerDomainAction.Method, m.mp.getURL(deleteResellerDomainAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Reseller) GetResellers() (result ResellerListItemResult, err error) {
	result = ResellerListItemResult{}

	response, err := m.mp.writeData(deleteResellerDomainAction.Method, m.mp.getURL(deleteResellerDomainAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}
