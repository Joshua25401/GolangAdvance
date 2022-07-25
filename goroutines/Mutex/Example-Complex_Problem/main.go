package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Pemasukan struct {
	SumberPemasukan string
	Jumlah          int
}

func main() {
	// Declare variabel Saldo Bank
	var saldoBank int
	var mtxSaldoBank sync.Mutex

	// Print nilai saldo bank awal
	fmt.Println(fmt.Sprintf("Nilai saldo bank awal Rp.%d,00", saldoBank))

	// Declare Pemasukan
	listPemasukan := []*Pemasukan{
		{
			SumberPemasukan: "Bekerja",
			Jumlah:          7_000,
		},
		{
			SumberPemasukan: "Hadiah",
			Jumlah:          3_000,
		},
		{
			SumberPemasukan: "Amplop Korupsi",
			Jumlah:          12_000,
		},
		{
			SumberPemasukan: "Investasi Kosan",
			Jumlah:          20_000,
		},
	}

	wg.Add(len(listPemasukan))

	// Simulate getting income
	for i, pemasukan := range listPemasukan {
		go func(i int, pemasukan *Pemasukan, mtx *sync.Mutex) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {

				// This mutex is to guarantee no race conditions happen!
				mtx.Lock()
				temp := saldoBank
				temp += pemasukan.Jumlah
				saldoBank = temp
				mtx.Unlock()

				fmt.Println(
					fmt.Sprintf(
						"Minggu: %v mendapat pemasukan sebesar Rp.%d,00 dari %s",
						week,
						pemasukan.Jumlah,
						pemasukan.SumberPemasukan,
					),
				)
			}
		}(i, pemasukan, &mtxSaldoBank)
	}
	// Wait for goroutines done executing job
	wg.Wait()

	fmt.Printf("Saldo akhir anda Rp.%d,00\n", saldoBank)
}
