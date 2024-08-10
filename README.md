### README.md

# chuckle-cli

`chuckle-cli` is a simple command-line application that fetches and displays a random joke. This README provides instructions for installing and using `chuckle-cli` on Windows and macOS/Unix systems.

## Installation

### Prerequisites

- **Go**: Ensure you have Go installed on your system. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/).
- **Git**: Ensure you have Git installed to clone the repository. You can download and install it from [https://git-scm.com/downloads](https://git-scm.com/downloads).

### Clone the Repository

Open your terminal or command prompt and clone the repository using Git:

```sh
git clone https://github.com/yourusername/chuckle-cli.git
cd chuckle-cli
```

### macOS/Unix Installation

1. Open your terminal and navigate to the cloned repository directory if you haven't already done so:

    ```sh
    cd chuckle-cli
    ```

2. Run the setup script:

    ```sh
    ./setup.sh
    ```

3. Restart your terminal to apply the changes.

### Windows Installation

1. Open PowerShell with administrator rights and navigate to the cloned repository directory if you haven't already done so:

    ```sh
    cd chuckle-cli
    ```

2. Run the setup script:

    ```sh
    .\setup.ps1
    ```

3. Restart your terminal to apply the changes.

## Usage

Once you have installed `chuckle-cli`, you can use the `chuckle` command to fetch and display a random joke:

```sh
chuckle
```

## Troubleshooting

- If the `chuckle` command is not recognized after installation, ensure that the PATH has been updated correctly by checking your shell configuration file (e.g., `.zshrc`, `.bash_profile`) and verifying that the `chuckle-cli` directory is included in your PATH.
- Ensure you have restarted your terminal or sourced the shell configuration file after running the setup script.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

```
The MIT License (MIT)

Copyright (c) 2024-present Sebastian Urban

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```