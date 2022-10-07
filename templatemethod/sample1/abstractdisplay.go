package sample1

type AbstractDisplayIn interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	AbstractDisplayIn
}

func (ab *AbstractDisplay) Display() {
	ab.open()
	for i := 0; i < 5; i++ {
		ab.print()
	}
	ab.close()
}
