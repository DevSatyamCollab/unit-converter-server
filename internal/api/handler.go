package api

import (
	"encoding/json"
	"net/http"
	"unit-converter/internal/services"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	// client Request
	params := Converter{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid payload Request")
		return
	}

	// server work
	ans := services.LengthConvert(params.Value, params.FromUnit, params.ToUnit)

	// server response to client
	response(w, http.StatusOK, Converter{
		Value:    params.Value,
		FromUnit: params.FromUnit,
		ToUnit:   params.ToUnit,
		Ans:      ans,
	})

}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	// client Request
	params := Converter{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid payload Request")
		return
	}

	// server work
	ans := services.WeightConvert(params.Value, params.FromUnit, params.ToUnit)

	// server response to client
	response(w, http.StatusOK, Converter{
		Value:    params.Value,
		FromUnit: params.FromUnit,
		ToUnit:   params.ToUnit,
		Ans:      ans,
	})

}

func TemperatureHanlder(w http.ResponseWriter, r *http.Request) {
	// client Request
	params := Converter{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid payload Request")
		return
	}

	// server work
	ans := services.TemperatureConvert(params.Value, params.FromUnit, params.ToUnit)

	// server response to client
	response(w, http.StatusOK, Converter{
		Value:    params.Value,
		FromUnit: params.FromUnit,
		ToUnit:   params.ToUnit,
		Ans:      ans,
	})
}
