// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors


// Supported partners with fURL integrations.
type Partner string

const (
	Partner_GITHUB Partner = "GITHUB"
	Partner_STRIPE Partner = "STRIPE"
	Partner_TWILIO Partner = "TWILIO"
)

