package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"111", "Bi", "HCM", "700000", "1996", "1"},
		{"112", "Bi2", "HCM", "700000", "1996", "1"},
	}
	return CustomerRepositoryStub{customers}
}

