package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarket_IsValid(t *testing.T) {
	market := Market{}
	market.ID = "1234567"
	market.Longitude = "-1111111"
	market.Latitude = "-9999999"
	market.CensusSector = "CensusSector"
	market.WeightingArea = "WeightingArea"
	market.TownshipCode = "123"
	market.Township = "Township"
	market.SubPrefectureCode = "00"
	market.SubPrefecture = "SubPrefecture"
	market.Region5 = "Region5"
	market.Region8 = "Region8"
	market.Name = "Name"
	market.Registry = "123456"
	market.Street = "Street"
	market.Number = "123"
	market.District = "District"
	market.Reference = "Reference"
	_, err := market.IsValid()
	require.Nil(t, err)

	market.ID = ""
	market.Number = ""
	market.Reference = ""
	_, err = market.IsValid()
	require.Nil(t, err)

	market.Longitude = ""
	market.Latitude = ""
	market.CensusSector = ""
	market.WeightingArea = ""
	market.TownshipCode = ""
	market.Township = ""
	market.SubPrefectureCode = ""
	market.SubPrefecture = ""
	market.Region5 = ""
	market.Region8 = ""
	market.Name = ""
	market.Registry = ""
	market.Street = ""
	market.District = ""
	_, err = market.IsValid()
	require.Equal(t,
		"CensusSector: non zero value required;"+
			"District: non zero value required;"+
			"Latitude: non zero value required;"+
			"Longitude: non zero value required;"+
			"Name: non zero value required;"+
			"Region5: non zero value required;"+
			"Region8: non zero value required;"+
			"Registry: non zero value required;"+
			"Street: non zero value required;"+
			"SubPrefecture: non zero value required;"+
			"SubPrefectureCode: non zero value required;"+
			"Township: non zero value required;"+
			"TownshipCode: non zero value required;"+
			"WeightingArea: non zero value required",
		err.Error())
}
