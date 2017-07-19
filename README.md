# What is it?

The `nullable` package simplifies working with `nil` pointers in Go.
You can e.g. inspect if a value is `nil` and return its zero value
or a default value instead.

Example:

```go
import (
    "flag"
    "fmt"
)

func main() {
	book := struct {
		Title string
		Year *int
	}{
		Title: "Harry Potter",
		Year: nil,
	}
	fmt.Println("%s was released in %d.",
		nullable.IntWithDefault(doc.Year, 1997)
	)
	// Output: Harry Potter was released in 1997.
}
```

# Prior art

The [AWS SDK for Go](https://github.com/aws/aws-sdk-go) uses this
mechanism to work with optional values and convert between pointer
and values types. See [here](https://github.com/aws/aws-sdk-go/blob/master/aws/convert_types.go) for examples. 

# License

MIT. See [LICENSE](https://github.com/olivere/nullable/blob/master/LICENSE) file.
