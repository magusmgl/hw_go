package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUerl, err := url.Parse("http://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUerl.RawQuery = params.Encode()
	resp, err := http.Get(baseUerl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)

}
