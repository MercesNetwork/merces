package keeper_test

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/json"
	"math/big"
	"time"

	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	rUsername = "Na_dsfF"
)

func (suite *IntegrationTestSuite) TestBookDNSRegistry() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	otherCerts, err := generateCAChain("other.com")
	suite.Require().Nil(err)
	mercesCerts, err := generateCAChain("mercestoken.com")
	suite.Require().Nil(err)

	x509.ParseCertificate(otherCerts.ca)

	certPool := x509.NewCertPool()
	xCa, err := x509.ParseCertificate(mercesCerts.ca)
	suite.Require().Nil(err)

	certPool.AddCert(xCa)

	data := keeper.DomainClaimingMessage{
		DomainClaiming: "mercestoken.com",
		PublicKey:      bob,
	}
	datab, _ := json.Marshal(data)
	digest := sha256.Sum256(datab)
	signature, err := rsa.SignPSS(rand.Reader, mercesCerts.certPrk, crypto.SHA256, digest[:], nil)
	if err != nil {
		suite.T().Fatal("cannot sign")
	}

	err = suite.app.PlanetKeeper.BookDNSRegistry(suite.ctx, certPool, mercesCerts.cert, data, signature)
	suite.Require().Nil(err)

	res, err := suite.app.PlanetKeeper.DNSRegistryAll(ctx, &types.QueryAllDNSRegistryRequest{})
	suite.Require().Nil(err)
	suite.Require().Len(res.DNSRegistry, 1)

	suite.Run("FailUnknownRoot", func() {
		certPool := x509.NewCertPool()
		xCa, err := x509.ParseCertificate(otherCerts.ca)
		suite.Require().Nil(err)
		certPool.AddCert(xCa)

		data := keeper.DomainClaimingMessage{
			DomainClaiming: "mercestoken.com",
			PublicKey:      bob,
		}
		datab, _ := json.Marshal(data)
		digest := sha256.Sum256(datab)
		signature, err := rsa.SignPSS(rand.Reader, mercesCerts.certPrk, crypto.SHA256, digest[:], nil)
		suite.Require().Nil(err)

		err = suite.app.PlanetKeeper.BookDNSRegistry(suite.ctx, certPool, mercesCerts.cert, data, signature)
		suite.Require().NotNil(err)
	})
	suite.Run("FailUnexpectedDomain", func() {
		certPool := x509.NewCertPool()
		xCa, err := x509.ParseCertificate(mercesCerts.ca)
		suite.Require().Nil(err)
		certPool.AddCert(xCa)

		data := keeper.DomainClaimingMessage{
			DomainClaiming: "app.mercestoken.com",
			PublicKey:      bob,
		}
		datab, _ := json.Marshal(data)
		digest := sha256.Sum256(datab)
		signature, err := rsa.SignPSS(rand.Reader, mercesCerts.certPrk, crypto.SHA256, digest[:], nil)
		suite.Require().Nil(err)

		err = suite.app.PlanetKeeper.BookDNSRegistry(suite.ctx, certPool, mercesCerts.cert, data, signature)
		suite.Require().NotNil(err)
	})
}

type ChainTest struct {
	ca      []byte
	caPk    *rsa.PrivateKey
	cert    []byte
	certPrk *rsa.PrivateKey
}

func generateCAChain(domain string) (ChainTest, error) {
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
		return ChainTest{}, err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return ChainTest{}, err
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 1),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		DNSNames:     []string{domain},
		Extensions: []pkix.Extension{pkix.Extension{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 17},
			Critical: false,
			Value:    []byte(`DNSName: mercestoken.com`),
		}},
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return ChainTest{}, err
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return ChainTest{}, err
	}

	return ChainTest{
		ca:      caBytes,
		caPk:    caPrivKey,
		cert:    certBytes,
		certPrk: certPrivKey,
	}, nil
}
