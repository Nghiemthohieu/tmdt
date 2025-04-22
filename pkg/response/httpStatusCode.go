package response

const (
	ErrCodeSuccess             = 20001
	ErrCodeBadRequest          = 40001
	ErrProductSizeCreateError  = 20011
	ErrProductSizeUpdateError  = 20012
	ErrProductSizeDeleteError  = 20013
	ErrProductSizeGetByIdError = 20014
	ErrProductSizeGetListError = 20015
	ErrProductColorCreateError  = 20021
	ErrProductColorUpdateError  = 20022
	ErrProductColorDeleteError  = 20023
	ErrProductColorGetByIdError = 20024
	ErrProductColorGetListError = 20025
	ErrCategoryCreateError  = 20031
	ErrCategoryUpdateError  = 20032
	ErrCategoryDeleteError  = 20033
	ErrCategoryGetByIdError = 20034
	ErrCategoryGetListError = 20035
)

var msg = map[int]string{
	ErrCodeSuccess:             "Success",                   // ✅ Sửa lỗi chính tả
	ErrProductSizeCreateError:  "Create product size error", // ✅ Sửa lỗi chính tả
	ErrProductSizeUpdateError:  "Update product size error",
	ErrProductSizeDeleteError:  "Delete product size error",
	ErrProductSizeGetByIdError: "Get product size by ID error",
	ErrProductSizeGetListError: "Get product size list error",
	ErrCodeBadRequest:          "Bad request error",
	ErrProductColorCreateError:  "Create product Color error", // ✅ Sửa lỗi chính tả
	ErrProductColorUpdateError:  "Update product Color error",
	ErrProductColorDeleteError:  "Delete product Color error",
	ErrProductColorGetByIdError: "Get product Color by ID error",
	ErrProductColorGetListError: "Get product Color list error",
	ErrCategoryCreateError:  "Create product Category error", // ✅ Sửa lỗi chính tả
	ErrCategoryUpdateError:  "Update product Category error",
	ErrCategoryDeleteError:  "Delete product Category error",
	ErrCategoryGetByIdError: "Get product Category by ID error",
	ErrCategoryGetListError: "Get product Category list error",
}
