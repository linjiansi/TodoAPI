package config

import (
	"testing"
)

func TestEnvOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return key value",
			args: args{
				key:          "testKey",
				defaultValue: "defaultValue",
			},
			want: "testValue",
		},
		{
			name: "return defaultValue value",
			args: args{
				key:          "",
				defaultValue: "defaultValue",
			},
			want: "defaultValue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.args.key, tt.want)
			got := envOrDefault(tt.args.key, tt.args.defaultValue)

			if tt.want != got {
				t.Errorf("envOrDefault = %v, want = %v", got, tt.want)
			}
		})
	}
}
