//Paquete con las funciones necesarias para ejecutar el programa
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

//Estructura para poder guardar la información necesaria sobre cada lugar incl
//uyendo clima, codigo, latitud y longitud
type location struct{
  name string
  lat float64
  lon float64
  weather string
}


//Dado un location actualiza el clima utilizando el api de openWeather
func setWeather(loc chan *location)  {
  l := <-loc
  json := apiCall(climateUrlFromLocation(l))
  l.weather = strings.Split(json, "\"")[17]
  loc <-l
}

//Dado un location te regresa el url con el que puedes utilizar el api de openWeather
func climateUrlFromLocation(l *location) string {
  const apiKey = "352ab509f62f8707940ca19d3ab12341"
  baseUrl :="http://api.openweathermap.org/data/2.5/weather"
  return fmt.Sprintf("%s?lang=es&lat=%f&lon=%f&APPID=%s", baseUrl, l.lat, l.lon, apiKey)
}


//Regresa una cadena qu es la respuesta de una llamada get al url pasado como parametro
func apiCall(url string) string{
  res, _ := http.Get(url)
  data, _ :=  ioutil.ReadAll(res.Body)
  return string(data)
}

//Imprime el clima de un vuelo, y toma como parametro 2 canales de location
func printWeather(loc0, loc1 chan *location){
    l0 := <-loc0
    l1 := <-loc1
    fmt.Printf("El clima en %s es %s, %s es %s \n", l0.name, l0.weather, l1.name, l1.weather)
    loc0 <-l0
    loc1 <-l1
}


//se ejucuta cuando se ejecuta el archivo
func main()  {
    f, _ := os.Open("dataset.csv")
    r := csv.NewReader(bufio.NewReader(f))
    lines, _ := r.ReadAll()

    m := make(map[string](chan *location))
   

    for _, line := range lines{
    for alfa := 0; alfa<=1; alfa++{
        lat, _ := strconv.ParseFloat((line[alfa*2+2]), 64)
        lon, _ := strconv.ParseFloat((line[alfa*2+3]), 64)
        m[line[alfa]] = make(chan *location, 1)
	m[line[alfa]] <- &location {
          name: line[alfa],
          lat: lat,
          lon: lon,
          weather: "",
        }
      }
    }

    for key, _ := range m {
      go setWeather(m[key])
    }
    setWeather(m["MEX"])
    for _, line := range lines{
       defer printWeather(m[line[0]], m[line[1]])
    }

}



