package fetcher

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/jaehoonkim/sentinel/pkg/agent/log"
	"github.com/jaehoonkim/sentinel/pkg/manager/model/auths/v2"
	sessionv1 "github.com/jaehoonkim/sentinel/pkg/manager/model/session/v1"
	"github.com/jaehoonkim/sentinel/pkg/version"
)

func (f *Fetcher) HandShake() error {
	body := &auths.HttpReqAuth{AuthProperty: auths.AuthProperty{
		ClusterUuid:   f.clusterId,
		Assertion:     f.bearerToken,
		ClientVersion: version.Version,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := f.sentinelAPI.Auth(ctx, body); err != nil {
		return err
	}
	sessionToken := f.sentinelAPI.GetToken()
	log.Debugf("Successed to handshake: received token(%s) for polling.", sessionToken)

	f.ChangeAgentConfigFromToken()

	// save session_uuid from token
	claims := new(sessionv1.ClientSessionPayload)
	jwt_token, _, err := jwt.NewParser().ParseUnverified(sessionToken, claims)
	if _, ok := jwt_token.Claims.(*sessionv1.ClientSessionPayload); !ok || err != nil {
		log.Warnf("Failed to bind payload : %v\n", err)
		return err
	}
	if err := writeFile(".sentinel", []byte(claims.Uuid)); err != nil {
		return err
	}

	return nil
}

func (f *Fetcher) RetryHandshake() {
	maxRetryCnt := 5

	ticker := time.NewTicker(time.Second * time.Duration(f.pollingInterval))
	defer ticker.Stop()

	for retry := 0; ; <-ticker.C {
		log.Debugf("retry handshake : count(%d)\n", retry+1)
		if err := f.HandShake(); err != nil {
			log.Warnf("Failed to Handshake Retry : count(%d), error(%v)\n", retry, err)
		} else {
			f.ticker.Reset(time.Second * time.Duration(f.pollingInterval))
			return
		}
		retry++

		if maxRetryCnt <= retry {
			f.Cancel()
			return
		}
	}
}

func writeFile(filename string, data []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
