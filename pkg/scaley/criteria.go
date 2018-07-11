package scaley

func CanScale(group Group, direction Direction) bool {
	if len(Candidates(group, direction)) > 0 {
		return true
	}

	return false
}

func Candidates(group Group, direction Direction) []Server {
	if direction == Up {
		return candidatesForUpscale(group)
	}

	return candidatesForDownscale(group)
}

func candidatesForUpscale(group Group) []Server {
	candidates := make([]Server, 0)

	for _, server := range group.Servers {
		if server.State == Stopped {
			candidates = append(candidates, server)
		}
	}

	return candidates
}

func candidatesForDownscale(group Group) []Server {
	candidates := make([]Server, 0)

	for _, server := range group.Servers {
		if server.State == Running {
			candidates = append(candidates, server)
		}
	}

	return candidates
}
