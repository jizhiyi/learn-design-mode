package factory

type idCardFactory struct {
	data map[string]*IDCard
}

func NewIDCardFactory() *Factory {
	factory := &idCardFactory{}
	factory.data = make(map[string]*IDCard)
	return &Factory{inner: factory}
}

func (f *idCardFactory) createProduct(owner string) Product {
	idCard, ok := f.data[owner]
	if ok {
		return idCard
	}
	return newIDCard(owner)
}

func (f *idCardFactory) registerProduct(product Product) {
	idCard, ok := product.(*IDCard)
	if !ok {
		return
	}
	f.data[idCard.GetOwner()] = idCard
}
