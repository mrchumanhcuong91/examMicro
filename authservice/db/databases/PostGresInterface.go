package databases

type PG interface{
	InsertKey(id string, value string) (error)
	GetKey(id string) (string, error)
	DeleteKey(id string) error
}