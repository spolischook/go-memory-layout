package memoryTests

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
	"unsafe"
)

var b bool
var i8 int8
var i16 int16
var i32 int32
var i64 int64
var s string
var str struct{}
var str2 struct {
	i1 int8
	i2 int32
	i3 int8
}


func AlignmentGuarantee(i ...interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "VarType\tAlignment of var")

	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(b).Kind(), unsafe.Alignof(b))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(i8).Kind(), unsafe.Alignof(i8))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(i16).Kind(), unsafe.Alignof(i16))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(i32).Kind(), unsafe.Alignof(i32))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(i64).Kind(), unsafe.Alignof(i64))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(s).Kind(), unsafe.Alignof(s))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(str).Kind(), unsafe.Alignof(str))
	fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(str2).Kind(), unsafe.Alignof(str2))
	//fmt.Fprintf(w, "%s\t%d\n", reflect.ValueOf(i[0]).Kind(), reflect.TypeOf(i[0]).Align())
	w.Flush()
}

func AlignmentGuaranteeInterface() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	//fmt.Fprintln(w, "VarType\tCT alignment")
	fmt.Fprintln(w, "VarType\tCT alignment\tRT alignment")
	tps := []interface{}{b, i8, i16, i32, i64, s}
	for _, t := range tps {
		v := reflect.ValueOf(t)
		fmt.Fprintf(w, "%s\t%d", v.Kind(), unsafe.Alignof(t))
		fmt.Fprintf(w, "\t%d", reflect.TypeOf(t).Align())
		fmt.Fprintln(w)
	}

	w.Flush()
}

func StructSize() {
	var str2 struct {
		i1 int8
		i2 int64
		i3 int8
	}
	var str3 struct {
		i1 int8
		i3 int8
		i2 int64
	}

	t2 := reflect.TypeOf(str2)
	t3 := reflect.TypeOf(str3)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "#\tstr2\tstr3\n")
	fmt.Fprintf(w, "Size (byte): \t%d\t%d\n", t2.Size(), t3.Size())

	w.Flush()
}

func StructFieldsOffset() {
	var str2 struct {
		i1 int8
		i2 int64
		i3 int8
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "FieldName\tSize\tOffset\tAllign\n")
	t := reflect.TypeOf(str2)
	n := t.NumField()

	for i := 0; i < n; i++ {
		f := t.Field(i)
		fmt.Fprintf(w,
			"%s\t%d\t%d\t%d\n",
			f.Name, f.Type.Size(), f.Offset, f.Type.Align(),
		)
	}

	w.Flush()
}

func StructFieldsContent() {
	var str2 = struct {
		i1 uint8
		i3 uint8
		i2 uint64
	}{i1: 1, i2: 1<<64 - 1, i3: 1}
	bs := (*[24]byte)(unsafe.Pointer(&str2))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "0\t1\t2\t3\t4\t5\t6\t7\n")
	for i, b := range bs {
		fmt.Fprintf(w, "%#v", b)
		if 0 == (i+1)%8 {
			fmt.Fprintf(w, "\n")
		} else {
			fmt.Fprintf(w, "\t")
		}
	}
	w.Flush()
	fmt.Printf("Bytes: %#v", bs)
}

type Animals struct {
	items []Animal
}

func (as Animals) Say() {
	for _, a := range as.items {
		a.Say()
	}
}
func (as *Animals) Add(a Animal) {
	as.items = append(as.items, a)
}

type Animal interface {
	Say()
}
type Cat struct {}
type Dog struct {}

func (d Dog) Bite() {
	fmt.Println("rrrrrr kus")
}
func (d Dog) Say() {
	fmt.Println("gav")
}

func (c Cat) Jump() {
	fmt.Println("Jump")
}

func (c Cat) Say() {
	fmt.Println("Meu")
}

func Process(as []Animal) {
	for _, a := range as {
		a.Say()
	}
}

func InterfaceIssue() {
	//var animals []Animal
	//e = fmt.Errorf("i'm error")
	//fmt.Println(e)


}
