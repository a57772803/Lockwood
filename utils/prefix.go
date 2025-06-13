package utils

import "fmt"

const (
	CategoryPassword = "password"
	CategorySetting  = "setting"
	CategoryNote     = "note"
	CategoryLog      = "log"
)

func GenerateKey(userID, category, entryID string) string {
	return fmt.Sprintf("%s:%s:%s", userID, category, entryID)
}

func Prefix(userID, category string) []byte {
	return []byte(fmt.Sprintf("%s:%s:", userID, category))
}
