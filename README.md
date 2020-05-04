go-hook
=======

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/moutend/go-wca/blob/master/LICENSE

`go-hook`provides low level keyboard and mouse hook for Windows. This package is written in pure Go, `cgo` is not required.

Note: The package is currently tested on Windows 10 64bit edition. Other versions of Windows are not guaranteed to work.

## Prerequisites

Go 1.13 or later

## Usage

Examples are stored in `examples` directory.

- `mouse`: Capturing mouse events.
- `keyboard`: Capturing keyboard events.
- `swapkeys`: Swapping the keyboard input 'A' and 'B".

## Contributing

1. Fork ([https://github.com/moutend/go-hook/fork](https://github.com/moutend/go-hook/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)

## LICENSE

MIT
