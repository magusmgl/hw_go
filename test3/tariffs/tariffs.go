package tariffs

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// TariffKey определяет уникальный составной ключ для тарифа.
// Структуры могут быть ключами в map, если они "сравниваемые".
type TariffKey struct {
	// Store      string
	City string
	// TariffZone string
	WeightType string
}
type TariffValue struct {
	TariffZone string
	Tariff     float64
}

// TariffStore инкапсулирует логику хранения и доступа к тарифам.
type TariffStore struct {
	// Ключ - это комбинация города, зоны и веса.
	// Значение - цена тарифа.
	tariffs map[TariffKey][]TariffValue
}

// NewTariffStore создает и инициализирует новое хранилище тарифов.
func NewTariffStore() *TariffStore {
	return &TariffStore{
		tariffs: make(map[TariffKey][]TariffValue),
	}
}

// Add добавляет новый тариф в хранилище.
// Возвращает ошибку, если тариф с таким ключом уже существует.
func (tariffStore *TariffStore) Add(city, weightType, tariffZone string, price float64) {
	key := TariffKey{
		City:       city,
		WeightType: weightType,
	}

	values, exists := tariffStore.tariffs[key]
	if exists {
		// Ключ существует, проверяем, есть ли уже тариф для этой зоны.
		for _, value := range values {
			if value.TariffZone == tariffZone {
				// Тариф для этой зоны уже существует, ничего не делаем.
				return
			}
		}
		// Зона новая для этого ключа, добавляем.
		values = append(values, TariffValue{TariffZone: tariffZone, Tariff: price})
		tariffStore.tariffs[key] = values
	} else {
		// Ключ не существует, создаем новую запись.
		tariffStore.tariffs[key] = []TariffValue{{TariffZone: tariffZone, Tariff: price}}

	}
}

// Get получает цену для заданного тарифа.
// Возвращает цену и флаг, указывающий, был ли найден тариф.
func (tariffStore *TariffStore) Get(city, weightType string) ([]TariffValue, bool) {
	key := TariffKey{
		City:       city,
		WeightType: weightType,
	}
	tariffs, found := tariffStore.tariffs[key]
	return tariffs, found
}

const (
	statusCol = 2
	// storeCol         = 20
	reportCityCol    = 4
	weightCol        = 6
	tariffZoneCol    = 7
	BasePriceCol     = 13
	tariffHeaderRows = 1
)

func getRows(fileName string, sheetIndex int) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", fileName, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Ошибка при закрытии файла %s: %v", fileName, err)
		}
	}()

	sheetName := f.GetSheetName(sheetIndex)
	if sheetName == "" {
		return nil, fmt.Errorf("лист с индексом %d не найден в файле %s", sheetIndex, fileName)
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить строки с листа %s: %w", sheetName, err)
	}
	return rows, nil
}

// LoadTariffsFromFile загружает тарифы из Excel-файла в TariffStore.
func LoadTariffsFromFile(fileName string) (*TariffStore, error) {
	// Здесь предполагается, что у вас есть лист с тарифами,
	// а не сводный отчет по заказам.
	// Индекс листа и колонки нужно будет настроить.
	const tariffSheetIndex = 2 // Пример, укажите правильный индекс листа

	rows, err := getRows(fileName, tariffSheetIndex)
	if err != nil {
		return nil, err
	}

	tariffStore := NewTariffStore()
	startTime := time.Now()

	for i, row := range rows {
		if i < tariffHeaderRows { // Пропускаем заголовок
			continue
		}

		if row[statusCol] != "Выдан" {
			continue
		}

		// Извлекаем данные из колонок. Убедитесь, что константы верны.
		// store := row[storeCol]
		city := row[reportCityCol]       // reportCityCol = 4
		tariffZone := row[tariffZoneCol] // tariffZoneCol = 7
		weightStr := row[weightCol]      // weightCol = 6
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			log.Printf("Не удалось сконвертировать вес '%s' в строке %d: %v", weightStr, i+1, err)
			continue
		}

		var weightType string
		if weight <= 5.00 {
			weightType = "less_5"
		} else {
			weightType = "more_5"
		}

		priceStr := row[BasePriceCol] // BasePriceCol = 13
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Не удалось сконвертировать цену '%s' в строке %d: %v", priceStr, i+1, err)
			continue
		}

		tariffStore.Add(city, weightType, tariffZone, price)
	}
	fmt.Println(time.Since(startTime))

	return tariffStore, nil
}
