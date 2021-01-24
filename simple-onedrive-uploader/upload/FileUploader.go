package upload

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"https://github.com/thehaohcm/simple-onedrive/config"
	"https://github.com/thehaohcm/simple-onedrive/models"
	"https://github.com/thehaohcm/simple-onedrive/token"
	"strconv"
	"sync"
	"time"
)

var (
	fileName  string
	blockSize = 0
	fileBytes []byte
	fragSize  = config.FragSize //the defautl value is 59,375 MB (get from config file)
	//59,375 MB is a largest fragment size we can split for uploading
	//fragment size must be a multiple of 320 KiB
	//more detail: https://docs.microsoft.com/en-us/onedrive/developer/rest-api/api/driveitem_createuploadsession?view=odsp-graph-online

	fileSize               = 0
	fileCapacity           string
	uploadFinishedResponse models.UploadFinishedResponse
)

func UploadFile(localFilePath string, uploadFolderPath string) {

	token.RefreshToken()
	loadAndGetFileInfo(localFilePath)

	uploadJSONResult := SendInitUploadRequest(fileName, uploadFolderPath)

	//init, send request and get response
	if uploadJSONResult.UploadUrl != "" {
		fmt.Println("Uploading the file " + fileName + " into OneDrive...")
		for i := 0; i < blockSize; i++ {
			isSuccess := false
			numberOfAttempt := -1
			finishedPercent := float64(float64(i+1)/float64(blockSize)) * 100
			finishedPercentText := fmt.Sprintf("%.2f", finishedPercent)
			for !isSuccess {
				fmt.Println("Uploading fragment number: " + strconv.Itoa(i+1) + "/" + strconv.Itoa(blockSize) + "....(" + finishedPercentText + "%)")
				numberOfAttempt++
				if numberOfAttempt > 0 {
					fmt.Println("Uploading was failed, trying to upload again (Number of attempting: " + strconv.Itoa(numberOfAttempt) + ")....")
				}

				//check and refresh token if it expired by time
				if time.Now().After(config.ExpiredTokenTime) || numberOfAttempt > 0 {
					fmt.Println("The Token is expired, refreshing...")
					token.RefreshToken()
				}

				isSuccess = true
				var byteBlock []byte
				maxBlockRange := len(fileBytes)
				if i < blockSize-1 {
					maxBlockRange = i*fragSize + fragSize
				}
				byteBlock = fileBytes[(i * fragSize):maxBlockRange]

				resp, err := SendUploadBlockRequest(uploadJSONResult.UploadUrl, i, maxBlockRange, byteBlock)
				if err != nil {
					panic(err)
				}

				if i == blockSize-1 {
					uploadFinishedResponse, _ := resp.(models.UploadFinishedResponse)
					fmt.Println("Uploading is finished, file: " + uploadFinishedResponse.Name + " (size: " + fileCapacity + ")")
					fmt.Println("Download link: " + uploadFinishedResponse.DownloadUrl)

					fmt.Println("Sharing the file...")
					linkWebURL := ShareLinkFunc(&uploadFinishedResponse)
					if linkWebURL != "" {
						fmt.Println("The file " + fileName + " (" + uploadFinishedResponse.Name + ", size: " + fileCapacity + ") has been shared for every one")
						fmt.Println(" link view URL: " + linkWebURL)
					}
				}
			}
		}
	}
}

func UploadFileInThread(localFilePath string, uploadFolderPath string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(localFilePath string, uploadFolderPath string) {
		defer wg.Done()
		UploadFile(localFilePath, uploadFolderPath)
		wg.Done()
	}(localFilePath, uploadFolderPath)

	wg.Wait()
	fmt.Println("the file " + localFilePath + " has been uploaded successfully")
}

func loadAndGetFileInfo(localFilePath string) (string, int) {
	fileName = filepath.Base(localFilePath)

	fi, err := os.Open(localFilePath)
	fileBytes, err = ioutil.ReadFile(localFilePath)
	if err != nil {
		panic(err)
	}

	//get file size
	fileData, err := fi.Stat()
	if err != nil {
		panic(err)
	}
	fileSize = int(fileData.Size())

	fileCapacity = getReadableCapacity(fileSize)
	fmt.Println("File's capacity: " + fileCapacity)

	// read file into bytes
	blockSize = (fileSize + fragSize - 1) / fragSize
	return fileName, blockSize
}
