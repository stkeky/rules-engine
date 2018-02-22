package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/MainfluxLabs/rules-engine/engine"
)

func retrieveRuleEndpoint(svc engine.Service) endpoint.Endpoint {
	return func(_ context.Context, body interface{}) (interface{}, error) {
		res := &viewRuleRes{}

		b := body.(viewRuleReq)

		if err := b.validate(); err != nil {
			return nil, err
		}

		rule, err := svc.ViewRule(b.userId, b.ruleId)
		if err != nil {
			return nil, err
		}

		res.fromDomain(*rule)
		return *res, nil
	}
}

func retrieveRulesEndpoint(svc engine.Service) endpoint.Endpoint {
	return func(_ context.Context, body interface{}) (interface{}, error) {
		res := &listRulesRes{}

		b := body.(listRulesReq)

		if err := b.validate(); err != nil {
			return nil, err
		}

		rulesList, err := svc.ListRules(b.userId)
		if err != nil {
			return nil, err
		}

		res.fromDomain(rulesList)
		return *res, nil
	}
}

func removeRuleEndpoint(svc engine.Service) endpoint.Endpoint {
	return func(_ context.Context, body interface{}) (interface{}, error) {
		b := body.(viewRuleReq)

		if err := b.validate(); err != nil {
			return nil, err
		}

		if err := svc.RemoveRule(b.userId, b.ruleId); err != nil {
			return nil, err
		}

		return removeRes{}, nil
	}
}
