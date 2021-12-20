package sqlutil

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/internal/database/online"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

func Get(ctx context.Context, db *sqlx.DB, opt online.GetOpt, backend types.BackendType) (dbutil.RowMap, error) {
	featureNames := opt.FeatureList.Names()
	tableName := OnlineTableName(opt.RevisionID)
	qt, err := dbutil.QuoteFn(backend)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = ?`, qt(featureNames...), qt(tableName), qt(opt.Entity.Name))

	record, err := db.QueryRowxContext(ctx, db.Rebind(query), opt.EntityKey).SliceScan()
	if err != nil {
		if err == sql.ErrNoRows || dbutil.IsTableNotFoundError(err, backend) {
			return nil, nil
		}
		return nil, err
	}

	rs, err := deserializeIntoRowMap(record, opt.FeatureList, backend)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

// response: map[entity_key]map[feature_name]feature_value
func MultiGet(ctx context.Context, db *sqlx.DB, opt online.MultiGetOpt, backend types.BackendType) (map[string]dbutil.RowMap, error) {
	featureNames := opt.FeatureList.Names()
	tableName := OnlineTableName(opt.RevisionID)
	qt, err := dbutil.QuoteFn(backend)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT %s, %s FROM %s WHERE %s in (?);`, qt(opt.Entity.Name), qt(featureNames...), qt(tableName), qt(opt.Entity.Name))
	sql, args, err := sqlx.In(query, opt.EntityKeys)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryxContext(ctx, db.Rebind(sql), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getFeatureValueMapFromRows(rows, opt.FeatureList, backend)
}

func getFeatureValueMapFromRows(rows *sqlx.Rows, features types.FeatureList, backend types.BackendType) (map[string]dbutil.RowMap, error) {
	featureValueMap := make(map[string]dbutil.RowMap)
	for rows.Next() {
		record, err := rows.SliceScan()
		if err != nil {
			return nil, err
		}
		entityKey, values := dbutil.DeserializeString(record[0], backend), record[1:]
		rowMap, err := deserializeIntoRowMap(values, features, backend)
		if err != nil {
			return nil, err
		}
		featureValueMap[entityKey] = rowMap
	}
	return featureValueMap, nil
}

func deserializeIntoRowMap(values []interface{}, features types.FeatureList, backend types.BackendType) (dbutil.RowMap, error) {
	rs := map[string]interface{}{}
	for i := range values {
		value := values[i]
		typedValue, err := deserializeByTag(value, features[i].ValueType, backend)
		if err != nil {
			return nil, err
		}
		rs[features[i].Name] = typedValue
	}
	return rs, nil
}
