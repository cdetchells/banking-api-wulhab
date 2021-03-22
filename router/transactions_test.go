package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	transactionmocks "git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/transactions/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestGetAccountTransactions(t *testing.T) {
	createdAt := time.Date(2012, 10, 10, 1, 2, 3, 4, time.Local)

	tests := []struct {
		customerid     string
		id             string
		name           string
		repoErr        error
		repoRes        []*model.Transaction
		wantStatusCode int
		wantBody       string
	}{
		{
			customerid:     "1",
			id:             "1",
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
		},
		{
			customerid:     "2",
			id:             "1",
			name:           "Test Not Found error when repository returns nil",
			repoErr:        nil,
			repoRes:        nil,
			wantStatusCode: 200,
			wantBody:       "null",
		},
		{
			customerid: "1",
			id:         "1",
			name:       "Valid Test",
			repoErr:    nil,
			repoRes: []*model.Transaction{
				{
					ID:          1,
					FromAccount: 1,
					ToAccount:   1,
					CreatedAt:   createdAt,
					Amount:      1,
				},
			},
			wantStatusCode: 200,
			wantBody:       `[{"id":1,"createdAt":"2012-10-10T01:02:03.000000004+01:00","fromAccount":1,"toAccount":1,"amount":1}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			request, _ := http.NewRequest(http.MethodGet, "customers/1/accounts/1/transactions", nil)
			request = mux.SetURLVars(request, map[string]string{
				"customerid": tt.customerid,
				"accountid":  tt.id,
			})

			rr := httptest.NewRecorder()
			transactionRepo := transactionmocks.NewMockRepository(mockCtrl)
			transactionRepo.EXPECT().GetAccountTransactions(gomock.Any(), gomock.Any()).Return(tt.repoRes, tt.repoErr)

			router := &router{transactions: transactionRepo}
			router.GetAccountTransactions(rr, request)

			if rr.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("GetAccountTransactions()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("GetAccountTransactions()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestCreateAccountTransfer(t *testing.T) {
	createdAt := time.Date(2012, 10, 10, 1, 2, 3, 4, time.Local)

	tests := []struct {
		customerid     string
		id             string
		name           string
		repoErr        error
		repoRes        *model.Transaction
		wantStatusCode int
		wantBody       string
		reqBody        string
	}{
		{
			customerid:     "1",
			id:             "1",
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
			reqBody:        `{"createdAt":"2012-10-10T01:02:03.000000004+01:00","toAccount":2,"amount":1}`,
		},
		{
			customerid:     "2",
			id:             "1",
			name:           "Test Not Found error when repository returns nil",
			repoErr:        nil,
			repoRes:        nil,
			wantStatusCode: 200,
			wantBody:       "null",
			reqBody:        `{"createdAt":"2012-10-10T01:02:03.000000004+01:00","toAccount":2,"amount":1}`,
		},
		{
			customerid: "1",
			id:         "1",
			name:       "Valid Test",
			repoErr:    nil,
			repoRes: &model.Transaction{
				ID:          1,
				FromAccount: 1,
				ToAccount:   2,
				CreatedAt:   createdAt,
				Amount:      1,
			},
			wantStatusCode: 200,
			wantBody:       `{"id":1,"createdAt":"2012-10-10T01:02:03.000000004+01:00","fromAccount":1,"toAccount":2,"amount":1}`,
			reqBody:        `{"createdAt":"2012-10-10T01:02:03.000000004+01:00","toAccount":2,"amount":1}`,
		},
	}
	for _, tt := range tests {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		request, _ := http.NewRequest(http.MethodGet, "customers/1/accounts/1/transactions", strings.NewReader(tt.reqBody))
		request.Header.Set("Content-Type", "text/plain")
		request = mux.SetURLVars(request, map[string]string{
			"customerid": tt.customerid,
			"accountid":  tt.id,
		})

		rr := httptest.NewRecorder()
		transactionRepo := transactionmocks.NewMockRepository(mockCtrl)
		transactionRepo.EXPECT().CreateAccountTransaction(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.repoRes, tt.repoErr)

		router := &router{transactions: transactionRepo}
		router.CreateAccountTransfer(rr, request)

		if rr.Result().StatusCode != tt.wantStatusCode {
			t.Errorf("CreateAccountTransfer()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
		}

		if rr.Body.String() != tt.wantBody {
			t.Errorf("CreateAccountTransfer()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
		}
	}
}
