package main

//import"bufio"
import"fmt"
//import"os"


func main() {
	// EX 1
	/*
	var nr int
	e_prim := true
	fmt.Print("Numarul de verificat este: ")
	fmt.Scan(&nr)
	if nr < 2 || (nr != 2 &&nr % 2 == 0){
		e_prim = false
	} else{
		for i :=3; i * i <= nr; i+=2{
			if nr % i == 0{
				e_prim = false
				break
			}
		}
	}
	if e_prim{
		fmt.Println("Nr e prim")
	} else {
		fmt.Println("Nr nu e prim")
	} 
	*/

	// EX 2
	/*
	intArr := [5]int{1, 4, 3, 5, 2}
	for i:=0; i < len(intArr) - 1; i++{
		for j := 0; j < len(intArr) - i - 1; j++{
			if intArr[j] > intArr[j + 1]{
				intArr[j], intArr[j + 1] = intArr[j + 1], intArr[j]
			}
		}
	}
	fmt.Println(intArr) 
	*/

	// EX 3
	/*
	intArr := [7]int{1, 4, 3, 5, 2, 6, 8}
	var intSlice []int
	for i := range len(intArr){
		if intArr[i] % 2 == 0{
			intSlice = append(intSlice, intArr[i])
		}
	}
	fmt.Println(intSlice)
	*/

	// EX 4
	/*
	floatSlice := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
	var indice int
	fmt.Print("Indicele elementului de eliminat este: ")
	fmt.Scan(&indice)
	floatSlice = append(floatSlice[:indice], floatSlice[indice + 1:]...)
	fmt.Println(floatSlice)
	*/

	// EX 5
	/*
	stringSlice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	for i := range len(stringSlice) / 2{
		stringSlice[i], stringSlice[len(stringSlice) - i - 1] = stringSlice[len(stringSlice) - i - 1], stringSlice[i]
	}
	fmt.Println(stringSlice)
	*/

	// EX 6
	/*
		inaltimiMap := map[string]int{
			"Andrei": 190,
			"Maria": 165,
			"Ion": 182,
			"Elena": 172,
			"George": 178,
			"Ana": 165,
		}

		var stringNume []string

		for nume := range inaltimiMap{
			stringNume = append(stringNume, nume)
		}

		for i:=0; i < len(stringNume) - 1; i++{
			for j := 0; j < len(stringNume) - i - 1; j++{
				if inaltimiMap[stringNume[j]] > inaltimiMap[stringNume[j + 1]]{
					stringNume[j], stringNume[j + 1] = stringNume[j + 1], stringNume[j]
				}
			}
		}

		for i := range stringNume{
			fmt.Println(stringNume[i])
		}
	*/

	// EX 7
	/*
	file, err := os.Open("Ex7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	stringFrecv := make(map[string]int)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		cuv := scanner.Text()
		stringFrecv[cuv]++
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire: ", err)
	}

	for cuv, frecv := range stringFrecv {
		fmt.Printf("%s: %d\n", cuv, frecv)
	}
	*/

	// EX 8
	/*
	numDepartament := map[string]int{
		"Ana Popescu": 1,
		"Mihai Ionescu": 2,
		"Bogdan Radu": 1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile": 2,
		"Ioana Enache": 1,
	}

	orgDepartament := make(map[int][]string)

	for nume, departament := range numDepartament {
		orgDepartament[departament] = append(orgDepartament[departament], nume)
	}

	for departament := range orgDepartament {
		fmt.Printf("%v\n", orgDepartament[departament])
	}
	*/


	// EX 9
	intArr := [6]int{2, 7, 11, 8, 4, 9}
	var s int
	fmt.Print("Suma cautata: ")
	fmt.Scan(&s)

	dejaViz := make(map[int]int)

	for i, val := range intArr {
		complement := s - val
		pozComplement, gasit := dejaViz[complement]

		if gasit {
			fmt.Println(pozComplement, i)
		}

		dejaViz[val] = i
	}


}
