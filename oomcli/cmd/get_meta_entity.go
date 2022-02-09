package cmd

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/oom-ai/oomstore/pkg/oomstore"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
	"github.com/oom-ai/oomstore/pkg/oomstore/types/apply"
)

type getMetaEntityOption struct {
	entityName *string
}

var getMetaEntityOpt getMetaEntityOption

var getMetaEntityCmd = &cobra.Command{
	Use:   "entity [entity_name]",
	Short: "Get existing entity given specific conditions",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			getMetaEntityOpt.entityName = &args[0]
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		oomStore := mustOpenOomStore(ctx, oomStoreCfg)
		defer oomStore.Close()

		var listEntityOpt types.ListEntityOpt
		if getMetaEntityOpt.entityName != nil {
			listEntityOpt.EntityNames = &[]string{*getMetaEntityOpt.entityName}
		}
		entities, err := oomStore.ListEntity(ctx, listEntityOpt)
		if err != nil {
			exit(err)
		}

		if err = serializeEntitiesToWriter(ctx, os.Stdout, oomStore, entities, *getMetaOutput); err != nil {
			exitf("failed printing entities, error: %+v\n", err)
		}
	},
}

func init() {
	getMetaCmd.AddCommand(getMetaEntityCmd)
}

func serializeEntitiesToWriter(ctx context.Context, w io.Writer, oomStore *oomstore.OomStore,
	entities types.EntityList, outputOpt string) error {

	switch outputOpt {
	case YAML:
		// TODO: Use entity ids to filter, rather than taking them all out
		groups, err := oomStore.ListGroup(ctx, nil)
		if err != nil {
			return err
		}

		groupItems, err := groupsToApplyGroupItems(ctx, oomStore, groups)
		if err != nil {
			return err
		}

		return serializeInYaml(w, apply.FromEntityList(entities, groupItems))
	default:
		return serializeMetadata(w, entities, outputOpt, *getMetaWide)
	}
}
