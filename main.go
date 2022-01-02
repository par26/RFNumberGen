package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rocketmind26/EntropyRandomNumberGen/rfsource"
)

func main() {

	sdrSource := rfsource.NewSdrSource("C:/test1.wav", 20480)
	sdrSource.GetWavData()
	outPutEntropy(sdrSource.SdrBuffer, "/data.txt")

	rfsource.ExtractEntropy(sdrSource.SdrBuffer, .20, "outputfilename")
}

/*func outputData(data []byte, directory string) {
	outputNameNum := rand.Int()
	outputName := strconv.Itoa(outputNameNum)
	outputFileName := directory + outputName + ".txt"
	fo, err := os.Create(outputFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	w := bufio.NewWriter(fo)

	if _, err := w.Write(data); err != nil {
		log.Fatal(err)
	}

	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}*/

func outPutEntropy(input []int, directory string) {
	file, err := os.Open(directory)
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(input))
	if err != nil {
		log.Panic(err)
	}

	err = file.Sync()
	if err != nil {
		log.Panic(err)
	}

	fmt.Print("Finished writing data to the file")
}

