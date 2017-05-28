# Exception

基于`panic` `recover` 实现类似`try...catch...`的异常处理机制。

# Usage

```golang
	tr := exception.New()
	tr.Try(
		func() {
			n1, err := strconv.Atoi("123a")
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
```

output

```shell
exception: strconv.ParseInt: parsing "123a": invalid syntax
```

change `123a` to `123`

output

```shell
exception: runtime error: integer divide by zero
```