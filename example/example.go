package example

func Run() {
	f := &Foo{
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       50,
	}

	f.reflect()

	Sum(1,1)

}
