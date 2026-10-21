clean:
	rm -rf api

gen: clean
	java -jar openapi-generator-cli.jar generate -g markdown -i ./openapi.json -o api/ --additional-properties enumClassPrefix=true,generateAliasAsModel=true \
		--inline-schema-name-mappings claimPromoLink_request=ClaimPromoLinkRequest \
		--inline-schema-name-mappings UpdateAffiliateProgramRequest_name=UpdateAffiliateProgramRequestName \
		--inline-schema-name-mappings UpdateAffiliateProgramRequest_description=UpdateAffiliateProgramRequestDescription \
		--inline-schema-name-mappings UpdateAffiliateProgramRequest_is_active=UpdateAffiliateProgramRequestIsActive \
		--inline-schema-name-mappings PromoClaimResponseEnvelope_allOf_data=PromoClaimResponseEnvelopeData \
		--inline-schema-name-mappings CreateOrdersBulkResponse_allOf_data=CreateOrdersBulkResponseData \
		--inline-schema-name-mappings IssuePromoLinkBatchRequest_oneOf=IssuePromoLinkBatchRequestSingleUse \
		--inline-schema-name-mappings IssuePromoLinkBatchRequest_oneOf_1=IssuePromoLinkBatchRequestReusable \
		--inline-schema-name-mappings IssuePromoLinkBatchResponse_allOf_data=IssuePromoLinkBatchResponseData	 \
		--inline-schema-name-mappings PromoLinkBatchListResponse_allOf_data=PromoLinkBatchListResponseData \
		--inline-schema-name-mappings PromoLinkListResponse_allOf_data=PromoLinkListResponseData \
		--inline-schema-name-mappings RevokePromoLinkResponse_allOf_data=RevokePromoLinkResponseData \
		--inline-schema-name-mappings PromoAttributionResponse_allOf_data=PromoAttributionResponseData \
		--inline-schema-name-mappings PromoCampaignSelectorResponse_data_campaigns_inner=PromoCampaignSelectorResponseDataCampaigns \
		--inline-schema-name-mappings PromoCampaignSelectorResponse_data=PromoCampaignSelectorResponseData

