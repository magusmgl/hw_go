package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "Lontisn", "Город пользователя")
	format := flag.Int("format", 1, "Формат вызова погоды")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(geoData)

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
	// r := strings.NewReader("dsdsd sds")
	// block := make([]byte, 4)
	// for {
	// 	_, err := r.Read(block)
	// 	fmt.Printf("%q\n", block)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

}
