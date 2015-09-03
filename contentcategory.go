package openrtb

import (
	"math/big"
)

var bigInt_0 = big.NewInt(0)
var bigInt_1 = big.NewInt(1)

type ContentCategory big.Int

func NewContentCategory(p string) *ContentCategory {
	cat := ContentCategories[p]
	if cat != nil {
		return cat
	}
	return big.NewInt(0)
}

func AssempleContentCategory(p ...string) *ContentCategory {
    c := big.NewInt(0)
	for _, v := range(p) {
	    cat := ContentCategories[v]
	    if cat != nil {
		    c.And(c, cat)
	    }
	}
	return c
}

func (c *ContentCategory) Check(cat *ContentCategory) bool {
	return (big.NewInt(0).And(c, cat).Cmp(bigInt_0) != 0)
}

var ContentCategories = map[string]*ContentCategory {
	"IAB1":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 0)),   // Arts & Entertainment
	"IAB1-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 1)),   // Books & Literature
	"IAB1-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 2)),   // Celebrity Fan/Gossip
	"IAB1-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 3)),   // Fine Art
	"IAB1-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 4)),   // Humor
	"IAB1-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 5)),   // Movies
	"IAB1-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 6)),   // Music
	"IAB1-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 7)),   // Television
	"IAB2":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 8)),   // Automotive
	"IAB2-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 9)),   // Auto Parts
	"IAB2-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 10)),  // Auto Repair
	"IAB2-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 11)),  // Buying/Selling Cars
	"IAB2-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 12)),  // Car Culture
	"IAB2-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 13)),  // Certified Pre-Owned
	"IAB2-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 14)),  // Convertible
	"IAB2-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 15)),  // Coupe
	"IAB2-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 16)),  // Crossover
	"IAB2-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 17)),  // Diesel
	"IAB2-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 18)),  // Electric Vehicle
	"IAB2-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 19)),  // Hatchback
	"IAB2-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 20)),  // Hybrid
	"IAB2-13":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 21)),  // Luxury
	"IAB2-14":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 22)),  // Minivan
	"IAB2-15":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 23)),  // Motorcycles
	"IAB2-16":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 24)),  // Off-Road Vehicles
	"IAB2-17":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 25)),  // Performance Vehicles
	"IAB2-18":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 26)),  // Pickup
	"IAB2-19":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 27)),  // Road-Side Assistance
	"IAB2-20":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 28)),  // Sedan
	"IAB2-21":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 29)),  // Trucks & Accessories
	"IAB2-22":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 30)),  // Vintage Cars
	"IAB2-23":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 31)),  // Wagon
	"IAB3":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 32)),  // Business
	"IAB3-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 33)),  // Advertising
	"IAB3-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 34)),  // Agriculture
	"IAB3-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 35)),  // Biotech/Biomedical
	"IAB3-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 36)),  // Business Software
	"IAB3-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 37)),  // Construction
	"IAB3-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 38)),  // Forestry
	"IAB3-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 39)),  // Government
	"IAB3-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 40)),  // Green Solutions
	"IAB3-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 41)),  // Human Resources
	"IAB3-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 42)),  // Logistics
	"IAB3-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 43)),  // Marketing
	"IAB3-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 44)),  // Metals
	"IAB4":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 45)),  // Careers
	"IAB4-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 46)),  // Career Planning
	"IAB4-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 47)),  // College
	"IAB4-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 48)),  // Financial Aid
	"IAB4-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 49)),  // Job Fairs
	"IAB4-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 50)),  // Job Search
	"IAB4-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 51)),  // Resume Writing/Advice
	"IAB4-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 52)),  // Nursing
	"IAB4-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 53)),  // Scholarships
	"IAB4-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 54)),  // Telecommuting
	"IAB4-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 55)),  // U.S. Military
	"IAB4-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 56)),  // Career Advice
	"IAB5":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 57)),  // Education
	"IAB5-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 58)),  // 7-12 Education
	"IAB5-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 59)),  // Adult Education
	"IAB5-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 60)),  // Art History
	"IAB5-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 61)),  // College Administration
	"IAB5-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 62)),  // College Life
	"IAB5-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 63)),  // Distance Learning
	"IAB5-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 64)),  // English as a 2nd Language
	"IAB5-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 65)),  // Language Learning
	"IAB5-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 66)),  // Graduate School
	"IAB5-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 67)),  // Homeschooling
	"IAB5-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 68)),  // Homework/Study Tips
	"IAB5-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 69)),  // K-6 Educators
	"IAB5-13":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 70)),  // Private School
	"IAB5-14":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 71)),  // Special Education
	"IAB5-15":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 72)),  // Studying Business
	"IAB6":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 73)),  // Family & Parenting
	"IAB6-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 74)),  // Adoption
	"IAB6-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 75)),  // Babies & Toddlers
	"IAB6-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 76)),  // Daycare/Pre School
	"IAB6-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 77)),  // Family Internet
	"IAB6-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 78)),  // Parenting - K-6 Kids
	"IAB6-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 79)),  // Parenting teens
	"IAB6-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 80)),  // Pregnancy
	"IAB6-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 81)),  // Special Needs Kids
	"IAB6-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 82)),  // Eldercare
	"IAB7":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 83)),  // Health & Fitness
	"IAB7-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 84)),  // Exercise
	"IAB7-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 85)),  // ADD
	"IAB7-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 86)),  // AIDS/HIV
	"IAB7-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 87)),  // Allergies
	"IAB7-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 88)),  // Alternative Medicine
	"IAB7-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 89)),  // Arthritis
	"IAB7-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 90)),  // Asthma
	"IAB7-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 91)),  // Autism/PDD
	"IAB7-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 92)),  // Bipolar Disorder
	"IAB7-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 93)),  // Brain Tumor
	"IAB7-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 94)),  // Cancer
	"IAB7-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 95)),  // Cholesterol
	"IAB7-13":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 96)),  // Chronic Fatigue Syndrome
	"IAB7-14":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 97)),  // Chronic Pain
	"IAB7-15":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 98)),  // Cold & Flu
	"IAB7-16":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 99)),  // Deafness
	"IAB7-17":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 100)), // Dental Care
	"IAB7-18":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 101)), // Depression
	"IAB7-19":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 102)), // Dermatology
	"IAB7-20":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 103)), // Diabetes
	"IAB7-21":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 104)), // Epilepsy
	"IAB7-22":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 105)), // GERD/Acid Reflux
	"IAB7-23":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 106)), // Headaches/Migraines
	"IAB7-24":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 107)), // Heart Disease
	"IAB7-25":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 108)), // Herbs for Health
	"IAB7-26":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 109)), // Holistic Healing
	"IAB7-27":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 110)), // IBS/Crohn’s Disease
	"IAB7-28":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 111)), // Incest/Abuse Support
	"IAB7-29":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 112)), // Incontinence
	"IAB7-30":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 113)), // Infertility
	"IAB7-31":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 114)), // Men’s Health
	"IAB7-32":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 115)), // Nutrition
	"IAB7-33":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 116)), // Orthopedics
	"IAB7-34":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 117)), // Panic/Anxiety Disorders
	"IAB7-35":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 118)), // Pediatrics
	"IAB7-36":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 119)), // Physical Therapy
	"IAB7-37":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 120)), // Psychology/Psychiatry
	"IAB7-38":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 121)), // Senior Health
	"IAB7-39":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 122)), // Sexuality
	"IAB7-40":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 123)), // Sleep Disorders
	"IAB7-41":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 124)), // Smoking Cessation
	"IAB7-42":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 125)), // Substance Abuse
	"IAB7-43":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 126)), // Thyroid Disease
	"IAB7-44":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 127)), // Weight Loss
	"IAB7-45":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 128)), // Women's Health
	"IAB8":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 129)), // Food & Drink
	"IAB8-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 130)), // American Cuisine
	"IAB8-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 131)), // Barbecues & Grilling
	"IAB8-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 132)), // Cajun/Creole
	"IAB8-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 133)), // Chinese Cuisine
	"IAB8-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 134)), // Cocktails/Beer
	"IAB8-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 135)), // Coffee/Tea
	"IAB8-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 136)), // Cuisine-Specific
	"IAB8-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 137)), // Desserts & Baking
	"IAB8-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 138)), // Dining Out
	"IAB8-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 139)), // Food Allergies
	"IAB8-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 140)), // French Cuisine
	"IAB8-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 141)), // Health/Low-Fat Cooking
	"IAB8-13":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 142)), // Italian Cuisine
	"IAB8-14":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 143)), // Japanese Cuisine
	"IAB8-15":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 144)), // Mexican Cuisine
	"IAB8-16":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 145)), // Vegan
	"IAB8-17":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 146)), // Vegetarian
	"IAB8-18":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 147)), // Wine
	"IAB9":     &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 148)), // Hobbies & Interests
	"IAB9-1":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 149)), // Art/Technology
	"IAB9-2":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 150)), // Arts & Crafts
	"IAB9-3":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 151)), // Beadwork
	"IAB9-4":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 152)), // Bird-Watching
	"IAB9-5":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 153)), // Board Games/Puzzles
	"IAB9-6":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 154)), // Candle & Soap Making
	"IAB9-7":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 155)), // Card Games
	"IAB9-8":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 156)), // Chess
	"IAB9-9":   &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 157)), // Cigars
	"IAB9-10":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 158)), // Collecting
	"IAB9-11":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 159)), // Comic Books
	"IAB9-12":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 160)), // Drawing/Sketching
	"IAB9-13":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 161)), // Freelance Writing
	"IAB9-14":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 162)), // Genealogy
	"IAB9-15":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 163)), // Getting Published
	"IAB9-16":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 164)), // Guitar
	"IAB9-17":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 165)), // Home Recording
	"IAB9-18":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 166)), // Investors & Patents
	"IAB9-19":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 167)), // Jewelry Making
	"IAB9-20":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 168)), // Magic & Illusion
	"IAB9-21":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 169)), // Needlework
	"IAB9-22":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 170)), // Painting
	"IAB9-23":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 171)), // Photography
	"IAB9-24":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 172)), // Radio
	"IAB9-25":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 173)), // Roleplaying Games
	"IAB9-26":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 174)), // Sci-Fi & Fantasy
	"IAB9-27":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 175)), // Scrapbooking
	"IAB9-28":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 176)), // Screenwriting
	"IAB9-29":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 177)), // Stamps & Coins
	"IAB9-30":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 178)), // Video & Computer Games
	"IAB9-31":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 179)), // Woodworking
	"IAB10":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 180)), // Home & Garden
	"IAB10-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 181)), // Appliances
	"IAB10-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 182)), // Entertaining
	"IAB10-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 183)), // Environmental Safety
	"IAB10-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 184)), // Gardening
	"IAB10-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 185)), // Home Repair
	"IAB10-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 186)), // Home Theater
	"IAB10-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 187)), // Interior Decorating
	"IAB10-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 188)), // Landscaping
	"IAB10-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 189)), // Remodeling & Construction
	"IAB11":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 190)), // Law, Government, & Politics
	"IAB11-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 191)), // Immigration
	"IAB11-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 192)), // Legal Issues
	"IAB11-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 193)), // U.S. Government Resources
	"IAB11-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 194)), // Politics
	"IAB11-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 195)), // Commentary
	"IAB12":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 196)), // News
	"IAB12-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 197)), // International News
	"IAB12-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 198)), // National News
	"IAB12-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 199)), // Local News
	"IAB13":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 200)), // Personal Finance
	"IAB13-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 201)), // Beginning Investing
	"IAB13-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 202)), // Credit/Debt & Loans
	"IAB13-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 203)), // Financial News
	"IAB13-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 204)), // Financial Planning
	"IAB13-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 205)), // Hedge Fund
	"IAB13-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 206)), // Insurance
	"IAB13-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 207)), // Investing
	"IAB13-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 208)), // Mutual Funds
	"IAB13-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 209)), // Options
	"IAB13-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 210)), // Retirement Planning
	"IAB13-11": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 211)), // Stocks
	"IAB13-12": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 212)), // Tax Planning
	"IAB14":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 213)), // Society
	"IAB14-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 214)), // Dating
	"IAB14-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 215)), // Divorce Support
	"IAB14-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 216)), // Gay Life
	"IAB14-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 217)), // Marriage
	"IAB14-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 218)), // Senior Living
	"IAB14-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 219)), // Teens
	"IAB14-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 220)), // Weddings
	"IAB14-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 221)), // Ethnic Specific
	"IAB15":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 222)), // Science
	"IAB15-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 223)), // Astrology
	"IAB15-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 224)), // Biology
	"IAB15-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 225)), // Chemistry
	"IAB15-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 226)), // Geology
	"IAB15-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 227)), // Paranormal Phenomena
	"IAB15-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 228)), // Physics
	"IAB15-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 229)), // Space/Astronomy
	"IAB15-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 230)), // Geography
	"IAB15-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 231)), // Botany
	"IAB15-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 232)), // Weather
	"IAB16":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 233)), // Pets
	"IAB16-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 234)), // Aquariums
	"IAB16-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 235)), // Birds
	"IAB16-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 236)), // Cats
	"IAB16-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 237)), // Dogs
	"IAB16-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 238)), // Large Animals
	"IAB16-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 239)), // Reptiles
	"IAB16-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 240)), // Veterinary Medicine
	"IAB17":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 241)), // Sports
	"IAB17-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 242)), // Auto Racing
	"IAB17-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 243)), // Baseball
	"IAB17-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 244)), // Bicycling
	"IAB17-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 245)), // Bodybuilding
	"IAB17-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 246)), // Boxing
	"IAB17-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 247)), // Canoeing/Kayaking
	"IAB17-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 248)), // Cheerleading
	"IAB17-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 249)), // Climbing
	"IAB17-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 250)), // Cricket
	"IAB17-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 251)), // Figure Skating
	"IAB17-11": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 252)), // Fly Fishing
	"IAB17-12": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 253)), // Football
	"IAB17-13": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 254)), // Freshwater Fishing
	"IAB17-14": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 255)), // Game & Fish
	"IAB17-15": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 256)), // Golf
	"IAB17-16": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 257)), // Horse Racing
	"IAB17-17": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 258)), // Horses
	"IAB17-18": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 259)), // Hunting/Shooting
	"IAB17-19": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 260)), // Inline Skating
	"IAB17-20": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 261)), // Martial Arts
	"IAB17-21": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 262)), // Mountain Biking
	"IAB17-22": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 263)), // NASCAR Racing
	"IAB17-23": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 264)), // Olympics
	"IAB17-24": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 265)), // Paintball
	"IAB17-25": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 266)), // Power & Motorcycles
	"IAB17-26": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 267)), // Pro Basketball
	"IAB17-27": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 268)), // Pro Ice Hockey
	"IAB17-28": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 269)), // Rodeo
	"IAB17-29": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 270)), // Rugby
	"IAB17-30": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 271)), // Running/Jogging
	"IAB17-31": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 272)), // Sailing
	"IAB17-32": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 273)), // Saltwater Fishing
	"IAB17-33": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 274)), // Scuba Diving
	"IAB17-34": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 275)), // Skateboarding
	"IAB17-35": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 276)), // Skiing
	"IAB17-36": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 277)), // Snowboarding
	"IAB17-37": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 278)), // Surfing/Body-Boarding
	"IAB17-38": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 279)), // Swimming
	"IAB17-39": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 280)), // Table Tennis/Ping-Pong
	"IAB17-40": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 281)), // Tennis
	"IAB17-41": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 282)), // Volleyball
	"IAB17-42": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 283)), // Walking
	"IAB17-43": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 284)), // Waterski/Wakeboard
	"IAB17-44": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 285)), // World Soccer
	"IAB18":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 286)), // Style & Fashion
	"IAB18-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 287)), // Beauty
	"IAB18-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 288)), // Body Art
	"IAB18-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 289)), // Fashion
	"IAB18-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 290)), // Jewelry
	"IAB18-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 291)), // Clothing
	"IAB18-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 292)), // Accessories
	"IAB19":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 293)), // Technology & Computing
	"IAB19-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 294)), // 3-D Graphics
	"IAB19-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 295)), // Animation
	"IAB19-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 296)), // Antivirus Software
	"IAB19-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 297)), // C/C++
	"IAB19-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 298)), // Cameras & Camcorders
	"IAB19-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 299)), // Cell Phones
	"IAB19-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 300)), // Computer Certification
	"IAB19-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 301)), // Computer Networking
	"IAB19-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 302)), // Computer Peripherals
	"IAB19-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 303)), // Computer Reviews
	"IAB19-11": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 304)), // Data Centers
	"IAB19-12": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 305)), // Databases
	"IAB19-13": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 306)), // Desktop Publishing
	"IAB19-14": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 307)), // Desktop Video
	"IAB19-15": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 308)), // Email
	"IAB19-16": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 309)), // Graphics Software
	"IAB19-17": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 310)), // Home Video/DVD
	"IAB19-18": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 311)), // Internet Technology
	"IAB19-19": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 312)), // Java
	"IAB19-20": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 313)), // JavaScript
	"IAB19-21": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 314)), // Mac Support
	"IAB19-22": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 315)), // MP3/MIDI
	"IAB19-23": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 316)), // Net Conferencing
	"IAB19-24": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 317)), // Net for Beginners
	"IAB19-25": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 318)), // Network Security
	"IAB19-26": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 319)), // Palmtops/PDAs
	"IAB19-27": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 320)), // PC Support
	"IAB19-28": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 321)), // Portable
	"IAB19-29": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 322)), // Entertainment
	"IAB19-30": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 323)), // Shareware/Freeware
	"IAB19-31": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 324)), // Unix
	"IAB19-32": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 325)), // Visual Basic
	"IAB19-33": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 326)), // Web Clip Art
	"IAB19-34": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 327)), // Web Design/HTML
	"IAB19-35": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 328)), // Web Search
	"IAB19-36": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 329)), // Windows
	"IAB20":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 330)), // Travel
	"IAB20-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 331)), // Adventure Travel
	"IAB20-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 332)), // Africa
	"IAB20-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 333)), // Air Travel
	"IAB20-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 334)), // Australia & New Zealand
	"IAB20-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 335)), // Bed & Breakfasts
	"IAB20-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 336)), // Budget Travel
	"IAB20-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 337)), // Business Travel
	"IAB20-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 338)), // By US Locale
	"IAB20-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 339)), // Camping
	"IAB20-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 340)), // Canada
	"IAB20-11": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 341)), // Caribbean
	"IAB20-12": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 342)), // Cruises
	"IAB20-13": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 343)), // Eastern Europe
	"IAB20-14": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 344)), // Europe
	"IAB20-15": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 345)), // France
	"IAB20-16": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 346)), // Greece
	"IAB20-17": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 347)), // Honeymoons/Getaways
	"IAB20-18": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 348)), // Hotels
	"IAB20-19": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 349)), // Italy
	"IAB20-20": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 350)), // Japan
	"IAB20-21": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 351)), // Mexico & Central America
	"IAB20-22": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 352)), // National Parks
	"IAB20-23": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 353)), // South America
	"IAB20-24": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 354)), // Spas
	"IAB20-25": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 355)), // Theme Parks
	"IAB20-26": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 356)), // Traveling with Kids
	"IAB20-27": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 357)), // United Kingdom
	"IAB21":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 358)), // Real Estate
	"IAB21-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 359)), // Apartments
	"IAB21-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 360)), // Architects
	"IAB21-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 361)), // Buying/Selling Homes
	"IAB22":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 362)), // Shopping
	"IAB22-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 363)), // Contests & Freebies
	"IAB22-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 364)), // Couponing
	"IAB22-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 365)), // Comparison
	"IAB22-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 366)), // Engines
	"IAB23":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 367)), // Religion & Spirituality
	"IAB23-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 368)), // Alternative Religions
	"IAB23-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 369)), // Atheism/Agnosticism
	"IAB23-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 370)), // Buddhism
	"IAB23-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 371)), // Catholicism
	"IAB23-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 372)), // Christianity
	"IAB23-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 373)), // Hinduism
	"IAB23-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 374)), // Islam
	"IAB23-8":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 375)), // Judaism
	"IAB23-9":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 376)), // Latter-Day Saints
	"IAB23-10": &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 377)), // Pagan/Wiccan
	"IAB24":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 378)), // Uncategorized
	"IAB25":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 379)), // Non-Standard Content
	"IAB25-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 380)), // Unmoderated UGC
	"IAB25-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 381)), // Extreme Graphic/Explicit Violence
	"IAB25-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 382)), // Pornography
	"IAB25-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 383)), // Profane Content
	"IAB25-5":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 384)), // Hate Content
	"IAB25-6":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 385)), // Under Construction
	"IAB25-7":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 386)), // Incentivized
	"IAB26":    &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 387)), // Illegal Content
	"IAB26-1":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 388)), // Illegal Content
	"IAB26-2":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 389)), // Warez
	"IAB26-3":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 390)), // Spyware/Malware
	"IAB26-4":  &ContentCategory(*big.NewInt(0).Lsh(bigInt_1, 391)), // Copyright Infringement
}
