package certificates

// certificate as byte
// message (public key) signed with certificate private key
//func GetDomainPublicKey(publicKey []byte, message []byte) {
//	block, _ := pem.Decode(publicKey)
//	var cert *x509.Certificate
//	cert, _ = x509.ParseCertificate(block.Bytes)
//	cert.Verify()
//	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
//
//
//	// rsa.VerifyPSS()
//}
//
//type DomainPublicKey struct {
//	Domain    string
//	PublicKey string
//}
//
//type Key struct {
//	publicKey  *rsa.PublicKey
//	privateKey *rsa.PrivateKey
//}
//
//func New() (Key, error) {
//	k, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		return Key{}, err
//	}
//	return Key{
//		publicKey:  &k.PublicKey,
//		privateKey: k,
//	}, nil
//}
//
//func Ti(k Key) {
//	k.privateKey.Sign()
//}
