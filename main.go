package main

import (
	"github.com/spolischook/go-memory-layout/memoryTests"
)

func main() {
	//memoryTests.AlignmentGuarantee()
	//memoryTests.AlignmentGuaranteeInterface()
	//memoryTests.StructSize()
	//memoryTests.StructFieldsOffset()
	//memoryTests.StructFieldsContent()

	//var i memoryTests.Animal
	//i = memoryTests.Cat{}

	d := (memoryTests.Animal)(memoryTests.Cat{})
	d.Say()

	//cat := memoryTests.Cat{}
	//cats := []memoryTests.Cat{cat, cat}
	//cats2 := cats.([]memoryTests.Animal)
	//memoryTests.Process(cats)
	//animals := memoryTests.Animals{}
	//animals.Add(cat)
	//animals.Add(cat)
	//animals.Say()
	//memoryTests.InterfaceIssue()
}

