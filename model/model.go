package model

import "time"

type GoDaddyDomain struct {
	AuthCode string `json:"authCode"`
	Contact  struct {
		AddressMailing struct {
			Address1   string `json:"address1"`
			Address2   string `json:"address2"`
			City       string `json:"city"`
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
			State      string `json:"state"`
		} `json:"addressMailing"`
		Email        string `json:"email"`
		Fax          string `json:"fax"`
		NameFirst    string `json:"nameFirst"`
		NameLast     string `json:"nameLast"`
		Organization string `json:"organization"`
		Phone        string `json:"phone"`
	} `json:"contactAdmin"`
	ContactBilling struct {
		AddressMailing struct {
			Address1   string `json:"address1"`
			Address2   string `json:"address2"`
			City       string `json:"city"`
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
			State      string `json:"state"`
		} `json:"addressMailing"`
		Email        string `json:"email"`
		Fax          string `json:"fax"`
		NameFirst    string `json:"nameFirst"`
		NameLast     string `json:"nameLast"`
		Organization string `json:"organization"`
		Phone        string `json:"phone"`
	} `json:"contactBilling"`
	ContactRegistrant struct {
		AddressMailing struct {
			Address1   string `json:"address1"`
			Address2   string `json:"address2"`
			City       string `json:"city"`
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
			State      string `json:"state"`
		} `json:"addressMailing"`
		Email        string `json:"email"`
		Fax          string `json:"fax"`
		NameFirst    string `json:"nameFirst"`
		NameLast     string `json:"nameLast"`
		Organization string `json:"organization"`
		Phone        string `json:"phone"`
	} `json:"contactRegistrant"`
	ContactTech struct {
		AddressMailing struct {
			Address1   string `json:"address1"`
			Address2   string `json:"address2"`
			City       string `json:"city"`
			Country    string `json:"country"`
			PostalCode string `json:"postalCode"`
			State      string `json:"state"`
		} `json:"addressMailing"`
		Email        string `json:"email"`
		Fax          string `json:"fax"`
		NameFirst    string `json:"nameFirst"`
		NameLast     string `json:"nameLast"`
		Organization string `json:"organization"`
		Phone        string `json:"phone"`
	} `json:"contactTech"`
	CreatedAt           time.Time `json:"createdAt"`
	Domain              string    `json:"domain"`
	DomainID            int       `json:"domainId"`
	ExpirationProtected bool      `json:"expirationProtected"`
	Expires             time.Time `json:"expires"`
	HoldRegistrar       bool      `json:"holdRegistrar"`
	Locked              bool      `json:"locked"`
	NameServers         []string  `json:"nameServers"`
	Privacy             bool      `json:"privacy"`
	RenewAuto           bool      `json:"renewAuto"`
	RenewDeadline       time.Time `json:"renewDeadline"`
	Renewable           bool      `json:"renewable"`
	Status              string    `json:"status"`
	TransferProtected   bool      `json:"transferProtected"`
}
