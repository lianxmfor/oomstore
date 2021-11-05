package metadata

import "github.com/oom-ai/oomstore/pkg/oomstore/types"

type CreateFeatureOpt struct {
	types.CreateFeatureOpt
	ValueType string
}

type CreateFeatureGroupOpt struct {
	types.CreateFeatureGroupOpt
	Category string
}

type CreateRevisionOpt struct {
	Revision    int64
	GroupName   string
	DataTable   string
	Description string
	Anchored    bool
}

type GetRevisionOpt struct {
	GroupName  *string
	Revision   *int64
	RevisionId *int32
}

type ListRevisionOpt struct {
	GroupName  *string
	DataTables []string
}
