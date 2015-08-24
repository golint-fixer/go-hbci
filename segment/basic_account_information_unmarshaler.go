package segment

import (
	"bytes"
	"fmt"

	"github.com/mitch000001/go-hbci/element"
)

func (a *AccountInformationSegment) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	header := &element.SegmentHeader{}
	err = header.UnmarshalHBCI(elements[0])
	if err != nil {
		return err
	}
	var segment accountInformationSegment
	switch header.Version.Val() {
	case 4:
		segment = &AccountInformationV4{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	case 5:
		segment = &AccountInformationV5{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	case 6:
		segment = &AccountInformationV6{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown segment version: %d", header.Version.Val())
	}
	a.accountInformationSegment = segment
	return nil
}

func (a *AccountInformationV4) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], a)
	if err != nil {
		return err
	}
	a.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		a.AccountConnection = &element.AccountConnectionDataElement{}
		err = a.AccountConnection.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		a.UserID = &element.IdentificationDataElement{}
		err = a.UserID.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		a.AccountCurrency = &element.CurrencyDataElement{}
		err = a.AccountCurrency.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		a.Name1 = &element.AlphaNumericDataElement{}
		err = a.Name1.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		a.Name2 = &element.AlphaNumericDataElement{}
		err = a.Name2.UnmarshalHBCI(elements[5])
		if err != nil {
			return err
		}
	}
	if len(elements) > 6 && len(elements[6]) > 0 {
		a.AccountProductID = &element.AlphaNumericDataElement{}
		err = a.AccountProductID.UnmarshalHBCI(elements[6])
		if err != nil {
			return err
		}
	}
	if len(elements) > 7 && len(elements[7]) > 0 {
		a.AccountLimit = &element.AccountLimitDataElement{}
		err = a.AccountLimit.UnmarshalHBCI(elements[7])
		if err != nil {
			return err
		}
	}
	if len(elements) > 8 && len(elements[8]) > 0 {
		a.AllowedBusinessTransactions = &element.AllowedBusinessTransactionsDataElement{}
		if len(elements)+1 > 8 {
			err = a.AllowedBusinessTransactions.UnmarshalHBCI(bytes.Join(elements[8:], []byte("+")))
		} else {
			err = a.AllowedBusinessTransactions.UnmarshalHBCI(elements[8])
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AccountInformationV5) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], a)
	if err != nil {
		return err
	}
	a.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		a.AccountConnection = &element.AccountConnectionDataElement{}
		err = a.AccountConnection.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		a.UserID = &element.IdentificationDataElement{}
		err = a.UserID.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		a.AccountType = &element.NumberDataElement{}
		err = a.AccountType.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		a.AccountCurrency = &element.CurrencyDataElement{}
		err = a.AccountCurrency.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		a.Name1 = &element.AlphaNumericDataElement{}
		err = a.Name1.UnmarshalHBCI(elements[5])
		if err != nil {
			return err
		}
	}
	if len(elements) > 6 && len(elements[6]) > 0 {
		a.Name2 = &element.AlphaNumericDataElement{}
		err = a.Name2.UnmarshalHBCI(elements[6])
		if err != nil {
			return err
		}
	}
	if len(elements) > 7 && len(elements[7]) > 0 {
		a.AccountProductID = &element.AlphaNumericDataElement{}
		err = a.AccountProductID.UnmarshalHBCI(elements[7])
		if err != nil {
			return err
		}
	}
	if len(elements) > 8 && len(elements[8]) > 0 {
		a.AccountLimit = &element.AccountLimitDataElement{}
		err = a.AccountLimit.UnmarshalHBCI(elements[8])
		if err != nil {
			return err
		}
	}
	if len(elements) > 9 && len(elements[9]) > 0 {
		a.AllowedBusinessTransactions = &element.AllowedBusinessTransactionsDataElement{}
		if len(elements)+1 > 9 {
			err = a.AllowedBusinessTransactions.UnmarshalHBCI(bytes.Join(elements[9:], []byte("+")))
		} else {
			err = a.AllowedBusinessTransactions.UnmarshalHBCI(elements[9])
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AccountInformationV6) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], a)
	if err != nil {
		return err
	}
	a.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		a.AccountConnection = &element.AccountConnectionDataElement{}
		err = a.AccountConnection.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		a.IBAN = &element.AlphaNumericDataElement{}
		err = a.IBAN.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		a.UserID = &element.IdentificationDataElement{}
		err = a.UserID.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		a.AccountType = &element.NumberDataElement{}
		err = a.AccountType.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		a.AccountCurrency = &element.CurrencyDataElement{}
		err = a.AccountCurrency.UnmarshalHBCI(elements[5])
		if err != nil {
			return err
		}
	}
	if len(elements) > 6 && len(elements[6]) > 0 {
		a.Name1 = &element.AlphaNumericDataElement{}
		err = a.Name1.UnmarshalHBCI(elements[6])
		if err != nil {
			return err
		}
	}
	if len(elements) > 7 && len(elements[7]) > 0 {
		a.Name2 = &element.AlphaNumericDataElement{}
		err = a.Name2.UnmarshalHBCI(elements[7])
		if err != nil {
			return err
		}
	}
	if len(elements) > 8 && len(elements[8]) > 0 {
		a.AccountProductID = &element.AlphaNumericDataElement{}
		err = a.AccountProductID.UnmarshalHBCI(elements[8])
		if err != nil {
			return err
		}
	}
	if len(elements) > 9 && len(elements[9]) > 0 {
		a.AccountLimit = &element.AccountLimitDataElement{}
		err = a.AccountLimit.UnmarshalHBCI(elements[9])
		if err != nil {
			return err
		}
	}
	if len(elements) > 10 && len(elements[10]) > 0 {
		a.AllowedBusinessTransactions = &element.AllowedBusinessTransactionsDataElement{}
		err = a.AllowedBusinessTransactions.UnmarshalHBCI(elements[10])
		if err != nil {
			return err
		}
	}
	if len(elements) > 11 && len(elements[11]) > 0 {
		a.AccountExtensions = &element.AlphaNumericDataElement{}
		if len(elements)+1 > 11 {
			err = a.AccountExtensions.UnmarshalHBCI(bytes.Join(elements[11:], []byte("+")))
		} else {
			err = a.AccountExtensions.UnmarshalHBCI(elements[11])
		}
		if err != nil {
			return err
		}
	}
	return nil
}