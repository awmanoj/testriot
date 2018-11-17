package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = riot.Engine{}
)

var blacklist = []string{
	"alat bantu wanita fsdkfsd fsjdkfjsd  fsjlkfjsdlf d",
	"alat bantu pria flskjflsd fjsdlk jflsdkj",
}

func main() {
	// Init
	searcher.Init(types.EngineOpts{
		// Using:             4,
		NotUseGse: true,
	})

	defer searcher.Close()
	for idx, keyword := range blacklist {
		searcher.Index(uint64(idx), types.DocData{Content: keyword})
	}

	// Wait for the index to refresh
	searcher.Flush()

	// The search output format is found in the types.SearchResp structure
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		if input == "exit" {
			log.Println("Exiting...")
			return
		}

		fmt.Println("**", input)
		resp := searcher.Search(types.SearchReq{Text: input})
		fmt.Println("***", resp)

		// docs := resp.Docs.(types.ScoredDocs)
		// fmt.Println("**", docs)
		// for i := 0; i < resp.NumDocs; i++ {
		// 	fmt.Printf("== [%d][%s]\n", docs[i].DocId, docs[i].Content)
		// }
	}
}
