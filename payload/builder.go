// Package payload is a helper package which contains a payload
// builder to make constructing notification payloads easier.
package payload

import "encoding/json"

// Payload represents a notification which holds the content that will be
// marshalled as JSON.
type Payload struct {
	content map[string]interface{}
}

type aps struct {
	Alert            interface{} `json:"alert,omitempty"`
	Badge            interface{} `json:"badge,omitempty"`
	Category         string      `json:"category,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
	URLArgs          []string    `json:"url-args,omitempty"`
	Sound            string      `json:"sound,omitempty"`
}

type alert struct {
	Action       string   `json:"action,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	Body         string   `json:"body,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	Title        string   `json:"title,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
}

// NewPayload returns a new Payload struct
func NewPayload() *Payload {
	return &Payload{
		map[string]interface{}{
			"aps": &aps{},
		},
	}
}

// Alert sets the aps alert on the payload.
// This will display a notification alert message to the user.
//
//	{"aps":{"alert":alert}}`
func (p *Payload) Alert(alert interface{}) *Payload {
	p.aps().Alert = alert
	return p
}

// Badge sets the aps badge on the payload.
// This will display a numeric badge on the app icon.
//
//	{"aps":{"badge":b}}
func (p *Payload) Badge(b int) *Payload {
	p.aps().Badge = b
	return p
}

// ZeroBadge sets the aps badge on the payload to 0.
// This will clear the badge on the app icon.
//
//	{"aps":{"badge":0}}
func (p *Payload) ZeroBadge() *Payload {
	p.aps().Badge = 0
	return p
}

// UnsetBadge removes the badge attribute from the payload.
// This will leave the badge on the app icon unchanged.
// If you wish to clear the app icon badge, use ZeroBadge() instead.
//
//	{"aps":{}}
func (p *Payload) UnsetBadge() *Payload {
	p.aps().Badge = nil
	return p
}

// Sound sets the aps sound on the payload.
// This will play a sound from the app bundle, or the default sound otherwise.
//
//	{"aps":{"sound":sound}}
func (p *Payload) Sound(sound string) *Payload {
	p.aps().Sound = sound
	return p
}

// ContentAvailable sets the aps content-available on the payload to 1.
// This will indicate to the app that there is new content available to download
// and launch the app in the background.
//
//	{"aps":{"content-available":1}}
func (p *Payload) ContentAvailable() *Payload {
	p.aps().ContentAvailable = 1
	return p
}

// Custom payload

// Custom sets a custom key and value on the payload.
// This will add custom key/value data to the notification payload at root level.
//
//	{"aps":{}, key:value}
func (p *Payload) Custom(key string, val interface{}) *Payload {
	p.content[key] = val
	return p
}

// Alert dictionary

// AlertTitle sets the aps alert title on the payload.
// This will display a short string describing the purpose of the notification.
// Apple Watch & Safari display this string as part of the notification interface.
//
//	{"aps":{"alert":"title"}}
func (p *Payload) AlertTitle(title string) *Payload {
	p.aps().alert().Title = title
	return p
}

// AlertTitleLocKey sets the aps alert title localization key on the payload.
// This is the key to a title string in the Localizable.strings file for the
// current localization. See Localized Formatted Strings in Apple documentation
// for more information.
//
//	{"aps":{"alert":{"title-loc-key":key}}}
func (p *Payload) AlertTitleLocKey(key string) *Payload {
	p.aps().alert().TitleLocKey = key
	return p
}

// AlertTitleLocArgs sets the aps alert title localization args on the payload.
// These are the variable string values to appear in place of the format
// specifiers in title-loc-key. See Localized Formatted Strings in Apple
// documentation for more information.
//
//	{"aps":{"alert":{"title-loc-args":args}}}
func (p *Payload) AlertTitleLocArgs(args []string) *Payload {
	p.aps().alert().TitleLocArgs = args
	return p
}

// AlertBody sets the aps alert body on the payload.
// This is the text of the alert message.
//
//	{"aps":{"alert":{"body":body}}}
func (p *Payload) AlertBody(body string) *Payload {
	p.aps().alert().Body = body
	return p
}

// AlertLaunchImage sets the aps launch image on the payload.
// This is the filename of an image file in the app bundle. The image is used
// as the launch image when users tap the action button or move the action
// slider.
//
//	{"aps":{"alert":{"launch-image":image}}}
func (p *Payload) AlertLaunchImage(image string) *Payload {
	p.aps().alert().LaunchImage = image
	return p
}

// AlertLocKey sets the aps alert localization key on the payload.
// This is the key to an alert-message string in the Localizable.strings file
// for the current localization. See Localized Formatted Strings in Apple
// documentation for more information.
//
//	{"aps":{"alert":{"loc-key":key}}}
func (p *Payload) AlertLocKey(key string) *Payload {
	p.aps().alert().LocKey = key
	return p
}

// AlertAction sets the aps alert action on the payload.
// This is the label of the action button, if the user sets the notifications
// to appear as alerts. This label should be succinct, such as “Details” or
// “Read more”. If omitted, the default value is “Show”.
//
//	{"aps":{"alert":{"action":action}}}
func (p *Payload) AlertAction(action string) *Payload {
	p.aps().alert().Action = action
	return p
}

// AlertActionLocKey sets the aps alert action localization key on the payload.
// This is the the string used as a key to get a localized string in the current
// localization to use for the notfication right button’s title instead of
// “View”. See Localized Formatted Strings in Apple documentation for more
// information.
//
//	{"aps":{"alert":{"action-loc-key":key}}}
func (p *Payload) AlertActionLocKey(key string) *Payload {
	p.aps().alert().ActionLocKey = key
	return p
}

// General

// Category sets the aps category on the payload.
// This is a string value that represents the identifier property of the
// UIMutableUserNotificationCategory object you created to define custom actions.
//
//	{"aps":{"alert":{"category":category}}}
func (p *Payload) Category(category string) *Payload {
	p.aps().Category = category
	return p
}

// Mdm sets the mdm on the payload.
// This is for Apple Mobile Device Management (mdm) payloads.
//
//	{"aps":{}:"mdm":mdm}
func (p *Payload) Mdm(mdm string) *Payload {
	p.content["mdm"] = mdm
	return p
}

// URLArgs sets the aps category on the payload.
// This specifies an array of values that are paired with the placeholders
// inside the urlFormatString value of your website.json file.
// See Apple Notification Programming Guide for Websites.
//
//	{"aps":{"url-args":urlArgs}}
func (p *Payload) URLArgs(urlArgs []string) *Payload {
	p.aps().URLArgs = urlArgs
	return p
}

// MarshalJSON returns the JSON encoded version of the Payload
func (p *Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.content)
}

func (p *Payload) aps() *aps {
	return p.content["aps"].(*aps)
}

func (a *aps) alert() *alert {
	if _, ok := a.Alert.(*alert); !ok {
		a.Alert = &alert{}
	}
	return a.Alert.(*alert)
}
