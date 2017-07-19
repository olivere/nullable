// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package nullable_test

import (
	"fmt"

	"github.com/olivere/nullable"
)

func Example() {
	book := struct {
		Title string
		Year  *int
	}{
		Title: "Harry Potter",
		Year:  nil,
	}
	fmt.Printf("%s was released in %d.",
		book.Title,
		nullable.IntWithDefault(book.Year, 1997))
	// Output: Harry Potter was released in 1997.
}
