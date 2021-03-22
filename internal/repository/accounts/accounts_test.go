package accounts

import (
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
		wantErr      error
	}{
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
			wantErr: nil,
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

			if err != nil && err != tt.wantErr {
				t.Errorf("GetCustomerAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(accounts, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", accounts, tt.want)
			}
		})
	}
}
