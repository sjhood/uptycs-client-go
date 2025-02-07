package uptycs

import "errors"

func (T AlertRule) GetID() string {
	return T.ID
}

func (T AlertRule) GetName() string {
	return T.Name
}

func (T AlertRule) KeysToDelete() []string {
	keysToDelete := []string{
		"seedId",
		"throttled",
		"lock",
		"alertTags",
		"links",
	}

	if T.Type != "sql" {
		keysToDelete = append(keysToDelete, "sqlConfig")
	}
	if T.Type != "javascript" {
		keysToDelete = append(keysToDelete, "scriptConfig")
	}

	return keysToDelete
}

func (c *Client) GetAlertRules() (AlertRules, error) {
	return doGetMany(c, AlertRules{}, "alertRules")
}

func (c *Client) GetAlertRule(alertRule AlertRule) (AlertRule, error) {
	if len(alertRule.ID) == 0 {
		all, _ := c.GetAlertRules()
		for _, _item := range all.Items {
			if _item.Name == alertRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, alertRule, "alertRules")
	}
	return alertRule, errors.New("alertRule was not found")
}

func (c *Client) DeleteAlertRule(alertRule AlertRule) (AlertRule, error) {
	return doDelete(c, alertRule, "alertRules")
}

func (c *Client) CreateAlertRule(alertRule AlertRule) (AlertRule, error) {
	return doCreate(c, alertRule, "alertRules")
}

func (c *Client) UpdateAlertRule(alertRule AlertRule) (AlertRule, error) {
	return doUpdate(c, alertRule, "alertRules")
}
