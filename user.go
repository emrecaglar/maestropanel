package maestropanel

import (
	"encoding/json"
)

type UserExecutionResult struct {
	Result
	Details UserInfo
}

type UserInfo struct {
	Id        string
	UserName  string
	Type      string
	Status    int32
	Email     string
	FirstName string
	LastName  string
}

type User struct {
	mp MaestroPanel
}

func (m *User) Whoami() (result UserExecutionResult, err error) {
	result = UserExecutionResult{}

	response, err := m.mp.writeData(whoamiUserAction.Method, m.mp.getURL(whoamiUserAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}

func (m *User) LogOff() (result Result, err error) {
	result = Result{}

	response, err := m.mp.writeData(logOffUserAction.Method, m.mp.getURL(logOffUserAction), nil)

	if err == nil {
		json.Unmarshal(response, &result)
	}

	return
}
