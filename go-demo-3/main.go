package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	m := map[string]string{
		"PurpleSchool": "http://purplescholl",
	}
	m["Google"] = "https://google.com"
	m["Yandex"] = "mhttps://yandex.ru"

	fmt.Println(m)

	delete(m, "Yandex")
	delete(m, "Y")

	fmt.Println(m)

	n := map[string]int{"a": 1, "b": 2}
	for key, value := range n {
		fmt.Println(key, value)
	}

	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Введите команду:")
	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newKey string
	var newValue string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&newKey)
	fmt.Print("Введите значение: ")
	fmt.Scan(&newValue)
	bookmarks[newKey] = newValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var keyToDelete string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&keyToDelete)
	delete(bookmarks, keyToDelete)
}
