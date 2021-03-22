package customers

import (
	"errors"
	"reflect"
	"testing"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/datastore/mocks"
	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	"github.com/golang/mock/gomock"
)

func TestGetCustomers(t *testing.T) {
	tests := []struct {
		name         string
		datastoreErr error
		datastoreRes []*model.Customer
		want         []*model.Customer
		wantErr      bool
	}{
		{
			name:         "Test that error returned when datastore returns error",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "Valid Test",
			datastoreErr: nil,
			datastoreRes: []*model.Customer{
				{
					ID:   1,
					Name: "Hello",
				},
			},
			want: []*model.Customer{
				{
					ID:   1,
					Name: "Hello",
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
			datastore.EXPECT().GetCustomers().Return(tt.datastoreRes, tt.datastoreErr)
			repository := &repository{datastore: datastore}

			customers, err := repository.GetCustomers()

			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(customers, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", customers, tt.want)
			}
		})
	}
}

func TestGetCustomer(t *testing.T) {
	tests := []struct {
		id           int
		name         string
		datastoreErr error
		datastoreRes []*model.Customer
		want         *model.Customer
		wantErr      bool
	}{
		{
			id:           1,
			name:         "Test that error returned when datastore returns error",
			datastoreErr: errors.New("Error"),
			datastoreRes: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			id:           1,
			name:         "Valid Test",
			datastoreErr: nil,
			datastoreRes: []*model.Customer{
				{
					ID:   1,
					Name: "Hello",
				},
			},
			want: &model.Customer{
				ID:   1,
				Name: "Hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			datastore := mocks.NewMockDataStore(mockCtrl)
			datastore.EXPECT().GetCustomers().Return(tt.datastoreRes, tt.datastoreErr)
			repository := &repository{datastore: datastore}

			accounts, err := repository.GetCustomer(tt.id)

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
