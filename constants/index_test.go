package constants

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {

	var setDataMsg = "An error occurred while adding data"
	var getDataMsg = "An error occurred while getting data"
	var getAllDataMsg = "An error occurred while getting all data"
	var notFoundMsg = "Data Not Found"
	var flushDataMsg = "An error occurred while flush data"
	var deleteSingleMsg = "An error occurred while deleting single data"

	var swaggerInfoTitleMsg = "Yemek Sepeti Swagger"
	var swaggerInfoDescriptionMsg = "This is a Golang In memory Services."
	var swagegrInfoVersion = "1.0"
	var swagegrInfoHost = ""
	var swaggerInfoBasePath = "/"

	assert.EqualValues(t, SETDATA_ERR_MSG, setDataMsg)
	assert.EqualValues(t, GETDATA_ERR_MSG, getDataMsg)
	assert.EqualValues(t, GETALLDATA_ERR_MSG, getAllDataMsg)
	assert.EqualValues(t, NOT_FOUND, notFoundMsg)
	assert.EqualValues(t, FLUSHDATA_ERR_MSG, flushDataMsg)
	assert.EqualValues(t, DELETESINGLE_ERR_MSG, deleteSingleMsg)

	assert.EqualValues(t, SWAGGERINFO_TITLE, swaggerInfoTitleMsg)
	assert.EqualValues(t, SWAGGERINFO_DESCRIPTION, swaggerInfoDescriptionMsg)
	assert.EqualValues(t, SWAGGERINFO_VERSION, swagegrInfoVersion)
	assert.EqualValues(t, SWAGGERINFO_HOST, swagegrInfoHost)
	assert.EqualValues(t, SWAGGERINFO_BASEPATH, swaggerInfoBasePath)

}
