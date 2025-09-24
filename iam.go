package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func getUserPermissions(client *iam.Client, userName string) ([]string, error) {
	var permissions []string

	attachedResp, err := client.ListAttachedUserPolicies(context.TODO(),
		&iam.ListAttachedUserPoliciesInput{
			UserName: aws.String(userName),
		})
	if err != nil {
		return nil, err
	}
	for _, policy := range attachedResp.AttachedPolicies {
		permissions = append(permissions, *policy.PolicyName)
	}

	inlineResp, err := client.ListUserPolicies(context.TODO(),
		&iam.ListUserPoliciesInput{
			UserName: aws.String(userName),
		})
	if err != nil {
		return nil, err
	}
	for _, policyName := range inlineResp.PolicyNames {
		permissions = append(permissions, policyName)
	}

	return permissions, nil
}

func listAllUsers() (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	client := iam.NewFromConfig(cfg)

	usersResp, err := client.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		return "", err
	}

	report := "User\t\tPermissions\n----------------------------------\n"
	for _, user := range usersResp.Users {
		perms, err := getUserPermissions(client, *user.UserName)
		if err != nil {
			return "", err
		}
		permsStr := "No Policies"
		if len(perms) > 0 {
			permsStr = strings.Join(perms, ", ")
		}
		report += fmt.Sprintf("%s\t\t%s\n", *user.UserName, permsStr)
	}

	return report, nil
}
