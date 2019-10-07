package setupclient

import (
	"context"
	"fmt"

	"github.com/skycoin/dmsg/cipher"
	"github.com/skycoin/skycoin/src/util/logging"

	"github.com/skycoin/skywire/pkg/routing"
	"github.com/skycoin/skywire/pkg/snet"
)

func DialRouteGroup(ctx context.Context, log *logging.Logger, n *snet.Network, nodes []cipher.PubKey,
	req routing.BidirectionalRoute) (routing.EdgeRules, error) {

	client, err := NewClient(ctx, log, n, nodes)
	if err != nil {
		return routing.EdgeRules{}, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Warn(err)
		}
	}()

	resp, err := client.DialRouteGroup(ctx, req)
	if err != nil {
		return routing.EdgeRules{}, fmt.Errorf("route setup: %s", err)
	}

	return resp, nil
}
