package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)


type AppList struct {
	Apps       []string
	SystemApps []string
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
    
   var systemApps []string 

    err = filepath.Walk(systemDir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error: %v", err)
        }

        if strings.HasSuffix(info.Name(), ".exe") {
            systemApps = append(systemApps, info.Name())
        }

        return nil
    })

    var apps []string

    err = filepath.Walk(appDir, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error: %v", err)
        }

        if strings.HasSuffix(info.Name(), ".exe") {
            apps = append(apps, info.Name())
        }

        return err
    })

    err = writeSSlice("systemApps.txt", systemApps)
    if err != nil {
    	return 
    }

    err = writeSSlice("Apps.txt", apps)
    if err != nil {
    	return 
    }

    return
}

func writeSSlice(filename string, strings []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	var fString string
	for _, s := range strings {
		fString = fString + s
	}
	file.WriteString(fString)
	return nil
}
