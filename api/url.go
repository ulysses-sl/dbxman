package dbxman

import "net/url"

const DbxOAuthBaseURL = "https://www.dropbox.com"
const DbxDefaultBaseURL = "https://api.dropboxapi.com"
const DbxContentBaseURL = "https://content.dropboxapi.com"
const DbxNotifyBaseURL = "https://notify.dropboxapi.com"

type BaseURL struct {
	Api     *url.URL
	Content *url.URL
	Notify  *url.URL
	OAuth   *url.URL
}

func NewBaseURL() *BaseURL {
	apiURL, err := url.Parse(DbxDefaultBaseURL)
	if err != nil {
		panic("Base URL for Dropbox default API is invalid.")
	}
	contentURL, err := url.Parse(DbxContentBaseURL)
	if err != nil {
		panic("Base URL for Dropbox content API is invalid.")
	}
	notifyURL, err := url.Parse(DbxNotifyBaseURL)
	if err != nil {
		panic("Base URL for Dropbox notify API is invalid.")
	}
	oauthURL, err := url.Parse(DbxOAuthBaseURL)
	if err != nil {
		panic("Base URL for Dropbox OAuth API is invalid.")
	}
	baseURL := BaseURL{
		Api:     apiURL,
		Content: contentURL,
		Notify:  notifyURL,
		OAuth:   oauthURL,
	}
	return &baseURL
}
