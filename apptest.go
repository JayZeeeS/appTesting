package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	Name string
	Path string
}

type AppList struct {
	Apps       []App
	SystemApps []App
}

func (*AppList)SyncList() error {
    return nil
}



func main() {
    getApps()
}


func getApps() (out AppList, err error) {
    
    systemDir := os.Getenv("SystemRoot")+`\System32`
    appDir := os.Getenv("ProgramFiles")
    
    err = filepath.Walk(systemDir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error: %v", err)
        }

        if strings.HasSuffix(info.Name(), ".exe") {
            fmt.Println(path)
        }

        return nil
    })

    err = filepath.Walk(appDir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error: %v", err)
        }

        if strings.HasSuffix(info.Name(), ".exe") {
            fmt.Println(path)
        }

        return nil
    })

    return out, nil
}
