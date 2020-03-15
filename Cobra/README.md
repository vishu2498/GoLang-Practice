>Use this command first to install the 'cobra' package: `go get -u github.com/spf13/cobra/cobra`

### Steps to implement this application made by Cobra package:

1. `go build`
2. `./using-cobra egmessage`

* `go build` will build the executable file for the application
* Using `./using-cobra` will print our help message and tell our available custom commands with available flags (like help).
* `./using-cobra egmessage` will print our custom message mentioned in the function.

_It is better to run `go clean` before using `go build`._