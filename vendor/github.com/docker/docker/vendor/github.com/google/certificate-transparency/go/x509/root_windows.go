// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"errors"
	"syscall"
	"unsafe"
)

// Creates a new *syscall.CertContext representing the leaf certificate in an in-memory
// certificate store containing itself and all of the intermediate certificates specified
// in the opts.Intermediates CertPool.
//
// A pointer to the in-memory store is available in the returned CertContext's Store field.
// The store is automatically freed when the CertContext is freed using
// syscall.CertFreeCertificateContext.
func createStoreContext(leaf *Certificate, opts *VerifyOptions) (*syscall.CertContext, error) ***REMOVED***
	var storeCtx *syscall.CertContext

	leafCtx, err := syscall.CertCreateCertificateContext(syscall.X509_ASN_ENCODING|syscall.PKCS_7_ASN_ENCODING, &leaf.Raw[0], uint32(len(leaf.Raw)))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer syscall.CertFreeCertificateContext(leafCtx)

	handle, err := syscall.CertOpenStore(syscall.CERT_STORE_PROV_MEMORY, 0, 0, syscall.CERT_STORE_DEFER_CLOSE_UNTIL_LAST_FREE_FLAG, 0)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer syscall.CertCloseStore(handle, 0)

	err = syscall.CertAddCertificateContextToStore(handle, leafCtx, syscall.CERT_STORE_ADD_ALWAYS, &storeCtx)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	if opts.Intermediates != nil ***REMOVED***
		for _, intermediate := range opts.Intermediates.certs ***REMOVED***
			ctx, err := syscall.CertCreateCertificateContext(syscall.X509_ASN_ENCODING|syscall.PKCS_7_ASN_ENCODING, &intermediate.Raw[0], uint32(len(intermediate.Raw)))
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***

			err = syscall.CertAddCertificateContextToStore(handle, ctx, syscall.CERT_STORE_ADD_ALWAYS, nil)
			syscall.CertFreeCertificateContext(ctx)
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***
		***REMOVED***
	***REMOVED***

	return storeCtx, nil
***REMOVED***

// extractSimpleChain extracts the final certificate chain from a CertSimpleChain.
func extractSimpleChain(simpleChain **syscall.CertSimpleChain, count int) (chain []*Certificate, err error) ***REMOVED***
	if simpleChain == nil || count == 0 ***REMOVED***
		return nil, errors.New("x509: invalid simple chain")
	***REMOVED***

	simpleChains := (*[1 << 20]*syscall.CertSimpleChain)(unsafe.Pointer(simpleChain))[:]
	lastChain := simpleChains[count-1]
	elements := (*[1 << 20]*syscall.CertChainElement)(unsafe.Pointer(lastChain.Elements))[:]
	for i := 0; i < int(lastChain.NumElements); i++ ***REMOVED***
		// Copy the buf, since ParseCertificate does not create its own copy.
		cert := elements[i].CertContext
		encodedCert := (*[1 << 20]byte)(unsafe.Pointer(cert.EncodedCert))[:]
		buf := make([]byte, cert.Length)
		copy(buf, encodedCert[:])
		parsedCert, err := ParseCertificate(buf)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		chain = append(chain, parsedCert)
	***REMOVED***

	return chain, nil
***REMOVED***

// checkChainTrustStatus checks the trust status of the certificate chain, translating
// any errors it finds into Go errors in the process.
func checkChainTrustStatus(c *Certificate, chainCtx *syscall.CertChainContext) error ***REMOVED***
	if chainCtx.TrustStatus.ErrorStatus != syscall.CERT_TRUST_NO_ERROR ***REMOVED***
		status := chainCtx.TrustStatus.ErrorStatus
		switch status ***REMOVED***
		case syscall.CERT_TRUST_IS_NOT_TIME_VALID:
			return CertificateInvalidError***REMOVED***c, Expired***REMOVED***
		default:
			return UnknownAuthorityError***REMOVED***c, nil, nil***REMOVED***
		***REMOVED***
	***REMOVED***
	return nil
***REMOVED***

// checkChainSSLServerPolicy checks that the certificate chain in chainCtx is valid for
// use as a certificate chain for a SSL/TLS server.
func checkChainSSLServerPolicy(c *Certificate, chainCtx *syscall.CertChainContext, opts *VerifyOptions) error ***REMOVED***
	servernamep, err := syscall.UTF16PtrFromString(opts.DNSName)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	sslPara := &syscall.SSLExtraCertChainPolicyPara***REMOVED***
		AuthType:   syscall.AUTHTYPE_SERVER,
		ServerName: servernamep,
	***REMOVED***
	sslPara.Size = uint32(unsafe.Sizeof(*sslPara))

	para := &syscall.CertChainPolicyPara***REMOVED***
		ExtraPolicyPara: uintptr(unsafe.Pointer(sslPara)),
	***REMOVED***
	para.Size = uint32(unsafe.Sizeof(*para))

	status := syscall.CertChainPolicyStatus***REMOVED******REMOVED***
	err = syscall.CertVerifyCertificateChainPolicy(syscall.CERT_CHAIN_POLICY_SSL, chainCtx, para, &status)
	if err != nil ***REMOVED***
		return err
	***REMOVED***

	// TODO(mkrautz): use the lChainIndex and lElementIndex fields
	// of the CertChainPolicyStatus to provide proper context, instead
	// using c.
	if status.Error != 0 ***REMOVED***
		switch status.Error ***REMOVED***
		case syscall.CERT_E_EXPIRED:
			return CertificateInvalidError***REMOVED***c, Expired***REMOVED***
		case syscall.CERT_E_CN_NO_MATCH:
			return HostnameError***REMOVED***c, opts.DNSName***REMOVED***
		case syscall.CERT_E_UNTRUSTEDROOT:
			return UnknownAuthorityError***REMOVED***c, nil, nil***REMOVED***
		default:
			return UnknownAuthorityError***REMOVED***c, nil, nil***REMOVED***
		***REMOVED***
	***REMOVED***

	return nil
***REMOVED***

// systemVerify is like Verify, except that it uses CryptoAPI calls
// to build certificate chains and verify them.
func (c *Certificate) systemVerify(opts *VerifyOptions) (chains [][]*Certificate, err error) ***REMOVED***
	hasDNSName := opts != nil && len(opts.DNSName) > 0

	storeCtx, err := createStoreContext(c, opts)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer syscall.CertFreeCertificateContext(storeCtx)

	para := new(syscall.CertChainPara)
	para.Size = uint32(unsafe.Sizeof(*para))

	// If there's a DNSName set in opts, assume we're verifying
	// a certificate from a TLS server.
	if hasDNSName ***REMOVED***
		oids := []*byte***REMOVED***
			&syscall.OID_PKIX_KP_SERVER_AUTH[0],
			// Both IE and Chrome allow certificates with
			// Server Gated Crypto as well. Some certificates
			// in the wild require them.
			&syscall.OID_SERVER_GATED_CRYPTO[0],
			&syscall.OID_SGC_NETSCAPE[0],
		***REMOVED***
		para.RequestedUsage.Type = syscall.USAGE_MATCH_TYPE_OR
		para.RequestedUsage.Usage.Length = uint32(len(oids))
		para.RequestedUsage.Usage.UsageIdentifiers = &oids[0]
	***REMOVED*** else ***REMOVED***
		para.RequestedUsage.Type = syscall.USAGE_MATCH_TYPE_AND
		para.RequestedUsage.Usage.Length = 0
		para.RequestedUsage.Usage.UsageIdentifiers = nil
	***REMOVED***

	var verifyTime *syscall.Filetime
	if opts != nil && !opts.CurrentTime.IsZero() ***REMOVED***
		ft := syscall.NsecToFiletime(opts.CurrentTime.UnixNano())
		verifyTime = &ft
	***REMOVED***

	// CertGetCertificateChain will traverse Windows's root stores
	// in an attempt to build a verified certificate chain.  Once
	// it has found a verified chain, it stops. MSDN docs on
	// CERT_CHAIN_CONTEXT:
	//
	//   When a CERT_CHAIN_CONTEXT is built, the first simple chain
	//   begins with an end certificate and ends with a self-signed
	//   certificate. If that self-signed certificate is not a root
	//   or otherwise trusted certificate, an attempt is made to
	//   build a new chain. CTLs are used to create the new chain
	//   beginning with the self-signed certificate from the original
	//   chain as the end certificate of the new chain. This process
	//   continues building additional simple chains until the first
	//   self-signed certificate is a trusted certificate or until
	//   an additional simple chain cannot be built.
	//
	// The result is that we'll only get a single trusted chain to
	// return to our caller.
	var chainCtx *syscall.CertChainContext
	err = syscall.CertGetCertificateChain(syscall.Handle(0), storeCtx, verifyTime, storeCtx.Store, para, 0, 0, &chainCtx)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer syscall.CertFreeCertificateChain(chainCtx)

	err = checkChainTrustStatus(c, chainCtx)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	if hasDNSName ***REMOVED***
		err = checkChainSSLServerPolicy(c, chainCtx, opts)
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
	***REMOVED***

	chain, err := extractSimpleChain(chainCtx.Chains, int(chainCtx.ChainCount))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	chains = append(chains, chain)

	return chains, nil
***REMOVED***

func initSystemRoots() ***REMOVED***
***REMOVED***
