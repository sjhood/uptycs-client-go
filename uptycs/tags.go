package uptycs

import (
	"errors"
)

func (T Tag) GetID() string {
	return T.ID
}

func (T Tag) GetName() string {
	return T.Name
}

func (T Tag) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateTag(tag Tag) (Tag, error) {
	for ind, fpg := range tag.FilePathGroups {
		if len(fpg.Name) > 0 && len(fpg.ID) == 0 {
			_fpg, _ := c.GetFilePathGroup(FilePathGroup{
				Name: fpg.Name,
			})
			tag.FilePathGroups[ind] = TagConfigurationObject{
				ID: _fpg.ID,
			}
		}
	}
	for ind, eep := range tag.EventExcludeProfiles {
		if len(eep.Name) > 0 && len(eep.ID) == 0 {
			_eep, _ := c.GetEventExcludeProfile(EventExcludeProfile{
				Name: eep.Name,
			})
			tag.EventExcludeProfiles[ind] = TagConfigurationObject{
				ID: _eep.ID,
			}
		}
	}
	for ind, rp := range tag.RegistryPaths {
		if len(rp.Name) > 0 && len(rp.ID) == 0 {
			_rp, _ := c.GetRegistryPath(RegistryPath{
				Name: rp.Name,
			})
			tag.RegistryPaths[ind] = TagConfigurationObject{
				ID: _rp.ID,
			}
		}
	}
	for ind, qp := range tag.Querypacks {
		if len(qp.Name) > 0 && len(qp.ID) == 0 {
			_qp, _ := c.GetQuerypack(Querypack{
				Name: qp.Name,
			})
			tag.Querypacks[ind] = TagConfigurationObject{
				ID: _qp.ID,
			}
		}
	}
	for ind, ygr := range tag.YaraGroupRules {
		if len(ygr.Name) > 0 && len(ygr.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ygr.Name,
			})
			tag.YaraGroupRules[ind] = TagConfigurationObject{
				ID: _ygr.ID,
			}
		}
	}
	for ind, ac := range tag.AuditConfigurations {
		if len(ac.Name) > 0 && len(ac.ID) == 0 {
			_ac, _ := c.GetAuditConfiguration(AuditConfiguration{
				Name: ac.Name,
			})
			tag.AuditConfigurations[ind] = TagConfigurationObject{
				ID: _ac.ID,
			}
		}
	}
	return doCreate(c, tag, "tags")
}

func (c *Client) UpdateTag(tag Tag) (Tag, error) {
	for ind, fpg := range tag.FilePathGroups {
		if len(fpg.Name) > 0 && len(fpg.ID) == 0 {
			_fpg, _ := c.GetFilePathGroup(FilePathGroup{
				Name: fpg.Name,
			})
			tag.FilePathGroups[ind] = TagConfigurationObject{
				ID: _fpg.ID,
			}
		}
	}
	for ind, eep := range tag.EventExcludeProfiles {
		if len(eep.Name) > 0 && len(eep.ID) == 0 {
			_eep, _ := c.GetEventExcludeProfile(EventExcludeProfile{
				Name: eep.Name,
			})
			tag.EventExcludeProfiles[ind] = TagConfigurationObject{
				ID: _eep.ID,
			}
		}
	}
	for ind, rp := range tag.RegistryPaths {
		if len(rp.Name) > 0 && len(rp.ID) == 0 {
			_rp, _ := c.GetRegistryPath(RegistryPath{
				Name: rp.Name,
			})
			tag.RegistryPaths[ind] = TagConfigurationObject{
				ID: _rp.ID,
			}
		}
	}
	for ind, qp := range tag.Querypacks {
		if len(qp.Name) > 0 && len(qp.ID) == 0 {
			_qp, _ := c.GetQuerypack(Querypack{
				Name: qp.Name,
			})
			tag.Querypacks[ind] = TagConfigurationObject{
				ID: _qp.ID,
			}
		}
	}
	for ind, ygr := range tag.YaraGroupRules {
		if len(ygr.Name) > 0 && len(ygr.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ygr.Name,
			})
			tag.YaraGroupRules[ind] = TagConfigurationObject{
				ID: _ygr.ID,
			}
		}
	}
	for ind, ac := range tag.AuditConfigurations {
		if len(ac.Name) > 0 && len(ac.ID) == 0 {
			_ac, _ := c.GetAuditConfiguration(AuditConfiguration{
				Name: ac.Name,
			})
			tag.AuditConfigurations[ind] = TagConfigurationObject{
				ID: _ac.ID,
			}
		}
	}
	return doUpdate(c, tag, "tags")
}

func (c *Client) GetTags() (Tags, error) {
	return doGetMany(c, Tags{}, "tags")
}

func (c *Client) GetTag(tag Tag) (Tag, error) {
	if len(tag.ID) == 0 {
		all, _ := c.GetTags()
		for _, _item := range all.Items {
			if _item.Key == tag.Key && _item.Value == tag.Value {
				return _item, nil
			}
		}
	} else {
		return doGet(c, tag, "tags")
	}
	return tag, errors.New("tag was not found")
}

func (c *Client) DeleteTag(tag Tag) (Tag, error) {
	return doDelete(c, tag, "tags")
}
