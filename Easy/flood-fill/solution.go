package floodfill

// Accepted. But with help from internet :(. Need to learn more about recursion. And DFS.

func floodFill(image [][]int, sr, sc, color int) [][]int {
	target := image[sr][sc]

	if target != color {
		fill(image, sr, sc, color, target)
	}

	return image
}

func fill(image [][]int, sr, sc, color, target int) {
	if sr < 0 ||
		sc < 0 ||
		sr >= len(image) ||
		sc >= len(image[0]) ||
		image[sr][sc] != target {
		return
	}

	image[sr][sc] = color

	fill(image, sr-1, sc, color, target)
	fill(image, sr+1, sc, color, target)
	fill(image, sr, sc-1, color, target)
	fill(image, sr, sc+1, color, target)
}
