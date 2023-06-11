package objstorage_test

import (
	"context"
	"testing"

	"github.com/srlk/go-commons/gcloud/objstorage"
	"github.com/stretchr/testify/require"
)

func Test_NewStorage(t *testing.T) {
	_, err := objstorage.NewStorage(context.Background(), nil, "buckyou", "pre")
	require.Error(t, err)
}
