package main

import (
	"crypto/sha256"
	"fmt"
	"loki/crypt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	banniere :=
		`
	
█    ████▄ █  █▀ ▄█     ▄███▄      ▄   ▄█▄    █▄▄▄▄ ▀▄    ▄ █ ▄▄     ▄▄▄▄▀ ████▄ █▄▄▄▄ 
█    █   █ █▄█   ██     █▀   ▀      █  █▀ ▀▄  █  ▄▀   █  █  █   █ ▀▀▀ █    █   █ █  ▄▀ 
█    █   █ █▀▄   ██     ██▄▄    ██   █ █   ▀  █▀▀▌     ▀█   █▀▀▀      █    █   █ █▀▀▌  
███▄ ▀████ █  █  ▐█     █▄   ▄▀ █ █  █ █▄  ▄▀ █  █     █    █        █     ▀████ █  █  
    ▀        █    ▐     ▀███▀   █  █ █ ▀███▀    █    ▄▀      █      ▀              █   
            ▀                   █   ██         ▀              ▀                   ▀    
                                                                                       
	A local file encryptor in AES. Designed by Yatoub42 under MIT Licence


	`

	usage :=
		`
	LOKI(1)                     User Manuals                    LOKI(1)


NAME
    Loki - AES Local file encryptor

SYNOPSIS
    loki OPTIONS path/to/file

DESCRIPTION
	 Loki encrypts a given file with a password using the AES algorithm. 
	 The password is not kept and files encrypted with Loki 
	 must be decrypted by him with the same password used for encryption.

OPTIONS
	 -h --help
	 	Display this help message

	 -e --encrypt
	 	Encrypt file given in argument and replace it by .loki extension

	 -d --decrypt
	 	Decrypt a loki file

FILES
     /path/to/file
        The given file must be with full path
	`
	if len(os.Args) == 1 {
		fmt.Println(banniere)
		fmt.Println("Help with -h argument")
		os.Exit(0)
	}
	fmt.Println(banniere)
	if os.Args[1] == "-h" || os.Args[1] == "--help" || len(os.Args) < 3 {
		fmt.Println(usage)
	} else if len(os.Args) == 3 {
		file := os.Args[2]
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err == nil {
			password := string(bytePassword)
			passwordHash := sha256.Sum256([]byte(password))
			crypt.InitializeBlock([]byte(passwordHash[:]))
			if strings.HasSuffix(file, ".loki") && os.Args[1] == "-d" || os.Args[1] == "--decrypt" {
				err := crypt.Decrypter(file)
				if err != nil {
					panic(err.Error())
				}
			} else if os.Args[1] == "-e" || os.Args[1] == "--encrypt" {
				err := crypt.Encrypter(file)
				if err != nil {
					panic(err.Error())
				}
			}
		}
	} else {
		fmt.Println("Error in given argument, maybe use -h")
	}
}
