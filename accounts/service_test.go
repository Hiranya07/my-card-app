package accounts

import (
	"context"
	"errors"
	"testing"

	"my-card-app/db/repositories"
	dbMock "my-card-app/db/repositories/mock"

	"github.com/golang/mock/gomock"
)

func TestAccountService_CreateAccount(t *testing.T) {

	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	dbM := dbMock.NewMockIRepository(ctrl)
	type args struct {
		ctx       context.Context
		accountId int
	}
	tests := []struct {
		name      string
		accServer AccountService
		args      args
		wantErr   bool
		setup     func() repositories.IRepository
	}{
		{

			name:      "tc01-error during create account",
			accServer: AccountService{},
			args: args{
				ctx:       context.Background(),
				accountId: 1234,
			},
			wantErr: true,
			setup: func() repositories.IRepository {
				dbM.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(errors.New("erro"))
				return dbM
			},
		},

		{

			name:      "tc02-successfully create account",
			accServer: AccountService{},
			args: args{
				ctx:       context.Background(),
				accountId: 1234,
			},
			wantErr: false,
			setup: func() repositories.IRepository {
				dbM.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)
				return dbM
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.accServer.repo = tt.setup
			if err := tt.accServer.CreateAccount(tt.args.ctx, tt.args.accountId); (err != nil) != tt.wantErr {
				t.Errorf("AccountService.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
