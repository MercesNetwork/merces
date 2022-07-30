package keeper

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var rootCertificates *x509.CertPool

type DomainClaimingMessage struct {
	DomainClaiming string
	PublicKey      string
}

func (k Keeper) BookDNSRegistry(ctx sdk.Context, certPool *x509.CertPool, certificate []byte, message DomainClaimingMessage, signature []byte) error {
	cert, err := x509.ParseCertificate(certificate)
	if err != nil {
		return err
	}
	if _, err := cert.Verify(x509.VerifyOptions{
		DNSName:     message.DomainClaiming,
		CurrentTime: time.Now(),
		Roots:       certPool,
	}); err != nil {
		return fmt.Errorf("domain certificate %v", err)
	}

	// certificate valid and containing claiming

	// b is digest
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	digest := sha256.Sum256(b)

	certPB := cert.PublicKey.(*rsa.PublicKey)
	err = rsa.VerifyPSS(certPB, crypto.SHA256, digest[:], signature, nil)
	if err != nil {
		return fmt.Errorf("signature %v", err)
	}

	// Private key of certificate has signed claiming + public key

	k.SetDNSRegistry(ctx, types.DNSRegistry{
		Domain:    message.DomainClaiming,
		PublicKey: message.PublicKey,
	})
	return nil
}
