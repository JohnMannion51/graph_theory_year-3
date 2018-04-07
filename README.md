# graph_theory_year-3

Hi my name is John Mannion and i'am a thrid year student in GMIT studying software development. As part of a module called 
graph theory our class was given a project to create a program using [GO](https://golang.org/)  .
This is a relatively new programming language that has been developed by google and was officially released as Go version 1.0 in March 2012. 

# Running the code

To run the code in this repository, the files must first be compiled. The Go [compiler](https://golang.org/doc/install) 
must first be installed on your machine. Once that is installed, the code can be compiled and run by following these steps. 
We assume you are using the command line or alternatively you can use [cmder](https://cmder.en.softonic.com/).

## Clone the repository from Github

1. git clone https://github.com/JohnMannion51/graph_theory_year-3.git

2. Change to the folder "graph_theory_year-3" 

3. Compile the file "go build rega1.go"

4. Run the executable "rega1.exe"

5. Choose an option 1(to start)or 2(to exit)

6. Enter the expression i.e. a*b

7. Enter the string to compare i.e. abb

8. The user will get an answer stating "true or false"


# The Project Goals

You must write a program in the Go programming language that can
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


# Helpful resources

Throughout the course of this semester our lecturer, Dr. Ian McLoughlin, provided some extremely 
helpful videos to aid us in our project. 

[Helpful video resources](https://www.youtube.com/watch?v=4bjqVsoy6bA)

[Shunting Yard Algorithm](https://brilliant.org/wiki/shunting-yard-algorithm/) 

[Deterministic Finite Automata Example](https://www.tutorialspoint.com/automata_theory/deterministic_finite_automaton.htm)

[Non-Deterministic Finite Automata NFA Example](https://www.tutorialspoint.com/automata_theory/non_deterministic_finite_automaton.htm)

[Thompson's Regular Expression Explaination](https://swtch.com/~rsc/regexp/regexp1.html)

[NFA using java](https://algs4.cs.princeton.edu/54regexp/NFA.java.html)

[Regular exp to nfa c++](http://cplusplus.happycodings.com/algorithms/code21.html)

[Python example](https://www.ics.uci.edu/~eppstein/PADS/Automata.py)
