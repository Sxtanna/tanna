package token

type Augmenter struct {
	IntoToken func(here, prev, next *Data) []*Data
	HereMatch func(data *Data) bool
	PrevMatch func(data *Data) bool
	NextMatch func(data *Data) bool
	SkipCount int
	BackTrack int
}

func (a *Augmenter) Matches(here, prev, next *Data) bool {
	return !((a.HereMatch != nil && (here == nil || !a.HereMatch(here))) ||
		(a.PrevMatch != nil && (prev == nil || !a.PrevMatch(prev))) ||
		(a.NextMatch != nil && (next == nil || !a.NextMatch(next))))
}

func ExecuteAugmentation(data []*Data, augmenters []*Augmenter) []*Data {
	result := make([]*Data, 0, len(data))

	i := 0
	for {
		if i >= len(data) {
			break
		}

		var here = data[i]
		var prev *Data
		var next *Data

		if i-1 < 0 {
			prev = nil
		} else {
			prev = data[i-1]
		}

		if i+1 >= len(data) {
			next = nil
		} else {
			next = data[i+1]
		}

		var augmenter *Augmenter

		for _, test := range augmenters {
			if !test.Matches(here, prev, next) {
				continue
			}

			augmenter = test
			break
		}

		if augmenter == nil {
			result = append(result, here)
		} else {
			i = i + augmenter.SkipCount

			if augmenter.BackTrack != 0 {
				result = append(result[:(len(result)-augmenter.BackTrack)], result[(len(result)-augmenter.BackTrack)+1:]...)
			}

			result = append(result, augmenter.IntoToken(here, prev, next)...)
		}

		i++
	}

	return result
}
