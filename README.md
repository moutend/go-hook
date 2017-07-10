# go-hook (beta)

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/moutend/go-wca/blob/master/LICENSE

Package `go-hook` is the package for hooking low-level keyboard or mouse input on Windows.

Note: The package is currently tested on Windows 10 64bit edition. Other versions of Windows are not supported yet.

## Prerequisites

- Go 1.8 or later

## Examples

The following example reports all keyboard input events.

```go
```

Open Command Prompt or PowerShell, and then run the example above. This program reports the keyboard input even if the focus is out of that window.

```console
C:\> go run examples/keyboard.go
```

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
