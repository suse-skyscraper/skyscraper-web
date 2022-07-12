package responses

import (
	"net/http"

	"github.com/suse-skyscraper/skyscraper/internal/db"
)

type CloudAccountTagsResponse struct {
	Tags []string `json:"tags"`
}

func (rd *CloudAccountTagsResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewCloudAccountTagsResponse(tags []string) *CloudAccountTagsResponse {
	return &CloudAccountTagsResponse{
		Tags: tags,
	}
}

type CloudAccountItemAttributes struct {
	CloudProvider     string      `json:"cloud_provider"`
	TenantID          string      `json:"tenant_id"`
	AccountID         string      `json:"account_id"`
	Name              string      `json:"name"`
	Active            bool        `json:"active"`
	TagsCurrent       interface{} `json:"tags_current"`
	TagsDesired       interface{} `json:"tags_desired"`
	TagsDriftDetected bool        `json:"tags_drift_detected"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
}

type CloudAccountItem struct {
	ID         string                     `json:"id"`
	Type       string                     `json:"type"`
	Attributes CloudAccountItemAttributes `json:"attributes"`
}

func newCloudAccount(account db.CloudAccount) CloudAccountItem {
	return CloudAccountItem{
		ID:   account.ID.String(),
		Type: "cloud_account",
		Attributes: CloudAccountItemAttributes{
			CloudProvider:     account.Cloud,
			TenantID:          account.TenantID,
			AccountID:         account.AccountID,
			Name:              account.Name,
			Active:            account.Active,
			TagsCurrent:       account.TagsCurrent.Get(),
			TagsDesired:       account.TagsDesired.Get(),
			TagsDriftDetected: account.TagsDriftDetected,
			CreatedAt:         account.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:         account.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
}

type CloudAccountResponse struct {
	Data CloudAccountItem `json:"data"`
}

func (rd *CloudAccountResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewCloudAccountResponse(account db.CloudAccount) *CloudAccountResponse {
	cloudAccount := newCloudAccount(account)
	return &CloudAccountResponse{
		Data: cloudAccount,
	}
}

type CloudAccountListResponse struct {
	Data []CloudAccountItem `json:"data"`
}

func (rd *CloudAccountListResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewCloudAccountListResponse(accounts []db.CloudAccount) *CloudAccountListResponse {
	list := make([]CloudAccountItem, len(accounts))
	for i, account := range accounts {
		list[i] = newCloudAccount(account)
	}

	return &CloudAccountListResponse{
		Data: list,
	}
}
