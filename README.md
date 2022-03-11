# Homework | Week 2

![Golang Image](golang.png)

-------------------------------------------------------------------

This is a very simple application about movies in order to practice and learn CLI (command line interface) in Go programming language. 



Firstly, the program has 2 commands which are **list** and **search moviename**. The first command **(list)** simply prints out the movies that are on the movie list. Moreover, the **search moviename** command checks the movie name in the movie list. If the movie name which is typed in the command line exists in the movie list, the program prints out some pieces of information (movie name, year, director, etc.) about the movie. Otherwise, the program tells that the movie does not exist.



* Usage of **list** command can be seen below.

```[console]
$ go run main.go list
```

```[echo]
--------------- Movies ----------------
Movie Name:       Lord of the Rings: Return of the King
Movie Director :  Peter Jackson
Movie Year :      2003
Movie Rating :    9.1
Movie Genre :     [Adventure Fantastic Drama Action]
---------------------------------------
Movie Name:       The Godfather
Movie Director :  Francis Ford Coppola
Movie Year :      1972
Movie Rating :    9.2
Movie Genre :     [Crime Drama]
---------------------------------------
Movie Name:       The Dark Knight
Movie Director :  Christopher Nolan
Movie Year :      2008
Movie Rating :    9.1
Movie Genre :     [Action Crime Drama]
---------------------------------------
Movie Name:       Fight Club
Movie Director :  David Fincher
Movie Year :      1999
Movie Rating :    8.8
Movie Genre :     [Crime Drama Action]
---------------------------------------
```

* Furthermore, the usage of **search moviename** can seen below.

```[console]
$ go run main.go search Lord of the Rings: Return of the King
```

```[console]
Movie Name:       Lord of the Rings: Return of the King
Movie Director :  Peter Jackson
Movie Year :      2003
Movie Rating :    9.1
Movie Genre :     [Adventure Fantastic Drama Action]
```

-------------------------------------------

## Diving into code

* The program uses command-line arguments which are one of the best ways to create parameters for programs during execution.
* Structs are used in order to create movies. Fields of the movies are **movie name (string)**, **movie director (string)**, **movie year (int)**, **movie rating (float64)** and **movie genre ([]string)**.
* Moreover, there are four function in program. Two of them are **list** and **search** which are explianed above in term of functionality.
* On the other hand, other two functions are **movieNameCreatorFromArgs** and **commandListOrSearch**. The first one takes command-line arguments from console and concatenate them to create single string value (movie name). In Go programming language, command-line arguments, aka. **os.Args**, are in list form. What we write in command-line as arguments are slice elements. 0th element of that list is nothing but name of the program. Therefore, we use the elements starting from 1st element which is 1st element in command-line arguments as well. Furthermore, first element of the arguments is **command** which can be **list** or **search**. The other elements will become movie name. The other function, **commandListOrSearch**, decides what command will be applied by taking first element of arguments.
* Finally, movies are stored in slice. However, there can be different types in order to store movies like array, dictionary or JSON file as well.

---------------------------------------------

## Conclusion

* The program is very simple which has 2 functionalities. There would be different features like add movie or delete movie. 
* Error cases are considered in the program. Any commands other than **list** and **search** will cause an error message.
* The command-line arguments are not case sensetive. For instance, **go run main.go search Fight Club** and **go run main.go search FiGhT ClUb** are same for the program.

---------------------------------------------
