package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var data [][]string
var contcorrectas = 0.0
var contadas = 0

type Resultado struct {
	Accuracy float64 `json:"acc"`
}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Distancia(row1 []string, row2 []string) float64 {
	dist := 0.0
	for idx, el1 := range row1 {
		if idx != 5 {
			val1, _ := strconv.ParseFloat(el1, 64)
			val2, _ := strconv.ParseFloat(row2[idx], 64)
			dist = dist + math.Pow(val1-val2, 2)
		}
	}
	dist = math.Sqrt(dist)
	return dist
}

func Ordenar(dists []float64, aclass []int) []int {
	for i := len(aclass); i > 0; i-- {
		for j := 1; j < i; j++ {
			if dists[j-1] > dists[j] {
				intermediated := dists[j]
				intermediatea := aclass[j]
				dists[j] = dists[j-1]
				aclass[j] = aclass[j-1]
				dists[j-1] = intermediated
				aclass[j-1] = intermediatea
			}
		}
	}
	return aclass
}

func Verificar(aclassord []int, row1 []string, contc int, k int, test [][]string, chresult chan float64) {
	for i := 0; i < k; i++ {
		if aclassord[i] == 0 {
			contc++
		}
	}
	if contc > (k - contc) {
		val, _ := strconv.Atoi(row1[5])
		if val == 0 {
			contcorrectas++
		}
	} else {
		val, _ := strconv.Atoi(row1[5])
		if val == 1 {
			contcorrectas++
		}
	}
	contadas++
	if contadas >= len(test) {
		chresult <- contcorrectas / float64(len(test))
		contcorrectas = 0.0
		contadas = 0
	}
}

func Predecir(row1 []string, train [][]string, k int, test [][]string, chresult chan float64) {
	dists := []float64{}
	aclass := []int{}
	for _, row2 := range train {
		d := Distancia(row1, row2)
		dists = append(dists, d)
		val, _ := strconv.Atoi(row2[5])
		aclass = append(aclass, val)
	}
	aclassord := Ordenar(dists, aclass)
	contc := 0
	go Verificar(aclassord, row1, contc, k, test, chresult)
}

func KNN(data [][]string, testp float64, k int, chresult chan float64) {
	test := [][]string{}
	train := [][]string{}
	for idx, row := range data {
		if idx == 0 {
			continue
		} else if idx <= int(float64(len(data))*testp) {
			test = append(test, row)
		} else if idx > int(float64(len(data))*testp) {
			train = append(train, row)
		}
	}
	for _, row1 := range test {
		go Predecir(row1, train, k, test, chresult)
	}
}

func cargarData() {
	url := "https://gist.githubusercontent.com/Hokaid/ccc8cab127b32ab08bcb427f8c01eb4c/raw/9d5cd07c12e3295a09ee4ec6d1f346d3cb1208f5/dtrafic.csv"
	data, _ = readCSVFromUrl(url)
}

func testKNN(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	sk := req.FormValue("k")
	k, _ := strconv.Atoi(sk)
	sptest := req.FormValue("ptest")
	ptest, _ := strconv.ParseFloat(sptest, 64)
	chresult := make(chan float64)
	go KNN(data, ptest, k, chresult)
	result := <-chresult
	accuracy := Resultado{result}
	ojsonBytes, _ := json.MarshalIndent(accuracy, "", " ")
	io.WriteString(res, string(ojsonBytes))
}

//manejador de peticiones
func handleRequest() {
	http.HandleFunc("/test", testKNN)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	cargarData()
	handleRequest()
}
