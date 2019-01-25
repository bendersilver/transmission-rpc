package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const (
	TorrentStart    = "torrent-start"
	TorrentStartNow = "torrent-start-now"
	TorrentVerify   = "torrent-verify"
)

type TorrentStartRequest struct {
	IDs []int `json:"ids,omitempty"`
}

// TorrentStart starts torrent with id.
func (t *Transmission) TorrentStart(ctx context.Context, id int) error {
	req := TorrentStartRequest{
		IDs: []int{id},
	}

	err := t.Do(ctx, TorrentStart, &req, nil)
	return errors.Wrap(err, "failed request")
}

// TorrentStartNow starts torrent with id.
func (t *Transmission) TorrentStartNow(ctx context.Context, id int) error {
	req := TorrentStartRequest{
		IDs: []int{id},
	}

	err := t.Do(ctx, TorrentStartNow, &req, nil)
	return errors.Wrap(err, "failed request")
}

// TorrentVerify starts torrent with id.
func (t *Transmission) TorrentVerify(ctx context.Context, id int) error {
	req := TorrentStartRequest{
		IDs: []int{id},
	}
	err := t.Do(ctx, TorrentVerify, &req, nil)
	return errors.Wrap(err, "failed request")
}
