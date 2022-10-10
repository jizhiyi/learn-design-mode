package sample

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("从早上至下午")
	d.builder.MakeItems([]string{
		"早上好",
		"下午好",
	})
	d.builder.MakeString("晚上")
	d.builder.MakeItems([]string{
		"晚上好",
		"晚安",
		"再见",
	})
	d.builder.Close()
}
