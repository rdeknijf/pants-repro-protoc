package repro_protoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	person := &Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "john.doe@example.com",
	}

	assert.Equal(t, "John Doe", person.Name)
	assert.Equal(t, int32(1234), person.Id)
	assert.Equal(t, "john.doe@example.com", person.Email)
}

func TestMother(t *testing.T) {
	mother := &Mother{
		Name: "Jane Doe",
	}

	assert.Equal(t, "Jane Doe", mother.Name)
}
