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

	var numBytesWritten int

	reader := bytes.NewReader([]byte(data))

	pr, _, _ := debias.Kaminsky(reader, false, 512)

	f, err := os.Create(outputName)

	if err != nil {
		log.Fatal(err)
	}

	for { //writes the conditioned data to the file
		var data = make([]byte, len(data))

		n, err := pr.Read(data)

		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				fmt.Println(err)
				break
			}
			log.Fatal(err)
		}

		data = data[:n]

		// write output buffer
		n, err = f.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		numBytesWritten += n
	}

	// close output file handle
	err = f.Close()
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
