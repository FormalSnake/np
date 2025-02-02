package main

func hyphen(uuid string) string {
	if len(uuid) != 32 {
		return ""
	}
	return uuid[:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:]
}

// │ ┌ ┐ └ ┘ ─ ┤ ├

func printInfo(profile *Profile) {
	m := map[string]string{
		"Track":  profile.Track.Name,
		"Artist": profile.Track.Artist.Name,
		"URL":    profile.Track.Url,
	}
	prettyPrint(m)

}

func line(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "─"
	}
	return s
}

const padding = 4

func prettyPrint(m map[string]string) {
	longestK := 0
	for k := range m {
		if len(k) > longestK {
			longestK = len(k)
		}
	}
	longestV := 0
	for _, v := range m {
		if len(v) > longestV {
			longestV = len(v)
		}
	}

	println("┌" + line(longestK+longestV+padding*2+1) + "┐")
	for k, v := range m {
		print("│ ", k)
		padDiff(k, longestK+padding)
		print(v)
		padDiff(v, longestV+padding)
		println("│")
	}
	println("└" + line(longestK+longestV+padding*2+1) + "┘")
}

func padDiff(s string, n int) {
	for i := 0; i < n-len(s); i++ {
		print(" ")
	}
}
