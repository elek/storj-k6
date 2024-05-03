package k6

import (
	"context"
	"github.com/pkg/errors"
	"storj.io/uplink"
)

func UplinkTest(accessGrant string) func() error {
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		panic(err)
	}

	project, err := uplink.OpenProject(context.Background(), access)
	if err != nil {
		panic(err)
	}
	return func() error {
		ctx := context.Background()
		info, err := project.UploadObject(ctx, "bucket1", "key1", nil)
		if err != nil {
			return errors.WithStack(err)
		}
		_, err = info.Write([]byte{})
		if err != nil {
			return errors.WithStack(err)
		}
		return info.Commit()
	}
}
