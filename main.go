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

	banner :=
		`
	
█    ████▄ █  █▀ ▄█     ▄███▄      ▄   ▄█▄    █▄▄▄▄ ▀▄    ▄ █ ▄▄     ▄▄▄▄▀ ████▄ █▄▄▄▄ 
█    █   █ █▄█   ██     █▀   ▀      █  █▀ ▀▄  █  ▄▀   █  █  █   █ ▀▀▀ █    █   █ █  ▄▀ 
█    █   █ █▀▄   ██     ██▄▄    ██   █ █   ▀  █▀▀▌     ▀█   █▀▀▀      █    █   █ █▀▀▌  
███▄ ▀████ █  █  ▐█     █▄   ▄▀ █ █  █ █▄  ▄▀ █  █     █    █        █     ▀████ █  █  
    ▀        █    ▐     ▀███▀   █  █ █ ▀███▀    █    ▄▀      █      ▀              █   
            ▀                   █   ██         ▀              ▀                   ▀    
                                                                                       
	A local file encryptor in AES-256. Designed by Yatoub42 under MIT License


	`

	usage :=
		`
	LOKI(1)                     User Manuals                    LOKI(1)


NAME
    Loki - AES Local file encryptor

SYNOPSIS
    loki OPTIONS path/to/file

DESCRIPTION
	 Loki encrypts a given file with a password using the AES-256 algorithm. 
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

	// Run without args
	if len(os.Args) == 1 {
		fmt.Println(banner)
		fmt.Println("Help with -h argument")
		os.Exit(0)
	}
	fmt.Println(banner)
	//check args
	switch {
		case os.Args[1] == "-h" || os.Args[1] == "--help" || len(os.Args) < 3:
			fmt.Println(usage)
		case len(os.Args) == 3:
			file := os.Args[2]
			// Read password
			fmt.Print("Enter Password: ")
			bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err == nil {
				// Hash given password in 32 byte slice to have AES-256 encryption
				passwordHash := sha256.Sum256([]byte(string(bytePassword)))
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
		default:
			fmt.Println("Error in given argument, maybe use -h")
	}
}
