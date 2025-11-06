package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	callback "test3/5post/callBack"
	"test3/5post/tariffs"
	"time"

	"github.com/xuri/excelize/v2"
)

type OrderInfo struct {
	OrderNumber     string `json:"OrderNumber"`
	DeliveryService string `json:"deliveryService"`
	BasePrice       struct {
		TotalSumPlan float64 `json:"totalSumPlan"`
		TotalSumFact float64 `json:"totalSumFact"`
	} `json:"basePrice"`
	FactRequest struct {
		DeliveryDate string `json:"deliveryDate"`
	} `json:"factRequest"`
}

type OrderData struct {
	OrderNumber     string
	Status          string
	BasePrice       float64
	BasePriceReturn float64
	InsurancePay    float64
}

type ListOrdersData struct {
	Orders []OrderData
}

func main() {

	//----------------------------------------------
	// Читаем колбэки и сравниваем с файлом
	callbacks, err := callback.GetCallBack("FivePostReport-2025-09-30.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении файла costreconciliation: %v", err)
	}
	fmt.Printf("Успешно прочитано %d записей из callback файла.\n", len(callbacks))
	fmt.Println(callbacks[2])
	///--------------------------

	// tariffs, err1 := tariffs.LoadTariffsFromFile("_08_свод+разбивка_5пост.xlsx")
	// if err1 != nil {
	// 	log.Fatalf("Ошибка при чтении файла _08_свод: %v", err1)
	// }

	// t, ok := tariffs.Get("Совхозный п", "less_5")
	// if !ok {
	// 	fmt.Println("Тариф не найден")
	// }
	// fmt.Println(t)

	// err = writeToExcelDiff("test.xlsx", &callbacks, tariffs)
	// if err != nil {
	// 	fmt.Println("Не удалось записать файл: ", err.Error())
	// }

	// Пример чтения данных из файла отчета API
	// result, err := getFactInsuarancePriceFromAPI("Сверочный_отчет_2.xlsx")
	// result, err := getFactBasePriceFromAPI("Сверочный_отчет_2.xlsx")

	// if err != nil {
	// log.Fatalf("Ошибка при чтении файла сверочного отчета: %v", err)
	// }
	// fmt.Println(len(result))

	dataFromReport, err := getBasePriceFromReportFile("Золотое яблоко сентябрь 2025 ОЧ.xlsx")

	if err != nil {
		log.Fatalf("Ошибка при чтении файла отчета ВКС: %v", err)
	}
	fmt.Println(len(dataFromReport))

	diffBasePrice := make(map[string][]float64, len(callbacks))

	for i := range len(callbacks) {
		orderNum := callbacks[i].OrderNumber
		callBackPrice := callbacks[i].BasePrice

		orderData, ok := dataFromReport[orderNum]
		if !ok {
			continue
		}

		diffBasePrice[orderNum] = []float64{
			callBackPrice,
			orderData.BasePrice,
		}

		// println(orderNum, factPrice, orderData.BasePrice)
	}

	// for key, basePriceFact := range result {
	// 	orderInfo, ok := dataFromReport[key]
	//  	if !ok {
	// 		continue
	// 	}

	// 	diffBasePrice[key] = []float64{
	// 		basePriceFact,
	// 		// orderInfo.InsurancePay,
	// 		orderInfo.BasePrice,
	// 	}
	// }

	fmt.Println(len(diffBasePrice))
	err = writeToExcel("order_prices.xlsx", diffBasePrice)
	if err != nil {
		log.Fatalf("Ошибка при записи в Excel: %v", err)
	}

}

func writeToExcelDiff(filename string, callback *[]callback.DataForTable, tarrifs *tariffs.TariffStore) error {
	// Создаем новый файл Excel.
	f := excelize.NewFile()
	defer func() {
		// Закрываем файл, чтобы избежать утечек ресурсов.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Создаем новый лист.
	sheetName := "Order Prices"
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("не удалось создать лист: %w", err)
	}

	// Удаляем лист по умолчанию "Sheet1".
	f.DeleteSheet("Sheet1")

	// Записываем заголовки.
	headers := []string{"№ отправления заказчика", "Город / населенный пункт", "Вес", "Тарифная зона (из отчета)", "Объявленная ценность", "Стоимость услуг по приему отправлений (страховка)", "Стоимость услуг по доставке (API)", "Тарифная зона (расчет)", "Тариф (расчет)"}
	// Устанавливаем стиль для заголовков (жирный шрифт)
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	if err != nil {
		return fmt.Errorf("не удалось создать стиль: %w", err)
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, style)
	}
	// Записываем данные из map.
	rowNum := 2 // Начинаем со второй строки, так как первая - заголовки.
Exit:
	for _, order := range *callback {
		weight := order.Weight / 1000
		var weightType string
		if weight <= 5.00 {
			weightType = "less_5"
		} else {
			weightType = "more_5"
		}

		planTarif, ok := tarrifs.Get(order.CityTo, weightType)

		if !ok || len(planTarif) == 0 {
			// Если для города/веса не найдено ни одного тарифа, можно либо пропустить,
			// либо записать строку с пустыми ячейками для тарифа. Пропускаем.
			continue
		}

		// Для каждого найденного тарифа для данного города/веса создаем отдельную строку
		for _, tariff := range planTarif {
			if tariff.Tariff == order.BasePrice {
				continue Exit
			}
		}
		for _, tariff := range planTarif {

			f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNum), order.OrderNumber)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNum), order.CityTo)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNum), order.Weight)
			f.SetCellValue(sheetName, "D"+strconv.Itoa(rowNum), order.Tariff)        // Тарифная зона из исходных данных
			f.SetCellValue(sheetName, "E"+strconv.Itoa(rowNum), order.AssessedValue) // ОЦ
			f.SetCellValue(sheetName, "F"+strconv.Itoa(rowNum), order.InsurancePay)  // Страховка
			f.SetCellValue(sheetName, "G"+strconv.Itoa(rowNum), order.BasePrice)     // Цена из API
			f.SetCellValue(sheetName, "H"+strconv.Itoa(rowNum), tariff.TariffZone)   // Расчетная тарифная зона
			f.SetCellValue(sheetName, "I"+strconv.Itoa(rowNum), tariff.Tariff)       // Расчетный тариф
			rowNum++
		}
	}

	// Устанавливаем автоширину для колонок
	f.SetColWidth(sheetName, "A", "G", 25)

	// Сохраняем файл.
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("не удалось сохранить файл: %w", err)
	}

	fmt.Printf("Данные успешно записаны в файл %s\n", filename)
	return nil
}

func writeToExcelDiff2(filename string, callback *[]callback.DataForTable, tarrifs *tariffs.TariffStore) error {
	// Создаем новый файл Excel.
	f := excelize.NewFile()
	defer func() {
		// Закрываем файл, чтобы избежать утечек ресурсов.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Создаем новый лист.
	sheetName := "Order Prices"
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("не удалось создать лист: %w", err)
	}

	// Удаляем лист по умолчанию "Sheet1".
	f.DeleteSheet("Sheet1")

	// Записываем заголовки.
	headers := []string{"№ отправления заказчика", "Город / населенный пункт", "Вес", "Тарифная зона (из отчета)", "Объявленная ценность", "Стоимость услуг по приему отправлений (страховка)", "Стоимость услуг по доставке (API)", "Тарифная зона (расчет)", "Тариф (расчет)"}
	// Устанавливаем стиль для заголовков (жирный шрифт)
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	if err != nil {
		return fmt.Errorf("не удалось создать стиль: %w", err)
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, style)
	}
	// Записываем данные из map.
	rowNum := 2 // Начинаем со второй строки, так как первая - заголовки.
Exit:
	for _, order := range *callback {
		weight := order.Weight / 1000
		var weightType string
		if weight <= 5.00 {
			weightType = "less_5"
		} else {
			weightType = "more_5"
		}

		planTarif, ok := tarrifs.Get(order.CityTo, weightType)

		if !ok || len(planTarif) == 0 {
			// Если для города/веса не найдено ни одного тарифа, можно либо пропустить,
			// либо записать строку с пустыми ячейками для тарифа. Пропускаем.
			continue
		}

		// Для каждого найденного тарифа для данного города/веса создаем отдельную строку
		for _, tariff := range planTarif {
			if tariff.Tariff == order.BasePrice {
				continue Exit
			}
		}
		for _, tariff := range planTarif {

			f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNum), order.OrderNumber)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNum), order.CityTo)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNum), order.Weight)
			f.SetCellValue(sheetName, "D"+strconv.Itoa(rowNum), order.Tariff)        // Тарифная зона из исходных данных
			f.SetCellValue(sheetName, "E"+strconv.Itoa(rowNum), order.AssessedValue) // ОЦ
			f.SetCellValue(sheetName, "F"+strconv.Itoa(rowNum), order.InsurancePay)  // Страховка
			f.SetCellValue(sheetName, "G"+strconv.Itoa(rowNum), order.BasePrice)     // Цена из API
			f.SetCellValue(sheetName, "H"+strconv.Itoa(rowNum), tariff.TariffZone)   // Расчетная тарифная зона
			f.SetCellValue(sheetName, "I"+strconv.Itoa(rowNum), tariff.Tariff)       // Расчетный тариф
			rowNum++
		}
	}

	// Устанавливаем автоширину для колонок
	f.SetColWidth(sheetName, "A", "G", 25)

	// Сохраняем файл.
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("не удалось сохранить файл: %w", err)
	}

	fmt.Printf("Данные успешно записаны в файл %s\n", filename)
	return nil
}

// func getInfoFromDHFiles(filename string) {
// 	jsonBytes, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var data []OrderInfo

// 	err = json.Unmarshal(jsonBytes, &data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, order := range data {
// 		if order.DeliveryService == "5Post" {
// 			if order.BasePrice.TotalSumFact > 0 &&
// 				order.BasePrice.TotalSumPlan != order.BasePrice.TotalSumFact &&
// 				order.FactRequest.DeliveryDate != "" {
// 				parsedTime, err := time.Parse("2006-01-02", order.FactRequest.DeliveryDate)
// 				if err != nil {

// 					log.Fatal(err)
// 				}
// 				if parsedTime.Month() != time.July {
// 					fmt.Printf("Fact: %f ", order.BasePrice.TotalSumFact)
// 					fmt.Printf("Plan: %f \n", order.BasePrice.TotalSumPlan)
// 					fmt.Println(order.OrderNumber + "/" + parsedTime.Month().String())
// 				}
// 			}
// 		}
// 	}
// }

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

const (
	// Константы для getFactBasePriceFromReportFile
	reportSheetIndex          = 0
	reportHeaderRows          = 7
	reportOrderNumberCol      = 0
	reportBasePriceFactCol    = 10
	reportFactInsuarancePrice = 16
	// Константы для getBasePriceFromReport
	svodSheetIndex         = 0
	svodHeaderRows         = 1
	svodOrderNumberCol     = 1
	svodStatusCol          = 2
	svodBasePriceCol       = 13
	svodInsurancePayCol    = 16
	svodBasePriceReturnCol = 17
)

func getFactBasePriceFromAPI(filename string) (map[string]float64, error) {
	rows, err := getRows(filename, reportSheetIndex)
	if err != nil {
		return nil, err
	}
	mapFactBasePrice := make(map[string]float64, len(rows))

	before := time.Now()
	for i, row := range rows {
		if i < reportHeaderRows { // Пропускаем заголовок
			continue
		}

		orderNumber := row[reportOrderNumberCol]
		basePriceFactStr := row[reportBasePriceFactCol]
		basePriceFactFloat, err := strconv.ParseFloat(basePriceFactStr, 64)
		if err != nil {
			log.Printf("Ошибка конвертации 'basePriceFact' в строке %d: %v. Пропускаем.", i+1, err)
			continue
		}
		if basePriceFactFloat != 0 {
			mapFactBasePrice[orderNumber] = basePriceFactFloat
		}
	}
	fmt.Printf("Время обработки файла отчета: %s\n", time.Since(before))

	return mapFactBasePrice, nil
}

func getFactInsuarancePriceFromAPI(filename string) (map[string]float64, error) {
	rows, err := getRows(filename, reportSheetIndex)
	if err != nil {
		return nil, err
	}
	mapFactBasePrice := make(map[string]float64, len(rows))

	before := time.Now()
	for i, row := range rows {
		if i < reportHeaderRows { // Пропускаем заголовок
			continue
		}

		orderNumber := row[reportOrderNumberCol]
		insuarancePriceFactStr := row[reportFactInsuarancePrice]
		insuarancePriceFactFloat, err := strconv.ParseFloat(insuarancePriceFactStr, 64)
		if err != nil {
			log.Printf("Ошибка конвертации 'basePriceFact' в строке %d: %v. Пропускаем.", i+1, err)
			continue
		}
		if insuarancePriceFactFloat != 0 {
			mapFactBasePrice[orderNumber] = insuarancePriceFactFloat
		}
	}
	fmt.Printf("Время обработки файла отчета: %s\n", time.Since(before))

	return mapFactBasePrice, nil
}

func getBasePriceFromReportFile(fileName string) (map[string]OrderData, error) {
	rows, err := getRows(fileName, svodSheetIndex)
	if err != nil {
		return nil, err
	}

	mapResult := make(map[string]OrderData, len(rows)-svodHeaderRows)
	ordersData := ListOrdersData{}

	before := time.Now()

	for i, row := range rows {
		if i < svodHeaderRows {
			continue
		}
		orderNumber := row[svodOrderNumberCol]

		basePriceStr := row[svodBasePriceCol]
		basePriceFloat, err := strconv.ParseFloat(basePriceStr, 64)
		if err != nil {
			log.Printf("Ошибка конвертации 'basePriceFloat' в строке %d: %v", i+1, err)
			continue
		}

		basePriceReturnStr := row[svodBasePriceReturnCol]
		basePriceReturnFloat, err := strconv.ParseFloat(basePriceReturnStr, 64)
		if err != nil {
			log.Printf("Ошибка конвертации 'basePriceReturnFloat' в строке %d: %v", i+1, err)
			continue
		}

		insurancePayStr := row[svodInsurancePayCol]
		insurancePayFloat, err := strconv.ParseFloat(insurancePayStr, 64)
		if err != nil {
			log.Printf("Ошибка конвертации 'insurancePayFloat' в строке %d: %v", i+1, err)
			continue
		}

		status := row[svodStatusCol]

		orderData := OrderData{
			OrderNumber:     orderNumber,
			Status:          status,
			BasePrice:       basePriceFloat,
			BasePriceReturn: basePriceReturnFloat,
			InsurancePay:    insurancePayFloat,
		}
		ordersData.Orders = append(ordersData.Orders, orderData)
		mapResult[orderNumber] = orderData

	}
	fmt.Printf("Время обработки сводного отчета: %s\n", time.Since(before))

	data, err := ToBytes(&ordersData)
	if err != nil {
		return nil, err
	}
	if err := Write(data, "orders_5post.json"); err != nil {
		return nil, err
	}

	return mapResult, nil
}

func ToBytes(listOrdersData *ListOrdersData) ([]byte, error) {
	file, err := json.Marshal(listOrdersData)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Write(content []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("не удалось записать в файл %s: %w", filename, err)
	}
	fmt.Printf("Данные успешно записаны в JSON файл %s\n", filename)
	return nil
}

// writeToExcel демонстрирует, как создать файл Excel и записать в него данные.
func writeToExcel(filename string, diffData map[string][]float64) error {
	// Создаем новый файл Excel.
	f := excelize.NewFile()
	defer func() {
		// Закрываем файл, чтобы избежать утечек ресурсов.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Создаем новый лист.
	sheetName := "Order Prices"
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("не удалось создать лист: %w", err)
	}

	// Удаляем лист по умолчанию "Sheet1".
	f.DeleteSheet("Sheet1")

	// Записываем заголовки.
	headers := []string{"Заказ", "API", "Отчет КС"}
	// Устанавливаем стиль для заголовков (жирный шрифт)
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	if err != nil {
		return fmt.Errorf("не удалось создать стиль: %w", err)
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, style)
	}

	// Записываем данные из map.
	rowNum := 2 // Начинаем со второй строки, так как первая - заголовки.
	for orderNumber, basePrices := range diffData {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNum), orderNumber)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNum), basePrices[0])
		f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNum), basePrices[1])

		rowNum++
	}

	// Устанавливаем автоширину для колонок
	f.SetColWidth(sheetName, "A", "C", 20)

	// Сохраняем файл.
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("не удалось сохранить файл: %w", err)
	}

	fmt.Printf("Данные успешно записаны в файл %s\n", filename)
	return nil
}
