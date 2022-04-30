package cmd

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add values passed to function",
	Long:  `Demo application to demonstrate cobra featues`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := rootCmd.Flags().GetString("username")
		password, _ := rootCmd.Flags().GetString("password")
		auth := fmt.Sprintf("%s:%s", username, password)
		authEncode := b64.StdEncoding.EncodeToString([]byte(auth))

		url := "https://api.github.com/user/repos"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authEncode))

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		var response []interface{}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(body, &response)

		for _, repoDetails := range response {
			repo := repoDetails.(map[string]interface{})
			fmt.Println(" name : ", repo["name"], " private :", repo["private"], "clone_url :", repo["clone_url"])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.PersistentFlags().StringP("username", "u", "", "the username of git")
	rootCmd.PersistentFlags().StringP("password", "p", "", "the access token of the git")
}
