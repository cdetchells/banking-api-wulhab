package transactions

import (
	"errors"
	"reflect"
	"testing"
	"time"

	mocks "git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore/mocks"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	accountmocks "git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/accounts/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetAccountTransactions(t *testing.T) {
	createdAt := time.Date(2012, 10, 10, 1, 2, 3, 4, time.Local)

	tests := []struct {
		customerid   int
		accountid    int
		name         string
		datastoreErr error
		datastoreRes []*model.Transaction
		accountsErr  error
		accountsRes  *model.Account
		want         []*model.Transaction
		wantErr      bool
	}{
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when accounts returns error",
			datastoreErr: nil,
			datastoreRes: nil,
			accountsErr:  errors.New("Error"),
			want:         nil,
			wantErr:      true,
		},
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when datastore returns error",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			accountsErr:  nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			customerid:  1,
			accountid:   1,
			name:        "Valid Test with Transaction out of account",
			accountsErr: nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			datastoreErr: nil,
			datastoreRes: []*model.Transaction{
				{
					ID:          1,
					CreatedAt:   createdAt,
					FromAccount: 1,
					ToAccount:   2,
					Amount:      3,
				},
			},
			want: []*model.Transaction{
				{
					ID:          1,
					CreatedAt:   createdAt,
					FromAccount: 1,
					ToAccount:   2,
					Amount:      -3,
				},
			},
			wantErr: false,
		},
		{
			customerid:  1,
			accountid:   1,
			name:        "Valid Test with Transaction into account",
			accountsErr: nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			datastoreErr: nil,
			datastoreRes: []*model.Transaction{
				{
					ID:          1,
					CreatedAt:   createdAt,
					FromAccount: 2,
					ToAccount:   1,
					Amount:      3,
				},
			},
			want: []*model.Transaction{
				{
					ID:          1,
					CreatedAt:   createdAt,
					FromAccount: 2,
					ToAccount:   1,
					Amount:      3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			if tt.datastoreRes != nil || tt.datastoreErr != nil {
				datastore.EXPECT().GetTransactions().Return(tt.datastoreRes, tt.datastoreErr)
			}

			accounts := accountmocks.NewMockRepository(mockCtrl)
			accounts.EXPECT().GetCustomerAccount(gomock.Any(), gomock.Any()).Return(tt.accountsRes, tt.accountsErr)

			repository := &repository{datastore: datastore, accounts: accounts}
			transactions, err := repository.GetAccountTransactions(tt.customerid, tt.accountid)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(transactions, tt.want) {
				t.Errorf("GetAccountTransactions()\ngot :%#v\nwant: %#v", transactions, tt.want)
			}
		})
	}
}

func TestCreateAccountTransaction(t *testing.T) {
	createdAt := time.Date(2012, 10, 10, 1, 2, 3, 4, time.Local)

	tests := []struct {
		customerid    int
		accountid     int
		name          string
		datastoreErr  error
		datastoreRes  []*model.Transaction
		accountsErr   error
		accountsRes   *model.Account
		getAccountErr error
		getAccountRes *model.Account
		want          *model.Transaction
		wantErr       bool
		input         *model.NewTransaction
	}{
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when accounts returns error",
			datastoreErr: nil,
			datastoreRes: nil,
			accountsErr:  errors.New("Error"),
			want:         nil,
			wantErr:      true,
			input:        nil,
		},
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when accounts returns nil",
			datastoreErr: nil,
			datastoreRes: nil,
			accountsErr:  nil,
			accountsRes:  nil,
			want:         nil,
			wantErr:      true,
			input:        nil,
		},
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when getAccount returns error",
			datastoreErr: nil,
			datastoreRes: nil,
			accountsErr:  nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			getAccountErr: errors.New("Error"),
			getAccountRes: nil,
			want:          nil,
			wantErr:       true,
			input:         &model.NewTransaction{},
		},
		{
			customerid:   1,
			accountid:    1,
			name:         "Test that error returned when Account Ids are the same",
			datastoreErr: nil,
			datastoreRes: nil,
			accountsErr:  nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			getAccountErr: nil,
			getAccountRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			want:    nil,
			wantErr: true,
			input: &model.NewTransaction{
				CreatedAt: createdAt,
				ToAccount: 1,
				Amount:    1,
			},
		},
		{
			customerid:   1,
			accountid:    1,
			name:         "Valid Test",
			datastoreErr: nil,
			datastoreRes: []*model.Transaction{
				{
					ID:          1,
					CreatedAt:   createdAt,
					FromAccount: 1,
					ToAccount:   2,
					Amount:      1,
				},
			},
			accountsErr: nil,
			accountsRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			getAccountErr: nil,
			getAccountRes: &model.Account{
				ID:       2,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			want: &model.Transaction{
				ID:          2,
				CreatedAt:   createdAt,
				FromAccount: 1,
				ToAccount:   2,
				Amount:      1,
			},
			wantErr: false,
			input: &model.NewTransaction{
				CreatedAt: createdAt,
				ToAccount: 2,
				Amount:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			if tt.datastoreRes != nil || tt.datastoreErr != nil {
				datastore.EXPECT().GetTransactions().Return(tt.datastoreRes, tt.datastoreErr)
			}

			accounts := accountmocks.NewMockRepository(mockCtrl)
			accounts.EXPECT().GetCustomerAccount(gomock.Any(), gomock.Any()).Return(tt.accountsRes, tt.accountsErr)
			if tt.getAccountRes != nil || tt.getAccountErr != nil {
				accounts.EXPECT().GetAccount(gomock.Any()).Return(tt.getAccountRes, tt.getAccountErr)
			}

			repository := &repository{datastore: datastore, accounts: accounts}
			transactions, err := repository.CreateAccountTransaction(tt.customerid, tt.accountid, tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(transactions, tt.want) {
				t.Errorf("GetAccountTransactions()\ngot :%#v\nwant: %#v", transactions, tt.want)
			}
		})
	}
}
