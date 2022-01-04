package rfsource

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dreadl0ck/debias"
)

func ExtractEntropy(rawdata []int, outputFile string) {
	WriteToFile(rawdata, outputFile)
}

func debaisData(rawdata []int) []byte {

	stringData := convertToString(rawdata)

	buffer := &bytes.Buffer{}

	gob.NewEncoder(buffer).Encode(stringData)
	byteSlice := buffer.Bytes()

	reader := bytes.NewReader(byteSlice)

	pr, _, _ := debias.Kaminsky(reader, false, int64(debias.MaxChunkSize)) //debaises data

	var data = make([]byte, len(rawdata))
	n, err := pr.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	data = data[:n]

	return data
}

func convertToString(input []int) []string {
	output := make([]string, len(input))

	var stringval string

	for i := 0; i < len(input); i++ {
		stringval = strconv.Itoa(input[i])
		output = append(output, stringval)
	}

	return output
}

func WriteToFile(input []int, directory string) { //writes to a bin file

	buf := new(bytes.Buffer)

	data := debaisData(input)

	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(directory, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	n, err := f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number Bytes outputed", n)

	defer f.Close()
}
