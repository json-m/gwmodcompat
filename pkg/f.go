package pkg

func init() {
	// start thing to watch for new versions to download of each mod
}

func Download(mod string) {
	// arcdps
	// https://www.deltaconnected.com/arcdps/x64/d3d11.dll
	// checksum: https://www.deltaconnected.com/arcdps/x64/d3d11.dll.md5sum

	// arcdps knowthyenemy switch
	switch mod {
	case "arcdps":
		modArcdps()
	case "knowthyenemy":
		modKnowthyenemy()
	}
}
