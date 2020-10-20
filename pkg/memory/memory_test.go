package memory

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestMemoryNew(t *testing.T) {

	_, err := New()

	if err != nil {
		t.Fatal(err)
	}
}

type TestMemorySuite struct {
	suite.Suite
	memory *Memory
}

func (m *TestMemorySuite) SetupTest() {
	m.memory, _ = New()
}

func (m TestMemorySuite) TestInsert() {

	err := m.memory.Insert(Line{
		Msg:       "asdasd",
		Level:     "asd",
		Timestamp: time.Now().String(),
	})

	m.Assert().Nil(err)

	err = m.memory.Insert(Line{
		Msg: "asdasd",
	})

	m.Assert().Nil(err)

	logs := make([]Line, 0)

	err = m.memory.db.Select(&logs, "select * from logs;")

	m.Assert().Nil(err)

	m.Assert().Len(logs, 2)
}

func (m TestMemorySuite) TestSelect() {
	stmt, err := m.memory.db.Prepare("insert into logs(logName, level, msg, timestamp) values (?, ?, ?, ?)")

	m.Assert().Nil(err)

	for i := 0; i < 110; i++ {
		_, _ = stmt.Exec("test", "info", "hekki", time.Now().String())
	}

	_ = stmt.Close()

	lines, err := m.memory.Last("test")

	m.Assert().Nil(err)

	m.Assert().Len(lines, 100)

	m.Assert().Equal(0, lines[len(lines)-1].ID)
}

func (m TestMemorySuite) TestSelectMatch() {
	stmt, err := m.memory.db.Prepare("insert into logs (logName, level, msg, timestamp) values (?, ?, ?, ?)")

	m.Assert().Nil(err)

	for i := 0; i < 110; i++ {
		_, _ = stmt.Exec("test", "info", "hekki", time.Now().String())
	}

	_ = stmt.Close()

	lines, err := m.memory.Match("test", "hekki")

	m.Assert().Nil(err)

	m.Assert().Len(lines, 110)

	m.Assert().Equal(0, lines[len(lines)-1].ID)
}

func TestMemory(t *testing.T) {
	suite.Run(t, new(TestMemorySuite))
}

func BenchmarkMemoryInsert(b *testing.B) {
	memory, _ := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		memory.Insert(Line{
			Msg:       "asdasd",
			Level:     "asd",
			Timestamp: time.Now().String(),
		})
	}
}
