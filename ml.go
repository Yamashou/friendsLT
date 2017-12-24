package main

import (
	"fmt"
	"image"
	"mime/multipart"
	"os"

	"github.com/Yamashou/elm"
	"gonum.org/v1/gonum/mat"
)

func ml(file multipart.File) int {
	w, err := os.Open("./w_test")
	if err != nil {
		panic(err)
	}
	b, err := os.Open("./beta_test")
	if err != nil {
		panic(err)
	}
	model, err := elm.UnmarshalBinaryFrom(w, b)

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	data := make([][]float64, 1)
	vec := elm.GetCharacteristic(elm.GetLBH(img))
	vec = append(vec, float64(1))
	data[0] = vec
	var testDataSet elm.DataSet

	testDataSet.Data = data
	testDataSet.XSize = len(data[0])
	testDataSet.YSize = 1
	var data2 mat.Dense
	testArray := mat.NewDense(1, 257, data[0])
	data2.Mul(&model.W, testArray.T())

	gData := elm.SetSigmoid(data2)
	var data3 mat.Dense
	data3.Mul(gData.T(), &model.Beta)
	if data3.At(0, 0) > 0 {
		fmt.Println("OK")
		return 1
	} else {
		return -1
	}

}
