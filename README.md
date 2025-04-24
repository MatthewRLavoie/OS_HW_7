# OS_HW_7

Running Instructions:
1. Place all files into a folder named RAID
2. In the command window, type: go mod init RAID
3. In the command window, type: go run main.go disk.go raid0.go raid1.go raid4.go raid5.go
4. Wait for completion

Analysis:
The measured performance trends generally align with expectations based on textbook RAID analysis. RAID 0 shows the fastest write time due to parallel striping with no redundancy overhead, while RAID 1 is the slowest because each block is written to all disks (mirroring). RAID 4 and RAID 5 both exhibit better performance than RAID 1, with similar write times, though slightly slower than RAID 0 due to parity calculations. Notably, RAID 5 performs marginally better than RAID 4 in reads, which matches textbook predictions since RAID 5 distributes parity across disks, reducing read bottlenecks on a dedicated parity disk. Minor discrepancies in absolute times can be attributed to OS-level file caching, file I/O overhead, or lack of simulated concurrent disk access.

AI Usage:
ChatGPT was used for help with making the benchmark test, and with how to write certain parts of the code.
