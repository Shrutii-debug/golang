package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	//slice
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

/*
	//type 1 syntax for FOR
	for d := 0; d < len(days); d++{
		fmt.Println(days[d])
	}
*/
// syntax 2
	for i := range days{
		fmt.Println(days[i])
	}

//syntax 3
for index, day := range days{
	fmt.Printf("index is %v and value is %v\n", index, day)

	//incase you want to ignore anything as studied previously, use an underscore
}

//break and continue
rogueValue := 1

for rogueValue < 10 {

	if rogueValue == 2{
		goto lco
	}

	/*
	if rogueValue == 5 {
		break
	}
	*/

	if rogueValue == 5 {
		rogueValue++
		continue
	}



	fmt.Println("value is :" , rogueValue)
	rogueValue++
	//in golang we only have name++ and not ++name
}


//goto

lco:
	fmt.Println("jumping at learncodeonline.com")

}