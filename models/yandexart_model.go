package models

import "fmt"

const (
	YandexARTLatest = "yandex-art/latest"
)

func GetARTModelURI(catalogID string) string {
	return fmt.Sprintf("art://%s/%s", catalogID, YandexARTLatest)
}
