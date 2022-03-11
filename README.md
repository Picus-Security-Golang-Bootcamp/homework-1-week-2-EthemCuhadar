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

The program uses command-line arguments which are one of the best ways to create parameters for programs during execution. 
