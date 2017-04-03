package maestropanel

import (
    "encoding/json"
)

type DNS struct {
    mp MaestroPanel
}


func (m *DNS) SetDnsZone(dnsZone DNSZone)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(setDNSZoneAction.Method, m.mp.getURL(setDNSZoneAction), dnsZone)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *DNS) AddDNSRecord(dnsRecord DNSRecord)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(addDNSRecordAction.Method, m.mp.getURL(addDNSRecordAction), dnsRecord)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *DNS) DeleteDnsRecord(dnsRecord DNSRecord)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(deleteDNSRecordAction.Method, m.mp.getURL(deleteDNSRecordAction), dnsRecord)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *DNS) GetDnsRecords(domainName string)(result GetDNSRecordsResult, err error) {
    result = GetDNSRecordsResult{}

    extra := struct {
        Name string `json:"name"`
    } {
        domainName,
    }

    response, err := m.mp.writeData(getDNSRecordsAction.Method, m.mp.getURL(getDNSRecordsAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}