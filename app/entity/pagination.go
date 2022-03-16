package entity

type Pagination struct {
  PageCount int32 `json:"page_count"`
  PageSize  int32 `json:"page_size"`
  Page      int32 `json:"page"`
  Total     int32 `json:"total"`
}
