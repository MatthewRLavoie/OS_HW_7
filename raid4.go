package main

// RAID4 has dedicated parity disk, rest store data
type RAID4 struct {
	DataDisks  []*Disk
	ParityDisk *Disk
	BlockSize  int
}

// XORs blocks for parity
func xorBlocks(blocks [][]byte) []byte {
	result := make([]byte, len(blocks[0]))
	for _, block := range blocks {
		for i := range result {
			result[i] ^= block[i]
		}
	}
	return result
}

// Write data and update parity on the stripe
func (r *RAID4) Write(blockNum int, data []byte) error {
	diskIndex := blockNum % len(r.DataDisks)
	stripe := blockNum / len(r.DataDisks)

	// Collect data blocks to recalculate parity
	dataBlocks := make([][]byte, len(r.DataDisks))
	for i := range r.DataDisks {
		if i == diskIndex {
			dataBlocks[i] = data
		} else {
			block, _ := r.DataDisks[i].ReadBlock(stripe, r.BlockSize)
			dataBlocks[i] = block
		}
	}
	parity := xorBlocks(dataBlocks)

	// Write new data and updated parity
	if err := r.DataDisks[diskIndex].WriteBlock(stripe, r.BlockSize, data); err != nil {
		return err
	}
	return r.ParityDisk.WriteBlock(stripe, r.BlockSize, parity)
}

// Read block from correct data disk
func (r *RAID4) Read(blockNum int) ([]byte, error) {
	diskIndex := blockNum % len(r.DataDisks)
	stripe := blockNum / len(r.DataDisks)
	return r.DataDisks[diskIndex].ReadBlock(stripe, r.BlockSize)
}
