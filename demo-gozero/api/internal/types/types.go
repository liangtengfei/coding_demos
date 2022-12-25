// Code generated by goctl. DO NOT EDIT.
package types

type AddReq struct {
	Book  string `form:"book"`
	Price int64  `form:"price"`
}

type AddResp struct {
	Ok bool `json:"ok"`
}

type CheckReq struct {
	Book string `form:"book"`
}

type CheckResp struct {
	Found bool  `json:"found"`
	Price int64 `json:"price"`
}

type BorrowReq struct {
	Book string `form:"book"`
}

type BorrowResp struct {
	Ok       bool  `json:"ok"`
	BorrowId int64 `json:"borrowId"`
}

type Record struct {
	Id        int64  `json:"id"`
	Book      string `json:"book"`
	CreatedAt string `json:"createdAt"`
}

type HistoryResp struct {
	List []Record `json:"list"`
}