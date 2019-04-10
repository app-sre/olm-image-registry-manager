package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

var currentCSV = &cobra.Command{
	Use:   "current-csv",
	Short: "get the current CSV",
	Run: func(cmd *cobra.Command, args []string) {
		gitRepoURL = "https://github.com/app-sre/saas-hive-operator-bundle.git"
		gitBranch = "staging"
		gitDir = "hive"

		fs := memfs.New()

		// Git objects storer based on memory
		storer := memory.NewStorage()

		// Clones the repository into the worktree (fs) and storer all the .git
		// content into the storer
		_, err := git.Clone(storer, fs, &git.CloneOptions{
			URL:           gitRepoURL,
			ReferenceName: plumbing.ReferenceName(gitBranch),
		})
		if err != nil {
			log.Fatal(err)
		}

		files, err := fs.ReadDir(gitDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
		// package_file := fmt.Sprintf("%s/")
		// // Prints the content of the CHANGELOG file from the cloned repository
		// changelog, err := fs.Open("README.md")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// io.Copy(os.Stdout, changelog)
	},
}

func init() {
	rootCmd.AddCommand(currentCSV)
}
