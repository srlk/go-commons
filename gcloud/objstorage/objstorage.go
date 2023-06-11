package objstorage

import (
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Storage struct {
	client        *storage.Client
	bucket        string
	predefinedAcl string
}

func NewStorage(ctx context.Context, credentialsOption option.ClientOption, bucket string, predefinedAcl string, opts ...option.ClientOption) (*Storage, error) {

	if credentialsOption == nil {
		return nil, errors.New("missing google credentials")
	}

	accessOpts := make([]option.ClientOption, 0)
	if len(opts) == 0 || opts == nil {
		accessOpts = append(accessOpts, credentialsOption)
	} else {
		accessOpts = append(opts, credentialsOption)
	}

	client, err := storage.NewClient(ctx, accessOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "create google client fail")
	}

	return &Storage{client: client, bucket: bucket, predefinedAcl: predefinedAcl}, nil
}

func (g *Storage) UploadWithRetries(ctx context.Context, name string, data []byte, retryCount int, waitDuration time.Duration) error {
	var err error
	for i := 0; i < retryCount; i++ {
		err = g.Upload(ctx, name, data)
		if err == nil {
			return nil
		}
		time.Sleep(waitDuration)
	}
	return err
}

func (g *Storage) Upload(ctx context.Context, name string, data []byte) (err error) {
	objectHandle := g.client.Bucket(g.bucket).Object(name)

	writer := objectHandle.NewWriter(ctx)
	defer func() {
		if writerErr := writer.Close(); err == nil && writerErr != nil {
			err = errors.Wrap(writerErr, "failed to close writer")
		}
	}()

	if g.predefinedAcl != "" {
		writer.PredefinedACL = g.predefinedAcl
	}

	if _, err := writer.Write(data); err != nil {
		return errors.Wrap(err, "failed to write data")
	}
	return nil
}

func (g *Storage) GetWriter(ctx context.Context, name, contentType string) io.WriteCloser {
	objectHandle := g.client.Bucket(g.bucket).Object(name)
	writer := objectHandle.NewWriter(ctx)
	if g.predefinedAcl != "" {
		writer.PredefinedACL = g.predefinedAcl
	}
	if contentType != "" {
		writer.ContentType = contentType
	}

	return writer
}

func (g *Storage) CopyWithRetries(ctx context.Context, source string, destination string, retryCount int, waitDuration time.Duration) error {
	var err error
	for i := 0; i < retryCount; i++ {
		err = g.Copy(ctx, source, destination)
		if err == nil {
			return nil
		}
		time.Sleep(waitDuration)
	}
	return err
}

func (g *Storage) Copy(ctx context.Context, source string, destination string) error {
	src := g.client.Bucket(g.bucket).Object(source)
	dst := g.client.Bucket(g.bucket).Object(destination)
	copier := dst.CopierFrom(src)
	if g.predefinedAcl != "" {
		copier.PredefinedACL = g.predefinedAcl
	}
	_, err := copier.Run(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to copy from %s to %s", source, destination)
	}
	return nil
}

func (g *Storage) Check(ctx context.Context, name string) (bool, error) {
	objectHandle := g.client.Bucket(g.bucket).Object(name)
	attr, err := objectHandle.Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, errors.Wrap(err, "failed to read data")
	}
	return attr != nil, nil
}

func (g *Storage) Download(ctx context.Context, name string) (data []byte, err error) {
	objectHandle := g.client.Bucket(g.bucket).Object(name)
	reader, err := objectHandle.NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read data")
	}
	defer func() {
		if readerErr := reader.Close(); err == nil && readerErr != nil {
			err = errors.Wrap(readerErr, "reader error")
		}
	}()
	if data, err = io.ReadAll(reader); err != nil {
		return nil, errors.Wrap(err, "failed to read stream")
	}
	return data, nil
}

func (g *Storage) Delete(ctx context.Context, bucket string, name string) error {
	objectHandle := g.client.Bucket(bucket).Object(name)
	if err := objectHandle.Delete(ctx); err != nil {
		return errors.Wrap(err, "failed to delete data")
	}
	return nil
}
