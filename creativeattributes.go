package openrtb

type CreativeAttribute struct {
	m uint
}

func AssembleCreativeAttribute(p ...int) *CreativeAttribute {
    c := &CreativeAttribute{}
	for _, v := range(p) {
	    if v < len(CreativeAttributes) && v > 0 {
	    	c.m = c.m & CreativeAttributes[v]
	    }
	}
	return c
}

func (c *CreativeAttribute) Check(cat *CreativeAttribute) bool {
	return ((c.m & cat.m) != 0)
}

var CreativeAttributes = []uint {
	0,     // 0  Empty attribute
	1,     // 1  Audio Ad (Auto-Play)                                   
	2,     // 2  Audio Ad (User Initiated)                              
	4,     // 3  Expandable (Automatic)                                 
	8,     // 4  Expandable (User Initiated - Click)                    
	16,    // 5  Expandable (User Initiated - Rollover)                 
	32,    // 6  In-Banner Video Ad (Auto-Play)                         
	64,    // 7  In-Banner Video Ad (User Initiated)                    
	128,   // 8  Pop (e.g., Over, Under, or Upon Exit)                  
	256,   // 9  Provocative or Suggestive Imagery                      
	512,   // 10 Shaky, Flashing, Flickering, Extreme Animation, Smileys
	1024,  // 11 Surveys                                                
	2048,  // 12 Text Only                                              
	4096,  // 13 User Interactive (e.g., Embedded Games)                
	8192,  // 14 Windows Dialog or Alert Style                          
	16384, // 15 Has Audio On/Off Button                                
	32768, // 16 Ad Can be Skipped (e.g., Skip Button on Pre-Roll Video)
}
