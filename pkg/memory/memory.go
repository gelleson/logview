package memory

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type LogInfo struct {
	LogName string `json:"logName" db:"logName"`
	Count   int    `json:"count" db:"count"`
}

type Memory struct {
	db       *sqlx.DB
	filename string
}

func New(filename ...string) (*Memory, error) {
	var db *sqlx.DB
	var err error

	if len(filename) == 0 {
		db, err = sqlx.Connect("sqlite3", ":memory:?cache=shared")
		if err != nil {
			return nil, err
		}
	} else {
		db, err = sqlx.Connect("sqlite3", filename[0])
		if err != nil {
			return nil, err
		}
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	memory := &Memory{
		db: db,
	}

	if err := memory.prepareTable(); err != nil {
		return &Memory{}, err
	}

	return memory, nil
}

func (m *Memory) prepareTable() error {
	tableSchema := `CREATE VIRTUAL TABLE logs using FTS5 (
	logName,
   	level,
	msg,
	timestamp,
)`

	_, err := m.db.Exec(tableSchema)

	if err != nil {
		return err
	}

	return nil
}

func (m *Memory) Insert(lines ...Line) error {
	stmt, err := m.db.Prepare("insert into logs(logName, level, msg, timestamp) values (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err := stmt.Exec(line.LogName, line.Level, line.Msg, line.Timestamp)

		if err != nil {
			return err
		}
	}

	return stmt.Close()
}

func (m *Memory) Last(logName string) ([]Line, error) {
	logs := make([]Line, 0)

	if err := m.db.Select(&logs, "select * from logs where logName = $1 order by rowId desc limit 100;", logName); err != nil {
		return nil, err
	}

	return logs, nil
}

func (m *Memory) Offset(logName string, limit, offset int) ([]Line, error) {
	logs := make([]Line, 0)

	if err := m.db.Select(&logs, "select * from logs where logName = $1 order by rowId desc limit $2 offset $3;", logName, limit, offset); err != nil {
		return nil, err
	}

	return logs, nil
}

func (m *Memory) LogsList() ([]LogInfo, error) {
	logs := make([]LogInfo, 0)

	if err := m.db.Select(&logs, "select logName, count(\"logName\") as count from logs group by logName;"); err != nil {
		return nil, err
	}

	return logs, nil
}

func (m *Memory) Match(logName, query string) ([]Line, error) {
	q := `
SELECT
*,
highlight(logs, 2, '<b>', '</b>') msg
FROM logs 

WHERE logs MATCH '%s' and logName = '%s' ORDER BY rank limit 100;
`

	lines := make([]Line, 0)

	if err := m.db.Select(&lines, fmt.Sprintf(q, query, logName)); err != nil {
		return nil, err
	}

	return lines, nil
}
