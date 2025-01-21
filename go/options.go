package kit

// ToolkitOption defines a function type for configuring a Toolkit
type ToolkitOption func(*kit)

func WithToolkitName(name string) KitOption {
	return func(t *Kit) {
		t.name = name
	}
}

func WithToolkitDescription(description string) KitOption {
	return func(t *Kit) {
		t.description = description
	}
}

func WithTools(tools ...Tool) KitOption {
	return func(t *Kit) {
		for _, tool := range tools {
			t.tools[tool.GetName()] = tool
		}
	}
}
