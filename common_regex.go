package ling

import (
	"regexp"
)

// https://github.com/mingrammer/commonregex
// Regular expression patterns
const (
	DatePattern = `(?i)(?:(?:tgl)?\d{1,2}[^0-9^:]\d{1,2}[^0-9^:](?:19|20)?\d{2})|(?:(?:19|20)?\d{2}[^0-9^:]\d{1,2}[^0-9^:]\d{1,2})`
	TimePattern = `(?:(?:0?|[12])\d\s*:+\s*[0-5]\d(?:\s*:+\s*[0-5]\d(?:\.\d+)?(?:\s*(?:\+|-)(?:0?\d|1[0-2]):?(?:0|3)0)?)?)`
	//TimePattern = `(?is)((?:0?|[12])\d\s*:+\s*[0-5]\d(?:\s*:+\s*[0-5]\d)?(?:\s*[,:.]*\s*(?:am|pm))?|(?:0?|[12])\d\s*[.\s]+\s*[0-5]\d(?:\s*[,:.]*\s*(?:am|pm))+)`
	PhonePattern          = `(?:(?:\+?\d{1,3}[-.\s*]?)?(?:\(?\d{3}\)?[-.\s*]?)?\d{3}[-.\s*]?\d{4,6})|(?:(?:(?:\(\+?\d{2}\))|(?:\+?\d{2}))\s*\d{2}\s*\d{3}\s*\d{4})`
	PhonesWithExtsPattern = `(?i)(?:(?:\+?1\s*(?:[.-]\s*)?)?(?:\(\s*(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9])\s*\)|(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9]))\s*(?:[.-]\s*)?)?(?:[2-9]1[02-9]|[2-9][02-9]1|[2-9][02-9]{2})\s*(?:[.-]\s*)?(?:[0-9]{4})(?:\s*(?:#|x\.?|ext\.?|extension)\s*(?:\d+)?)`
	LinkPattern           = `(?i)(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w\.-]*)*\/?`
	EmailPattern          = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	IPv4Pattern           = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
	CreditCardPattern     = `(?i)(?:(?:(?:[\d\*x]{4}[- ]?){3}[\d\*x]{4}|[\d\*x]{15,16}))`
	VISACreditCardPattern = `4\d{3}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	MCCreditCardPattern   = `5[1-5]\d{2}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	BtcAddressPattern     = `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`
	SSNPattern            = `(?:\d{3}-\d{2}-\d{4})`
	MD5HexPattern         = `[0-9a-fA-F]{32}`
	SHA1HexPattern        = `[0-9a-fA-F]{40}`
	SHA256HexPattern      = `[0-9a-fA-F]{64}`
	GUIDPattern           = `[0-9a-fA-F]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}`
	ISBN13Pattern         = `(?:[\d]-?){12}[\dxX]`
	ISBN10Pattern         = `(?:[\d]-?){9}[\dxX]`
	MACAddressPattern     = `(([a-fA-F0-9]{2}[:-]){5}([a-fA-F0-9]{2}))`
	IBANPattern           = `[A-Z]{2}\d{2}[A-Z0-9]{4}\d{7}([A-Z\d]?){0,16}`
	NumericPattern        = `([+\-]?((\d{1,3}(,\d{3})+))|((?:0|[1-9]\d*)(?:\.\d*)?(?:[eE][+\-]?\d+)?))`
	DigitsPattern         = `\d+`
)

// Regexes is the compiled regular expressions
var Regexes = map[string]*regexp.Regexp{
	"date":             regexp.MustCompile(DatePattern),
	"time":             regexp.MustCompile(TimePattern),
	"phone":            regexp.MustCompile(PhonePattern),
	"phones_with_exts": regexp.MustCompile(PhonesWithExtsPattern),
	"link":             regexp.MustCompile(LinkPattern),
	"email":            regexp.MustCompile(EmailPattern),
	"ipv4":             regexp.MustCompile(IPv4Pattern),
	"credit_card":      regexp.MustCompile(CreditCardPattern),
	"btc_address":      regexp.MustCompile(BtcAddressPattern),
	"ssn":              regexp.MustCompile(SSNPattern),
	"md5_hex":          regexp.MustCompile(MD5HexPattern),
	"sha1_hex":         regexp.MustCompile(SHA1HexPattern),
	"sha256_hex":       regexp.MustCompile(SHA256HexPattern),
	"guid":             regexp.MustCompile(GUIDPattern),
	"isbn13":           regexp.MustCompile(ISBN13Pattern),
	"isbn10":           regexp.MustCompile(ISBN10Pattern),
	"visa_credit_card": regexp.MustCompile(VISACreditCardPattern),
	"mc_credit_card":   regexp.MustCompile(MCCreditCardPattern),
	"mac_address":      regexp.MustCompile(MACAddressPattern),
	"iban":             regexp.MustCompile(IBANPattern),
	"numeric":          regexp.MustCompile(NumericPattern),
	"digits":           regexp.MustCompile(DigitsPattern),
}
