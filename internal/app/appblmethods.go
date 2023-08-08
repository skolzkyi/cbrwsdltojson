package app

import (
	"context"
	"strings"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
)

func (a *App) GetCursOnDateXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.GetCursOnDateXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "GetCursOnDateXML"
		startNodeName := "ValuteData"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return nil, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.GetCursOnDateXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}
		inputAsserted, ok := input.(*datastructures.GetCursOnDateXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		res, err := a.soapSender.SoapCall(ctx, SOAPMethod, *inputAsserted)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.XMLToStructDecoder(res, startNodeName, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		for i := range response.ValuteCursOnDate {
			response.ValuteCursOnDate[i].Vname = strings.TrimSpace(response.ValuteCursOnDate[i].Vname)
			response.ValuteCursOnDate[i].Vname = strings.Trim(response.ValuteCursOnDate[i].Vname, "\r\n")
		}
		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) BiCurBaseXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.BiCurBaseXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "BiCurBaseXML"
		startNodeName := "BiCurBase"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.BiCurBaseXMLResult{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.BiCurBaseXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.BiCurBaseXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return nil, err
		}

		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) BliquidityXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.BliquidityXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "BliquidityXML"
		startNodeName := "Bliquidity"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.BliquidityXMLResult{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.BliquidityXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.BliquidityXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) DepoDynamicXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.DepoDynamicXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "DepoDynamicXML"
		startNodeName := "DepoDynamic"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.DepoDynamicXML{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.DepoDynamicXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.DepoDynamicXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) DragMetDynamicXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.DragMetDynamicXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "DragMetDynamicXML"
		startNodeName := "DragMetall"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.DragMetDynamicXML{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.DragMetDynamicXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.DragMetDynamicXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) DVXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.DVXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "DVXML"
		startNodeName := "DV_base"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.DVXML{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.DVXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.DVXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) EnumReutersValutesXML(ctx context.Context) (interface{}, error) {
	var err error
	var response datastructures.EnumReutersValutesXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "EnumReutersValutesXML"
		startNodeName := "ReutersValutesList"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.EnumReutersValutesXMLResult{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.Appmemcache.GetCacheDataInCache(SOAPMethod)
		if ok {
			response, ok = cachedData.Payload.(datastructures.EnumReutersValutesXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted := datastructures.EnumReutersValutesXML{}
		inputAsserted.Init()

		res, err := a.soapSender.SoapCall(ctx, SOAPMethod, inputAsserted)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.XMLToStructDecoder(res, startNodeName, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		a.Appmemcache.AddOrUpdatePayloadInCache(SOAPMethod, response)
	}
	return response, nil
}

func (a *App) EnumValutesXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.EnumValutesXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "EnumValutesXML"
		startNodeName := "ValuteData"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.EnumValutesXML{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.EnumValutesXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.EnumValutesXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		res, err := a.soapSender.SoapCall(ctx, SOAPMethod, *inputAsserted)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.XMLToStructDecoder(res, startNodeName, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		for i := range response.EnumValutes {
			response.EnumValutes[i].Vcode = strings.TrimSpace(response.EnumValutes[i].Vcode)
			response.EnumValutes[i].Vname = strings.TrimSpace(response.EnumValutes[i].Vname)
			response.EnumValutes[i].VEngname = strings.TrimSpace(response.EnumValutes[i].VEngname)
			response.EnumValutes[i].VcommonCode = strings.TrimSpace(response.EnumValutes[i].VcommonCode)
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}

func (a *App) KeyRateXML(ctx context.Context, input interface{}, rawBody string) (interface{}, error) {
	var err error
	var response datastructures.KeyRateXMLResult
	select {
	case <-ctx.Done():
		err = ErrContextWSReqExpired
		a.logger.Error(err.Error())
		return response, err
	default:
		SOAPMethod := "KeyRateXML"
		startNodeName := "KeyRate"
		if a.permittedRequests.PermittedRequestMapLength() > 0 {
			if a.permittedRequests.IsPermittedRequestInMap(SOAPMethod) {
				return datastructures.KeyRateXML{}, ErrMethodProhibited
			}
		}

		cachedData, ok := a.GetDataInCacheIfExisting(SOAPMethod, rawBody)
		if ok {
			response, ok = cachedData.(datastructures.KeyRateXMLResult)
			if !ok {
				err = ErrAssertionAfterGetCacheData
				a.logger.Error(err.Error())
			} else {
				return response, nil
			}
		}

		inputAsserted, ok := input.(*datastructures.KeyRateXML)
		if !ok {
			err = ErrAssertionOfInputData
			a.logger.Error(err.Error())
			return response, err
		}
		err = a.ProcessRequest(ctx, SOAPMethod, startNodeName, *inputAsserted, &response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}

		err = a.AddOrUpdateDataInCache(SOAPMethod, input, response)
		if err != nil {
			a.logger.Error(err.Error())
			return response, err
		}
	}
	return response, nil
}
