package handler

import (
	"net/http"
	"strconv"
	"text/template"
	"unit-converter/internal/services"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./internal/templates/index.html"))
	temp.Execute(w, nil)
}

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	path := "./internal/templates/length.html"

	if r.Method == http.MethodPost {
		value, _ := strconv.ParseFloat(r.FormValue("lengthValue"), 64)
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")

		result := services.LengthConvert(value, fromUnit, toUnit)
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, map[string]interface{}{
			"Result": result,
			"ToUnit": toUnit,
		})
	} else {
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, nil)
	}
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	path := "./internal/templates/weight.html"

	if r.Method == http.MethodPost {
		value, _ := strconv.ParseFloat(r.FormValue("weightVlaue"), 64)
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")

		result := services.LengthConvert(value, fromUnit, toUnit)
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, map[string]interface{}{
			"Result": result,
			"ToUnit": toUnit,
		})
	} else {
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, nil)
	}
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	path := "./internal/templates/temperature.html"

	if r.Method == http.MethodPost {
		value, _ := strconv.ParseFloat(r.FormValue("temperatureValue"), 64)
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")

		result := services.LengthConvert(value, fromUnit, toUnit)
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, map[string]interface{}{
			"Result": result,
			"ToUnit": toUnit,
		})
	} else {
		tmpl := template.Must(template.ParseFiles(path))
		tmpl.Execute(w, nil)
	}
}
