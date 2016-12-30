package maestropanel

import (
    "encoding/json"
)

type FileManager struct {
    mp MaestroPanel
}

func (m *FileManager) SetWriteAccess(path Path)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(setWriteAccessAction.Method, m.mp.getURL(setWriteAccessAction), path)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *FileManager) RevokeWriteAccess(path Path)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(revokeWriteAccessAction.Method, m.mp.getURL(revokeWriteAccessAction), path)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *FileManager) CreateDirectory(path Path)(result DomainExecutionResult, err error) {
	result = DomainExecutionResult{}

	response, err := m.mp.writeData(createDirectoryAction.Method, m.mp.getURL(createDirectoryAction), path)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *FileManager) GetItems(path Path)(result GetItemsResult, err error) {
    result = GetItemsResult{}

    response, err := m.mp.writeData(getItemsAction.Method,m.mp.getURL(getItemsAction),path)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *FileManager) DeleteItems(domainName string, items ...string)(result GetItemsResult, err error) {
    result = GetItemsResult{}

    extra := struct {
        Name string `json:"name"`
        Items []string `json:"item"`
    } {
        domainName,
        items,
    }

    response, err := m.mp.writeData(deleteItemsAction.Method,m.mp.getURL(deleteItemsAction),extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

func (m *FileManager) ZipItem(domainName string, zipFilePath string,  items ...string)(result DomainExecutionResult, err error) {
    result = DomainExecutionResult{}

    extra := struct {
        Name string `json:"name"`
        Items []string `json:"item"`
        Path string `json:"path"`
    } {
        domainName,
        items,
        zipFilePath,
    }

    response, err := m.mp.writeData(zipItemAction.Method,m.mp.getURL(zipItemAction),extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}