package moex

import (
	"encoding/json"
	"go-bonds-scanner/internal/config"
	"log"
	"net/http"
)

// Структура ответа
type SecurityResponse struct {
	Securities SecuritiesMap `json:"securities"`
}

// Структура ответа
type SecuritiesMap struct {
	Metadata map[string]any `json:"metadata"`
	Columns  []string       `json:"columns"`
	Data     []interface{}  `json:"data"`
}

// Метод запроса ценных бумаг
func GetSecurities() (*SecurityResponse, error) {
	// Возможно вынести общий запрос ценных бумаг в приватный метод
	resp, err := http.Get(config.API_ISS_URL + "securities.json")

	if err != nil {
		log.Fatalln("Во время запроса ценных бумаг произошла ошибка.")
		return nil, err
	}

	defer resp.Body.Close()

	security := &SecurityResponse{}
	err = json.NewDecoder(resp.Body).Decode(&security)

	if err != nil {
		log.Fatalln("Во время преобразования ответа произошла ошибка")
		return nil, err
	}

	return security, nil
}
