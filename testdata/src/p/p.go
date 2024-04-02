package main

type a struct {
	a string
}

func cha() {
	ch := make(chan interface{})
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	select {
	case <-ch:
		return
	case <-ch1:
		{
			a := 1
			_ = a
			{
				a := 1
				_ = a

				return
			}
			return
		}
		return
	case <-ch2:
		{
			a := 1
			_ = a
			return // want "no blank line before"
		}

		return // want "no blank line needed"
	}
}

func baz() {
	switch 0 {
	case 0:
		a := 1
		_ = a
		fallthrough
	case 1:
		a := 1
		_ = a
		break
	case 2:
		break
	}
}

func foo() int {
	v := []int{}
	for range v {
		return 0
	}

	for range v {
		for range v {
			return 0
		}
		return 0
	}

	o := []int{
		0, 1,
	}
	return o[0]
}

func fooa() int {
	o := []int{0, 1}
	return o[0] // want "no blank line before"
}

func foob(s string) *a {
	o := &a{
		a: s,
	}
	return o
}

func fooc() int {
	defer bar()

	return 0
}

func fooe(s string) interface{} {
	o := foob(
		s,
	)
	return o
}

func foof() interface{} {
	o := fooe(
		"a",
	)
	switch s := o.(type) {
	case *a:
		return s
	default:
	}
	return o
}

func bar() int {
	o := 1
	if o == 1 {
		if o == 0 {
			return 1
		}
		return 0
	}

	return o // want "no blank line needed"
}

func main() {
	return
}

func bugNoAssignSmthHandling() string {
	switch 0 {
	case 0:
		o := struct {
			foo string
		}{
			"foo",
		}
		return o.foo
	case 1:
		o := struct {
			foo string
		}{
			"foo",
		}

		return o.foo // want "no blank line needed"
	}
	return ""
}

func bugNoExprSmthHandling(string) {
	switch 0 {
	case 0:
		bugNoExprSmthHandling("")
		return // want "no blank line before"
	case 1:
		bugNoExprSmthHandling(
			"",
		)

		return // want "no blank line needed"
	}
}

func bugNoDeferSmthHandling(string) {
	switch 0 {
	case 0:
		defer bugNoDeferSmthHandling(
			"",
		)
		return
	case 1:
		defer bugNoDeferSmthHandling(
			"",
		)

		return // want "no blank line needed"
	}
}

func bugNoGoSmthHandling(string) {
	switch 0 {
	case 0:
		go bugNoGoSmthHandling(
			"",
		)
		return
	case 1:
		go bugNoGoSmthHandling(
			"",
		)

		return // want "no blank line needed"
	}
}
