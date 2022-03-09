package universe

import "testing"

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	{
		"version",
		"47c133e0cc85556e2472318809678e868b7a99f6",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Version(); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModule(t *testing.T) {
	type args struct {
		module string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "00",
			args: args{
				module: "download_report_vehicle_activation",
			},
			want:  "GFc68qWy",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetModule(tt.args.module)
			if got != tt.want {
				t.Errorf("GetModule() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetModule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}