package main

import "fmt"

const (
	french = "French"
	spanish = "Espanol"

	engprefix= "Hello"
	frenchPrefix="bonj"
	spanishprefix="hola"
)



func strin (myarg string, lang string) string{
	
	if myarg == "" {
		myarg = "fuz"
	}
	return greet(lang) +", " + myarg 
}

func greet(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishprefix
	default:
		prefix = engprefix
	}
	return
}

func  main(){
	fmt.Println(strin("fuz",""))
}









