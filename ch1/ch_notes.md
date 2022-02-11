# The go command

The Go toolchain is responsible for converting the source code into machine language instructions. 

Those tools are acessed trough th *go* command, which has a number of subcommands.

## run

The run command compiles the source code, links the libraries and runs the executable file. 

It does not save the compiled file.

## build

The build command is much like run, but saves an executable binary.

# Imports

Unlike other languages, go code will not compile if there are unecesssary package declarations.
