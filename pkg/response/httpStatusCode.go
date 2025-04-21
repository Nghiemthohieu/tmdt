package response

const (
	ErrCodeSuccess             = 20001
	ErrCodeBadRequest          = 40001
	ErrProductSizeCreateError  = 20011
	ErrProductSizeUpdateError  = 20012
	ErrProductSizeDeleteError  = 20013
	ErrProductSizeGetByIdError = 20014
	ErrProductSizeGetListError = 20015
)

var msg = map[int]string{
	ErrCodeSuccess:             "Success",                   // ✅ Sửa lỗi chính tả
	ErrProductSizeCreateError:  "Create product size error", // ✅ Sửa lỗi chính tả
	ErrProductSizeUpdateError:  "Update product size error",
	ErrProductSizeDeleteError:  "Delete product size error",
	ErrProductSizeGetByIdError: "Get product size by ID error",
	ErrProductSizeGetListError: "Get product size list error",
	ErrCodeBadRequest:          "Bad request error",
}
