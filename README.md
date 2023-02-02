# x509 PrintableString HACK

# To Use

`import _ "github.com/fredwangwang/go-x509hack"`

# Rathional

Go's certificate parsing follows strict asn1 standard, where `_` is not considered a valid character set in [PrintableString](https://en.wikipedia.org/wiki/PrintableString).
Its rule can be seen in the [source code](https://github.com/golang/go/blob/88a36c9e9a511ec6ad218633bce1e82f25e54d35/src/crypto/x509/parser.go#L59))

However, there are cases of invalid use of PrintableString to include characters that is not part of asn1 standard. This seems especially true with certificate generated from `.netframework` projects, as I encountered recently. This causes critical issues in integrating with existing infrastructures, where fixing those certificate is just cost prohibitive.

Nontheless, other languages (Java/C#, etc) choose to _play nicely_ with this, and standard tools
like OpenSSL has no issues to use the certificate neither.

Issues \([1](https://github.com/golang/go/issues/14017) [2](https://github.com/golang/go/issues/36044)\) has been raised to Golang in the past without resolution. So not optmistic about it being addressed in the upstream timely in spite of a real product issue.

Although I do like when implementation follows the standard closely, as it avoids some pitfalls in unexpected bugs and even security issues. This is one of the cases I think "be liberal in what you accept" makes sense. As accepting _non-standard_ chars is de facto the standard.

So here we are. A hack to bypass parsing checks in asn1 strings.

# NOTE

USE YOUR OWN JUDGEMENT in using this library. This library removes all checks in parsing asn1 string. This might or might not work for your use case. Similar patching strategy can be used to override certain checks without giving up others.
