# Applicator

Applicator is a sandboxing tool that allows you to run applications in a controlled environment. It is designed to be lightweight and easy to use, providing a simple way to isolate applications from the rest of the system. 

**This project should not be used for production purposes yet. It is still under severe development and may not be 100% functional or secure.**

## Usage

```shell
./Applicator --help
```

## Running an application

```shell
./Applicator run --app <application> [--args <args>]
```

## Running the test application

```shell
./Applicator run --app Testing/com.arvasyn.testing --verbose
```

## Building

To build Applicator, you need to have Go installed on your system. Once you have Go installed, you can build Applicator by running the following command:

```shell
go build -o Applicator Source/Main.go
```

## Documentation

There is no documentation available yet as this project is still under development. You can find commands and command arguments using the `--help` flag.
