package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const fileConfigName 		= "stash-notes.toml"	
const defaultConfigContent 	= `[storage]
workspace = "./"
use_current_directory = false

[editor]
auto_save = true
auto_save_delay_ms = 300

[interface]
theme = "dark"
show_preview = true
`

type (
	config struct {
		Storage storage
		Editor editor
		EditorInterface editorInterface `toml:"interface"`
	}

	storage struct {
		Workspace string 
		UseCurrentDirectory bool `toml:"use_current_directory"`
	}

	editor struct {
		AutoSave bool `toml:"auto_save"`
		AutoSaveDelayms int32 `toml:"auto_save_delay_ms"`

	}

	editorInterface struct {
		Theme string
		ShowPreview bool `toml:"show_preview"`
	}
)

func ApplyConfig() {

}

func hasConfigFile() bool {	
	if _, err := os.Stat(fileConfigName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
		return false
	}

	return true
}

func CreateDefaultConfigFile() {
	err := os.WriteFile(fileConfigName, []byte(defaultConfigContent), 0644)
	if err != nil {
		panic(err)
	}
}

func LoadConfig() config {

	var conf config

	hasConfigFile := hasConfigFile()
	if !hasConfigFile {
		CreateDefaultConfigFile()
		fmt.Println("not has config file")
	}

	_, err := toml.DecodeFile(fileConfigName, &conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("configs acessadas com sucesso")
	fmt.Println(conf.Storage.Workspace)
	fmt.Println(conf.Editor.AutoSaveDelayms)

	return conf
}

