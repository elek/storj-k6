package k6

import (
	"context"
	"go.uber.org/zap"
	"storj.io/common/storj"
	"storj.io/common/testrand"
	"storj.io/common/uuid"
	"storj.io/storj/satellite/metabase"
	"time"
)

func MetainfoTest(connection string) func() error {
	ctx := context.Background()
	log, _ := zap.NewDevelopment()

	expire := time.Now().Add(time.Hour)

	metabaseDB, err := metabase.Open(ctx, log.Named("metabase"), connection, metabase.Config{
		ApplicationName:  "k6",
		MaxNumberOfParts: 10,
	})
	if err != nil {
		panic(err)
	}
	projectId := uuid.UUID{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xCA, 0xFE, 0xBA, 0xBE}

	return func() error {
		objectStream := metabase.ObjectStream{
			ProjectID:  projectId,
			BucketName: "bucket",
			ObjectKey:  metabase.ObjectKey("test/" + testrand.UUID().String()),
			Version:    metabase.NextVersion,
			StreamID:   testrand.UUID(),
		}

		_, err = metabaseDB.BeginObjectNextVersion(ctx, metabase.BeginObjectNextVersion{
			ObjectStream: objectStream,
			ExpiresAt:    &expire,
			Encryption: storj.EncryptionParameters{
				CipherSuite: storj.EncAESGCM,
				BlockSize:   256,
			},
		})

		return err

	}
}
