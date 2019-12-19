package repository

type Repository interface {
	Run (f func(con Connectable)) error
}

type Connectable interface {
	Connect() interface{}
}
