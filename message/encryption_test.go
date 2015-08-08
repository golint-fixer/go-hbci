package message

import (
	"fmt"
	"testing"

	"github.com/mitch000001/go-hbci/domain"
)

func TestEncryptedPinTanMessageDecrypt(t *testing.T) {
	keyName := domain.NewPinTanKeyName(domain.BankId{CountryCode: 280, ID: "1"}, "userID", "V")
	pinKey := domain.NewPinKey("abcde", keyName)

	provider := NewPinTanCryptoProvider(pinKey, "clientSystemID")

	syncSegment := "HISYN:2:3:8+newClientSystemID'"
	acknowledgement := "HIRMG:2:2:1+0100::Dialog beendet'"

	body := fmt.Sprintf("%s%s", acknowledgement, syncSegment)

	encryptedMessage := NewEncryptedPinTanMessage("clientSystemID", *keyName, []byte(body))

	decryptedMessage, err := encryptedMessage.Decrypt(provider)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	actualSyncSegment := decryptedMessage.FindSegment("HISYN")

	if syncSegment != string(actualSyncSegment) {
		t.Logf("Expected decrypted message to include SynchronisationResponse, but had not\n")
		t.Fail()
	}
}