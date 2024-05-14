package main

import (
	"fmt"

	"github.com/otaxhu/trielut"
)

func main() {
	// We instantiate a Trie like this
	t := &trielut.Trie{}

	// We insert some values to it
	t.Insert([]byte("cat"))
	t.Insert([]byte("catty"))
	t.Insert([]byte("carl"))
	t.Insert([]byte("carlson"))

	// * Trying to insert a already inserted entry should return false (failure), prints 'false'
	fmt.Println(t.Insert([]byte("cat")))

	// We chech if the trie has it
	fmt.Println("\nHas:")
	fmt.Println(t.Has([]byte("cat")))
	fmt.Println(t.Has([]byte("catty")))
	fmt.Println(t.Has([]byte("carl")))
	fmt.Println(t.Has([]byte("carlson")))

	// * Trying to search for a inexistent entry should fail, prints 'false'
	fmt.Println(t.Has([]byte("catterpillar")))

	// We delete them
	//
	// * This method is not neccessary to call if you want to free up all of the trie, just throw the
	// reference away and let the GC collect it for you ;D
	fmt.Println("\nDelete:")
	fmt.Println(t.Delete([]byte("cat")))
	fmt.Println(t.Delete([]byte("catty")))
	fmt.Println(t.Delete([]byte("carl")))
	fmt.Println(t.Delete([]byte("carlson")))

	// * Trying to Delete a inexistent entry should fail, prints 'false'
	fmt.Println(t.Delete([]byte("caterpillar")))
}
