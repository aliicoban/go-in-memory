package services

import (
	"github.com/alicobanserver/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mID = ""
var Db = models.Database{}
var St = Store{Db: Db}

func TestSetData(t *testing.T) {
	id, err := St.Set("setDataValue")
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	mID = id
	assert.NotNil(t, id)
}

func TestGetData(t *testing.T) {
	//First Scenario
	value, err := St.Get(mID)
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	assert.EqualValues(t, Db[mID], value)

	//Second Scenario
	value2, err2 := St.Get("")
	if err2 != nil {
		assert.NotNil(t, err2)
		return
	}
	assert.Nil(t, value2)
}

func TestDeleteSingle(t *testing.T) {
	id, err := St.Set("value")
	if err != nil {
		assert.NotNil(t, err)
		return
	}

	assert.NotNil(t, id)

	St.DeleteSingle(id)
	assert.Equal(t, Db[id], "")
}

func TestFlush(t *testing.T) {
	id, err := St.Set("test2")
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	assert.NotNil(t, id)

	St.FlushData()
	assert.Equal(t, Db[mID], "")
}

func TestGetAll(t *testing.T) {
	id, err := St.Set("test2")
	if err != nil {
		assert.NotNil(t, err)
		return
	}
	assert.NotNil(t, id)

	mapVal, err2 := St.GetAll()
	if err2 != nil {
		assert.NotNil(t, err2)
		return
	}
	assert.NotNil(t, mapVal)

}
