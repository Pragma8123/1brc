package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Station struct {
	Min   float64
	Sum   float64
	Max   float64
	Count float64
}

var globalStations = make(map[string]*Station)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	chunkSize := 1_000_000 // Smaller chunk size for better load distribution
	chunks := make(chan []string, runtime.NumCPU()*2)

	// Start worker pool
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(chunks, &wg)
	}

	// Read file in chunks and send to workers
	var chunk []string
	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())
		if len(chunk) >= chunkSize {
			chunks <- chunk
			chunk = make([]string, 0, chunkSize)
		}
	}

	// Send any remaining lines as a final chunk
	if len(chunk) > 0 {
		chunks <- chunk
	}
	close(chunks)

	// Wait for all workers to finish
	wg.Wait()

	printData()
}

func worker(chunks chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for chunk := range chunks {
		localStations := processChunk(chunk)
		mergeStations(localStations)
	}
}

func processChunk(chunk []string) map[string]*Station {
	localStations := make(map[string]*Station)
	for _, value := range chunk {
		station, temp := parseLine(value)
		if i, ok := localStations[station]; !ok {
			localStations[station] = &Station{
				Min:   temp,
				Sum:   temp,
				Max:   temp,
				Count: 1,
			}
		} else {
			if i.Min > temp {
				i.Min = temp
			}
			if i.Max < temp {
				i.Max = temp
			}
			i.Count++
			i.Sum += temp
		}
	}
	return localStations
}

func parseLine(data string) (string, float64) {
	split := strings.Split(data, ";")
	station := split[0]
	temp, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	return station, temp
}

func mergeStations(localStations map[string]*Station) {
	for station, localStation := range localStations {
		if globalStation, ok := globalStations[station]; !ok {
			globalStations[station] = localStation
		} else {
			if globalStation.Min > localStation.Min {
				globalStation.Min = localStation.Min
			}
			if globalStation.Max < localStation.Max {
				globalStation.Max = localStation.Max
			}
			globalStation.Count += localStation.Count
			globalStation.Sum += localStation.Sum
		}
	}
}

func printData() {
	keys := make([]string, 0, len(globalStations))
	for k := range globalStations {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("{")
	for _, k := range keys {
		value := globalStations[k]
		fmt.Printf("%v=%.1f/%.1f/%.1f", k, value.Min, value.Sum/value.Count, value.Max)
		if k != keys[len(keys)-1] {
			fmt.Print(",")
		}
	}
	fmt.Print("}")
}
