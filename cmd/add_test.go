package cmd

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func Test_parseLabels(t *testing.T) {
	t.Run("should success parse", func(t *testing.T) {
		type args struct {
			labels []string
		}
		tests := []struct {
			name string
			args args
			want map[string]string
		}{
			{
				name: "given null string arrays when parse should get empty labels",
				args: args{labels: nil},
				want: map[string]string{},
			},
			{
				name: "given empty string arrays when parse should get empty labels",
				args: args{labels: []string{}},
				want: map[string]string{},
			},
			{
				name: "given label string arrays when parse should get empty labels",
				args: args{labels: []string{"name1:value1", "name2:value2"}},
				want: map[string]string{
					"name1": "value1",
					"name2": "value2",
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := parseLabels(tt.args.labels); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseLabels() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("given wrong string arrays when parse should exit program", func(t *testing.T) {
		if os.Getenv("TEST_SUM_COMMAND_Test_parseLabels") == "true" {
			parseLabels([]string{"test-test1"})
		} else {
			cmd := exec.Command(os.Args[0], "-test.run=Test_parseLabels")
			cmd.Env = append(os.Environ(), "TEST_SUM_COMMAND_Test_parseLabels=true")
			err := cmd.Run()

			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want exit status 1", err)
		}
	})
}
