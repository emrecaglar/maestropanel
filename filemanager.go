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



// func (m *Web) GetDotNetRuntimeVersion(domainName string)(result NETRuntimeResult, err error) {
// 	result = NETRuntimeResult{}

// 	extra := struct {
// 		Name string `json:"name"`
// 	} {
// 		domainName,
// 	}

// 	response, err := m.mp.writeData(getNETRuntimeVersionAction.Method, m.mp.getURL(getNETRuntimeVersionAction), extra)

// 	if err == nil {
// 		json.Unmarshal(response, &result)
// 	}

// 	return
// }
