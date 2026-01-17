package main

import (
	"fmt"
	"os"

	lnk "github.com/cuhsat/go-lnk"
)

func main() {

	f, err := os.Open("remote.directory.xp.test")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// // lnk files are small-ish, no reason not to read everything at once.
	// // lnkBytes, err := ioutil.ReadAll(f)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Printf("Read %d bytes.\n", len(lnkBytes))

	// h, err := lnk.Header(f)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = h

	// // fmt.Println(lnk.StructToJSON(h, true))
	// fmt.Println(h)

	// fmt.Println(h.LinkFlags)

	// lt, err := lnk.LinkTarget(f)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = lt
	// fmt.Println(lnk.StructToJSON(lt, true))

	// li, err := lnk.LinkInfo(f)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = li
	// fmt.Println(lnk.StructToJSON(li, true))

	// st, err := lnk.StringData(f, h.LinkFlags)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = st
	// fmt.Println(lnk.StructToJSON(st, true))

	// edb, err := lnk.DataBlock(f)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = edb
	// // fmt.Println(lnk.StructToJSON(edb, true))

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	ln, err := lnk.Read(f, uint64(fi.Size()))
	if err != nil {
		panic(err)
	}
	fmt.Println(ln.Header)

}
