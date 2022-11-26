package queryParameter

import (
	"smartdom/pkg/type/pagination"
	"smartdom/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
