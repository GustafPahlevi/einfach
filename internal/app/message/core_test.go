package message

import (
	"einfach-msg/contract"
	"einfach-msg/test/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	collectionMock := mock.NewMockCollection(ctrl)
	type args struct {
		collection contract.Collector
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{
			name: "instantiate message dependency",
			args: args{
				collection: collectionMock,
			},
			want: New(collectionMock),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
