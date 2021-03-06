package openrtb

// The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type NativeVideo struct {
	Mimes          []string   `json:"mimes"`                    // Content MIME types supported.
	Minduration    *int       `json:"minduration"`              // Minimum video ad duration in seconds
	Maxduration    *int       `json:"maxduration"`              // Maximum video ad duration in seconds
	Protocol       *int       `json:"protocol,omitempty"`       // Video bid response protocols
	Protocols      []int      `json:"protocols"`                // Video bid response protocols
	W              *int       `json:"w,omitempty"`              // Width of the player in pixels
	H              *int       `json:"h,omitempty"`              // Height of the player in pixels
	Startdelay     *int       `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Linearity      *int       `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Sequence       *int       `json:"sequence,omitempty"`       // Default: 1
	Battr          []int      `json:"battr,omitempty"`          // Blocked creative attributes
	Maxextended    *int       `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	Minbitrate     *int       `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	Maxbitrate     *int       `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	Boxingallowed  *int       `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	Playbackmethod []int      `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery       []int      `json:"delivery,omitempty"`       // List of supported delivery methods
	Pos            *int       `json:"pos,omitempty"`            // Ad Position
	Companionad    []Banner   `json:"companionad,omitempty"`
	Api            []int      `json:"api,omitempty"` // List of supported API frameworks
	Companiontype  []int      `json:"companiontype,omitempty"`
	Ext            Extensions `json:"ext,omitempty"`
}

// Returns the sequence number, with default fallback
func (v *NativeVideo) Seq() int {
	if v.Sequence != nil {
		return *v.Sequence
	}
	return 1
}

// Returns the boxing permission status, with default fallback
func (v *NativeVideo) IsBoxingAllowed() bool {
	if v.Boxingallowed != nil {
		return *v.Boxingallowed == 1
	}
	return true
}

// Returns the position, with default fallback
func (v *NativeVideo) Position() int {
	if v.Pos != nil {
		return *v.Pos
	}
	return AD_POS_UNKNOWN
}

// Validates the object
func (v *NativeVideo) Valid() (bool, error) {
	if len(v.Mimes) == 0 {
		return false, ErrInvalidVideoMimes
	} else if v.Linearity == nil {
		return false, ErrInvalidVideoLinearity
	} else if v.Minduration == nil {
		return false, ErrInvalidVideoMinduration
	} else if v.Maxduration == nil {
		return false, ErrInvalidVideoMaxduration
	} else if v.Protocol == nil {
		return false, ErrInvalidVideoProtocol
	}
	return true, nil
}

// Applies defaults
func (v *NativeVideo) WithDefaults() *NativeVideo {
	if v.Sequence == nil {
		v.Sequence = new(int)
		*v.Sequence = 1
	}
	if v.Boxingallowed == nil {
		v.Boxingallowed = new(int)
		*v.Boxingallowed = 1
	}
	if v.Pos == nil {
		v.Pos = new(int)
		*v.Pos = AD_POS_UNKNOWN
	}
	return v
}

// Set the Linearity
func (v *NativeVideo) SetLinearity(lin int) *NativeVideo {
	if v.Linearity == nil {
		v.Linearity = new(int)
	}
	*v.Linearity = lin
	return v
}

// Set the Minduration
func (v *NativeVideo) SetMinduration(mdur int) *NativeVideo {
	if v.Minduration == nil {
		v.Minduration = new(int)
	}
	*v.Minduration = mdur
	return v
}

// Set the Maxduration
func (v *NativeVideo) SetMaxduration(mdur int) *NativeVideo {
	if v.Maxduration == nil {
		v.Maxduration = new(int)
	}
	*v.Maxduration = mdur
	return v
}

// Set the Protocol
func (v *NativeVideo) SetProtocol(p int) *NativeVideo {
	if v.Protocol == nil {
		v.Protocol = new(int)
	}
	*v.Protocol = p
	return v
}
