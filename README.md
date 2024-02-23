# CGo Project Example

This project demonstrates how to build and run a Go application that uses a C library for both static and dynamic linking. The application increments a counter in a loop, using a function from the C library to manipulate the counter values.

## Project Structure

- `lib/`: Contains the C library source (`addlib.c`) and header (`addlib.h`) files.
- `main.go`: The Go application source file that uses the C library.
- `Makefile`: Contains commands to build the static and dynamic versions of the application and to clean up build artifacts.

## Prerequisites

- Go (1.x) [Installation guide](https://golang.org/doc/install)
- GCC (for compiling C code) [Installation guide](https://gcc.gnu.org/install/)
- Make [Installation guide](https://www.gnu.org/software/make/)

## Building the Application

### Static Linking

To build the application with static linking:

```sh
make static
```

This command compiles the C library into a static library and builds the Go application with the static library linked.

### Dynamic Linking

To build the application with dynamic linking:

```sh
make dynamic
```

This command compiles the C library into a shared library and builds the Go application with the shared library linked.

## Running the Application

After building the application, you can run it using the following command:

```sh
./build/cgo_static
./build/cgo_dynamic
```

## Checking Dynamic Linking 

To check the dynamic linking of the application, you can use the following command: 

```sh
make ldd 
```

This command prints the shared library dependencies of the application.
