package metadatav2

import "github.com/oom-ai/oomstore/pkg/oomstore/types"

type CreateFeatureOpt struct {
	types.CreateFeatureOpt
	ValueType string
}

type CreateFeatureGroupOpt struct {
	Name        string
	EntityID    int16
	Description string
	Category    string
}

type CreateRevisionOpt struct {
	Revision    int64
	GroupName   string
	DataTable   string
	Description string
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
