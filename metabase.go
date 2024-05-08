package k6

import (
	"context"
	"fmt"
	"storj.io/common/grant"
	"storj.io/common/identity"
	"storj.io/common/pb"
	"storj.io/common/peertls/tlsopts"
	"storj.io/common/rpc"
	"storj.io/common/storj"
	"storj.io/storj/cmd/uplink/ulloc"
	"time"

	"math/rand"
)

func NewBeginObject(uplinkAccess string, key string) *BeginObject {

	accessGrant, err := grant.ParseAccess(uplinkAccess)
	if err != nil {
		panic(err)
	}

	return &BeginObject{
		Key:    key,
		access: accessGrant,
	}
}

type BeginObject struct {
	Key    string `help:"key in the form of sj://bucket/encryptedpath" arg:""`
	dialer rpc.Dialer
	access *grant.Access
	client pb.DRPCMetainfoClient
	bucket string
	key    string
	conn   *rpc.Conn
}

func (b *BeginObject) Init() error {
	var err error
	var ident *identity.FullIdentity

	ident, err = identity.NewFullIdentity(context.Background(), identity.NewCAOptions{
		Difficulty:  0,
		Concurrency: 1,
	})

	tlsConfig := tlsopts.Config{
		UsePeerCAWhitelist: false,
		PeerIDVersions:     "0",
	}

	tlsOptions, err := tlsopts.NewOptions(ident, tlsConfig, nil)
	if err != nil {
		panic(err)
	}

	b.dialer = rpc.NewDefaultDialer(tlsOptions)

	satelliteURL, err := storj.ParseNodeURL(b.access.SatelliteAddress)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	b.conn, err = b.dialer.DialNode(ctx, satelliteURL, rpc.DialOptions{})
	if err != nil {
		panic(err)
	}
	b.client = pb.NewDRPCMetainfoClient(b.conn)

	loc, err := ulloc.Parse(b.Key)
	if err != nil {
		panic(err)
	}

	var ok bool
	b.bucket, b.key, ok = loc.RemoteParts()
	if !ok {
		panic("Wrong url " + b.Key)
	}
	return nil

}

func (b *BeginObject) Run() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := b.client.BeginObject(ctx, &pb.BeginObjectRequest{
		Header: &pb.RequestHeader{
			ApiKey: []byte(b.access.APIKey.SerializeRaw()),
		},
		Bucket:             []byte(b.bucket),
		EncryptedObjectKey: []byte(fmt.Sprintf("%s-%d-%d", b.key, rand.Int(), rand.Int())),
		EncryptionParameters: &pb.EncryptionParameters{
			CipherSuite: pb.CipherSuite_ENC_AESGCM,
			BlockSize:   256,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *BeginObject) Close() {
	_ = b.conn.Close()
}
