package tagger

import (
	"github.com/liuzl/xurls"
	//"mvdan.cc/xurls"
	"regexp"
)

// https://github.com/mingrammer/commonregex

// Regular expression patterns
const (
	DatePattern           = `(?i)(?:[0-3]?\d(?:st|nd|rd|th)?\s+(?:of\s+)?(?:jan\.?|january|feb\.?|february|mar\.?|march|apr\.?|april|may|jun\.?|june|jul\.?|july|aug\.?|august|sep\.?|september|oct\.?|october|nov\.?|november|dec\.?|december)|(?:jan\.?|january|feb\.?|february|mar\.?|march|apr\.?|april|may|jun\.?|june|jul\.?|july|aug\.?|august|sep\.?|september|oct\.?|october|nov\.?|november|dec\.?|december)\s+[0-3]?\d(?:st|nd|rd|th)?)(?:\,)?\s*(?:\d{4})?|[0-3]?\d[-\./][0-3]?\d[-\./]\d{2,4}`
	TimePattern           = `(?i)\d{1,2}:\d{2} ?(?:[ap]\.?m\.?)?|\d[ap]\.?m\.?`
	PhonePattern          = `(?:(?:\+?\d{1,3}[-.\s*]?)?(?:\(?\d{3}\)?[-.\s*]?)?\d{3}[-.\s*]?\d{4,6})|(?:(?:(?:\(\+?\d{2}\))|(?:\+?\d{2}))\s*\d{2}\s*\d{3}\s*\d{4})`
	PhonesWithExtsPattern = `(?i)(?:(?:\+?1\s*(?:[.-]\s*)?)?(?:\(\s*(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9])\s*\)|(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9]))\s*(?:[.-]\s*)?)?(?:[2-9]1[02-9]|[2-9][02-9]1|[2-9][02-9]{2})\s*(?:[.-]\s*)?(?:[0-9]{4})(?:\s*(?:#|x\.?|ext\.?|extension)\s*(?:\d+)?)`
	EmailPattern          = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	IPv4Pattern           = `(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`
	IPv6Pattern           = `(?:(?:(?:[0-9A-Fa-f]{1,4}:){7}(?:[0-9A-Fa-f]{1,4}|:))|(?:(?:[0-9A-Fa-f]{1,4}:){6}(?::[0-9A-Fa-f]{1,4}|(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){5}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,2})|:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){4}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,3})|(?:(?::[0-9A-Fa-f]{1,4})?:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){3}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,4})|(?:(?::[0-9A-Fa-f]{1,4}){0,2}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){2}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,5})|(?:(?::[0-9A-Fa-f]{1,4}){0,3}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){1}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,6})|(?:(?::[0-9A-Fa-f]{1,4}){0,4}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?::(?:(?:(?::[0-9A-Fa-f]{1,4}){1,7})|(?:(?::[0-9A-Fa-f]{1,4}){0,5}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(?:%.+)?\s*`
	IPPattern             = IPv4Pattern + `|` + IPv6Pattern
	PricePattern          = `[$]\s?[+-]?[0-9]{1,3}(?:(?:,?[0-9]{3}))*(?:\.[0-9]{1,2})?`
	HexColorPattern       = `(?:#([0-9a-fA-F]{6}|[0-9a-fA-F]{3}))`
	CreditCardPattern     = `(?:(?:(?:\d{4}[- ]?){3}\d{4}|\d{15,16}))`
	VISACreditCardPattern = `4\d{3}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	MCCreditCardPattern   = `5[1-5]\d{2}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	BtcAddressPattern     = `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`
	StreetAddressPattern  = `\d{1,4} [\w\s]{1,20}(?:street|st|avenue|ave|road|rd|highway|hwy|square|sq|trail|trl|drive|dr|court|ct|park|parkway|pkwy|circle|cir|boulevard|blvd)\W?`
	ZipCodePattern        = `\b\d{5}(?:[-\s]\d{4})?\b`
	PoBoxPattern          = `(?i)P\.? ?O\.? Box \d+`
	SSNPattern            = `(?:\d{3}-\d{2}-\d{4})`
	MD5HexPattern         = `[0-9a-fA-F]{32}`
	SHA1HexPattern        = `[0-9a-fA-F]{40}`
	SHA256HexPattern      = `[0-9a-fA-F]{64}`
	GUIDPattern           = `[0-9a-fA-F]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}`
	ISBN13Pattern         = `(?:[\d]-?){12}[\dxX]`
	ISBN10Pattern         = `(?:[\d]-?){9}[\dxX]`
	MACAddressPattern     = `(([a-fA-F0-9]{2}[:-]){5}([a-fA-F0-9]{2}))`
	IBANPattern           = `[A-Z]{2}\d{2}[A-Z0-9]{4}\d{7}([A-Z\d]?){0,16}`
	GitRepoPattern        = `((git|ssh|http(s)?)|(git@[\w\.]+))(:(\/\/)?)([\w\.@\:/\-~]+)(\.git)(\/)?`
)

// Compiled regular expressions
var (
	DateRegex           = regexp.MustCompile(DatePattern)
	TimeRegex           = regexp.MustCompile(TimePattern)
	PhoneRegex          = regexp.MustCompile(PhonePattern)
	PhonesWithExtsRegex = regexp.MustCompile(PhonesWithExtsPattern)
	UrlRegex            = xurls.Relaxed()
	EmailRegex          = regexp.MustCompile(EmailPattern)
	IPv4Regex           = regexp.MustCompile(IPv4Pattern)
	IPv6Regex           = regexp.MustCompile(IPv6Pattern)
	IPRegex             = regexp.MustCompile(IPPattern)
	PriceRegex          = regexp.MustCompile(PricePattern)
	HexColorRegex       = regexp.MustCompile(HexColorPattern)
	CreditCardRegex     = regexp.MustCompile(CreditCardPattern)
	BtcAddressRegex     = regexp.MustCompile(BtcAddressPattern)
	StreetAddressRegex  = regexp.MustCompile(StreetAddressPattern)
	ZipCodeRegex        = regexp.MustCompile(ZipCodePattern)
	PoBoxRegex          = regexp.MustCompile(PoBoxPattern)
	SSNRegex            = regexp.MustCompile(SSNPattern)
	MD5HexRegex         = regexp.MustCompile(MD5HexPattern)
	SHA1HexRegex        = regexp.MustCompile(SHA1HexPattern)
	SHA256HexRegex      = regexp.MustCompile(SHA256HexPattern)
	GUIDRegex           = regexp.MustCompile(GUIDPattern)
	ISBN13Regex         = regexp.MustCompile(ISBN13Pattern)
	ISBN10Regex         = regexp.MustCompile(ISBN10Pattern)
	VISACreditCardRegex = regexp.MustCompile(VISACreditCardPattern)
	MCCreditCardRegex   = regexp.MustCompile(MCCreditCardPattern)
	MACAddressRegex     = regexp.MustCompile(MACAddressPattern)
	IBANRegex           = regexp.MustCompile(IBANPattern)
	GitRepoRegex        = regexp.MustCompile(GitRepoPattern)
)

var regexes = map[string]*regexp.Regexp{
	"date":             DateRegex,
	"time":             TimeRegex,
	"phone":            PhoneRegex,
	"phones_with_exts": PhonesWithExtsRegex,
	"url":              UrlRegex,
	"email":            EmailRegex,
	"ipv4":             IPv4Regex,
	"ipv6":             IPv6Regex,
	"ip":               IPRegex,
	"price":            PriceRegex,
	"hex_color":        HexColorRegex,
	"credit_card":      CreditCardRegex,
	"btc_address":      BtcAddressRegex,
	"street_address":   StreetAddressRegex,
	"zip_code":         ZipCodeRegex,
	"po_box":           PoBoxRegex,
	"ssn":              SSNRegex,
	"md5_hex":          MD5HexRegex,
	"sha1_hex":         SHA1HexRegex,
	"sha256_hex":       SHA256HexRegex,
	"guid":             GUIDRegex,
	"isbn13":           ISBN13Regex,
	"isbn10":           ISBN10Regex,
	"visa_credit_card": VISACreditCardRegex,
	"mc_credit_card":   MCCreditCardRegex,
	"mac_address":      MACAddressRegex,
	"iban":             IBANRegex,
	"git_repo":         GitRepoRegex,
}
