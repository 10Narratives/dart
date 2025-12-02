package pgcl

type componentOptions struct {
}

type ComponentOption func(co *componentOptions)

func defaultComponentOptions() *componentOptions {
	return &componentOptions{}
}
