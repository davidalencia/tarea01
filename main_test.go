package main

import (
	"testing"
	"strings"
)

var l = &location{
        lat:19.4322,
        lon:-99.1772,
	name:"name",
	weather: "bad",
}

func TestClimateUrlFromLocation (t *testing.T){
    s := climateUrlFromLocation(l)
    if(s!="http://api.openweathermap.org/data/2.5/weather?lang=es&lat=19.432200&lon=-99.177200&APPID=352ab509f62f8707940ca19d3ab12341"){
        t.Error("aaaaahhhhh!!! url wrong")
    }
}

func TestSetWeatherChanged(t *testing.T){
    c := make(chan *location, 1)
    c <- l
    setWeather(c)
    loc := <-c
    if(loc.weather=="bad"){
	    t.Error("api failed")
    }
}

func TestSetWeatherSeemsToWork(t *testing.T){
    c := make(chan *location, 1)
    c <- l
    setWeather(c)
    loc := <-c
    if(len(strings.Split(loc.weather, " "))<=1){
            t.Error("api failed")
    }
}
