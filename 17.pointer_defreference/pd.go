package main

import ("fmt")

func change(str *string){
	*str = "changed!"
}


func change2(str string) string{
	str = "changed!"
	return str
}


func main() {
	// & - to get pointers,  * - Dereference
	x := 10
	y := &x  // y is equal  to poinnter x
	fmt.Println(x,y)
	*y = 20  // derefernce the pointer and change the value
	fmt.Println(x,y)

	str :=  "Hello"
	fmt.Println(str)
	change(&str)
	fmt.Println(str)

	sstr := "HHello"
	fmt.Println(sstr)
	change2(sstr)
	fmt.Println(sstr)
	fmt.Println("its not changing the value, ist copying the new value and returning: ",change2(sstr))

	tststr := "teststring"
	var v *string = &tststr
	fmt.Println(v,&v,*v)


}