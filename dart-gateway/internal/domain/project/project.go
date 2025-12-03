package projectdomain

import (
	"time"

	"github.com/10Narratives/dart/pkg/garp/resource"
)

type Project struct {
	name        *resource.ResourceName
	displayName string
	description string
	createTime  time.Time
	updateTime  time.Time
}

var _ resource.Resource = &Project{}

func (p *Project) ResourceName() *resource.ResourceName {
	return p.name
}

func (p *Project) DisplayName() string {
	return p.displayName
}

func (p *Project) Description() string {
	return p.description
}

func (p *Project) CreateTime() time.Time {
	return p.createTime
}

func (p *Project) UpdateTime() time.Time {
	return p.updateTime
}

func NewProject(name *resource.ResourceName, displayName, description string, createTime, updateTime time.Time) *Project {
	return &Project{
		name:        name,
		displayName: displayName,
		description: description,
		createTime:  createTime,
		updateTime:  updateTime,
	}
}
