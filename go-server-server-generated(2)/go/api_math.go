/*
 * Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var oper Operation
	json.Unmarshal(reqBody, &oper)

	var result Result
	result.Result = oper.Number1 + oper.Number2
	_, err := json.Marshal(result)
	fmt.Fprintln(w, result.Result)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)
	}
}

func DivNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var oper Operation
	json.Unmarshal(reqBody, &oper)

	if oper.Number2 == 0 {
		json.NewEncoder(w).Encode("ERROR")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var result Result

		result.Result = oper.Number1 / oper.Number2
		_, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		} else {
			json.NewEncoder(w).Encode(result)
			w.WriteHeader(http.StatusOK)
		}

	}
}

func Multi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var oper Operation
	json.Unmarshal(reqBody, &oper)

	var result Result
	result.Result = oper.Number1 * oper.Number2
	_, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)
	}
}

func SubNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var oper Operation
	json.Unmarshal(reqBody, &oper)

	var result Result
	result.Result = oper.Number1 - oper.Number2
	_, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)
	}
}
