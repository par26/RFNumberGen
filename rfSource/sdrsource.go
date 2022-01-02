package rfsource

import (
	"io"
	"log"
	"os"

	"github.com/youpy/go-wav"
)

//var RtlCmd string = "/usr/bin/rtl_fm -g 50 -f%s -M wfm -s 180k -E deemp | usr/bin/tee %s | play -q -r 180k -t raw -e s -b 16 -c 1 -V1 - lowpass 16k"

type SdrSource struct {
	NoiseFile string

	SdrBuffer []int

	ReadPos int
}

func NewSdrSource(noiseFile string, bufferSize uint16) *SdrSource {
	return &SdrSource{
		NoiseFile: noiseFile,
		SdrBuffer: make([]int, 0, bufferSize),
	}
}

/*func (sdr *SdrSource) GetNoise() {
	file, err := os.Open(sdr.NoiseDirectory)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	n, err := r.Read(sdr.SdrBuffer[:cap(sdr.SdrBuffer)]) // read bytes
	if err != nil {
		log.Fatal(err)
	}
	sdr.ReadPos += n //advance the read position for the next read
}*/

func (sdr *SdrSource) GetWavData() { // reads wav source and store into the sdr buffer

	file, err := os.Open(sdr.NoiseFile)

	if err != nil {
		log.Panic(err)
	}

	reader := wav.NewReader(file)

	defer file.Close() //close the file when done

	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			sdr.SdrBuffer = append(sdr.SdrBuffer, reader.IntValue(sample, 0)) //adds sample value to the
		}

	}
}
