package message

import (
	"go-simple-svc/contract"
	"go-simple-svc/model"
	"go-simple-svc/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
)

func TestMessage_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	collectionMock := mock.NewMockCollection(ctrl)

	type fields struct {
		collection contract.Collector
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type mockChecker struct {
		checkGet bool
	}
	type getScenario struct {
		messages []*model.Model
		err      error
	}
	type mockScenario struct {
		mockChecker mockChecker
		getScenario getScenario
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockScenario mockScenario
	}{
		{
			name: "1. negative | error while get all document",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/v1/message", nil),
			},
			fields: fields{
				collection: collectionMock,
			},
			mockScenario: mockScenario{
				getScenario: getScenario{
					err: errors.New("error"),
				},
				mockChecker: mockChecker{
					checkGet: true,
				},
			},
		},
		{
			name: "2. positive | successfully get all document",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/v1/message", nil),
			},
			fields: fields{
				collection: collectionMock,
			},
			mockScenario: mockScenario{
				getScenario: getScenario{
					messages: []*model.Model{
						{
							Subject: "msg",
							Message: "hello world!",
						},
					},
				},
				mockChecker: mockChecker{
					checkGet: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				collection: tt.fields.collection,
			}
			if tt.mockScenario.mockChecker.checkGet {
				collectionMock.EXPECT().Get().Return(tt.mockScenario.getScenario.messages, tt.mockScenario.getScenario.err)
			}
			m.Get(tt.args.w, tt.args.r)
		})
	}
}
