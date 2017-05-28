package exception_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/qianlnk/gocoding/exception"
)

func throw(err error) {
	if err != nil {
		panic(err)
	}
}

func TestTryCatch(t *testing.T) {
	tr := exception.New()
	tr.Try(
		func() {
			n1, err := strconv.Atoi("123")
			tr.Throw(err)
			n2, err := strconv.Atoi("0")
			tr.Throw(err)
			res := n1 / n2
			fmt.Println(res)
		},
	).Catch(
		func(e exception.Exception) {
			fmt.Println("exception:", e)
		},
	)
}
