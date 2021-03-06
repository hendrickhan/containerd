package boltdb

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/docker/containerd/snapshot/storage"
	"github.com/docker/containerd/snapshot/storage/testsuite"

	// Does not require root but flag must be defined for snapshot tests
	_ "github.com/docker/containerd/testutil"
)

func TestBoltDB(t *testing.T) {
	testsuite.MetaStoreSuite(t, "BoltDB", func(ctx context.Context, root string) (storage.MetaStore, error) {
		return NewMetaStore(ctx, filepath.Join(root, "metadata.db"))
	})
}

func BenchmarkSuite(b *testing.B) {
	testsuite.Benchmarks(b, "BoltDBBench", func(ctx context.Context, root string) (storage.MetaStore, error) {
		return NewMetaStore(ctx, filepath.Join(root, "metadata.db"))
	})
}
