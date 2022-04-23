package transactions

import (
	"context"
	"errors"
	"my-card-app/models"
	"my-card-app/transactions/db/repositories"
	tMock "my-card-app/transactions/db/repositories/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTransactionService_Create(t *testing.T) {

	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	trxMock := tMock.NewMockIRepository(ctrl)
	type args struct {
		ctx  context.Context
		trnx models.Transactions
	}
	tests := []struct {
		name    string
		trx     TransactionService
		args    args
		wantErr bool
		setup   func() repositories.IRepository
	}{
		{
			name: "tc_01:failure in creation of transaction",
			trx:  TransactionService{},
			args: args{
				ctx:  context.Background(),
				trnx: models.Transactions{},
			},
			wantErr: true,
			setup: func() repositories.IRepository {
				trxMock.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("error"))
				return trxMock
			},
		},
		{
			name: "tc_02:succesful in creation of transaction",
			trx:  TransactionService{},
			args: args{
				ctx:  context.Background(),
				trnx: models.Transactions{},
			},
			wantErr: true,
			setup: func() repositories.IRepository {
				trxMock.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(errors.New("error"))
				return trxMock
			},
		},
	}
	for _, tt := range tests {
		tt.trx.repo = tt.setup
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.trx.Create(tt.args.ctx, tt.args.trnx); (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
