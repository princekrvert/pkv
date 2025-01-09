/*
Copyright © 2025 prince kumar
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// make a function to check if the file is there
func fileIsThere(filePath string) bool {
	// now check the file with os stat
	file, err := os.Stat(filePath)
	if err == nil {
		filepath.Match(filePath, file.Name())
		if err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}

}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This commnad is used to get the value",
	Long:  `This commnad is used to extract the value form the pkv file`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("file")
		if fileIsThere(fileName) {
			// now read the file
		} else {
			fmt.Println("File not found")
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
	getCmd.Flags().StringP("file", "f", "*.pkv", "File path of pkv file")
}
