package main

import (
	"fmt"
	"os"
	"strings"
)

// Movie struct to create movies with relative fields.
type Movie struct{
	Name		string
	Director	string
	Year		int
	Rating		float32
	Genre		[]string
}


// movieNameCreatorFromArgs takes Command Line Args
// and concatenate them to create movie name.
func movieNameCreatorFromArgs(l []string)string{
	sep := " "
	var s string
	
	for i:=2; i<len(l); i++ {
		if i == len(l) - 1{
			s += l[i]
		}else{
		s += l[i] + sep
		}
	}
	return s
}

// search checks typed movie name in movie list.
func search(s string, l []Movie)bool{
	for _, v := range l{
		if strings.ToLower(v.Name) == strings.ToLower(s){
			fmt.Println("Movie Name:      ", v.Name)
			fmt.Println("Movie Director : ", v.Director)
			fmt.Println("Movie Year :     ", v.Year)
			fmt.Println("Movie Rating :   ", v.Rating)
			fmt.Println("Movie Genre :    ", v.Genre)
			return true
		}
	}
	fmt.Println("The movies does not exist in the list.")
	return false
}

// list simply prints movies that are in list.
func list(l []Movie){
	
	fmt.Println("--------------- Movies ----------------")
	for _, v := range l{
		fmt.Println("Movie Name:      ", v.Name)
		fmt.Println("Movie Director : ", v.Director)
		fmt.Println("Movie Year :     ", v.Year)
		fmt.Println("Movie Rating :   ", v.Rating)
		fmt.Println("Movie Genre :    ", v.Genre)
		fmt.Println("---------------------------------------")
	}
}

// commandListOrSearch simply decides what command will be applied.
// c: command(list or search)
// s: movie name from args
// l: movie list that is created
func commandListOrSearch(c string, s string, l []Movie){
	if strings.ToLower(c) == "list"{
		list(l)
	}else if strings.ToLower(c) == "search"{
		search(s, l)
	}else{
		fmt.Println("Unknown command, please check the command again!")
		fmt.Println("go run main.go search <movie_name> --> to search a specific movie in the list.")
		fmt.Println("go run main.go list  --> to print all movies in the list.")
	}
}

func main(){
	
	// Define movies
	movie001 := Movie{
		Name: "Lord of the Rings: Return of the King",
		Director: "Peter Jackson",
		Year: 2003,
		Rating: 9.1,
		Genre: []string{"Adventure", "Fantastic", "Drama", "Action"},
	}
	
	movie002 := Movie{
		Name: "The Godfather",
		Director: "Francis Ford Coppola",
		Year: 1972,
		Rating: 9.2,
		Genre: []string{"Crime", "Drama"},
	}
	
	movie003 := Movie{
		Name: "The Dark Knight",
		Director: "Christopher Nolan",
		Year: 2008,
		Rating: 9.1,
		Genre: []string{"Action", "Crime", "Drama"},
	}
	
	movie004 := Movie{
		Name: "Fight Club",
		Director: "David Fincher",
		Year: 1999,
		Rating: 8.8,
		Genre: []string{"Crime", "Drama", "Action"},
	}
	
	// Define movie list, command and movie name.
	movieList := []Movie{movie001, movie002, movie003, movie004}
	movieCommand := os.Args[1]
	movieName := movieNameCreatorFromArgs(os.Args)
	
	// Apply the function commandListOrSearch.
	commandListOrSearch(movieCommand, movieName, movieList)

}
