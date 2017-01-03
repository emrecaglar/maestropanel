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

func (m *Database) DeleteDatabase(domainName string, dbType string, databaseName string)(result DomainExecutionResult, err error)  {
    result = DomainExecutionResult{}

    extra := struct {
        Name string `json:"name"`
        DBType string `json:"dbtype"`
        Database string `json:"database"`
    } {
        domainName,
        dbType,
        databaseName,
    }

    response, err := m.mp.writeData(deleteDatabaseAction.Method, m.mp.getURL(deleteDatabaseAction), extra)

    if err == nil {
        json.Unmarshal(response, &result)
    }

    return
}

