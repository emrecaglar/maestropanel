package maestropanel

import (
	"encoding/json"
)

type Server struct {
	mp MaestroPanel
}

type IPAddressModel struct {
	ServerName  string `json:"serverName"`
	NicName     string `json:"nicName"`
	IPAddr      string `json:"ipAddr"`
	SubNet      string `json:"subNet"`
	IsShared    bool   `json:"isShared"`
	IsDedicated bool   `json:"isDedicated"`
	IsExclusive bool   `json:"isExclusive"`
}

type ServerIPAddress struct {
	Id          int32
	ServerId    int32
	NicId       int32
	NicName     string
	IpAddr      string
	Subnet      string
	isExclusive bool
	isDedicated bool
	isShared    bool
	domainCount int32
}

type IPAddressListResult struct {
	Result
	Details []ServerIPAddress
}

type MethodResult struct {
	Status bool
	Msg    string
}

type ServerExecutionResult struct {
	Result
	Details MethodResult
}

type ServerItem struct {
	Id              int32
	Name            string
	Host            string
	ComputerName    string
	OperatingSystem string
	Version         string
	Cpu             string
	Status          bool
}

type ServerListExecutionResult struct {
	Result
	Details []ServerItem
}

type NicItem struct {
	Id            int32
	Name          string
	Identifier    string
	Description   string
	InterfaceType string
	ServerId      int32
}

type NicListExecutionResult struct {
	Result
	Details []NicItem
}

type ServerModel struct {
	ServerName string `json:"servername"`
	Host       string `json:"host"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	AgentPort  string `json:"agentport"`
}

type ResourceItem struct {
	resourceType string `json:"resourceType"`
	resourceName string `json:"resourceName"`
	Total        string `json:"Total"`
	Used         string `json:"Used"`
}

type GetResourcesExecutionResult struct {
	Result
	Details []ResourceItem
}

func (m *Server) GetIpAddrList() (result IPAddressListResult, err error) {
	result = IPAddressListResult{}

	response, err := m.mp.writeData(getServerIPAdressListAction.Method, m.mp.getURL(getServerIPAdressListAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) AddIpAddr(ipaddr IPAddressModel) (result ServerExecutionResult, err error) {
	result = ServerExecutionResult{}

	response, err := m.mp.writeData(addServerIPAddrAction.Method, m.mp.getURL(addServerIPAddrAction), ipaddr)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) DeleteIpAddr(serverName string, nicName string, ipAddress string) (result ServerExecutionResult, err error) {
	result = ServerExecutionResult{}

	extra := struct {
		ServerName string `json:"serverName"`
		NicName    string `json:"nicName"`
		IPAddress  string `json:"ipAddr"`
	}{
		serverName,
		nicName,
		ipAddress,
	}

	response, err := m.mp.writeData(deleteServerIPAddrAction.Method, m.mp.getURL(deleteServerIPAddrAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) GetServerList() (result ServerListExecutionResult, err error) {
	result = ServerListExecutionResult{}

	response, err := m.mp.writeData(getServerListAction.Method, m.mp.getURL(getServerListAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) GetNicList() (result NicListExecutionResult, err error) {
	result = NicListExecutionResult{}

	response, err := m.mp.writeData(getServerNicListAction.Method, m.mp.getURL(getServerNicListAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) AddServer(server ServerModel) (result ServerExecutionResult, err error) {
	result = ServerExecutionResult{}

	response, err := m.mp.writeData(addServerAction.Method, m.mp.getURL(addServerAction), server)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) DeleteServer(serverName string) (result ServerExecutionResult, err error) {
	result = ServerExecutionResult{}

	extra := struct {
		ServerName string `json:"servername"`
	}{
		serverName,
	}

	response, err := m.mp.writeData(deleteServerAction.Method, m.mp.getURL(deleteServerAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *Server) GetResources(serverName string) (result GetResourcesExecutionResult, err error) {
	result = GetResourcesExecutionResult{}

	extra := struct {
		ServerName string `json:"servername"`
	}{
		serverName,
	}

	response, err := m.mp.writeData(getServerResourcesAction.Method, m.mp.getURL(getServerResourcesAction), extra)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}
