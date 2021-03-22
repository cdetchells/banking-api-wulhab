package accounts

import (
	"errors"
	"reflect"
	"testing"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore/mocks"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"github.com/golang/mock/gomock"
)

func TestGetCustomerAccounts(t *testing.T) {
	tests := []struct {
		id           int
		name         string
		datastoreErr error
		datastoreRes []*model.Account
		want         []*model.Account
		wantErr      bool
	}{
		{
			id:           1,
			name:         "Get account for customer 1",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			id:           1,
			name:         "Get accounts for customer 1",
			datastoreErr: nil,
			datastoreRes: []*model.Account{
				{
					ID:       1,
					Type:     1,
					Customer: 1,
					Balance:  0,
				},
			},
			want: []*model.Account{
				{
					ID:       1,
					Type:     1,
					Customer: 1,
					Balance:  0,
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
			datastore.EXPECT().GetAccounts().Return(tt.datastoreRes, tt.datastoreErr)
			repository := &repository{datastore: datastore}

			accounts, err := repository.GetCustomerAccounts(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(accounts, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", accounts, tt.want)
			}
		})
	}
}

func TestGetCustomerAccount(t *testing.T) {
	tests := []struct {
		customerid   int
		id           int
		name         string
		datastoreErr error
		datastoreRes []*model.Account
		want         *model.Account
		wantErr      bool
	}{
		{
			customerid:   1,
			id:           1,
			name:         "Get account for customer 1",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			customerid:   1,
			id:           1,
			name:         "Get account for customer 1",
			datastoreErr: nil,
			datastoreRes: []*model.Account{
				{
					ID:       1,
					Type:     1,
					Customer: 1,
					Balance:  0,
				},
			},
			want: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			datastore.EXPECT().GetAccounts().Return(tt.datastoreRes, tt.datastoreErr)
			repository := &repository{datastore: datastore}

			accounts, err := repository.GetCustomerAccount(tt.customerid, tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(accounts, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", accounts, tt.want)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		id           int
		name         string
		datastoreErr error
		datastoreRes []*model.Account
		want         *model.Account
		wantErr      bool
	}{
		{
			id:           1,
			name:         "Get account for customer 1",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			id:           1,
			name:         "Get account for customer 1",
			datastoreErr: nil,
			datastoreRes: []*model.Account{
				{
					ID:       1,
					Type:     1,
					Customer: 1,
					Balance:  0,
				},
			},
			want: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 1,
				Balance:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			datastore.EXPECT().GetAccounts().Return(tt.datastoreRes, tt.datastoreErr)
			repository := &repository{datastore: datastore}

			accounts, err := repository.GetAccount(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(accounts, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", accounts, tt.want)
			}
		})
	}
}

func TestCreateCustomerAccount(t *testing.T) {

	tests := []struct {
		name         string
		datastoreErr error
		datastoreRes []*model.Account
		want         *model.Account
		wantErr      bool
		input        *model.NewAccount
	}{
		{
			name:         "Test that error returned when datastore returns error",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
			input:        &model.NewAccount{},
		},
		{
			name:         "Valid Test",
			datastoreErr: nil,
			datastoreRes: []*model.Account{
				{
					ID:       1,
					Customer: 1,
					Type:     1,
					Balance:  1,
				},
			},
			want: &model.Account{
				ID:       2,
				Customer: 1,
				Type:     1,
				Balance:  1,
			},
			wantErr: false,
			input: &model.NewAccount{
				Customer: 1,
				Type:     1,
				Balance:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			if tt.datastoreRes != nil || tt.datastoreErr != nil {
				datastore.EXPECT().GetAccounts().Return(tt.datastoreRes, tt.datastoreErr)
			}

			repository := &repository{datastore: datastore}
			transactions, err := repository.CreateCustomerAccount(tt.input)

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
