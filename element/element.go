package element

type Element interface {
	GetName() string
	CanRemove() bool
	Remove()
	Replace(map[string]string)
	Check() error
	Parse() error
	Do(map[string]string) error
	Val(map[string]string) []byte
}

type Parent struct {
	Name     string
	Enable   bool
	Children []Element
}

func (p *Parent) GetName() string {
	return p.Name
}

func (p *Parent) CanRemove() bool {
	return p == nil || !p.Enable
}

//Remove removes disable elements
func (p *Parent) Remove() {
	var children []Element
	for _, child := range p.Children {
		if child.CanRemove() {
			continue
		}
		child.Remove()
		children = append(children, child)
	}
	p.Children = children
}

//Replace replaces variables
func (p *Parent) Replace(vars map[string]string) {
	for _, child := range p.Children {
		child.Replace(vars)
	}
}

//Check checks settings
func (p *Parent) Check() (err error) {
	for _, child := range p.Children {
		if err = child.Check(); err != nil {
			return
		}
	}
	return
}

//Parse parses files/variables
func (p *Parent) Parse() (err error) {
	for _, child := range p.Children {
		if err = child.Parse(); err != nil {
			return
		}
	}
	return
}

//Val returns a byte stream, the purpose is to be used for sampler parameters
func (p *Parent) Val(map[string]string) []byte {
	return nil
}
