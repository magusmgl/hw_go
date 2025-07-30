package output

import "github.com/fatih/color"

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки %d", intValue)
		return
	}

	stringValue, ok := value.(string)
	if ok {
		color.Red(stringValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}

	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Код ошибки %d", t)
	// case error:
	// 	color.Red(t.Error())
	// default:
	// 	color.Red("Неизвестный тип ошибки")
	// }
	color.Red("Неизвестный тип ошибки")
}
