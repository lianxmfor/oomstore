package postgres

import (
	"context"
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
	"github.com/onestore-ai/onestore/internal/database"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

func (db *DB) ListRevision(ctx context.Context, groupName *string) ([]*types.Revision, error) {
	query := "SELECT * FROM feature_group_revision"
	var cond []interface{}
	if groupName != nil {
		query += " WHERE group_name = $1"
		cond = append(cond, *groupName)
	}
	revisions := make([]*types.Revision, 0)

	if err := db.SelectContext(ctx, &revisions, query, cond...); err != nil {
		return nil, err
	}
	return revisions, nil
}

func (db *DB) GetRevision(ctx context.Context, groupName string, revision int64) (*types.Revision, error) {
	query := "SELECT * FROM feature_group_revision WHERE group_name = $1 and revision = $2"
	var rs types.Revision
	if err := db.GetContext(ctx, rs, query, groupName, revision); err != nil {
		return nil, err
	}
	return &rs, nil
}

func (db *DB) InsertRevision(ctx context.Context, opt types.InsertRevisionOpt) error {
	return database.WithTransaction(db.DB, ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		query := "INSERT INTO feature_group_revision(group_name, revision, data_table, description) VALUES ($1, $2, $3, $4)"
		_, err := tx.ExecContext(ctx, query, opt.GroupName, opt.Revision, opt.DataTable, opt.Description)
		if err != nil {
			return err
		}

		if opt.UpdateGroupInfo {
			query := "UPDATE feature_group SET revision = $1, data_table = $2 WHERE name = $3"
			if _, err := tx.ExecContext(ctx, query, opt.Revision, opt.DataTable, opt.GroupName); err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *DB) BuildRevisionRanges(ctx context.Context, groupName string) ([]*types.RevisionRange, error) {
	query := fmt.Sprintf(`
		SELECT
			revision AS min_revision,
			LEAD(revision, 1, %d) OVER (ORDER BY revision) AS max_revision,
			data_table
		FROM feature_group_revision
		WHERE group_name = $1
	`, math.MaxInt64)

	var ranges []*types.RevisionRange
	if err := db.SelectContext(ctx, &ranges, query, groupName); err != nil {
		return nil, err
	}
	return ranges, nil
}
