package morestubs

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type EntitiesStub struct {
	Entities
}

func (es EntitiesStub) GetPets(userId string) ([]Pet, error) {
	switch userId {
	case "1":
		return []Pet{{Name: "Mittens"}}, nil
	case "2":
		return []Pet{{Name: "Fluffy"}, {Name: "Fang"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userId)
	}
}

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userId   string
		petNames []string
	}{
		{"case1", "1", []string{"Mittens"}},
		{"case2", "2", []string{"Fluffy", "Fang"}},
		{"case3", "3", nil},
	}
	l := Logic{EntitiesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			expected := d.petNames
			actual, err := l.GetPetNames(d.userId)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
