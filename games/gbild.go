import (
    "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/archive"
    "context"
)

func main() {
            tar, err := archive.TarWithOptions("node-hello/", &archive.TarOptions{})
            resp, err := cli.ImageBuild(context.Background(), file, options)
            if err != nil {
		        fmt.Println("Error:", err)
		        return
	        }
            defer resp.Body.Close()
}
