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
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"gopkg.in/ini.v1"
	"path/filepath"
)

// sessionCmd represents the login command
var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Get and set session token with STS",
	Long: `Get and set session token AWS with AWS Security Token Service (STS).

This command set credentials to credentials file.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		prof, err := cmd.Flags().GetString("profile")
		if err != nil {
			return xerrors.Errorf("Failed to parse flag '%s': %w", prof, err)
		}

		newp, err := cmd.Flags().GetString("session-profile")
		if err != nil {
			return xerrors.Errorf("Failed to parse flag '%s': %w", prof, err)
		}

		sess, err := session.NewSessionWithOptions(session.Options{Profile: prof})
		if err != nil {
			return xerrors.Errorf("Failed to start new session: %w", err)
		}

		arn, err := cmd.Flags().GetString("arn-mfa")
		if err != nil {
			return xerrors.Errorf("Failed to parse flag: %w", arn, err)
		}

		code, err := stscreds.StdinTokenProvider()
		if err != nil {
			return xerrors.Errorf("Failed to get token code: %w", err)
		}

		svc := sts.New(sess)
		input := &sts.GetSessionTokenInput{
			SerialNumber: &arn,
			TokenCode:    &code,
		}

		result, err := svc.GetSessionToken(input)
		if err != nil {
			return xerrors.Errorf("Failed to get session token: %w", err)
		}

		home, err := homedir.Dir()
		if err != nil {
			return xerrors.Errorf("Failed to get home directory: %w", err)
		}
		src := filepath.Join(home, ".aws/credentials")
		cfg, err := ini.Load(src)
		if err != nil {
			return xerrors.Errorf("Failed to get credentials file: %w", err)
		}
		sec, err := cfg.NewSection(newp)
		if err != nil {
			return xerrors.Errorf("Failed to create new section: %w", err)
		}
		sec.Key("aws_access_key_id").SetValue(*result.Credentials.AccessKeyId)
		sec.Key("aws_secret_access_key").SetValue(*result.Credentials.SecretAccessKey)
		sec.Key("aws_session_token").SetValue(*result.Credentials.SessionToken)

		if err := cfg.SaveTo(src); err != nil {
			return xerrors.Errorf("Failed to save credentials file: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)

	sessionCmd.Flags().StringP("serial-number", "a", "myRoleArn", "An arn of the MFA device")
	sessionCmd.Flags().StringP("profile", "p", "default", "A name of profile use to get session token")
	sessionCmd.Flags().StringP("session-profile", "s", "", "A name of profile to set credentials")
}
