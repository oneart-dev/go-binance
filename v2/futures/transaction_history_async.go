package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetTransactionHistoryService get transaction history file id
type GetTransactionHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

// StartTime set startTime
func (s *GetTransactionHistoryService) StartTime(startTime int64) *GetTransactionHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetTransactionHistoryService) EndTime(endTime int64) *GetTransactionHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *GetTransactionHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*TransactionHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/income/asyn",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*TransactionHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IncomeHistory define position margin history info
type TransactionHistory struct {
	ExpectedWaitTIme int64  `json:"avgCostTimestampOfLast30d"`
	DownloadID       string `json:"downloadId"`
}
