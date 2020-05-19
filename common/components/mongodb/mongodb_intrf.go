package mongodb

//MDBInterface interface
type MDBInterface interface {
	Init(MDBConfig) error
	Insert(string, interface{}) error
	Update(string, interface{}) error
	Remove(string, interface{}) error
	FindOne(string, map[string]interface{}) (interface{}, error)
	FindAll(string, map[string]interface{}) ([]interface{}, error)
	Count(string) (int, error)
	Find(string, map[string]interface{}, map[string]interface{}, int, int) ([]interface{}, error)
}
