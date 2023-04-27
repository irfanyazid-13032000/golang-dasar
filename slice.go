package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	months[5] = "haduh"
	fmt.Println(slice)

	slice[0] = "mei diubah gak nih"
	fmt.Println(months)

	var winter = months[10:]
	fmt.Println(cap(winter))

	// ini seakan akan membuar array baru (jika sudah tidak sanggup menampung)
	var libur = append(winter,"yudi")
	fmt.Println(libur)

	// ini mengganti array asli saja
	winter[1] = "bukan desember"
	fmt.Println(winter)
	fmt.Println(libur)
	fmt.Println(months)

	// ketika merubah slice, maka merubah array asli juga
	// ketika merubah append maka merubah append saja, tidak sampai array aslinya dirubah (jika membuat array baru)


	newSlice := make([]string,2,5)

	newSlice[0] = "yudi"
	newSlice[1] = "tabudi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string,1,cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)


	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	
}