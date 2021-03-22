package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/model"
	customermocks "git.codesubmit.io/terem-technologies/banking-api-wulhab/internal/repository/customers/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestGetCustomer(t *testing.T) {

	tests := []struct {
		id             string
		name           string
		repoErr        error
		repoRes        *model.Customer
		wantStatusCode int
		wantBody       string
	}{
		{
			id:             "1",
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
		},
		{
			id:      "1",
			name:    "Valid Test",
			repoErr: nil,
			repoRes: &model.Customer{
				ID:   1,
				Name: "Test",
			},
			wantStatusCode: 200,
			wantBody:       `{"id":1,"name":"Test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			customerRepo := customermocks.NewMockRepository(mockCtrl)
			customerRepo.EXPECT().GetCustomer(gomock.Any()).Return(tt.repoRes, tt.repoErr)

			request, _ := http.NewRequest(http.MethodGet, "customers/1", nil)
			request = mux.SetURLVars(request, map[string]string{
				"id": tt.id,
			})

			rr := httptest.NewRecorder()

			router := &router{customers: customerRepo}
			router.GetCustomer(rr, request)

			if rr.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("GetCustomer()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("GetCustomer()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestGetCustomers(t *testing.T) {

	tests := []struct {
		name           string
		repoErr        error
		repoRes        []*model.Customer
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "Test that error returned when repository returns error",
			repoErr:        errors.New("Error"),
			repoRes:        nil,
			wantStatusCode: 500,
			wantBody:       "",
		},
		{
			name:    "Valid Test",
			repoErr: nil,
			repoRes: []*model.Customer{
				{
					ID:   1,
					Name: "Test",
				},
			},
			wantStatusCode: 200,
			wantBody:       `[{"id":1,"name":"Test"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			customerRepo := customermocks.NewMockRepository(mockCtrl)
			customerRepo.EXPECT().GetCustomers().Return(tt.repoRes, tt.repoErr)

			request := httptest.NewRequest(http.MethodGet, "/customers", nil)
			rr := httptest.NewRecorder()

			router := &router{customers: customerRepo}
			router.GetCustomers(rr, request)

			if rr.Result().StatusCode != tt.wantStatusCode {
				t.Errorf("GetCustomers()\ngot :%#v\nwant: %#v", rr.Result().StatusCode, tt.wantStatusCode)
			}

			if rr.Body.String() != tt.wantBody {
				t.Errorf("GetCustomers()\ngot :%#v\nwant: %#v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}
