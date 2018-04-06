# graph_theory_year-3

Hi my name is John Mannion and i'am a thrid year student in GMIT studying software development. As part of a module called 
graph theory our class was given a project to create a program using [GO](https://golang.org/)  .
This is a relatively new programming language that has been developed by google and was officially released as Go version 1.0 in March 2012. 

# Running the code

To run the code in this repository, the files must first be compiled. The Go [compiler](https://golang.org/doc/install) 
must first be installed on your machine. Once that is installed, the code can be compiled and run by following these steps. 
We assume you are using the command line or alternatively you can use [cmder](https://cmder.en.softonic.com/).

## Clone the repository from Github
1. git clone https://github.com/JohnMannion51/graph_theory_year-3.git.
Change into the folder.
2. cd rega1.go .
Compile the first file with the following command.
3. go build rega1.go .
(this will produce an executable file)
Run the executable produced.
4. rega1.exe .

# The Project Outlay

You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.

A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and "*” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1* means any number of 1’s. These special characters must be used in
your submission.

You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. 

## You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.

