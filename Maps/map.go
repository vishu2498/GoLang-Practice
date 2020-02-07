package main

import "fmt"

func main() {
	//Maps are basically key-value stores
	//A 'nil' or a map with no key-value can be created but it can’t be used to store key/value pairs.
	//1st data-type of map is the data-type for key
	//2nd data-type of map is the data-type for value
	//Putting ',' after every map value is compulsory
	map1:=map[string]int {
		"science": 55,
		"maths": 45,
	}
	fmt.Println(map1)

	//Slice can't be a data-type of key
	//Array can be a data-type of key
	map2:=map[[3]string]int{}
	fmt.Println("map2:",map2)

	//But slice can be a data-type of a value
	mapn:=map[string][]int{}
	fmt.Println(mapn)

	//Using make() function for creating maps
	//This method is generally used when the values of map are unknown at the time of declaring map or at the runtime
	//Further arguments can also be defined in the make() function after map for defining length & capacity but they are of no use
	map3:=make(map[string]int) //values of maps can't be defined with the make() function
	map3=map[string]int {
		"1st": 30,
		"2nd": 95,
	}
	fmt.Println(map3)

	//Getting value of key from a map
	map4:=make(map[string]int)
	map4=map[string]int {
		"1st": 80,
		"2nd": 15,
	}
	fmt.Println(map4["1st"]) //we have to mention the name of the key

	//Adding key-value to the map
	map5:=make(map[string]int)
	map5=map[string]int {
		"1st": 98,
		"2nd": 16,
	}
	map5["3rd"]=23
	fmt.Println(map5)

	//The return order of map is not guaranteed. It means that maps will be unsorted (even when they are printed).
	//The reason for this is that maps are implemented using hash table.
	//But they will be only sorted on the basis of the keys. Maps will always be sorted on the alphabetical and numerical order or keys.
	//Unlike maps, arrays and slices will always return in the order they are stored.
	map6:=make(map[string]int)
	map6=map[string]int {
		"zharka": 98,
		"pune": 16,
		"rose": 65,
		"beach": 27,
	}
	fmt.Println(map6)

	/*The map’s hash table contains a collection of buckets. When we’re storing, removing,
	or looking up a key/value pair, everything starts with selecting a bucket. This is
	performed by passing the key—specified in our map operation—to the map’s hash
	function. The purpose of the hash function is to generate an index that evenly distributes
	key/value pairs across all available buckets. */

	//Deleting key-value from map. delete() function will be used.
	map7:=make(map[string]int)
	map7=map[string]int {
		"1st": 95,
		"2nd": 24,
		"3rd": 17,
	}
	delete(map7,"2nd") //1st argument will be name of the map, 2nd argument will be name of key.
	fmt.Println(map7)

	//If we print any key that is not present in the map, then it will return '0'.
	map8:=make(map[string]int)
	map8=map[string]int {
		"1st": 98,
		"2nd": 16,
	}
	fmt.Println(map8["3rd"])
	//GoLang has no particular reason that why result like this arrives.
	//To overcome this situation, Comma-Ok syntax is used.
	//In Comma-Ok syntax, the 1st variable will contain the actual data whereas the 2nd variable (ok) will store the boolean value telling that actual data is present or not.
	//2nd variable (ok) can be declared,defined and used multiple times.
	data,ok:=map8["3rd"]
	fmt.Println(data,ok)
	data1,ok:=map8["2nd"]
	fmt.Println(data1,ok)
	//Blank Identifier (_) can also be used for Comma-Ok syntax.
	//Black Identifier (_) is the special keyword & used when we receive a value from something but don't want to allot a variable to store it.
	//Rather the value is stored temporarily by blank identifier (_) and we don't need to mention it afterwards.
	//While using blank identifier (_), don't use short-declaration syntax (:=).
	_,ok=map8["4th"]
	fmt.Println(ok)

	//Getting the length of the map
	fmt.Println(len(map8))

	//Maps are also reference types for data just like slices.
	//So, if a variable is equal to another variable that is holding the map and we change/delete/add any key-value of/to map from this variable, the actual data (map) will get changed.
	map9:=make(map[string]int)
	map9=map[string]int {
		"1st": 78,
		"2nd": 47,
	}
	map10:=map9
	delete(map10,"1st")
	fmt.Println(map9) //deleting value from map10 also deleted value from map9 since both are pointing to the same map
	fmt.Println(map10)
}
