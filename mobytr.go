package main

import (
	_ "context"
	"fmt"
	_ "github.com/docker/docker/api/types"
	_ "github.com/docker/docker/client"
    "os"
    "archive/zip"
    _ "bufio"
    _ "io"
)

func main() {
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	//if err != nil {
	//	panic(err)
	//}

    //file, err := os.Open("./Dockerfile")
    //if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//defer file.Close() 
    //reader := io.Reader(file)

    zeper, _:= os.Open("/submission.zip")
    zipr, _ := zip.NewReader(zeper, header.Size)
    for _, z := range zipr.File {

        fmt.Printf("type: %T\n", z.FileHeader)
        fmt.Printf("type: %s\n", z.FileHeader.Name)
    }
    fmt.Printf("AAAAAAAAAAAAAAAAA")
    //dockerBuildContext, err := os.Open("/subm.tar")
    //defer dockerBuildContext.Close()

    //options := types.ImageBuildOptions{
    //    Tags: []string{"fromdec"},
    //    SuppressOutput: true,
    //    Dockerfile: "submission/Dockerfile",
    //}
    //resp, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
    //if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
    //defer resp.Body.Close()
    ////respData, err := io.ReadAll(resp.Body)
    //list_opt := types.ImageListOptions{All: true}
    //summary, err := cli.ImageList(context.Background(), list_opt)
    //if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
    //for _, img := range summary {
    //    for _,v := range img.RepoTags {
    //        fmt.Printf("%s \n", v)
    //    }
    //}
}
