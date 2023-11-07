package service

type product struct{}

func New() *product {
	return &product{}
}

func (srv *product) Get(id string) error {
	return nil
}
