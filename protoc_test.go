package createpb

import (
	"testing"
)

func Test_Generate(t *testing.T) {
	if err := Generate(); err != nil {
		println(err.Error())
	}
}

func Test_recursionReadFile(t *testing.T) {
	type args struct {
		dirname string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "递归读取所有文件",
			args: args{dirname: "."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recursionReadFile(tt.args.dirname)
			for _, v := range files {
				println(v)
			}
		})
	}
}
