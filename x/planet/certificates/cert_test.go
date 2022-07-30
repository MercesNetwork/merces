package certificates

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidDomain(t *testing.T) {
	ca := &x509.Certificate{
		SerialNumber:          big.NewInt(2022),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 1),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("%v", err)
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		t.Fatalf("%v", err)
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 1),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		DNSNames:     []string{"mercestoken.com"},
		Extensions: []pkix.Extension{pkix.Extension{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 17},
			Critical: false,
			Value:    []byte(`DNSName: mercestoken.com`),
		}},
	}
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("%v", err)
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		t.Fatalf("%v", err)
	}

	xCa, _ := x509.ParseCertificate(caBytes)
	xCert, _ := x509.ParseCertificate(certBytes)

	t.Run("VerifyDomainAndCert", func(t *testing.T) {
		certPool := x509.NewCertPool()
		certPool.AddCert(xCa)
		if _, err := xCert.Verify(x509.VerifyOptions{
			DNSName:     "mercestoken.com",
			CurrentTime: time.Now(),
			Roots:       certPool,
		}); err != nil {
			t.Fatalf("failed to validate cert %v", err)
		}

		if _, err := xCert.Verify(x509.VerifyOptions{
			DNSName:     "app.mercestoken.com",
			CurrentTime: time.Now(),
			Roots:       certPool,
		}); err == nil {
			t.Fatalf("failed to validate cert app.mercestoken.com should fail")
		}

		if _, err := xCert.Verify(x509.VerifyOptions{
			DNSName:     "app.com",
			CurrentTime: time.Now(),
			Roots:       certPool,
		}); err == nil {
			t.Fatalf("failed to validate cert app.com should fail")
		}
	})

	t.Run("VerifySignature", func(t *testing.T) {
		data := []byte("My Message!")

		digest := sha256.Sum256(data)
		signature, err := rsa.SignPSS(rand.Reader, certPrivKey, crypto.SHA256, digest[:], nil)
		if err != nil {
			t.Fatalf("cannot sign")
		}

		pb := xCert.PublicKey.(*rsa.PublicKey)
		err = rsa.VerifyPSS(pb, crypto.SHA256, digest[:], signature, nil)
		require.Nil(t, err)

		err = rsa.VerifyPSS(pb, crypto.SHA256, []byte("My Bad Message"), signature, nil)
		require.NotNil(t, err)
	})
}
