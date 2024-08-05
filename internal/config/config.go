package config

// API ценных бумаг МОЭКС
const (
	API_ISS_URL          string = "https://iss.moex.com/iss/"
	API_BOND_BOARDGROUPS string = "https://iss.moex.com/iss/engines/stock/markets/bonds/boardgroups/"
)

// Массив групп облигаций
var BONDS_BOARDGROUPS = [7]uint8{58, 193, 105, 77, 207, 167, 245}

// Параметры выборки облигаций
const (
	YIELD_MORE             = 15    // Доходность больше этой цифры
	YIELD_LESS             = 25    // Доходность меньше этой цифры
	PRICE_MORE             = 60    // Цена больше этой цифры
	PRICE_LESS             = 110   // Цена меньше этой цифры
	DURATION_MORE          = 3     // Дюрация больше этой цифры
	DURATION_LESS          = 18    // Дюрация меньше этой цифры
	VOLUME_MORE            = 1500  // Объем сделок в каждый из n дней, шт. больше этой цифры
	BOND_VOLUME_MORE       = 30000 // Совокупный объем сделок за n дней, шт. больше этой цифры
	KNOWN_OFFER      uint8 = 1     //Учитывать, чтобы денежные выплаты были известны до самого погашения?
	// 1 - ДА - облигации только с известными цифрами выплаты купонов
	// 0 - НЕТ - не важно, пусть в какие-то даты вместо выплаты прочерк
)
