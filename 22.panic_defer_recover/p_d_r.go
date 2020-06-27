package main

import "fmt"

func main(){
	fmt.Println("===============================================================")
	fmt.Println("Example : panic defer recover")	
	DB_PDR()

	fmt.Println("===============================================================")
	fmt.Println("Example : panic defer")
	
	DB_PD()
	fmt.Println("===============================================================")

}

func DB_PDR() {
	defer fmt.Println("Please check the DB connection")
	defer func () {
		if r :=recover(); r !=nil{
			fmt.Println("Trying to fix")
		}
	}()

	fmt.Println("Creating DB connection")
	panic("Issue when connecting DB")
	
}


func DB_PD() {
	defer fmt.Println("Please check the DB connection")	

	fmt.Println("Creating DB connection")
	panic("Issue when connecting DB")
	
}