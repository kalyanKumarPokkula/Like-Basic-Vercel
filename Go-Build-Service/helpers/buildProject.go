package helpers

import (
	"log"
	"os"
	"os/exec"
)

func npmInstall(folderPath string) error{
	cmd := exec.Command("npm" , "install")
	cmd.Dir = folderPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = cmd.Wait()
    if err != nil {
        log.Fatal(err)
		return err
    }

	return nil

}

func BuildProject(folderPath string) error{
	// commands := []string{"npm install" , "npm build"}
	var err error
	err = npmInstall(folderPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("npm" , "run" , "build")
	cmd.Dir = folderPath
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err= cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = cmd.Wait()
    if err != nil {
        log.Fatal(err)
		return err
    }

	return nil

}