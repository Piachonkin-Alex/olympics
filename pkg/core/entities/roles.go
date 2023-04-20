package entities

import (
	"log"

	"github.com/pkg/errors"
)

type Role int

const (
	NoneRole Role = iota
	ReaderRole
	WriteRole
	AdminRole
)

func (r Role) String() string {
	switch r {
	case ReaderRole:
		return "reader"
	case WriteRole:
		return "writer"
	case AdminRole:
		return "admin"
	}
	panic("invalid role provided")
}

func RoleFromString(role string) (Role, error) {
	switch role {
	case "reader":
		return ReaderRole, nil
	case "writer":
		return WriteRole, nil
	case "admin":
		return AdminRole, nil
	}
	return Role(0), errors.Errorf("invalid role provided: %s", role)
}

type HandlerDescription struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type HandlerDescriptionConfig struct {
	Description HandlerDescription `yaml:",inline"`
	Role        string             `yaml:"role"`
}

func ConstructHandlersMap(handlers []HandlerDescriptionConfig) map[HandlerDescription]Role {
	m := make(map[HandlerDescription]Role, len(handlers))
	for _, handler := range handlers {
		role, err := RoleFromString(handler.Role)
		if err != nil {
			log.Printf("handler %s %s failed auth registry")
		}
		m[handler.Description] = role
	}
	return m
}
