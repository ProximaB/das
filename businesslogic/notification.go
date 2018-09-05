// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package businesslogic

import "time"

const (
	// NotificationCategoryNewPartnershipRequestReceived is the value for New Partnership Request Received notification
	NotificationCategoryNewPartnershipRequestReceived = 1
	NotificationCategoryPartnershipRequestResponded   = 2
	NotificationCategoryRoleApplicationResponded      = 3
	NotificationCategoryRegistrationOpened            = 4
)

// NotificationCategory defines different categories of notifications
type NotificationCategory struct {
	ID              int
	Name            string
	Abbreviation    string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// INotificationCategoryRepository specifies the interface that a NotificationCategoryRepository should implement
type INotificationCategoryRepository interface {
	GetAllNotificationCategories() ([]NotificationCategory, error)
}

// NotificationPreference stores the preference of how user would like to receive system-generated notification
type NotificationPreference struct {
	ID              int
	AccountID       int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// SearchNotificationPreferenceCriteria specifies the parameters that can be used to search notification preferences in a repo
type SearchNotificationPreferenceCriteria struct {
	AccountID int
}

// INotificationPreferenceRepository specifies the interface that a Notification Preference repository should implement
type INotificationPreferenceRepository interface {
	CreateNotificationPreference(pref *NotificationPreference) error
	DeleteNotificationPreference(pref NotificationPreference) error
	SearchNotificationPreference(criteria SearchNotificationPreferenceCriteria) ([]NotificationPreference, error)
	UpdateNotificationPreference(pref NotificationPreference) error
}

// Notification stores the content of a notification and the status of it
type Notification struct {
	ID                     int
	AccountID              int
	NotificationCategoryID int
	HasRead                bool
	DateTimeCreated        time.Time
}

// SearchNotificationCriteria specifies the parameters that can be used to search notifications in a rep
type SearchNotificationCriteria struct {
	AccountID              int
	NotificationCategoryID int
}

// INotificationRepository specifies the interface that a Notification Repository should implement
type INotificationRepository interface {
	CreateNotification(notification *Notification) error
	DeleteNotification(notification Notification) error
	SearchNotification(criteria SearchNotificationCriteria) ([]Notification, error)
	UpdateNotification(notification Notification) error
}
