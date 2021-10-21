package speedy

const version = "1.0.0"

// Speedy App
type Speedy struct {
	AppName string
	Debug   bool //dev o production
	Version string
}

// Path Application Start
func (s *Speedy) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

// Initialize Folder
func (s *Speedy) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder
		err := s.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
