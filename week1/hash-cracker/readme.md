# Hash Cracker

CLI tool takes a hash value input and a file compares the hash value of the input to the values in the file and provides output if the hash value matches the hash of any of the words in the file.

The script was further enhanced for user to dictate wether the hash value being checked is md5 or SHA256.


# how to run it

go run main.go -algo sha256 *hashvalue* wordlist.txt

# Example Ouput 

go run main.go -algo sha256 *hashvalue* wordlist.txt
Algorithm: sha256
Cracked :->  root