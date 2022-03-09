package universe

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		module string
	}
	tests := []struct {
		name string
		args args
		want Access
	}{
		{
			"#00",
			args{"delete_vehicle_group"},
			Access{"ByXZ80FF", "delete"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseModule(tt.args.module); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
