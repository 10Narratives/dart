package resource

type ResourceName struct {
	name   string
	tokens []string
}

func NewResourceName(name string) (*ResourceName, error) {
	return nil, nil
}

func (r *ResourceName) Value() string

func (r *ResourceName) ID() string

func (r *ResourceName) Collection() string

func (r ResourceName) ParentName() *ResourceName

func (r ResourceName) ChildName(collection, id string) *ResourceName
