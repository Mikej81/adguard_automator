package main

// DNSRewrite represents a DNS rewrite rule
type DNSRewrite struct {
	Domain string `json:"domain"`
	Answer string `json:"answer"`
}
