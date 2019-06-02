package pixiv

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kanosaki/pixiv_oauth2"
)

// CacheTokenSource will cache token to specified path as json file, and retrieve from original token source when cache fails
// An error during saving token will return error, but an error during reading token will just print warning
// and proceed retrieving from original token source.
// Standard "log" module is used to warning, so please log.Output to modify output or disable warning
type CachedTokenSource struct {
	cachePath string
	origin    pixiv_oauth2.TokenSource
}

func (c *CachedTokenSource) Token() (*pixiv_oauth2.Token, error) {
	rf, err := os.Open(c.cachePath)
	if err == nil {
		ret := &pixiv_oauth2.Token{}
		dec := json.NewDecoder(rf)
		if err := dec.Decode(ret); err != nil {
			if ret.Valid() { // return only if token is valid
				rf.Close()
				return ret, nil
			}
		} else {
			log.Printf("failed to parse token from cache(%s): %v", c.cachePath, err)
		}
	} else {
		if !os.IsNotExist(err) { // suppress warning for just cache does not exist
			log.Printf("failed to read token from cache(%s): %v", c.cachePath, err)
		}
	}
	rf.Close()

	// retrieve new token
	tok, err := c.origin.Token()
	if err != nil {
		return nil, err
	}

	// save new token
	wf, err := os.Create(c.cachePath)
	if err != nil {
		return nil, err
	}
	defer wf.Close()
	if err := json.NewEncoder(wf).Encode(tok); err != nil {
		return nil, err
	}
	return tok, err
}
