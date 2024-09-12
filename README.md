# About the project #
- The program is a  simple text completion/editing/auto-correction tool that reads a file, process it according to specific rules given and writes the output given to a new file.
- This program is written in Go language .
## package and imports #
* package main declares that this go file is part of the main package.
* package reload is a package containing the functions responsibles of modifying the wanted text.
### imports used
* fmt - provides functions for input and output formats.
* os - provide interface to operating system functions like reading and writing files.
* strconv - Used for converting string to numeric type.
* strings - Used for manipulating strings.

## Main function
* Reads input arguments from commandline using os.Args
* Reads the file specified by the first argument.
## Process function 
* Transforms the wanted text following some given rules by the help of the functions in modification directory.

## What happens during text processing
There are specific rules that we are given to modify the text. The following is to be achieved:
* Capitalizing : Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
* Converting to uppercase : Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")
* Converting to lowercase: Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting").
* For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")
 -- 
* Converting hexadecimal and binary number to decimal number: Every instance of (hex) or (bin) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number or a binary number).
* Modifying vowels : Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h
* Handling punctuations: Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one.
--
* The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "In my opinion: ' Messi is the best football player in history '" -> "In my opinion : 'Messi is the best football player in history'")
## After processing
* After transforming  according to the above rules,it concatenates them to one single string and then it writes the modified text to a new file as specified by the second argument.