package arguments_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/lauri-lyytikainen/composemap/arguments"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name              string
		args              []string
		want              arguments.Arguments
		wantUsageContains string
		wantErr           bool
	}{
		{
			name: "positional compose file",
			args: []string{"docker-compose.yaml"},
			want: arguments.Arguments{
				ComposeFilePath: "docker-compose.yaml",
				Flags:           arguments.Flags{Help: false, Input: ""},
			},
			wantErr: false,
		},
		{
			name: "input flag takes precedence over positional",
			args: []string{"-i", "custom.yaml", "docker-compose.yaml"},
			want: arguments.Arguments{
				ComposeFilePath: "custom.yaml",
				Flags:           arguments.Flags{Help: false, Input: "custom.yaml"},
			},
			wantErr: false,
		}, {
			name: "no args gives empty path",
			args: []string{},
			want: arguments.Arguments{
				ComposeFilePath: "",
				Flags:           arguments.Flags{Help: false, Input: ""},
			},
			wantErr: false,
		},
		{
			name: "help flag shows usage",
			args: []string{"-h"},
			want: arguments.Arguments{
				ComposeFilePath:   "",
				Flags:             arguments.Flags{Help: true, Input: ""},
				UsageMessageShown: false,
			},
			wantUsageContains: "Usage: composemap",
			wantErr:           false,
		},
		{
			name:              "unknown flag shows usage",
			args:              []string{"-bogus"},
			want:              arguments.Arguments{UsageMessageShown: true},
			wantUsageContains: "Usage: composemap",
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			got, got2, gotErr := arguments.ParseArgs(tt.args, &buf)

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseArgs() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseArgs() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
			if tt.wantUsageContains != "" {
				got2()
				if !strings.Contains(buf.String(), tt.wantUsageContains) {
					t.Errorf("usage func output = %q, want it to contain %q", buf.String(), tt.wantUsageContains)
				}
			}
		})
	}
}

func TestHandleArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      arguments.Arguments
		usageFunc func()
		want      arguments.RunSettings
	}{
		{
			name: "filepath",
			args: arguments.Arguments{ComposeFilePath: "docker-compose.yaml", UsageMessageShown: false, Flags: arguments.Flags{Help: false, Input: ""}},
			want: arguments.RunSettings{
				ComposeFilePath: "docker-compose.yaml",
				CanReturnEarly:  false,
			},
		},
		{
			name: "usage message",
			args: arguments.Arguments{ComposeFilePath: "docker-compose.yaml", UsageMessageShown: true, Flags: arguments.Flags{Help: false, Input: ""}},
			want: arguments.RunSettings{
				ComposeFilePath: "docker-compose.yaml",
				CanReturnEarly:  true,
			},
		},
		{
			name: "can return early because of h flag",
			args: arguments.Arguments{ComposeFilePath: "docker-compose.yaml", UsageMessageShown: false, Flags: arguments.Flags{Help: true, Input: ""}},
			want: arguments.RunSettings{
				ComposeFilePath: "docker-compose.yaml",
				CanReturnEarly:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arguments.HandleArgs(tt.args, func() {})
			if tt.want != got {
				t.Errorf("HandleArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
