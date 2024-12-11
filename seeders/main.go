package main

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func trafficSources(db *gorm.DB) {
	trafficSources := []models.TrafficSource{
		{
			Name: "Google",
			Clid: sql.NullString{String: "gclid", Valid: true},
			Domains: []models.TrafficSourceDomain{
				{
					Domain: "google.com",
				},
			},
		},
		{
			Name: "Bing",
			Clid: sql.NullString{String: "msclkid", Valid: true},
			Domains: []models.TrafficSourceDomain{
				{
					Domain: "bing.com",
				},
			},
		},
		{
			Name: "Other",
			Domains: []models.TrafficSourceDomain{
				{
					Domain: "facebook.com",
				},
				{
					Domain: "instagram.com",
				},
			},
		},
		{
			Name: "Direct",
		},
	}

	for _, trafficSource := range trafficSources {
		db.Where(models.TrafficSource{Name: trafficSource.Name}).FirstOrCreate(&trafficSource)
	}
}

func publishersSeed(db *gorm.DB) {
	publishers := []models.Publisher{
		models.Publisher{
			Name:   "DogGear",
			Hash:   getMD5Hash("DogGear"),
			Domain: "*",
			Urls: []models.Url{
				models.Url{
					Url_Path:      "/best-glucosamine-supplement",
				},
				models.Url{
					Url_Path:      "/best-dog-washer",
				},
				models.Url{
					Url_Path:      "/best-dog-brush",
				},
				models.Url{
					Url_Path:      "/best-bicycle-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-kitchen-lock-dog-parent",
				},
				models.Url{
					Url_Path:      "/best-dog-stethoscope",
				},
				models.Url{
					Url_Path:      "/best-dog-nose-balm",
				},
				models.Url{
					Url_Path:      "/best-dog-casket",
				},
				models.Url{
					Url_Path:      "/best-dog-urn",
				},
				models.Url{
					Url_Path:      "/best-dog-washers-reviews",
				},
				models.Url{
					Url_Path:      "/best-food-for-puppy",
				},
				models.Url{
					Url_Path:      "/best-vegetarian-dog-food",
				},
				models.Url{
					Url_Path:      "/best-calming-supplement-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-frisbee-dog",
				},
				models.Url{
					Url_Path:      "/best-puppy-toy",
				},
				models.Url{
					Url_Path:      "/best-collapsible-dog-bowl",
				},
				models.Url{
					Url_Path:      "/best-blowdryer-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-grooming-glove",
				},
				models.Url{
					Url_Path:      "/best-handheld-vacuum-dog-fur",
				},
				models.Url{
					Url_Path:      "/best-stick-vacuum-dog-hair",
				},
				models.Url{
					Url_Path:      "/best-dog-camera",
				},
				models.Url{
					Url_Path:      "/best-food-dogs-allergies",
				},
				models.Url{
					Url_Path:      "/best-rawhide-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-carrier-for-bike-riding",
				},
				models.Url{
					Url_Path:      "/best-dog-insurance",
				},
				models.Url{
					Url_Path:      "/best-dog-crate-tray",
				},
				models.Url{
					Url_Path:      "/best-bow-tie-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-treat-bag",
				},
				models.Url{
					Url_Path:      "/best-canned-dog-food",
				},
				models.Url{
					Url_Path:      "/best-diabetic-dog-food",
				},
				models.Url{
					Url_Path:      "/best-carpet-shampooer",
				},
				models.Url{
					Url_Path:      "/best-dog-sunglasses",
				},
				models.Url{
					Url_Path:      "/best-vacuum-dog-hair",
				},
				models.Url{
					Url_Path:      "/best-mop-dog-owner",
				},
				models.Url{
					Url_Path:      "/best-broom-for-dog-hair",
				},
				models.Url{
					Url_Path:      "/best-dog-tug-toy",
				},
				models.Url{
					Url_Path:      "/best-cheese-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-anti-diarrhea-medication",
				},
				models.Url{
					Url_Path:      "/best-retractable-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-dog-feeding-reminder",
				},
				models.Url{
					Url_Path:      "/amazons-best-glucosamine-supplements",
				},
				models.Url{
					Url_Path:      "/best-dog-sweater",
				},
				models.Url{
					Url_Path:      "/best-flea-treatment-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-car-barrier",
				},
				models.Url{
					Url_Path:      "/best-lint-roller-dog-hair",
				},
				models.Url{
					Url_Path:      "/best-dog-toothpaste",
				},
				models.Url{
					Url_Path:      "/best-bandana-for-dogs",
				},
				models.Url{
					Url_Path:      "/wireless-dog-fence",
				},
				models.Url{
					Url_Path:      "/best-raw-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-food",
				},
				models.Url{
					Url_Path:      "/best-soft-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-food-brand",
				},
				models.Url{
					Url_Path:      "/best-grain-free-dog-food",
				},
				models.Url{
					Url_Path:      "/best-vegan-dog-food",
				},
				models.Url{
					Url_Path:      "/best-blue-dog-food",
				},
				models.Url{
					Url_Path:      "/best-waterless-dog-shampoo",
				},
				models.Url{
					Url_Path:      "/best-dog-toy",
				},
				models.Url{
					Url_Path:      "/best-kong-dog-toy",
				},
				models.Url{
					Url_Path:      "/best-calming-toy-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-crate-pad",
				},
				models.Url{
					Url_Path:      "/best-dog-backpack",
				},
				models.Url{
					Url_Path:      "/best-dog-ramp",
				},
				models.Url{
					Url_Path:      "/best-dog-treat",
				},
				models.Url{
					Url_Path:      "/best-dog-food-topper",
				},
				models.Url{
					Url_Path:      "/best-dog-dental-chew-toy",
				},
				models.Url{
					Url_Path:      "/best-dog-bone",
				},
				models.Url{
					Url_Path:      "/best-dog-handling-glove",
				},
				models.Url{
					Url_Path:      "/best-dog-fur-rake",
				},
				models.Url{
					Url_Path:      "/best-dog-training-clicker",
				},
				models.Url{
					Url_Path:      "/best-flea-tick-prevention-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-blanket",
				},
				models.Url{
					Url_Path:      "/best-heated-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-dog-training-leash",
				},
				models.Url{
					Url_Path:      "/best-dog-nail-clipper",
				},
				models.Url{
					Url_Path:      "/best-dog-conditioner",
				},
				models.Url{
					Url_Path:      "/best-dog-gate",
				},
				models.Url{
					Url_Path:      "/best-dog-agility-equipment",
				},
				models.Url{
					Url_Path:      "/best-dog-couch-cover",
				},
				models.Url{
					Url_Path:      "/best-sports-apparel-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-training-book",
				},
				models.Url{
					Url_Path:      "/best-dog-whistle",
				},
				models.Url{
					Url_Path:      "/best-training-treat-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dry-dog-food",
				},
				models.Url{
					Url_Path:      "/best-natural-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-dental-chew",
				},
				models.Url{
					Url_Path:      "/best-washable-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-outdoor-dog-kennel",
				},
				models.Url{
					Url_Path:      "/best-outdoor-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-calming-hemp-for-dogs-on-amazon",
				},
				models.Url{
					Url_Path:      "/best-puppy-pad",
				},
				models.Url{
					Url_Path:      "/best-pet-travel-bag",
				},
				models.Url{
					Url_Path:      "/best-dog-pool",
				},
				models.Url{
					Url_Path:      "/best-pee-pad-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-treat-container",
				},
				models.Url{
					Url_Path:      "/best-dog-door",
				},
				models.Url{
					Url_Path:      "/best-doorbell-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-shampoo-conditioner",
				},
				models.Url{
					Url_Path:      "/best-dog-shampoo",
				},
				models.Url{
					Url_Path:      "/best-dog-slicker",
				},
				models.Url{
					Url_Path:      "/best-dog-grooming-scissors",
				},
				models.Url{
					Url_Path:      "/best-waterproof-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-flea-spray-dog",
				},
				models.Url{
					Url_Path:      "/best-flea-tick-home-spray",
				},
				models.Url{
					Url_Path:      "/best-birthday-cake-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-food-weight-loss",
				},
				models.Url{
					Url_Path:      "/best-hemp-dog-treat",
				},
				models.Url{
					Url_Path:      "/best-human-grade-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-shock-collar",
				},
				models.Url{
					Url_Path:      "/best-jerky-dog-treat",
				},
				models.Url{
					Url_Path:      "/best-organix-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-gps-tracker",
				},
				models.Url{
					Url_Path:      "/best-dog-crate",
				},
				models.Url{
					Url_Path:      "/best-travel-dog-crate",
				},
				models.Url{
					Url_Path:      "/best-dog-carrier",
				},
				models.Url{
					Url_Path:      "/best-car-seat-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-seat-cover-car",
				},
				models.Url{
					Url_Path:      "/best-rawhide-free-dog-bone",
				},
				models.Url{
					Url_Path:      "/best-travel-dog-food-container",
				},
				models.Url{
					Url_Path:      "/best-all-natural-dog-shampoo",
				},
				models.Url{
					Url_Path:      "/best-raised-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-cooling-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-elevated-dog-bowl",
				},
				models.Url{
					Url_Path:      "/best-dog-cologne-spray",
				},
				models.Url{
					Url_Path:      "/best-dog-massager",
				},
				models.Url{
					Url_Path:      "/best-dog-crate-bed",
				},
				models.Url{
					Url_Path:      "/best-seat-belt-dogs",
				},
				models.Url{
					Url_Path:      "/best-calming-hemp-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-memorial-stone",
				},
				models.Url{
					Url_Path:      "/best-ventilated-dog-backpack",
				},
				models.Url{
					Url_Path:      "/best-flea-comb",
				},
				models.Url{
					Url_Path:      "/best-heated-outdoor-dog-house",
				},
				models.Url{
					Url_Path:      "/best-chew-deterrent-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-cranberry-supplement-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-anti-itch-spray",
				},
				models.Url{
					Url_Path:      "/best-hemp-oil-for-dog",
				},
				models.Url{
					Url_Path:      "/best-repellent-dogs",
				},
				models.Url{
					Url_Path:      "/best-suture-scissor",
				},
				models.Url{
					Url_Path:      "/best-dog-door-mat",
				},
				models.Url{
					Url_Path:      "/best-dog-leash-wall-holder",
				},
				models.Url{
					Url_Path:      "/best-service-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-snuffle-mat-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-interactive-dog-dish",
				},
				models.Url{
					Url_Path:      "/best-custom-dog-bowl",
				},
				models.Url{
					Url_Path:      "/best-baking-mold-dog-treat",
				},
				models.Url{
					Url_Path:      "/best-collapsible-dog-crate",
				},
				models.Url{
					Url_Path:      "/best-outlet-cover-dog-owner",
				},
				models.Url{
					Url_Path:      "/best-dog-travel-bowl",
				},
				models.Url{
					Url_Path:      "/best-slow-feeder-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-treat-cannister",
				},
				models.Url{
					Url_Path:      "/best-dog-gravity-bowl",
				},
				models.Url{
					Url_Path:      "/best-dog-water-fountain",
				},
				models.Url{
					Url_Path:      "/best-dog-bath-towel",
				},
				models.Url{
					Url_Path:      "/best-electrolyte-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-crate-cover",
				},
				models.Url{
					Url_Path:      "/best-dog-kennel-tray",
				},
				models.Url{
					Url_Path:      "/best-dress-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-collar",
				},
				models.Url{
					Url_Path:      "/best-martingale-collar",
				},
				models.Url{
					Url_Path:      "/best-dog-id-tag",
				},
				models.Url{
					Url_Path:      "/best-emotional-support-animal-id-card",
				},
				models.Url{
					Url_Path:      "/best-service-dog-vest",
				},
				models.Url{
					Url_Path:      "/best-tactical-dog-vest",
				},
				models.Url{
					Url_Path:      "/best-hands-free-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-dog-hair-bow-tie",
				},
				models.Url{
					Url_Path:      "/best-dog-collar-with-bowtie",
				},
				models.Url{
					Url_Path:      "/best-spray-bottle-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-toothbrush",
				},
				models.Url{
					Url_Path:      "/best-dog-jacket",
				},
				models.Url{
					Url_Path:      "/best-winter-jacket-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-leather-dog-collar",
				},
				models.Url{
					Url_Path:      "/best-light-up-dog-collar",
				},
				models.Url{
					Url_Path:      "/best-light-up-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-circular-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-dog-necktie",
				},
				models.Url{
					Url_Path:      "/best-birthday-hat-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-orthopedic-dog-bed",
				},
				models.Url{
					Url_Path:      "/best-dog-cave",
				},
				models.Url{
					Url_Path:      "/best-mat-for-dog-bowl",
				},
				models.Url{
					Url_Path:      "/best-led-dog-collar",
				},
				models.Url{
					Url_Path:      "/best-emotional-support-dog-tag",
				},
				models.Url{
					Url_Path:      "/best-ruffwear-dog-coat",
				},
				models.Url{
					Url_Path:      "/best-dog-leash",
				},
				models.Url{
					Url_Path:      "/best-dog-harness",
				},
				models.Url{
					Url_Path:      "/best-dog-diaper",
				},
				models.Url{
					Url_Path:      "/best-swimsuit-dogs",
				},
				models.Url{
					Url_Path:      "/best-wet-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-carrier-sling",
				},
				models.Url{
					Url_Path:      "/best-dog-dna-kit",
				},
				models.Url{
					Url_Path:      "/best-dog-ear-cleaning-solution",
				},
				models.Url{
					Url_Path:      "/best-dog-ear-wipe",
				},
				models.Url{
					Url_Path:      "/best-pig-ear-for-dogs",
				},
				models.Url{
					Url_Path:      "/your-pup-will-love-these-calming-supplements-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-tear-stain-remover",
				},
				models.Url{
					Url_Path:      "/best-dog-wipe",
				},
				models.Url{
					Url_Path:      "/best-dog-eye-wash-wipe",
				},
				models.Url{
					Url_Path:      "/best-do-not-pet-patch-dogs",
				},
				models.Url{
					Url_Path:      "/best-paper-towel-dog-owner",
				},
				models.Url{
					Url_Path:      "/best-dog-runner",
				},
				models.Url{
					Url_Path:      "/best-playpen-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-treadmill",
				},
				models.Url{
					Url_Path:      "/best-dog-supplement",
				},
				models.Url{
					Url_Path:      "/best-multivitamin-dogs",
				},
				models.Url{
					Url_Path:      "/best-fish-oil-dogs",
				},
				models.Url{
					Url_Path:      "/best-probiotic-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-puppy-formula",
				},
				models.Url{
					Url_Path:      "/best-hip-joint-supplement-dogs",
				},
				models.Url{
					Url_Path:      "/best-heartworm-medication-dog",
				},
				models.Url{
					Url_Path:      "/best-rachael-ray-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-dewormer",
				},
				models.Url{
					Url_Path:      "/best-couch-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-stroller",
				},
				models.Url{
					Url_Path:      "/best-dog-house",
				},
				models.Url{
					Url_Path:      "/best-dog-toy-basket",
				},
				models.Url{
					Url_Path:      "/best-interactive-dog-toy",
				},
				models.Url{
					Url_Path:      "/best-floating-dog-toy",
				},
				models.Url{
					Url_Path:      "/best-dog-food-storage-container",
				},
				models.Url{
					Url_Path:      "/best-dog-toy-box",
				},
				models.Url{
					Url_Path:      "/best-dog-nail-polish",
				},
				models.Url{
					Url_Path:      "/best-scoop-for-dog-food",
				},
				models.Url{
					Url_Path:      "/best-dog-stain-odor-remover",
				},
				models.Url{
					Url_Path:      "/best-dog-activity-monitor",
				},
				models.Url{
					Url_Path:      "/best-service-dog-patch",
				},
				models.Url{
					Url_Path:      "/best-dog-poop-scooper",
				},
				models.Url{
					Url_Path:      "/best-car-air-freshener-for-dog-parents",
				},
				models.Url{
					Url_Path:      "/best-dog-deodorizer",
				},
				models.Url{
					Url_Path:      "/best-breath-freshener-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-safe-air-freshener",
				},
				models.Url{
					Url_Path:      "/best-dog-poop-bag",
				},
				models.Url{
					Url_Path:      "/best-dog-clipper",
				},
				models.Url{
					Url_Path:      "/best-dog-thinning-shear",
				},
				models.Url{
					Url_Path:      "/best-dog-nail-dremel",
				},
				models.Url{
					Url_Path:      "/best-dog-grooming-system",
				},
				models.Url{
					Url_Path:      "/best-dog-antiseptic",
				},
				models.Url{
					Url_Path:      "/best-dog-hair-remover-for-laundry",
				},
				models.Url{
					Url_Path:      "/best-dog-detangler",
				},
				models.Url{
					Url_Path:      "/best-dog-collar-leash-set",
				},
				models.Url{
					Url_Path:      "/best-dog-collar-battery",
				},
				models.Url{
					Url_Path:      "/best-prong-collar",
				},
				models.Url{
					Url_Path:      "/best-flea-and-tick-collar-for-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-recovery-cone",
				},
				models.Url{
					Url_Path:      "/best-sunscreen-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-paw-cleaner",
				},
				models.Url{
					Url_Path:      "/best-dog-paw-protection",
				},
				models.Url{
					Url_Path:      "/best-itch-relief-for-dog",
				},
				models.Url{
					Url_Path:      "/best-allergy-relief-supplement-dog",
				},
				models.Url{
					Url_Path:      "/best-dog-first-aid-kit-essential",
				},
				models.Url{
					Url_Path:      "/best-anxiety-relief-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-pain-reliever",
				},
				models.Url{
					Url_Path:      "/best-dog-grooming-table",
				},
				models.Url{
					Url_Path:      "/best-boots-for-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-bowl",
				},
				models.Url{
					Url_Path:      "/best-dog-wheelchair",
				},
				models.Url{
					Url_Path:      "/best-automatic-dog-feeder",
				},
				models.Url{
					Url_Path:      "/best-dog-socks",
				},
				models.Url{
					Url_Path:      "/best-hind-leg-knee-brace-dogs",
				},
				models.Url{
					Url_Path:      "/best-dog-muzzle",
				},
				models.Url{
					Url_Path:      "/best-dog-head-halter",
				},
				models.Url{
					Url_Path:      "/best-dog-anxiety-jacket",
				},
				models.Url{
					Url_Path:      "/best-dog-life-jacket",
				},
				models.Url{
					Url_Path:      "/best-dog-costume",
				},
				models.Url{
					Url_Path:      "/best-dog-bed",
				},
			},
		},
	}

	for _, publisher := range publishers {
		db.Where(models.Publisher{Name: publisher.Name}).FirstOrCreate(&publisher)
	}
}

func tagsSeed(db *gorm.DB) {
	rows, _ := db.Model(models.Publisher{}).Rows()
	defer rows.Close()

	for rows.Next() {
		var publisher models.Publisher
		db.ScanRows(rows, &publisher)

		for i := 1; i <= 2000; i++ {
			var tag models.Tag
			db.Where(models.Tag{
				Tag:          fmt.Sprintf("%v-%v", publisher.Hash, i),
				Publisher_Id: publisher.ID,
			}).FirstOrCreate(&tag)
		}
	}
}

func publisherUrlsSeed(db *gorm.DB) {
	rows, _ := db.Model(models.Publisher{}).Rows()
	defer rows.Close()

	for rows.Next() {
		var publisher models.Publisher
		db.ScanRows(rows, &publisher)

		for i := 1; i <= 10; i++ {
			var url models.Url
			db.Where(models.Url{
				Publisher_Id: publisher.ID,
				Url_Path:     fmt.Sprintf("/review-%v", i),
			}).FirstOrCreate(&url)
		}
	}
}

func attachTagsToPublisherUrlAndTrafficSources(db *gorm.DB) {
	publisherRows, _ := db.Model(models.Publisher{}).Rows()
	defer publisherRows.Close()

	for publisherRows.Next() {
		var publisher models.Publisher
		db.ScanRows(publisherRows, &publisher)

		var urls []models.Url
		db.Model(&publisher).Association("Urls").Find(&urls)

		trafficSourcesRows, _ := db.Model(models.TrafficSource{}).Rows()
		defer trafficSourcesRows.Close()
		for trafficSourcesRows.Next() {
			var trafficSource models.TrafficSource
			db.ScanRows(trafficSourcesRows, &trafficSource)

			for _, url := range urls {
				var tag models.Tag
				db.Model(models.Tag{}).Where("used", false).Where("publisher_id = ?", publisher.ID).First(&tag)

				if tag.ID == 0 {
					fmt.Errorf("Free tag not found. Skip publisher")
					break
				}

				publisherUrl := models.PublisherUrl{
					Publisher_Hash:    publisher.Hash,
					Publisher_Id:      sql.NullInt64{Int64: int64(publisher.ID), Valid: publisher.ID != 0},
					Url_Id:            sql.NullInt64{Int64: int64(url.ID), Valid: url.ID != 0},
					TagStr:            tag.Tag,
					Tag_Id:            sql.NullInt64{Int64: int64(tag.ID), Valid: tag.ID != 0},
					Traffic_Source_Id: sql.NullInt64{Int64: int64(trafficSource.ID), Valid: trafficSource.ID != 0},
				}

				db.Where(models.PublisherUrl{
					Publisher_Hash:    publisher.Hash,
					Url_Id:            sql.NullInt64{Int64: int64(url.ID), Valid: url.ID != 0},
					Traffic_Source_Id: sql.NullInt64{Int64: int64(trafficSource.ID), Valid: trafficSource.ID != 0},
				}).FirstOrCreate(&publisherUrl)

				if publisherUrl.Tag_Id.Int64 == int64(tag.ID) {
					db.Model(&tag).Updates(models.Tag{Used: true, Traffic_Type: sql.NullString{String: "organic", Valid: true}})
				}
			}
		}
	}
}

func main() {
	db := db.GetDBInstance().GetDB()

	publishersSeed(db)
	trafficSources(db)
	tagsSeed(db)
	//publisherUrlsSeed(db)
	attachTagsToPublisherUrlAndTrafficSources(db)
}
