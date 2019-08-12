package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
    "strconv"
)

type Location struct{
  name string
  lat float64
  lon float64
  weather string
}

func (l *Location) setWeather()  {
  l.weather = "new weather"
}


func main()  {
    f, _ := os.Open("dataset.csv")
    r := csv.NewReader(bufio.NewReader(f))

    m := make(map[string]*Location)

    for line, e := r.Read(); e != io.EOF; line, e = r.Read(){
      for alfa := 0; alfa<=1; alfa++{
        lon, _ := strconv.ParseFloat((line[alfa*2+2]), 64)
        lat, _ := strconv.ParseFloat((line[alfa*2+3]), 64)
        m[line[alfa]] = &Location {
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
    fmt.Println(m["MEX"].weather)
}
