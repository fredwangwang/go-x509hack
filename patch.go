package x509hack

import (
	"crypto/x509"
	"unsafe"

	"github.com/agiledragon/gomonkey/v2"
)

// HACK
// this package patches the x509 private function parseASN1String
// to allow characters like `_` that used in the wild but not supported by asn1 standard.

var _ = x509.ExtKeyUsageAny
var _ = unsafe.Pointer(nil)

//go:noinline
//go:linkname x509parseASN1String crypto/x509.parseASN1String
func x509parseASN1String(tag uint8, value []byte) (string, error)

func init() {
	gomonkey.ApplyFunc(x509parseASN1String, func(tag uint8, value []byte) (string, error) {
		return string(value), nil
	})

	// log.Println(x509parseASN1String(cryptobyte_asn1.PrintableString, []byte("!@#$%^&*()hello_world™👍")))
}

// issue:
// https://github.com/golang/go/issues/36044
// https://github.com/golang/go/issues/14017
//
// impl context:
// linkname: https://www.pixelstech.net/article/1649596852-The-magic-of-go%3Alinkname, https://sitano.github.io/2016/04/28/golang-private/
// noinline: https://stackoverflow.com/questions/68280753/forbid-inlining-in-golang
// patching: https://bou.ke/blog/monkey-patching-in-go/
