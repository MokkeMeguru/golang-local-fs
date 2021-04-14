package env

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	LocalFileRoot string `required:"true" envconfig:"LOCAL_FILE_ROOT"`
	OverWriteFile bool   `default:"false" envconfig:"OW_FILE"`
}

func NewEnv() (*Env, error) {
	// localFileRoot := os.Getenv("LOCAL_FILE_ROOT")
	// if localFileRoot == "" {
	// 	return nil, errors.New("Please set env LOCAL_FILE_ROOT")
	// }
	// overWriteFile := os.Getenv("OVER_WRITE_FILE")

	// return &Env{LocalFileRoot: localFileRoot}, nil

	var env Env
	envconfig.Process("", &env)
	fmt.Println("Env:")
	fmt.Printf(spew.Sdump(env))
	return &env, nil
}
