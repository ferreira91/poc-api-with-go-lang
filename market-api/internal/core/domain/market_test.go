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
	market.SubprefectureCode = "00"
	market.Subprefecture = "Subprefecture"
	market.Region5 = "Region"
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
	market.District = ""
	_, err = market.IsValid()
	require.Nil(t, err)

	market.Registry = "00000000"
	_, err = market.IsValid()
	require.Equal(t, err.Error(), "Registry: 00000000 does not validate as maxstringlength(6)")

	market.Longitude = ""
	market.Latitude = ""
	market.CensusSector = ""
	market.WeightingArea = ""
	market.TownshipCode = ""
	market.Township = ""
	market.SubprefectureCode = ""
	market.Subprefecture = ""
	market.Region5 = ""
	market.Region8 = ""
	market.Name = ""
	market.Registry = ""
	market.Street = ""
	market.District = ""
	_, err = market.IsValid()
	require.Equal(t,
		err.Error(),
		"CensusSector: non zero value required;"+
			"Latitude: non zero value required;"+
			"Longitude: non zero value required;"+
			"Name: non zero value required;"+
			"Region5: non zero value required;"+
			"Region8: non zero value required;"+
			"Registry: non zero value required;"+
			"Street: non zero value required;"+
			"Subprefecture: non zero value required;"+
			"SubprefectureCode: non zero value required;"+
			"Township: non zero value required;"+
			"TownshipCode: non zero value required;"+
			"WeightingArea: non zero value required",
	)
}
