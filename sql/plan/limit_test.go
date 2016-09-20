package plan

import (
	"fmt"
	"io"
	"testing"

	"github.com/mvader/gitql/mem"
	"github.com/mvader/gitql/sql"
	"github.com/stretchr/testify/assert"
)

var testingTable *mem.Table
var testingTableSize int

func TestLimitPlan(t *testing.T) {
	assert := assert.New(t)
	table, _ := getTestingTable()
	limitPlan := NewLimit(0, table)
	assert.Equal(1, len(limitPlan.Children()))

	iterator, err := getLimitedIterator(1)
	assert.Nil(err)
	assert.NotNil(iterator)
}

func TestLimit0(t *testing.T) {
	_, testingTableSize := getTestingTable()
	testingLimit := 0
	iterator, _ := getLimitedIterator(int64(testingLimit))
	testLimitOverflow(t, iterator, testingLimit, testingTableSize)
}

func TestLimitLessThanTotal(t *testing.T) {
	_, testingTableSize := getTestingTable()
	testingLimit := testingTableSize - 1
	iterator, _ := getLimitedIterator(int64(testingLimit))
	testLimitOverflow(t, iterator, testingLimit, testingTableSize)
}

func TestLimitEqualThanTotal(t *testing.T) {
	_, testingTableSize := getTestingTable()
	testingLimit := testingTableSize
	iterator, _ := getLimitedIterator(int64(testingLimit))
	testLimitOverflow(t, iterator, testingLimit, testingTableSize)
}

func TestLimitGreaterThanTotal(t *testing.T) {
	_, testingTableSize := getTestingTable()
	testingLimit := testingTableSize + 1
	iterator, _ := getLimitedIterator(int64(testingLimit))
	testLimitOverflow(t, iterator, testingLimit, testingTableSize)
}

func testLimitOverflow(t *testing.T, iter sql.RowIter, limit int, dataSize int) {
	assert := assert.New(t)
	for i := 0; i < limit+1; i++ {
		row, err := iter.Next()
		hint := fmt.Sprintf("Iter#%d : size.%d : limit.%d", i, dataSize, limit)
		if i >= limit || i >= dataSize {
			assert.Nil(row, hint)
			assert.Equal(io.EOF, err, hint)
		} else {
			assert.NotNil(row, hint)
			assert.Nil(err, hint)
		}
	}
}

func getTestingTable() (*mem.Table, int) {

	if &testingTable == nil {
		return testingTable, testingTableSize
	}

	childSchema := sql.Schema{
		sql.Field{"col1", sql.String},
	}
	testingTable = mem.NewTable("test", childSchema)
	testingTable.Insert("11a")
	testingTable.Insert("22a")
	testingTable.Insert("33a")
	testingTableSize = 3

	return testingTable, testingTableSize
}

func getLimitedIterator(limitSize int64) (sql.RowIter, error) {
	table, _ := getTestingTable()
	limitPlan := NewLimit(limitSize, table)
	return limitPlan.RowIter()
}
