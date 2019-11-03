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
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"path/filepath"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login AWS with STS",
	Long: `login AWS with AWS Security Token Service (STS)

This command set credentials to environment variables.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		prof, err := cmd.Flags().GetString("profile")
		if err != nil {
			return fmt.Errorf("Failed to parse flag '%s': %+v", prof, err)
		}

		newp, err := cmd.Flags().GetString("session-profile")
		if err != nil {
			return fmt.Errorf("Failed to parse flag '%s': %+v", prof, err)
		}

		sess, err := session.NewSessionWithOptions(session.Options{Profile: prof})
		if err != nil {
			return fmt.Errorf("Failed to start new session: %+v", err)
		}

		arn, err := cmd.Flags().GetString("arn-mfa")
		if err != nil {
			return fmt.Errorf("Failed to parse flag: %+v", arn, err)
		}

		code, err := stscreds.StdinTokenProvider()
		if err != nil {
			return fmt.Errorf("Failed to get token code: %+v", err)
		}

		svc := sts.New(sess)
		input := &sts.GetSessionTokenInput{
			SerialNumber: &arn,
			TokenCode:    &code,
		}

		result, err := svc.GetSessionToken(input)
		if err != nil {
			return fmt.Errorf("Failed to get session token: %+v", err)
		}

		home, err := homedir.Dir()
		if err != nil {
			return fmt.Errorf("Failed to get home directory: %+v", err)
		}
		src := filepath.Join(home, ".aws/credentials")
		cfg, err := ini.Load(src)
		if err != nil {
			return fmt.Errorf("Failed to get credentials file: %+v", err)
		}
		sec, err := cfg.NewSection(newp)
		if err != nil {
			return fmt.Errorf("Failed to create new section: %+v", err)
		}
		sec.Key("aws_access_key_id").SetValue(*result.Credentials.AccessKeyId)
		sec.Key("aws_secret_access_key").SetValue(*result.Credentials.SecretAccessKey)
		sec.Key("aws_session_token").SetValue(*result.Credentials.SessionToken)

		if err := cfg.SaveTo(src); err != nil {
			return fmt.Errorf("Failed to save credentials file: %+v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("serial-number", "a", "myRoleArn", "An arn of the MFA device")
	loginCmd.Flags().StringP("profile", "p", "default", "A name of profile use to get session token")
	loginCmd.Flags().StringP("session-profile", "s", "", "A name of profile to set credentials")
}
