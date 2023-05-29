package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const gopherDirectory = "gophers/"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName string

		// check if the user has provided a gopher name as argument or not, if not, throw an error
		if len(args) < 1 {
			fmt.Println("Error: You must provide a Gopher name!")
			return
		}

		// Error for too many arguments
		if len(args) > 1 && args[0] != "" {
			fmt.Println("Error: Too many arguments!")
			return
		}

		// If the user has provided a gopher name as argument, use it
		if args[0] != "" {
			gopherName = args[0]
		}

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Println("Trying to get '" + gopherName + "' Gopher...")

		// Create the directory if it doesn't exist
		err := os.MkdirAll("gophers", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get the data
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Specify the full path to the file in the "gophers" directory
			filePath := filepath.Join(gopherDirectory, gopherName+".png")

			// Create the file
			out, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()

			// Writer the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Perfect! Just saved in " + out.Name() + "!")
		} else {
			fmt.Println("Error: " + gopherName + " not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
