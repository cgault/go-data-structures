package hashtable_test

import (
	"testing"

	"github.com/cgault/go-data-structures/hashtable"
)

func Test(t *testing.T) {
	ht := hashtable.Init()

	ht.Insert("ERIC", map[string]string{"firstName": "Eric", "lastName": "Cartman"})
	ht.Insert("KENNY", map[string]string{"firstName": "Kenny", "lastName": "McCormick"})
	ht.Insert("KYLE", map[string]string{"firstName": "Kyle", "lastName": "Broflovski"})
	ht.Insert("STAN", map[string]string{"firstName": "Stan", "lastName": "Marsh"})
	ht.Insert("RANDY", map[string]string{"firstName": "Randy", "lastName": "Marsh"})
	ht.Insert("BUTTERS", map[string]string{"firstName": "Butters", "lastName": "Stotch"})

	// search for valid entry
	v, ok := ht.Search("ERIC")
	if !ok {
		t.Errorf("Search failed!")
		return
	}

	if fn := v.(map[string]string)["firstName"]; fn != "Eric" {
		t.Errorf("firstName got: %q, want: 'Eric'", fn)
	}

	if ln := v.(map[string]string)["lastName"]; ln != "Cartman" {
		t.Errorf("firstName got: %q, want: 'Cartman'", ln)
	}

	// search for invalid entry
	if _, ok := ht.Search("BOB"); ok {
		t.Error("Search for BOB: got true, expected false")
	}

	// delete valid entry
	ht.Delete("ERIC")

	if _, ok := ht.Search("ERIC"); ok {
		t.Error("Delete ERIC failed")
	}
}
