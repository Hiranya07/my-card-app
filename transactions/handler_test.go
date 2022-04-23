package transactions

import (
	"bytes"
	"encoding/json"
	"errors"
	"my-card-app/models"
	mock "my-card-app/transactions/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

var defaultModel = models.Transactions{
	AccountId:       12,
	OperationTypeId: 1,
	Amount:          49,
}

func TestTransaction_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	trxMock := mock.NewMockITransaction(ctrl)

	bodyInvalid := []byte{}

	body, err := json.Marshal(defaultModel)
	if err != nil {
		t.Errorf("Can't encode body")
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name  string
		trx   Transaction
		args  args
		setup func() ITransaction
	}{
		{
			name: "tc01-invalid ",
			trx:  Transaction{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(bodyInvalid)),
			},
			setup: func() ITransaction {

				return trxMock

			},
		},
		{
			name: "tc02-failure ",
			trx:  Transaction{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body)),
			},
			setup: func() ITransaction {
				trxMock.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("error"))
				return trxMock

			},
		},

		{
			name: "tc03-Success",
			trx:  Transaction{},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body)),
			},
			setup: func() ITransaction {
				trxMock.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
				return trxMock

			},
		},
	}
	for _, tt := range tests {
		tt.trx.trxServ = tt.setup
		t.Run(tt.name, func(t *testing.T) {
			tt.trx.Create(tt.args.w, tt.args.r)
		})
	}
}
