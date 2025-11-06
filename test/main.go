package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// )

type Data struct {
	Results []struct {
		Times []float64 `json:"times"`
	} `json:"results"`
	Name string `json:"name"`
}

// ResponseStats содержит статистику времени ответа
type ResponseStats struct {
	Count       int
	Min         float64
	Max         float64
	Average     float64
	Percentiles map[string]float64
}

// processCSV читает CSV файл, ищет значение из колонки dateColumnName
// внутри строки из колонки jsonColumnName для каждой строки.
// Выводит общее количество обработанных строк и количество совпадений.
func processCSV(filePath string, dateColumnName string, jsonColumnName string) error {
	// Открываем CSV файл
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла %s: %w", filePath, err)
	}
	defer file.Close() // Гарантируем закрытие файла при выходе из функции

	reader := csv.NewReader(file)
	// Можно настроить разделитель, если он не запятая (например, reader.Comma = ';')

	// Читаем заголовок CSV файла, чтобы определить индексы колонок
	headers, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("файл %s пуст или содержит только заголовок", filePath)
		}
		return fmt.Errorf("ошибка при чтении заголовка CSV: %w", err)
	}

	// Находим индексы нужных колонок по их именам
	dateColIndex := -1
	jsonColIndex := -1
	for i, header := range headers {
		if header == dateColumnName {
			dateColIndex = i
		}
		if header == jsonColumnName {
			jsonColIndex = i
		}
	}

	// Проверяем, найдены ли обе колонки
	if dateColIndex == -1 {
		return fmt.Errorf("колонка '%s' не найдена в CSV файле", dateColumnName)
	}
	if jsonColIndex == -1 {
		return fmt.Errorf("колонка '%s' не найдена в CSV файле", jsonColumnName)
	}

	totalRows := 0  // Общее количество обработанных строк данных (без заголовка)
	matchCount := 0 // Количество строк, где найдено совпадение

	// Читаем CSV файл построчно
	for {
		record, err := reader.Read() // record - это срез строк, представляющий одну строку CSV
		if err != nil {
			if err == io.EOF {
				break // Достигнут конец файла
			}
			return fmt.Errorf("ошибка при чтении строки CSV: %w", err)
		}

		totalRows++ // Увеличиваем счетчик обработанных строк

		// Проверяем, что в текущей строке достаточно полей для доступа к нужным колонкам
		if len(record) <= dateColIndex || len(record) <= jsonColIndex {
			fmt.Printf("Предупреждение: строка %d (данные) имеет недостаточно полей, пропущена.\n", totalRows)
			continue // Пропускаем эту строку и переходим к следующей
		}

		dateValue := record[dateColIndex]  // Значение из колонки с датой
		jsonString := record[jsonColIndex] // Строка из колонки с JSON

		// Проверяем, содержится ли dateValue как подстрока в jsonString
		if strings.Contains(jsonString, dateValue) {
			matchCount++ // Увеличиваем счетчик совпадений
		}
	}

	// Выводим итоговую статистику
	fmt.Printf("--- Результаты анализа CSV ---\n")
	fmt.Printf("Файл: %s\n", filePath)
	fmt.Printf("Всего строк данных обработано: %d\n", totalRows)
	fmt.Printf("Количество совпадений: %d\n", matchCount)

	return nil
}

// createTestCSV создает простой тестовый CSV файл для демонстрации работы функции.
// В реальном приложении этот файл должен существовать заранее.

func main() {
	csvFilePath := "csvFile3.csv" // Имя вашего CSV файла
	dateColumn := "date"          // Имя колонки, содержащей дату для поиска
	jsonColumn := "json"          // Имя колонки, содержащей JSON-строку

	// Создаем тестовый CSV файл. В реальном сценарии этот шаг не нужен.
	// createTestCSV(csvFilePath)

	// Вызываем функцию для обработки CSV
	err := processCSV(csvFilePath, dateColumn, jsonColumn)
	if err != nil {
		fmt.Printf("Произошла ошибка: %v\n", err)
	}
}

// 	// Перцентили для расчета
// 	percentiles := []float64{50, 75, 90, 95, 99}
// 	// Открываем JSON файл
// 	// Обратите внимание: имя файла может потребовать корректировки для вашей системы
// 	jsonBytes, err := os.ReadFile("DPD изменение ДД ТЕСТ.postman_test_run")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Определяем структуру
// 	var data Data

// 	// Распаковываем json в структуру
// 	err = json.Unmarshal(jsonBytes, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Test Run Name: %s\n\n", data.Name)
// 	for i, result := range data.Results {
// 		fmt.Printf("--- Statistics for Result #%d ---\n", i+1)
// 		stats := CalculateStats(result.Times, percentiles)
// 		printStats(stats, percentiles)
// 		fmt.Println()
// 	}
// }

// // CalculateStats вычисляет полную статистику
// func CalculateStats(times []float64, percentiles []float64) ResponseStats {
// 	if len(times) == 0 {
// 		return ResponseStats{}
// 	}

// 	// Сортируем значения для расчета перцентилей
// 	sortedTimes := make([]float64, len(times))
// 	copy(sortedTimes, times)
// 	sort.Float64s(sortedTimes)

// 	// Вычисляем базовую статистику
// 	stats := ResponseStats{
// 		Count:       len(times),
// 		Min:         sortedTimes[0],
// 		Max:         sortedTimes[len(sortedTimes)-1],
// 		Percentiles: make(map[string]float64),
// 	}

// 	// Вычисляем сумму для среднего значения
// 	sum := 0.0
// 	for _, t := range times {
// 		sum += t
// 	}
// 	stats.Average = sum / float64(len(times))

// 	// Вычисляем перцентили
// 	for _, p := range percentiles {
// 		if p < 0 || p > 100 {
// 			continue
// 		}

// 		index := (p / 100) * float64(len(sortedTimes)-1)

// 		if index == float64(int64(index)) {
// 			stats.Percentiles[fmt.Sprintf("p%.0f", p)] = sortedTimes[int(index)]
// 		} else {
// 			lowerIndex := int(index)
// 			upperIndex := lowerIndex + 1
// 			fraction := index - float64(lowerIndex)

// 			interpolated := sortedTimes[lowerIndex] + fraction*(sortedTimes[upperIndex]-sortedTimes[lowerIndex])
// 			stats.Percentiles[fmt.Sprintf("p%.0f", p)] = interpolated
// 		}
// 	}

// 	return stats
// }

// func printStats(stats ResponseStats, percentiles []float64) {
// 	fmt.Println("API Response Time Statistics:")
// 	fmt.Printf("Request count: %d\n", stats.Count)
// 	fmt.Printf("Min: %.2f ms\n", stats.Min)
// 	fmt.Printf("Max: %.2f ms\n", stats.Max)
// 	fmt.Printf("Average: %.2f ms\n", stats.Average)

// 	fmt.Println("\nPercentiles:")
// 	for _, p := range percentiles {
// 		key := fmt.Sprintf("p%.0f", p)
// 		if value, exists := stats.Percentiles[key]; exists {
// 			fmt.Printf("%s: %.2f ms\n", key, value)
// 		}
// 	}

// 	// Анализ результатов
// 	fmt.Println("\nAnalysis:")
// 	if stats.Percentiles["p95"] > 200 {
// 		fmt.Println("⚠️  Warning: p95 response time is above 200ms")
// 	} else {
// 		fmt.Println("✅ p95 response time is within acceptable limits")
// 	}

// 	if stats.Percentiles["p99"] > 300 {
// 		fmt.Println("🚨 Critical: p99 response time is above 300ms")
// 	}
// }
