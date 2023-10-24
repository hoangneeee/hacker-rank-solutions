
func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		lineText := ""
		spaceCount := int(n) - 1 - i

		for space := 0; space < spaceCount; space++ {
			lineText += " "
		}

		for stair := 0; stair < i+1; stair++ {
			lineText += "#"
		}

		fmt.Println(lineText)
	}
}