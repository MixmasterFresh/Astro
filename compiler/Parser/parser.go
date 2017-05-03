package parser

import "log"

func Parse(buffer string) {
	astro := &Astro{Buffer: string(buffer)}
	astro.Init()

	if err := astro.Parse(); err != nil {
		log.Fatal(err)
	}
	astro.PrintSyntaxTree()
}
