/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// CompanyProfile2 struct for CompanyProfile2
type CompanyProfile2 struct {
	// Country of company's headquarter.
	Country string `json:"country,omitempty"`
	// Currency used in company filings.
	Currency string `json:"currency,omitempty"`
	// Listed exchange.
	Exchange string `json:"exchange,omitempty"`
	// Company name.
	Name string `json:"name,omitempty"`
	// Company symbol/ticker as used on the listed exchange.
	Ticker string `json:"ticker,omitempty"`
	// IPO date.
	Ipo string `json:"ipo,omitempty"`
	// Market Capitalization.
	MarketCapitalization float32 `json:"marketCapitalization,omitempty"`
	// Number of oustanding shares.
	ShareOutstanding float32 `json:"shareOutstanding,omitempty"`
	// Logo image.
	Logo string `json:"logo,omitempty"`
	// Company phone number.
	Phone string `json:"phone,omitempty"`
	// Company website.
	Weburl string `json:"weburl,omitempty"`
	// Finnhub industry classification.
	FinnhubIndustry string `json:"finnhubIndustry,omitempty"`
}
