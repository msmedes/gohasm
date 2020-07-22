# gohasm

This is an implemenation of the assembler for the Hack computer from a certain MOOC which shall remain nameless (they don't want people posting code so others can do it themselves which I realize I'm not honoring by posting this don't @ me).  It's 20x faster than my [python](https://github.com/msmedes/pyhasm) implementation. 

To run:

`make run`
will run 
`go run <file_to_convert>` 
which includes the compilation step, or build your own binary (`go build ./cmd/hasm/hasm.go`) and run it that way (~20x faster).