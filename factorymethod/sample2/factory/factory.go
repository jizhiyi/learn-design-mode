package factory

type FactoryIn interface {
	createProduct(owner string) Product
	registerProduct(product Product)
	Create(owner string) Product
}

type factory struct {
	inner FactoryIn
}

func (factory *factory) Create(owner string) Product {
	product := factory.inner.createProduct(owner)
	factory.inner.registerProduct(product)
	return product
}
