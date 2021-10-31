package mode

import (
	"reflect"
	"testing"
)

func TestUserFilter_Filter(t *testing.T) {
	type fields struct {
		Expression IExpression
	}
	type args struct {
		users []*User
	}
	tests := []struct {
		name string
		args args
		want []*User
	}{
		{
			name: "切片过滤",
			args: args{users: []*User{
				{1, "静静", 20},
				{2, "琳琳", 25},
				{3, "佳佳", 26},
				{4, "露露", 27},
			}},
			want: []*User{
				{3, "佳佳", 26},
				{4, "露露", 27},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUserFilter("age > 25")
			if got := uf.Filter(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserFilter.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
