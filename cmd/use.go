/*
Copyright © 2019 shunsuke maeda

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use specified profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		prof, err := cmd.Flags().GetString("profile")
		if err != nil {
			return xerrors.Errorf("Failed to parse flag '%s': %w", prof, err)
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return xerrors.Errorf("Failed to parse flag '%s': %w", file, err)
		}

		cred, err := credentials.NewSharedCredentials(file, prof).Get()
		if err != nil {
			return xerrors.Errorf("Failed to get credentials '%s': %w", prof, err)
		}

		shell := exec.Command(os.Getenv("SHELL"), "-l")

		shell.Stdin = os.Stdin
		shell.Stdout = os.Stdout
		shell.Stderr = os.Stderr

		if err := os.Setenv("AWS_ACCESS_KEY_ID", cred.AccessKeyID); err != nil {
			return xerrors.Errorf("Failed to set environment variable: %w", err)
		}
		if err := os.Setenv("AWS_SECRET_ACCESS_KEY", cred.SecretAccessKey); err != nil {
			return xerrors.Errorf("Failed to set environment variable: %w", err)
		}
		if err := os.Setenv("AWS_SESSION_TOKEN", cred.SessionToken); err != nil {
			return xerrors.Errorf("Failed to set environment variable: %w", err)
		}

		shell.Env = os.Environ()

		return shell.Run()
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	useCmd.Flags().StringP("profile", "p", "default", "A name of profile use to get session token")
	useCmd.Flags().StringP("file", "f", filepath.Join(os.Getenv("HOME"), ".aws/credentials"), "Path to shared credentials file")

	if err := mfaCmd.MarkFlagRequired("profile"); err != nil {
		log.Fatalf("Failed to mark flag required: %+v\n", err)
	}
}
