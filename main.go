package main

import (
	"fmt"
	"joaopaulo-creator/monkey-lang/compiler"
	"joaopaulo-creator/monkey-lang/lexer"
	"joaopaulo-creator/monkey-lang/object"
	"joaopaulo-creator/monkey-lang/parser"
	"joaopaulo-creator/monkey-lang/vm"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: monkey <source-file>")
		os.Exit(1)
	}

	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(1)
	}

	// ── Lex + Parse ──────────────────────────────────────────────────────────
	l := lexer.New(string(src))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Fprintln(os.Stderr, "Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, "\t%s\n", msg)
		}
		os.Exit(1)
	}

	// ── Compile ───────────────────────────────────────────────────────────────
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	constants := []object.Object{}
	comp := compiler.NewWithState(symbolTable, constants)

	if err := comp.Compile(program); err != nil {
		fmt.Fprintf(os.Stderr, "Compilation error: %s\n", err)
		os.Exit(1)
	}

	// ── Run on VM ─────────────────────────────────────────────────────────────
	globals := make([]object.Object, vm.GlobalSize)
	machine := vm.NewWithGlobalsStore(comp.ByteCode(), globals)

	if err := machine.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %s\n", err)
		os.Exit(1)
	}

	// ── Print result ──────────────────────────────────────────────────────────
	lastPopped := machine.LastPoppedStackElemen()
	if lastPopped != nil {
		fmt.Println(lastPopped.Inspect())
	}
}
