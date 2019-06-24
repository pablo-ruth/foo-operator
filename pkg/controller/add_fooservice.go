package controller

import (
	"github.com/pablo-ruth/foo-operator/pkg/controller/fooservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, fooservice.Add)
}
