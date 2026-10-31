package game

func (state *GameState) GallowsArt() string {
	return getGallowsAsciiArts()[state.mistakes]
}

func getGallowsAsciiArts() []string {
	return []string{
		`
__________
|/       |
|
|
|
|
|
----------`,
		`
__________
|/       |
|       (_)
|
|
|
|
----------`,
		`
__________
|/       |
|       (_)
|        |
|        |
|
|
----------`,
		`
__________
|/       |
|       (_)
|      / |
|        |
|
|
----------`,
		`                
__________
|/       |
|       (_)
|      / | \
|        |
|
|
----------`,
		`
__________
|/       |
|       (_)
|      / | \
|        |
|       /
|
----------`,
		`
__________
|/       |
|       (_)
|      / | \
|        |
|       / \
|
----------`,
	}
}
