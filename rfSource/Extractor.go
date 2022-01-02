package rfsource

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/dreadl0ck/debias"
)

func ExtractEntropy(rawData []int, LowPosEntrop float64, outputFile string) {
	var bitString string

	var minEntropy float64

	bitString = convertBinary(rawData)

	minEntropy = MinEntropy(bitString)

	if minEntropy < LowPosEntrop {
		log.Panic("Error: Insufficent Min-Entropy\nShutting Program down")
	}

	debaisData(bitString, outputFile)
}

func MinEntropy(data string) float64 { // calculates the min-entropy of a bit string

	var numberOnes int
	var numberBits int
	var OnesToZero float64

	for _, value := range data { //find the number of ones in the bit string

		numberBits++

		if value == '1' {
			numberOnes++
		}
	}

	OnesToZero = float64(numberOnes / numberBits)

	return -math.Log2(OnesToZero)
}

func debaisData(data string, outputName string) {
	var cxt context.Context
	
	reader := bytes.NewReader([]byte(rawdata))

	pr, cxt, _ := debias.Kaminsky(reader, false, 512) //debaises data


	<- cxt.Done()   //once its done debaising

	f, err := os.Open(outputName)

	if err != nil {
		log.Fatal(err)
	}
	
	data, err := ioutil.ReadAll(pr)

	if err != nil {
		log.Fatal(err)
	}

	n, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Wrote:%s bytes",  n)

	err = f.Close() //close file
	if err != nil {
		log.Fatal(err)
	}
}

func convertBinary(input []int) string {

	var (
		output string
	)

	for _, value := range input {
		output += strconv.FormatInt(int64(value), 2)
	}

	return output
}
