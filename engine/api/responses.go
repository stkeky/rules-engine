package api

import (
	"fmt"
	"net/http"

	"github.com/MainfluxLabs/rules-engine/engine"
)

const contentType = "application/json; charset=utf-8"

type apiRes interface {
	code() int
	headers() map[string]string
	empty() bool
}

type viewRuleRes struct {
	ruleRes
}

type ruleRes struct {
	ID         string          `json:"id"`
	Name       string          `json:"name,omitempty"`
	Conditions []conditionRes  `json:"conditions"`
	Actions    []engine.Action `json:"actions"`
}

type conditionRes struct {
	DeviceID string          `json:"deviceId"`
	Property string          `json:"property"`
	Operator engine.Operator `json:"operator"`
	Value    interface{}     `json:"value"`
}

func (res *ruleRes) fromDomain(rule engine.Rule) {
	var cnds []conditionRes
	cndRes := &conditionRes{}

	res.ID = rule.ID
	res.Name = rule.Name
	res.Actions = rule.Actions

	for _, c := range rule.Conditions {
		cndRes.fromDomain(c)
		cnds = append(cnds, *cndRes)
	}

	res.Conditions = cnds
}

func (res *conditionRes) fromDomain(cnd engine.Condition) {
	res.DeviceID = cnd.DeviceID
	res.Property = cnd.Property
	res.Operator = cnd.Operator
	res.Value = cnd.Value
}

func (res viewRuleRes) code() int {
	return http.StatusOK
}

func (res viewRuleRes) headers() map[string]string {
	return map[string]string{}
}

func (res viewRuleRes) empty() bool {
	return false
}

type listRulesRes struct {
	Rules []ruleRes `json:"rules"`
	count int
}

func (res *listRulesRes) fromDomain(rules []engine.Rule) {
	ruleResps := make([]ruleRes, 0)
	ruleRes := &ruleRes{}

	for _, r := range rules {
		ruleRes.fromDomain(r)
		ruleResps = append(ruleResps, *ruleRes)
	}

	res.Rules = ruleResps
	res.count = len(rules)
}

func (res listRulesRes) code() int {
	return http.StatusOK
}

func (res listRulesRes) headers() map[string]string {
	return map[string]string{
		"X-Count": fmt.Sprintf("%d", res.count),
	}
}

func (res listRulesRes) empty() bool {
	return false
}

type removeRes struct{}

func (res removeRes) code() int {
	return http.StatusNoContent
}

func (res removeRes) headers() map[string]string {
	return map[string]string{}
}

func (res removeRes) empty() bool {
	return false
}
