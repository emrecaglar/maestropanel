package maestropanel

import (
    "encoding/json"
)

type Database struct {
    mp MaestroPanel
}

func (m *Database) AddDatabase(db DatabaseInfo)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    response, err := m.mp.writeData(addDatabaseAction.Method, m.mp.getURL(addDatabaseAction), db)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

