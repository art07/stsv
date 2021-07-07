package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type Decoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})

	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

/*type Service struct {
	log Logger
}*/

/*-------------------------------------------------------------*/

type MyDecoder struct {
	logger Logger
}

func (d MyDecoder) Decode(data []byte) (persons []Person, places []Place) {
	var allThings AllThings
	if err := json.Unmarshal(data, &allThings); err != nil {
		log.Println(err)
	}
	for _, item := range allThings.Things {
		if item.Name != "" {
			persons = append(persons, Person{Name: item.Name, Age: item.Age})
		} else {
			places = append(places, Place{City: item.City, Country: item.Country})
		}
	}
	return
}

func (d MyDecoder) Sort(dataToSort interface{}) {
	_, ok := dataToSort.([]Person)
	if ok {
		persons := dataToSort.([]Person)
		sort.Slice(persons, func(i, j int) bool {
			return persons[i].Age < persons[j].Age
		})
	} else {
		places := dataToSort.([]Place)
		sort.Slice(places, func(i, j int) bool {
			return len(places[i].City) < len(places[j].City)
		})
	}
}

func (d MyDecoder) Print(arr interface{}) {
	d.logger.Println(fmt.Sprintf("%v", arr))
}

func (d MyDecoder) Printlen(persons []Person, places []Place) {
	d.logger.Println(fmt.Sprintf("%d %d", len(persons), len(places)))
}

/*-------------------------------------------------------------*/

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type AllThings struct {
	Things []struct {
		Name    string `json:"name,omitempty"`
		Age     int    `json:"age,omitempty"`
		City    string `json:"city,omitempty"`
		Country string `json:"country,omitempty"`
	} `json:"things"`
}
