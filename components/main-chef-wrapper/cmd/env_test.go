package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/chef/chef-workstation/components/main-chef-wrapper/dist"
	"github.com/spf13/cobra"
)

func NewEnvCmd(s []string) *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "Prints environment variables used by %s",
		RunE: func(cmd *cobra.Command, args []string) error {
			return passThroughCommand(dist.WorkstationExec, "", s)
		},
	}
}

func Test_EnvCommand(t *testing.T) {
	s := []string{"env"}
	cmd := NewEnvCmd(s)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	// fmt.Println("x is ...", x)
	out, err := ioutil.ReadAll(b)
	fmt.Println("out is ...", string(out))
	// fmt.Println("error is ...", err)
	// fmt.Println(reflect.TypeOf(err))
	fmt.Println(reflect.TypeOf(string(out)))
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `` {
		t.Fatalf("expected \"%s\" got \"%s\"", ``, string(out))
	}
}
