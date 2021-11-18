package types

import (
	"io"
)

type CreateFeatureOpt struct {
	FeatureName string
	GroupName   string
	DBValueType string
	Description string
}

type ListFeatureOpt struct {
	EntityName   *string
	GroupName    *string
	FeatureNames *[]string
}

type UpdateFeatureOpt struct {
	FeatureName    string
	NewDescription string
}

type CreateEntityOpt struct {
	EntityName  string
	Length      int
	Description string
}

type CreateGroupOpt struct {
	GroupName   string
	EntityName  string
	Description string
}

type ExportFeatureValuesOpt struct {
	RevisionID   int
	FeatureNames []string
	Limit        *uint64
}

type ImportOpt struct {
	GroupName   string
	Description string
	DataSource  CsvDataSource
	Revision    *int64
}

type CsvDataSource struct {
	Reader    io.Reader
	Delimiter string
}

type OnlineGetOpt struct {
	FeatureNames []string
	EntityKey    string
}

type OnlineMultiGetOpt struct {
	FeatureNames []string
	EntityKeys   []string
}

type JoinOpt struct {
	FeatureNames []string
	EntityRows   <-chan EntityRow
}

type UpdateEntityOpt struct {
	EntityName     string
	NewDescription string
}

type UpdateGroupOpt struct {
	GroupName           string
	NewDescription      *string
	NewOnlineRevisionID *int
}

type SyncOpt struct {
	RevisionID int
}
