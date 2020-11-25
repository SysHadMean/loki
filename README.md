# Loki

    █    ████▄ █  █▀ ▄█     ▄███▄      ▄   ▄█▄    █▄▄▄▄ ▀▄    ▄ █ ▄▄     ▄▄▄▄▀ ████▄ █▄▄▄▄ 
    █    █   █ █▄█   ██     █▀   ▀      █  █▀ ▀▄  █  ▄▀   █  █  █   █ ▀▀▀ █    █   █ █  ▄▀ 
    █    █   █ █▀▄   ██     ██▄▄    ██   █ █   ▀  █▀▀▌     ▀█   █▀▀▀      █    █   █ █▀▀▌  
    ███▄ ▀████ █  █  ▐█     █▄   ▄▀ █ █  █ █▄  ▄▀ █  █     █    █        █     ▀████ █  █  
        ▀        █    ▐     ▀███▀   █  █ █ ▀███▀    █    ▄▀      █      ▀              █   
                ▀                   █   ██         ▀              ▀                   ▀    
------------
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Yatoub42/loki?style=for-the-badge)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-coffee.svg)](https://forthebadge.com)  
![GitHub top language](https://img.shields.io/github/languages/top/Yatoub42/loki?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/Yatoub42/loki?style=for-the-badge)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Yatoub42/loki?style=for-the-badge)

## Description

A local file encryptor in AES. Designed by Yatoub42

Loki encrypts a given file with a password using the AES algorithm. The password is not kept and files encrypted with Loki must be decrypted by him with the same password used for encryption.

## Getting start

To get a local copy up and running follow these simple steps.

### Installation

1. Clone the repo

```bash
    git clone ...
```

1. Build for your OS

```bash
   go build main.go
```

2. (Optionnal) Add the executable file to your PATH

### Usage

```bash
    loki OPTIONS path/to/file
```

#### OPTIONS

    -h --help
	Display this help message

	-e --encrypt
	Encrypt file given in argument and replace it by .loki extension

	-d --decrypt
	Decrypt a loki file

#### FILE

    /path/to/file
    The given file must be with full path

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/Yatoub42/loki/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
