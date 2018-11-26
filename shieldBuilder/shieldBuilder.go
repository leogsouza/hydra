package shieldBuilder

import (
	"strings"
)

type shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

type shBuilder struct {
	code string
}

// NewShieldBuilder constructor for a new shield builder
func NewShieldBuilder() *shBuilder {
	return new(shBuilder)
}

// RaiseFront enables shield on the front
func (sh *shBuilder) RaiseFront() *shBuilder {
	sh.code += "F"
	return sh
}

// RaiseFront enables shield on the back
func (sh *shBuilder) RaiseBack() *shBuilder {
	sh.code += "B"
	return sh
}

// RaiseFront enables shield on the right
func (sh *shBuilder) RaiseRight() *shBuilder {
	sh.code += "R"
	return sh
}

// RaiseFront enables shield on the left
func (sh *shBuilder) RaiseLeft() *shBuilder {
	sh.code += "L"
	return sh
}

func (sh *shBuilder) Build() *shield {
	code := sh.code
	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
