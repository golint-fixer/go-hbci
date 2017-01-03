package bankinfo

import (
	Csv "encoding/csv"
	"io"
	"strings"

	"github.com/mitch000001/go-hbci/internal"
	"github.com/wildducktheories/go-csv"
)

const (
	BANK_IDENTIFIER_HEADER = "BLZ"
	BANK_INSTITUTE_HEADER  = "Institut"
	VERSION_NUMBER_HEADER  = "HBCI-Version"
	URL_HEADER             = "PIN/TAN-Zugang URL"
	VERSION_NAME_HEADER    = "Version"
	CITY_HEADER            = "Ort"
)

const (
	BIC_BANK_IDENTIFIER = "Bank-leitzahl"
	BIC_IDENTIFIER      = "BIC"
)

func ParseBankInfos(reader io.Reader) ([]BankInfo, error) {
	CsvReader := Csv.NewReader(reader)
	CsvReader.Comma = ';'
	CsvReader.FieldsPerRecord = -1
	CsvReader.TrimLeadingSpace = true
	csvReader := csv.WithCsvReader(CsvReader, nil)
	records, err := csv.ReadAll(csvReader)
	if err != nil {
		return nil, err
	}
	var bankInfos []BankInfo
	for _, record := range records {
		if record.Get(BANK_IDENTIFIER_HEADER) == "" {
			internal.Debug.Printf("No BankIdentifier found for record:\n%#v\n", record.AsMap())
			continue
		}
		bankInfo := BankInfo{
			BankId:        strings.TrimSpace(record.Get(BANK_IDENTIFIER_HEADER)),
			VersionNumber: strings.TrimSpace(record.Get(VERSION_NUMBER_HEADER)),
			URL:           strings.TrimSpace(record.Get(URL_HEADER)),
			VersionName:   strings.TrimSpace(record.Get(VERSION_NAME_HEADER)),
			Institute:     strings.TrimSpace(record.Get(BANK_INSTITUTE_HEADER)),
			City:          strings.TrimSpace(record.Get(CITY_HEADER)),
		}
		bankInfos = append(bankInfos, bankInfo)
	}
	return bankInfos, nil
}

func ParseBicData(reader io.Reader) ([]BicInfo, error) {
	CsvReader := Csv.NewReader(reader)
	CsvReader.Comma = ';'
	CsvReader.FieldsPerRecord = -1
	CsvReader.TrimLeadingSpace = true
	csvReader := csv.WithCsvReader(CsvReader, nil)
	records, err := csv.ReadAll(csvReader)
	if err != nil {
		return nil, err
	}
	var bicInfos []BicInfo
	for _, record := range records {
		if record.Get(BIC_BANK_IDENTIFIER) == "" {
			internal.Debug.Printf("No BankIdentifier found for record:\n%#v\n", record.AsMap())
			continue
		}
		bicInfo := BicInfo{
			BankId: strings.TrimSpace(record.Get(BIC_BANK_IDENTIFIER)),
			BIC:    strings.TrimSpace(record.Get(BIC_IDENTIFIER)),
		}
		bicInfos = append(bicInfos, bicInfo)
	}
	return bicInfos, nil
}
