package cas

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"go-service-template/config"
	"go-service-template/internal/infrastructure/adapter/cas/casDTO"
	"go-service-template/pkgs/gplog"
	"go-service-template/pkgs/httpCaller"
	"go-service-template/pkgs/utils/mapper"
)

type casAdapterImpl struct {
	cfg        *config.AppConfig
	httpCaller httpCaller.Client
}

func NewCasAdapter(cfg *config.AppConfig) CasAdapter {
	httpClient := httpCaller.New()

	httpClient.Resty().
		SetDebug(cfg.Server.Debug).
		SetTimeout(cfg.Server.CtxDefaultTimeout).
		SetBaseURL(cfg.Adapter.CentralAuthService.URL).
		SetHeader("X-API-KEY", cfg.Adapter.CentralAuthService.APIKey)

	return &casAdapterImpl{
		cfg:        cfg,
		httpCaller: httpClient,
	}
}

func (c casAdapterImpl) VerifyToken(ctx context.Context, req *casDTO.CASVerifyTokenReq) (*casDTO.CASVerifyTokenResp, error) {
	resp, err := c.httpCaller.Resty().R().
		SetBody(req).
		Post(req.URL())
	if err != nil {
		gplog.Errorf("failed to call CAS service: %v", err)
		return nil, err
	}

	//check status code 200
	if resp.StatusCode() != 200 {
		gplog.Errorf("failed to call CAS service: %v", err)
		return nil, errors.New("failed to call CAS service")
	}

	//parse response
	baseResponse := new(casDTO.CasBaseResp)
	if err := sonic.Unmarshal(resp.Body(), baseResponse); err != nil {
		gplog.Errorf("failed to unmarshal response: %v", err)
		return nil, err
	}

	if baseResponse.Code != casDTO.OkResponseCode {
		gplog.Errorf("CAS service return error: %v", baseResponse.Message)
		return nil, errors.New(baseResponse.Message)
	}

	//parse destination response
	destinationResponse := new(casDTO.CASVerifyTokenResp)
	if err := mapper.BindingStruct(baseResponse.Data, destinationResponse); err != nil {
		gplog.Errorf("failed to bind response: %v", err)
		return nil, err
	}

	return destinationResponse, nil
}
