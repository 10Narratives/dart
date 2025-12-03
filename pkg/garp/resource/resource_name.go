package resource

import (
	"errors"
	"strings"
)

type ResourceName struct {
	name   string
	tokens []string
}

func NewResourceName(name string) (*ResourceName, error) {
	if name == "" {
		return nil, errors.New("resource name cannot be empty")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 0 || len(tokens)%2 == 0 {
		return nil, errors.New("invalid resource name format: must be collection/id[/collection/id...]")
	}

	for i := 0; i < len(tokens); i += 2 {
		if tokens[i] == "" { // collection name
			return nil, errors.New("collection name cannot be empty")
		}
		if i+1 >= len(tokens) || tokens[i+1] == "" { // id
			return nil, errors.New("resource ID cannot be empty")
		}
	}

	return &ResourceName{
		name:   name,
		tokens: tokens,
	}, nil
}

func (r *ResourceName) Value() string {
	if r == nil {
		return ""
	}
	return r.name
}

func (r *ResourceName) ID() string {
	if r == nil || len(r.tokens) == 0 {
		return ""
	}
	return r.tokens[len(r.tokens)-1]
}

func (r *ResourceName) Collection() string {
	if r == nil || len(r.tokens) < 2 {
		return ""
	}
	return r.tokens[len(r.tokens)-2]
}

func (r *ResourceName) ParentName() *ResourceName {
	if len(r.tokens) == 2 {
		return nil
	}
	parentTokens := r.tokens[:len(r.tokens)-2]
	parentName := strings.Join(parentTokens, "/")

	parent, _ := NewResourceName(parentName)
	return parent
}

func (r *ResourceName) ChildName(collection, id string) *ResourceName {
	if collection == "" || id == "" {
		return nil
	}

	var newName string
	if r.name != "" {
		newName = r.name + "/" + collection + "/" + id
	} else {
		newName = collection + "/" + id
	}

	child, _ := NewResourceName(newName)
	return child
}
