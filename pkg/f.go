package pkg

func init() {

}

func Download(mod string) {
	switch mod {
	case "arcdps":
		modArcdps()
	case "knowthyenemy":
		modKnowthyenemy()
	}
}
