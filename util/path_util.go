package util

import (
	"cybercoin/dal/const_dal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetFullPath(path string) string {
	file, _ := exec.LookPath(os.Args[0])
	executePath, _ := filepath.Abs(file)
	index := strings.Index(executePath, const_dal.PROJECT_NAME)
	executePath = executePath[:index] + const_dal.PROJECT_NAME + string(os.PathSeparator) + path
	return executePath
}
