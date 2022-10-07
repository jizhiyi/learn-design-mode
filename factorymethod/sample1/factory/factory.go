package factory

type factoryIn interface {
	createProduct(owner string) Product
	registerProduct(product Product)
}

type Factory struct {
	inner factoryIn
}

func (f *Factory) Create(owner string) Product {
	p := f.inner.createProduct(owner)
	f.inner.registerProduct(p)
	return p
}
