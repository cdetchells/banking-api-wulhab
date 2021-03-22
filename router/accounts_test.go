package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	accountmocks "git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/accounts/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestGetCustomerAccount(t *testing.T) {

	tests := []struct {
		customerid     string
		id             string
		name           string
		repoErr        error
		repoRes        *model.Account
		repoCall       bool
		wantStatusCode int
		wantBody       string
	}{
		{
			customerid:     "",
			id:             "1",
			name:           "Test Invalid Customer Id",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			repoCall:       false,
			wantStatusCode: 400,
			wantBody:       "Invalid Customer Id",
		},
		{
			customerid:     "1",
			id:             "",
			name:           "Test Invalid Account Id",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			repoCall:       false,
			wantStatusCode: 400,
			wantBody:       "Invalid Account Id",
		},
		{
			customerid:     "1",
			id:             "1",
			name:           "Test that error returned when datastore returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			repoCall:       true,
			wantStatusCode: 500,
			wantBody:       "",
		},
		{
			customerid:     "1",
			id:             "1",
			name:           "Test that error returned when datastore returns error",
			repoErr:        nil,
			repoRes:        nil,
			repoCall:       true,
			wantStatusCode: 404,
			wantBody:       "Account Not Found",
		},
		{
			customerid: "1",
			id:         "1",
			name:       "Valid Test",
			repoErr:    nil,
			repoRes: &model.Account{
				ID:       1,
				Customer: 1,
				Type:     1,
				Balance:  1,
			},
			repoCall:       true,
			wantStatusCode: 200,
			wantBody:       `{"id":1,"customer":1,"type":1,"balance":1}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			request, _ := http.NewRequest(http.MethodGet, "customers/1/accounts/1", nil)
			request = mux.SetURLVars(request, map[string]string{
				"customerid": tt.customerid,
				"accountid":  tt.id,
			})

			rr := httptest.NewRecorder()

			accountRepo := accountmocks.NewMockRepository(mockCtrl)
			if tt.repoCall {
				accountRepo.EXPECT().GetCustomerAccount(gomock.Any(), gomock.Any()).Return(tt.repoRes, tt.repoErr)
			}

			router := &router{accounts: accountRepo}
			router.GetCustomerAccount(rr, request)

			if rr.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("GetCustomerAccount()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("GetCustomerAccount()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestGetCustomerAccounts(t *testing.T) {

	tests := []struct {
		customerid     string
		name           string
		repoErr        error
		repoRes        []*model.Account
		repoCall       bool
		wantStatusCode int
		wantBody       string
	}{
		{
			customerid:     "",
			name:           "Test Invalid Account Id",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			repoCall:       false,
			wantStatusCode: 400,
			wantBody:       "Invalid Account Id",
		},
		{
			customerid:     "1",
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			repoCall:       true,
			wantStatusCode: 500,
			wantBody:       "",
		},
		{
			customerid:     "2",
			name:           "Test Not Found error when repository returns nil",
			repoErr:        nil,
			repoRes:        nil,
			repoCall:       true,
			wantStatusCode: 404,
			wantBody:       "",
		},
		{
			customerid: "1",
			name:       "Valid Test",
			repoErr:    nil,
			repoRes: []*model.Account{
				{
					ID:       1,
					Customer: 1,
					Type:     1,
					Balance:  1,
				},
			},
			repoCall:       true,
			wantStatusCode: 200,
			wantBody:       `[{"id":1,"customer":1,"type":1,"balance":1}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			request, _ := http.NewRequest(http.MethodGet, "customers/1/accounts", nil)
			request = mux.SetURLVars(request, map[string]string{
				"id": tt.customerid,
			})

			rr := httptest.NewRecorder()
			accountRepo := accountmocks.NewMockRepository(mockCtrl)
			if tt.repoCall {
				accountRepo.EXPECT().GetCustomerAccounts(gomock.Any()).Return(tt.repoRes, tt.repoErr)
			}

			router := &router{accounts: accountRepo}
			router.GetCustomerAccounts(rr, request)

			if rr.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("GetCustomerAccounts()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestCreateCustomerAccount(t *testing.T) {

	tests := []struct {
		id             string
		name           string
		repoErr        error
		repoRes        *model.Account
		wantStatusCode int
		wantBody       string
		reqBody        string
	}{
		{
			id:             "1",
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
			reqBody:        `{"type":1,"customer":2,"balance":1}`,
		},
		{
			id:             "1",
			name:           "Test error when repository returns nil",
			repoErr:        nil,
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
			reqBody:        `{"type":1,"customer":2,"balance":1}`,
		},
		{
			id:      "1",
			name:    "Valid Test",
			repoErr: nil,
			repoRes: &model.Account{
				ID:       1,
				Type:     1,
				Customer: 2,
				Balance:  1,
			},
			wantStatusCode: 200,
			wantBody:       `{"id":1,"customer":2,"type":1,"balance":1}`,
			reqBody:        `{"type":1,"customer":2,"balance":1}`,
		},
	}
	for _, tt := range tests {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		request, _ := http.NewRequest(http.MethodGet, "customers/1/accounts", strings.NewReader(tt.reqBody))
		request.Header.Set("Content-Type", "text/plain")
		request = mux.SetURLVars(request, map[string]string{
			"id": tt.id,
		})

		rr := httptest.NewRecorder()
		accountRepo := accountmocks.NewMockRepository(mockCtrl)
		accountRepo.EXPECT().CreateCustomerAccount(gomock.Any()).Return(tt.repoRes, tt.repoErr)

		router := &router{accounts: accountRepo}
		router.CreateCustomerAccount(rr, request)

		if rr.Result().StatusCode != tt.wantStatusCode {
			t.Errorf("CreateCustomerAccount()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
		}

		if rr.Body.String() != tt.wantBody {
			t.Errorf("CreateCustomerAccount()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
		}
	}
}
