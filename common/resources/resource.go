package resources

import "sync"

const (
	file     ResourceType = iota //0
	service                      //1
	redis                        //2
	domain                       //3
	mysql                        //4
	postgres                     //5
	cvm                          //6
)

var (
	ResourceInfo = []string{
		"FILE",       //0
		"SERVICE",    //1
		"REDIS",      //2
		"DOMAIN",     //3
		"MYSQL",      //4
		"POSTGRESQL", //5
		"CVM",        //6
	}
)

type ResourceType int8

// Resourcer
type Resourcer interface {
	// resource's name
	Name() string
	// resource's locate ip or path or else
	Locate() map[string]string
	// check resource's health
	// health check use io.Reader?
	IsHealth() bool
	// resource's type
	Type() ResourceType
	// update resource information
	// TODO: how to update?
	Update(interface{}) error
	// resource dependency
	DependBy() []*Resource
	DependOn() []*Resource
}

// Resource
type Resource struct {
	mu           sync.RWMutex
	name         string
	locate       map[string]string
	health       bool
	resourceType ResourceType
	content      interface{}
	dependOn     []*Resource
	dependBy     []*Resource
}

// Name return resource's name
func (r *Resource) Name() string {
	return r.name
}

// Locate resource's locate ip or path or else
func (r *Resource) Locate() map[string]string {
	return r.locate
}

// IsHealth health check use io.Reader?
// check resource's health
func (r *Resource) IsHealth() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.health
}

// resource's type
func (r *Resource) Type() ResourceType {
	return r.resourceType
}

// Update update resource information
func (r *Resource) Update(interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return nil
}

// DependOn resource depend on
func (r *Resource) DependOn() []*Resource {
	return r.dependOn
}

// DependBy resource depend by
func (r *Resource) DependBy() []*Resource {
	return r.dependBy
}
