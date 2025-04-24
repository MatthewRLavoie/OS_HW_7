package main

import (
	"fmt"
	"os"
)

// Disk wraps an OS file to simulate a physical disk
type Disk struct {
	File *os.File
	Path string
}

// Opens or creates a disk file
func OpenDisk(path string) (*Disk, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &Disk{File: f, Path: path}, nil
}

// Writes a block to disk at a specific offset
func (d *Disk) WriteBlock(blockNum int, blockSize int, data []byte) error {
	if len(data) != blockSize {
		return fmt.Errorf("invalid block size")
	}
	offset := int64(blockNum * blockSize)
	_, err := d.File.WriteAt(data, offset)
	if err != nil {
		return err
	}
	return d.File.Sync() // flush to simulate disk behavior
}

// Reads a block from disk at a specific offset
func (d *Disk) ReadBlock(blockNum int, blockSize int) ([]byte, error) {
	offset := int64(blockNum * blockSize)
	buf := make([]byte, blockSize)
	_, err := d.File.ReadAt(buf, offset)
	return buf, err
}
