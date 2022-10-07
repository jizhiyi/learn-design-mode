package factory

type idCardFactory struct {
	*factory
	data map[string]*IDCard
}

func NewIDCardFactory() FactoryIn {
	f := &idCardFactory{}
	f.data = make(map[string]*IDCard)
	f.factory = &factory{inner: f}
	return f
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
