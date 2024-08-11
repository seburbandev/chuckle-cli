# chuckle-cli

chuckle-cli is a command-line application that does only one thing: prints out a random joke. It utilises an API from [15DKatz/offcial_joke_api](https://github.com/15Dkatz/official_joke_api). You can install it on Windows, Mac or Linux. Once installed, everytime you're in terminal, you can type `chuckle` and it will print out a joke, so you can have a chuckle.

> [!NOTE]
> This app was created as a fun project to learn more about the Go programming language. I don't expect anyone to make real use of it nor do I plan to add many other features in the future.


## Installation

### Dependencies

- **Go**: Make sure you have Go installed on your system. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/).
- **Git**: Make sure you have Git installed to clone the repository. You can download and install it from [https://git-scm.com/downloads](https://git-scm.com/downloads).

### Clone the repository

Open your terminal or command prompt and clone the repository using Git:

```sh
git clone https://github.com/seburbandev/chuckle-cli.git
```

### Mac/Linux installation

1. Open your terminal and navigate to the cloned repository directory if you haven't already done so:

    ```sh
    cd chuckle-cli
    ```

2. Run the setup script:

    ```sh
    ./unix-setup.sh
    ```

3. Restart your terminal to apply changes.

### Windows installation

1. Open command prompt or powershell with administrator rights and navigate to the cloned repository directory if you haven't already done so:

    ```sh
    cd chuckle-cli
    ```

2. Run the setup script:

    ```sh
    .\windows-setup.ps1
    ```

3. Restart your terminal to apply changes.

## Usage

Once you have installed `chuckle-cli`, you can use the `chuckle` command to fetch and display a random joke:

```sh
chuckle
```

You should see joke printed in a format like below:

```
Have you ever heard of a music group called Cellophane?
...

...

...
They mostly wrap.
```

## Troubleshooting

- If the `chuckle` command is not recognized after installation, ensure that the PATH has been updated correctly by checking your shell configuration file (e.g., `.zshrc`, `.bash_profile`) and verifying that the `chuckle-cli` directory is included in your PATH.
- Ensure you have restarted your terminal or sourced the shell configuration file after running the setup script.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request for any improvements or bug fixes.

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