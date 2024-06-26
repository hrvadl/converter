package rate

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/hrvadl/converter/gw/internal/transport/http/handlers/rate/mocks"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()
	type args struct {
		rg  Getter
		log *slog.Logger
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Should create new handler correctly",
			args: args{
				rg:  mocks.NewMockGetter(gomock.NewController(t)),
				log: slog.Default(),
			},
			want: &Handler{
				log: slog.Default(),
				rg:  mocks.NewMockGetter(gomock.NewController(t)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewHandler(tt.args.rg, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerGetRate(t *testing.T) {
	t.Parallel()
	type fields struct {
		log *slog.Logger
		rg  Getter
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func(t *testing.T, getter Getter)
		want   int
	}{
		{
			name: "Should return 200 when rate getter succeeded",
			fields: fields{
				log: slog.Default(),
				rg:  mocks.NewMockGetter(gomock.NewController(t)),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			setup: func(t *testing.T, getter Getter) {
				t.Helper()
				g, ok := getter.(*mocks.MockGetter)
				if !ok {
					t.Fatal("Failed to cast getter to mock")
				}

				g.EXPECT().GetRate(gomock.Any()).Times(1).Return(float32(39.8), nil)
			},
			want: http.StatusOK,
		},
		{
			name: "Should return 400 when rate getter succeeded",
			fields: fields{
				log: slog.Default(),
				rg:  mocks.NewMockGetter(gomock.NewController(t)),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			setup: func(t *testing.T, getter Getter) {
				t.Helper()
				g, ok := getter.(*mocks.MockGetter)
				if !ok {
					t.Fatal("Failed to cast getter to mock")
				}

				g.EXPECT().
					GetRate(gomock.Any()).
					Times(1).
					Return(float32(0), errors.New("failed to get rate"))
			},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.setup(t, tt.fields.rg)
			h := &Handler{
				log: tt.fields.log,
				rg:  tt.fields.rg,
			}
			h.GetRate(tt.args.w, tt.args.r)
			if got := tt.args.w.Result().StatusCode; got != tt.want {
				t.Errorf("GetRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
