# One Billion Row Challenge

An attempt at the [One Billion Row Challenge](https://github.com/gunnarmorling/1brc) written in Go.

## Running the challenge

Follow the guide in the README for the official challenge using the link above to generate the measurement data.

1. Compile: `CGO=1 go build main.go`
2. Run: `time ./main measurements.txt`

## Benchmark History

### System Specs:

**OS**: WSL2 on Win11

**CPU**: Intel i5-12500

**RAM**: 32GB DDR4@3200GHz

**Disk**: Seagate FireCuda 520 NVMe

| Date      | Benchmark |                                         Commit SHA                                          |
| --------- | :-------: | :-----------------------------------------------------------------------------------------: |
| 6/20/2024 |   38.8s   | [f2f61b8](https://github.com/Pragma8123/1brc/tree/f2f61b85a2f0d65419613c7d0ceb966fec66e22f) |
