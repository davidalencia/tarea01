package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "strconv"
    "net/http"
    "io/ioutil"
    "strings"
)

type location struct{
  name string
  lat float64
  lon float64
  weather string
}


func (l *location) setWeather()  {
  json := apiCall(climateUrlFromLocation(l))
  l.weather = strings.Split(json, "\"")[17]
}

func climateUrlFromLocation(l *location) string {
  const apiKey = "352ab509f62f8707940ca19d3ab12341"
  baseUrl :="http://api.openweathermap.org/data/2.5/weather"
  return fmt.Sprintf("%s?lang=es&lat=%f&lon=%f&APPID=%s", baseUrl, l.lat, l.lon, apiKey)
}

func apiCall(url string) string{
  res, _ := http.Get(url)
  data, _ :=  ioutil.ReadAll(res.Body)
  return string(data)
}

func main()  {
    f, _ := os.Open("dataset.csv")
    r := csv.NewReader(bufio.NewReader(f))
    lines, _ := r.ReadAll()

    m := make(map[string]*location)

    for _, line := range lines{
      for alfa := 0; alfa<=1; alfa++{
        lat, _ := strconv.ParseFloat((line[alfa*2+2]), 64)
        lon, _ := strconv.ParseFloat((line[alfa*2+3]), 64)
        m[line[alfa]] = &location {
          name: line[alfa],
          lat: lat,
          lon: lon,
          weather: "",
        }
      }
    }

    for key, _ := range m {
      m[key].setWeather()
    }
    r = csv.NewReader(bufio.NewReader(f))
    for _, line := range lines{
	loc0 := m[line[0]]
	loc1 := m[line[1]]
    	fmt.Printf("El clima en %s es %s, %s es %s \n", loc0.name, loc0.weather, loc1.name, loc1.weather)
    }
}
