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
	"github.com/spf13/cobra"
	"os"
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
			return fmt.Errorf("Failed to parse profile '%s': %+v", prof, err)
		}

		sess, err := session.NewSessionWithOptions(session.Options{Profile: prof})
		if err != nil {
			return fmt.Errorf("Failed to start new session: %+v", err)
		}

		arn, err := cmd.Flags().GetString("arn-mfa")
		if err != nil {
			return fmt.Errorf("Failed to parse arn: %+v", arn, err)
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

		if err := os.Setenv("AWS_ACCESS_KEY_ID", *result.Credentials.AccessKeyId); err != nil {
			return fmt.Errorf("Failed to set environment variable: %+v", err)
		}
		if err := os.Setenv("AWS_SECRET_ACCESS_KEY", *result.Credentials.SecretAccessKey); err != nil {
			return fmt.Errorf("Failed to set environment variable: %+v", err)
		}
		if err := os.Setenv("AWS_SESSION_TOKEN", *result.Credentials.SessionToken); err != nil {
			return fmt.Errorf("Failed to set environment variable: %+v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("arn-mfa", "m", "myRoleArn", "An arn of the MFA device")
	loginCmd.Flags().StringP("profile", "p", "default", "A name of profile use to get session token")
}
