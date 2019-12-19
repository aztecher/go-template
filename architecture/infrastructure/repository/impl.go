package repository

type RepositoryImpl struct {
}

func NewRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (impl *RepositoryImpl) Run(f func(con Connectable)) error {
	con := NewConnectableImpl()
	f(con)
	return nil
}

type ConnectableImpl struct {
}

func NewConnectableImpl() *ConnectableImpl {
	return &ConnectableImpl{}
}

func (impl *ConnectableImpl) Connect() interface{} {
	return "hogehoge"
}


