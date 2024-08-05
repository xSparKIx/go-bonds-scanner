package moex

import (
	"encoding/json"
	"go-bonds-scanner/internal/config"
	"log"
	"net/http"
	"strconv"
)

// Структура ответа на запрос списка облигаций
type GetBondsResponse struct {
	Securities BondsList       `json:"securities"`
	Marketdata BondsMarketData `json:"marketdata"`
}

// Структура списка облигаций
type BondsList struct {
	Metadata map[string]any  `json:"metadata"`
	Columns  []string        `json:"columns"`
	Data     [][]interface{} `json:"data"`
}

// Структура информации о рынке ценных бумаг
// todo: Возможно объединить с BondsList и в целом с ответами, но вообще они различаются в данных
type BondsMarketData struct {
	Metadata map[string]any  `json:"metadata"`
	Columns  []string        `json:"columns"`
	Data     [][]interface{} `json:"data"`
}

// Тип облигации
type Bond = []interface{}

// Метод получения полного списка облигаций
// todo: Попробовать распихать в отдельный горутины
func GetAllBonds() ([]GetBondsResponse, error) {
	result := []GetBondsResponse{}

	for _, id := range config.BONDS_BOARDGROUPS {
		bondsResponse, err := GetBonds(id)

		if err != nil {
			log.Fatalf("Во время получения облигаций группы %d произошла ошибка\n", id)
			return nil, err
		}

		result = append(result, bondsResponse)
		log.Printf("Данные по группе %d получены\n", id)
	}

	return result, nil
}

// Метод получения списка облигаций
func GetBonds(boardId uint8) (GetBondsResponse, error) {
	url := config.API_BOND_BOARDGROUPS + strconv.Itoa(int(boardId)) + "/securities.json?iss.dp=comma&iss.meta=off&iss.only=securities,marketdata&securities.columns=SECID,SECNAME,PREVLEGALCLOSEPRICE&marketdata.columns=SECID,YIELD,DURATION"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("Во время получения списка облигаций произошла ошибка")
		return GetBondsResponse{}, err
	}

	defer resp.Body.Close()

	bondsResponse := GetBondsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&bondsResponse)

	if err != nil {
		log.Fatalln("Во время преобразования ответа произошла ошибка")
	}

	return bondsResponse, nil
}
