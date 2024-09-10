package app

import "strings"

// Scan represents the results of a scan.
type Scan struct {
	Settings      ScanSettings `json:"scanSettings"`
	ScanTarget    string       `json:"scanTarget"`
	CommandLine   string       `json:"commandLine"`
	DurationSecs  float64      `json:"durationSecs"`
	Error         string       `json:"error"`
	ExitCode      int          `json:"exitCode"`
	Ok            bool         `json:"ok"`
	NumCritical   int          `json:"numCritical"`
	NumHigh       int          `json:"numHigh"`
	NumMedium     int          `json:"numMedium"`
	NumLow        int          `json:"numLow"`
	NumNegligible int          `json:"numNegligible"`
	NumUnknown    int          `json:"numUnknown"`
	NumIgnored    int          `json:"numIgnored"`
	S3URL         string       `json:"s3URL"`
	stdout        []byte
}

// Failed returns true if the scan failed severity threshold.
func (s *Scan) Failed() bool {
	switch strings.ToLower(s.Settings.severity) {
	case "critical":
		return s.NumCritical > 0
	case "high":
		return s.NumHigh+s.NumCritical > 0
	case "medium":
		return s.NumMedium+s.NumHigh+s.NumCritical > 0
	case "low":
		return s.NumLow+s.NumMedium+s.NumHigh+s.NumCritical > 0
	default:
		return false
	}
}

// Score scores the scan based on the severity of the vulnerability.
func (s *Scan) Score(severity string) {
	switch strings.ToLower(severity) {
	case "critical":
		s.NumCritical++
	case "high":
		s.NumHigh++
	case "medium":
		s.NumMedium++
	case "low":
		s.NumLow++
	case "negligible":
		s.NumNegligible++
	default:
		s.NumUnknown++
	}
}

// NumTotal returns the total number of vulnerabilities.
func (s *Scan) NumTotal() int {
	return s.NumCritical + s.NumHigh + s.NumMedium + s.NumLow + s.NumNegligible + s.NumUnknown
}
