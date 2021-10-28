package oomstore

import (
	"context"
	"fmt"

	"github.com/oom-ai/oomstore/internal/database/metadata"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

// GetFeature: get feature by featureName
func (s *OomStore) GetFeature(ctx context.Context, featureName string) (*types.Feature, error) {
	feature, err := s.metadata.GetFeature(ctx, featureName)
	if err != nil {
		return nil, err
	}
	return feature, nil
}

func (s *OomStore) ListFeature(ctx context.Context, opt types.ListFeatureOpt) (types.FeatureList, error) {
	richFeatures, err := s.metadata.ListRichFeature(ctx, opt)
	if err != nil {
		return nil, err
	}
	return richFeatures.ToFeatureList(), nil
}

func (s *OomStore) UpdateFeature(ctx context.Context, opt types.UpdateFeatureOpt) error {
	return s.metadata.UpdateFeature(ctx, opt)
}

func (s *OomStore) CreateBatchFeature(ctx context.Context, opt types.CreateFeatureOpt) (*types.Feature, error) {
	valueType, err := s.offline.TypeTag(opt.DBValueType)
	if err != nil {
		return nil, err
	}
	group, err := s.metadata.GetFeatureGroup(ctx, opt.GroupName)
	if err != nil {
		return nil, err
	}
	if group.Category != types.BatchFeatureCategory {
		return nil, fmt.Errorf("expected batch feature group, got %s feature group", group.Category)
	}
	if err := s.metadata.CreateFeature(ctx, metadata.CreateFeatureOpt{
		CreateFeatureOpt: opt,
		ValueType:        valueType,
	}); err != nil {
		return nil, err
	}
	return s.metadata.GetFeature(ctx, opt.FeatureName)
}
