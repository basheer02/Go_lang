package lists

import "fmt"


/*type product struct{
	title string
	id int
	price float64
}


func main() {
	
	hobbies := [3]string{"football", "music", "movies"}

	fmt.Println(" Hobbies :", hobbies)
	fmt.Println(" first hobby :", hobbies[0])
	fmt.Println(" second and third hobby :", hobbies[1:3])

	featuredHobbies := hobbies[:2]

	fmt.Println(" featured hobbies :",featuredHobbies)

	featuredHobbies = featuredHobbies[1:3]

	fmt.Println(" featured hobbies updated :",featuredHobbies)

	courseGoals := []string{"Understand basics", "Be a pro"}
	fmt.Println(" Course goals :", courseGoals)

	courseGoals[1] = "Be good at it"
	courseGoals = append(courseGoals, "Be a pro")
	fmt.Println(" Updated course goals :", courseGoals)

	productList := []product{
		{
			title: "Book",
			id: 1234,
			price: 9.99,
		},
		{
			title: "Pen",
			id: 1235,
			price: 4.99,
		},
	}

	fmt.Println(" Product details :", productList)

	newProduct := []product{ {"Pencil", 1236, 2.99}, {"Paper", 1231, 0.99}}

	productList = append(productList, newProduct...)

	fmt.Println(" Updated product details :", productList)

}*/






func main() {
	prices := []float64{9.99, 10.99}

	prices = append(prices, 5.99)

	fmt.Println(prices)

	//prices = prices[1:] // it removes first element
	newPrices := []float64{4.99, 6.99, 15.99}

	prices = append(prices, newPrices...)
	
	fmt.Println(prices)
}


/*func main() {
	var productNames [4]string = [4]string{"Book"}

	prices := [4]float64{10.9, 5.9, 30.9, 50.9}

	productNames[1] = "Pen"
	productNames[2] = "Carpet"

	fmt.Println(productNames)
	fmt.Println(prices)

	featuredPrices := prices[:1]
	//featuredPrices[0] = 100.0 it changes the original array
	fmt.Println(featuredPrices)

	fmt.Println(" len of featured :",len(featuredPrices ))
	fmt.Println(" capacity of featured :", cap(featuredPrices)) 
	//from the left to the end of the original array is the capacity of sub array
	featuredPrices = featuredPrices[:3]
	fmt.Println(featuredPrices)

}*/