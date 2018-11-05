package db

type Spec struct {
	ID   int
	Cpu  int
	Mem  int
	Disk int
}

type Host struct {
	ID       int
	HostName string
	IP       string
	EcsID    string
	SpecID   int
}

type File struct {
	ID      int
	Name    string
	Path    string
	Version string
}
