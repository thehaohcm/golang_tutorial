package demo

import (
	"context"
	"fmt"
	"os"

	"https://github.com/thehaohcm/simple-onedrive/config"
	"https://github.com/thehaohcm/simple-onedrive/token"
	"https://github.com/thehaohcm/simple-onedrive/upload"

	"github.com/goh-chunlin/go-onedrive/onedrive"
	"golang.org/x/oauth2"
)

func getInstance() (context.Context, *onedrive.Client) {
	token.RefreshToken()
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.SavedToken.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := onedrive.NewClient(tc)

	return ctx, client
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 1 {
		fmt.Println("Please add or check a specific file path as 1 agurement")
		return
	}
	filePath := args[0]

	upload.UploadFile(filePath, config.UploadFolderPath)
	ctx, client := getInstance()

	drives, err := client.Drives.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, drive := range drives.Drives {
		fmt.Printf("Results: %v\n", drive.Owner.User.DisplayName)
	}

	//get list item of root
	driveItems, err := client.DriveItems.List(ctx, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Items of root folder: ")
		for _, driveItem := range driveItems.DriveItems {
			fmt.Printf(" -%v ", driveItem.Name)
			// if driveItem.Folder != nil {
			// 	fmt.Printf("- Folder\n")
			// }
		}
	}
}
