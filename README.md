# One Billion Row Challenge
An attempt at the [One Billion Row Challenge](https://github.com/gunnarmorling/1brc) written in Go.

## Running the challenge
Follow the guide in the README for the official challenge using the link above to generate the measurement data.

1) Compile: `CGO=1 go build main.go`
2) Run: `time ./main measurements.txt`

## Benchmark History
### System Specs:
__OS__: WSL2 on Win11

__CPU__: Intel i5-12500

__RAM__: 32GB DDR4@3200GHz

__Disk__: Seagate FireCuda 520 NVMe

| Date          | Benchmark     | Sha                                      |
| ------------- |:-------------:| -----------------------------------------|
| 6/20/2024     | 38.8s         | f2f61b85a2f0d65419613c7d0ceb966fec66e22f |