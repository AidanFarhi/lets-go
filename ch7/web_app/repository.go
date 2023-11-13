package main

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Concrete data store
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// A factory function to create an instance of a SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Joe",
			"3": "Jim",
		},
	}
}
