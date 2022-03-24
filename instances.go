package migration

var instances map[string]*Instance

func init() {
	instances = make(map[string]*Instance)
}

// AddInstance add instance
func AddInstance(key string, instance *Instance) {
	instances[key] = instance
}

// GetInstance get instance
func GetInstance(key string) *Instance {
	return instances[key]
}
