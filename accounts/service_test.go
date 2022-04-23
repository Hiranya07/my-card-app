package accounts

import (
	"context"
	"errors"
	"my-card-app/db/repositories"
	dbMock "my-card-app/db/repositories/mock"
	"my-card-app/models"
	"reflect"
	"testing"

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

func TestAccountService_GetAccount(t *testing.T) {

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
		want      models.AccountDetails
		wantErr   bool
		setup     func() repositories.IRepository
	}{
		{
			name:      "Failed to retreive record",
			accServer: AccountService{},
			args: args{
				ctx:       context.Background(),
				accountId: 2,
			},
			want:    models.AccountDetails{},
			wantErr: true,
			setup: func() repositories.IRepository {

				dbM.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(models.AccountDetails{}, errors.New("erro"))
				return dbM

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.accServer.repo = tt.setup
			got, err := tt.accServer.GetAccount(tt.args.ctx, tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
