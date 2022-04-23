package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	mock "my-card-app/accounts/mock"
	"my-card-app/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var defaultModel = models.AccountInput{
	AccId: 1234,
}

func TestAccount_CreateAccount(t *testing.T) {

	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := mock.NewMockIAccountService(ctrl)

	body, err := json.Marshal(defaultModel)
	if err != nil {
		t.Errorf("Can't encode body")
	}

	bodyInvalid := []byte{}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name  string
		acc   Account
		args  args
		setup func() IAccountService
	}{
		{
			name: "tc01-invalid ",
			acc:  Account{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(bodyInvalid)),
			},
			setup: func() IAccountService {
				//m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return m

			},
		},
		{
			name: "tc02-valid request with failure response",
			acc:  Account{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body)),
			},
			setup: func() IAccountService {
				m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("error occurrred"))
				return m

			},
		},
		{
			name: "tc03-valid request with successful response",
			acc:  Account{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body)),
			},
			setup: func() IAccountService {
				m.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return m

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.acc.accServe = tt.setup
			tt.acc.CreateAccount(tt.args.w, tt.args.r)
		})
	}
}

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name string
		want Account
	}{
		{
			name: "tc_01",
			want: Account{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
