package core

var RootDir string = "."

func GetPath(path string) string {
	return RootDir + path
}