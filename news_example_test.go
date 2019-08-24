package news_test

import (
	"fmt"
	"log"

	"github.com/tj/go-news"
)

func Example() {
	list := news.New("news_test")

	emails := []string{
		"tobi@apex.sh",
		"loki@apex.sh",
		"jane@apex.sh",
		"manny@apex.sh",
		"luna@apex.sh",
	}

	for _, email := range emails {
		err := list.AddSubscriber("product_updates", email)
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}
	}

	subscribers, err := list.GetSubscribers("product_updates")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	fmt.Printf("%#v\n", subscribers)
	// Output:
	// []string{"jane@apex.sh", "loki@apex.sh", "luna@apex.sh", "manny@apex.sh", "tobi@apex.sh"}
}
