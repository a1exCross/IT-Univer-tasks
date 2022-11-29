package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"microservice/solutions"
	"net/http"
	"os"
	"reflect"
)

const serviceUrl = "https://kuvaev-ituniversity.vps.elewise.com/tasks/"

type handler struct {
	UserName string
}

const (
	WeirdArrayEntry    string = "Чудные вхождения в массив"
	SequenceCheck      string = "Проверка последовательности"
	FindMissingElement string = "Поиск отсутствующего элемента"
	CyclicRotation     string = "Циклическая ротация"
)

func NewHandler(username string) *handler {
	return &handler{
		UserName: username,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	switch r.URL.Path {
	case "/task/" + CyclicRotation:
		res := h.taskOne(w, r)

		_, err := w.Write([]byte(res))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		return
	case "/task/" + WeirdArrayEntry:
		res := h.task(w, r, WeirdArrayEntry)

		_, err := w.Write([]byte(res))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		return
	case "/task/" + SequenceCheck:
		res := h.task(w, r, SequenceCheck)

		_, err := w.Write([]byte(res))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		return
	case "/task/" + FindMissingElement:
		res := h.task(w, r, FindMissingElement)

		_, err := w.Write([]byte(res))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		return
	case "/tasks":
		var result []interface{}

		result = append(result, string(h.taskOne(w, r)))
		result = append(result, string(h.task(w, r, WeirdArrayEntry)))
		result = append(result, string(h.task(w, r, SequenceCheck)))
		result = append(result, string(h.task(w, r, FindMissingElement)))

		str := fmt.Sprintf("%v", result)

		_, err := w.Write([]byte(str))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func getDataSet(url string) (*[]TasksResult, [][]interface{}, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	var data [][]interface{}

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, nil, err
	}

	var results []TasksResult

	for _, arr := range data {
		var res TasksResult

		for _, val := range arr {
			v := reflect.ValueOf(val)

			switch v.Type().Kind() {
			case reflect.Slice:
				for i := 0; i < v.Len(); i++ {
					res.Slice = append(res.Slice, int(v.Index(i).Elem().Float()))
				}
			default:
				log.Println("any!")
			}
		}
		results = append(results, res)
	}

	return &results, data, nil
}

type Results struct {
	Payload interface{} `json:"payload"`
	Results interface{} `json:"results"`
}

type Response struct {
	UserName string  `json:"user_name"`
	Task     string  `json:"task"`
	Result   Results `json:"results"`
}

type TasksResult struct {
	Slice []int
}

func (h *handler) task(w http.ResponseWriter, r *http.Request, taskName string) []byte {
	dataset, data, err := getDataSet(serviceUrl + taskName)

	if err != nil {
		log.Println(err)
		return nil
	}

	var result []interface{}

	switch taskName {
	case WeirdArrayEntry:
		for _, field := range *dataset {
			result = append(result, solutions.WeirdArrayEntry(field.Slice))
		}
	case SequenceCheck:
		for _, field := range *dataset {
			result = append(result, solutions.SequenceCheck(field.Slice))
		}
	case FindMissingElement:
		for _, field := range *dataset {
			result = append(result, solutions.FindMissingElement(field.Slice))
		}
	}

	response := Response{
		UserName: h.UserName,
		Task:     taskName,
		Result: Results{
			Payload: data,
			Results: result,
		},
	}

	body, err := json.Marshal(response)
	if err != nil {
		return nil
	}

	answer, err := sendResultOnCheck("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", body)
	if err != nil {
		return nil
	}

	log.Println(string(answer))

	return answer
}

func sendResultOnCheck(url string, data []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getDataSetTaskOne(url string) (*[]TaskOneResult, [][]interface{}, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	var data [][]interface{}

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var results []TaskOneResult

	for _, arr := range data {
		var res TaskOneResult

		for _, val := range arr {
			v := reflect.ValueOf(val)

			switch v.Type().Kind() {
			case reflect.Slice:
				for i := 0; i < v.Len(); i++ {
					res.Slice = append(res.Slice, int(v.Index(i).Elem().Float()))
				}
			case reflect.Float64:
				res.Shifts = int(v.Float())
			default:
				log.Println("any!")
			}
		}
		results = append(results, res)
	}

	return &results, data, nil
}

type TaskOneResult struct {
	Slice  []int
	Shifts int
}

func (h *handler) taskOne(w http.ResponseWriter, r *http.Request) []byte {
	dataset, data, err := getDataSetTaskOne(serviceUrl + CyclicRotation)
	if err != nil {
		log.Println(err)
		return nil
	}

	var result []interface{}

	for _, field := range *dataset {
		result = append(result, solutions.CyclicRotation(field.Slice, field.Shifts))
	}

	response := Response{
		UserName: h.UserName,
		Task:     CyclicRotation,
		Result: Results{
			Payload: data,
			Results: result,
		},
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return nil
	}

	answer, err := sendResultOnCheck("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", body)
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println(string(answer))

	return answer
}
