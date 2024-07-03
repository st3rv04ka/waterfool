
# Waterfool

Waterfool is a Go program that converts images into radio signals for drawing on a waterfall display. This tool is ideal for amateur radio enthusiasts and anyone interested in signal processing and visual communication through radio waves.

## Features

- Convert any image format to a waterfall display signal
- Easy-to-use command line interface

## Installation

First, ensure you have git installed on your machine. Then, you can download and build Waterfool by running:

```sh
git clone https://github.com/st3rv04ka/waterfool
cd waterfool
go get .
go build cmd/waterfool/main.go
```

## Usage

### Basic Usage

To convert an image to a radio signal, use the following command:

```sh
./waterfool -image image.png
```

### Command Line Options

- `-image` : Specify the input image file.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Special thanks to the open-source community for their valuable tools and resources.

---