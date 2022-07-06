package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetTransactionHistoryLinkService get transaction history file id
type GetTransactionHistoryLinkService struct {
	c          *Client
	downloadId *string
}

// StartTime set startTime
func (s *GetTransactionHistoryLinkService) DownloadID(downloadId string) *GetTransactionHistoryLinkService {
	s.downloadId = &downloadId
	return s
}

// Do send request
func (s *GetTransactionHistoryLinkService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionHistoryLink, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/income/asyn/id",
		secType:  secTypeSigned,
	}
	if s.downloadId != nil {
		r.setParam("downloadId", *s.downloadId)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = &TransactionHistoryLink{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IncomeHistory define position margin history info
type TransactionHistoryLink struct {
	ExpectedWaitTIme    int64  `json:"avgCostTimestampOfLast30d"`
	DownloadID          string `json:"downloadId"`
	Url                 string `json:"url"`
	Notified            bool   `json:"notified"`
	IsExpired           bool   `json:"isExpired"`
	ExpirationTimestamp int64  `json:"expirationTimestamp"`
	Status              string `json:"status"` // completed / processing
}

/*
{
    "downloadId":"545923594199212032",
    "status":"completed",     // Enum：completed，processing
    "url":"www.binance.com",  // The link is mapped to download id
    "notified":true,          // ignore
    "expirationTimestamp":1645009771000,  // The link would expire after this timestamp
    "isExpired":null,
}
*/
