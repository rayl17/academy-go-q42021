package router

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/rayl17/academy-go-q42021/controller"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		c controller.ControllerInterface
	}
	tests := []struct {
		name string
		args args
		want *http.Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
