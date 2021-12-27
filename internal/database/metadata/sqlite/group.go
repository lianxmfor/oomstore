package sqlite

import (
	"context"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/oom-ai/oomstore/internal/database/metadata"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

func createGroup(ctx context.Context, sqlxCtx metadata.SqlxContext, opt metadata.CreateGroupOpt) (int, error) {
	if opt.Category != types.CategoryBatch && opt.Category != types.CategoryStream {
		return 0, fmt.Errorf("illegal category '%s', should be either 'stream' or 'batch'", opt.Category)
	}

	query := "INSERT INTO feature_group(name, entity_id, category, description) VALUES(?, ?, ?, ?)"
	res, err := sqlxCtx.ExecContext(ctx, query, opt.GroupName, opt.EntityID, opt.Category, opt.Description)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return 0, fmt.Errorf("feature group %s already exists", opt.GroupName)
			}
		}
		return 0, err
	}

	groupID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(groupID), err
}
