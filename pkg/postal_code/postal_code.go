package postalcode

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"
)

// validation for postal code in indonesia
type PostalCode struct {
	Province string
	City     string
	// SubDistrict string
	// Urban       string
	PostalCode string
}

func ValidationPostalCode(pc *PostalCode) (bool, error) {
	if err := pc.matchingPostalCode(); err != nil {
		return false, err
	}
	return true, nil
}

// This "function" will validation address that input from user according actual/valid.
//
// Requirement :
//
//	-'Province' as "province" params
//	-'City' as "city" params
//	-'Sub' District as "subDistrict" params
//	-'Urban' as "urban" params
//
// Error Chance :
//
//	-"postal code do not match"
//	-"province not found"
func (reg *PostalCode) matchingPostalCode() error {
	pc := StockPostalCode()
	// p := "jawa tengah"
	// c := "purbalingga"
	// sd := "bukateja"
	// u := "kembangan"
	var provinceCode string
	for _, code := range pc.Province {
		if strings.EqualFold(code.ProvinceName, reg.Province) {
			provinceCode = code.ProvinceCode
			break
		}
		provinceCode = ""
		// return errors.New("province not found")
	}
	if provinceCode == "" {
		return errors.New("province not found")
	}

	stock := pc.Postal[provinceCode]
	for _, x := range stock {
		if strings.EqualFold(reg.City, x.City) && reg.PostalCode == x.PostalCode {
			// strings.ToLower(x.Urban) == reg.Urban {
			// fmt.Println(x.City)
			// fmt.Println(x.ProvinceCode)
			// fmt.Println(x.SubDistrict)
			// fmt.Println(x.Urban)
			// fmt.Printf("[%v]\n", x.PostalCode)
			// fmt.Println("---------------------------")
			return nil
		}
	}
	return errors.New("postal code do not match")
}

//go:embed merge_province_postal_full.json
var MergePostalCodeFull []byte

// var ProvCode []string

// func init() {
// 	for i := 0; i > 93; i++ {
// 		if i > 10 {
// 			pc := strconv.Itoa(i)
// 			ProvCode = append(ProvCode, pc)
// 		}
// 	}
// }

type PostalCodeModels struct {
	Province map[string]struct {
		ProvinceCode   string `json:"province_code"`
		ProvinceName   string `json:"province_name"`
		ProvinceNameEn string `json:"province_name_en"`
	} `json:"province"`
	Postal map[string][]struct {
		Urban        string `json:"urban"`
		SubDistrict  string `json:"sub_district"`
		City         string `json:"city"`
		ProvinceCode string `json:"province_code"`
		PostalCode   string `json:"postal_code"`
	}
}

func StockPostalCode() *PostalCodeModels {
	postalCode := new(PostalCodeModels)
	decodeErr := json.Unmarshal(MergePostalCodeFull, postalCode)
	if decodeErr != nil {
		panic("error postal code")
	}
	return postalCode
}
