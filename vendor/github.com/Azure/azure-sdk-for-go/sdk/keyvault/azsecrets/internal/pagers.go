//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// KeyVaultClientGetDeletedSecretsPager provides operations for iterating over paged responses.
type KeyVaultClientGetDeletedSecretsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedSecretsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedSecretsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetDeletedSecretsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSecretListResult.NextLink == nil || len(*p.current.DeletedSecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetDeletedSecretsPager) NextPage(ctx context.Context) (KeyVaultClientGetDeletedSecretsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetDeletedSecretsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetDeletedSecretsResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetDeletedSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetDeletedSecretsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.GetDeletedSecretsHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetDeletedSecretsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// KeyVaultClientGetSecretVersionsPager provides operations for iterating over paged responses.
type KeyVaultClientGetSecretVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetSecretVersionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetSecretVersionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetSecretVersionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetSecretVersionsPager) NextPage(ctx context.Context) (KeyVaultClientGetSecretVersionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetSecretVersionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetSecretVersionsResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetSecretVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetSecretVersionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.GetSecretVersionsHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetSecretVersionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// KeyVaultClientGetSecretsPager provides operations for iterating over paged responses.
type KeyVaultClientGetSecretsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetSecretsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetSecretsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *KeyVaultClientGetSecretsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *KeyVaultClientGetSecretsPager) NextPage(ctx context.Context) (KeyVaultClientGetSecretsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return KeyVaultClientGetSecretsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return KeyVaultClientGetSecretsResponse{}, err
	}
	resp, err := p.client.Pl.Do(req)
	if err != nil {
		return KeyVaultClientGetSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return KeyVaultClientGetSecretsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.GetSecretsHandleResponse(resp)
	if err != nil {
		return KeyVaultClientGetSecretsResponse{}, err
	}
	p.current = result
	return p.current, nil
}