package accounts

import (
	"reflect"
	"testing"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
)

func TestGetCustomerAccounts(t *testing.T) {
	tests := []struct {
		id   int
		name string
		want []*model.Account
	}{
		{
			id:   1,
			name: "Get accounts for customer 1",
			want: []*model.Account{
				{
					ID:       1,
					Type:     1,
					Customer: 1,
					Balance:  0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repository := &repository{}

			accounts, _ := repository.GetCustomerAccounts(tt.id)

			if !reflect.DeepEqual(accounts, tt.want) {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", accounts, tt.want)
			}
		})
	}
}
