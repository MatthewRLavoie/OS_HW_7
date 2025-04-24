package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

// RAID interface that all RAID types implement
type RAID interface {
	Write(blockNum int, data []byte) error
	Read(blockNum int) ([]byte, error)
}

// Runs write and read benchmarks on a RAID level
func BenchmarkRAID(name string, r RAID, totalMB int, blockSize int) {
	fmt.Printf("Benchmarking %s\n", name)
	blocks := (totalMB * 1024 * 1024) / blockSize
	data := make([]byte, blockSize)
	rand.Read(data) // fill with random data

	// Benchmark write
	start := time.Now()
	for i := 0; i < blocks; i++ {
		_ = r.Write(i, data)
	}
	writeDuration := time.Since(start)

	// Benchmark read
	start = time.Now()
	for i := 0; i < blocks; i++ {
		_, _ = r.Read(i)
	}
	readDuration := time.Since(start)

	// Output results
	fmt.Printf("Write time: %v (%.2f µs/block)\n", writeDuration, float64(writeDuration.Microseconds())/float64(blocks))
	fmt.Printf("Read time:  %v (%.2f µs/block)\n", readDuration, float64(readDuration.Microseconds())/float64(blocks))
	fmt.Println()
}

func main() {
	const blockSize = 4096 // 4 KB blocks
	const numDisks = 5
	const totalMB = 10 // total size to write/read

	// Open or create simulated disks
	var disks []*Disk
	for i := 0; i < numDisks; i++ {
		d, _ := OpenDisk(fmt.Sprintf("disk%d.dat", i))
		disks = append(disks, d)
	}

	// Run benchmarks for each RAID type
	BenchmarkRAID("RAID 0", &RAID0{Disks: disks, BlockSize: blockSize}, totalMB, blockSize)
	BenchmarkRAID("RAID 1", &RAID1{Disks: disks, BlockSize: blockSize}, totalMB, blockSize)
	BenchmarkRAID("RAID 4", &RAID4{DataDisks: disks[:4], ParityDisk: disks[4], BlockSize: blockSize}, totalMB, blockSize)
	BenchmarkRAID("RAID 5", &RAID5{Disks: disks, BlockSize: blockSize}, totalMB, blockSize)
}
