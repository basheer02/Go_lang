package maps

import "fmt"

func main() {
	websites := map[any]any {
		"Google": "https://google.com",
		"AWS": "https://aws.com",
	}

	fmt.Println(websites)

	websites[123] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}