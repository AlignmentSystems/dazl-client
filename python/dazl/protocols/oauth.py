# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

from ..model.network import OAuthSettings
import requests


async def oauth_flow(settings: OAuthSettings) -> OAuthSettings:
    """
    Implementation of an OAuth flow that ensures that a token is provided.

    :param settings:
    :return:
    """
    if not settings.token:
        headers = { 'Accept': 'application/json' }
        data = {
            'client_id': settings.client_id,
            'client_secret': settings.client_secret,
            'audience': "https://daml.com/ledger-api",
            'grant_type': "client_credentials"
        }

        response = None
        try:
            response = requests.post(settings.token_uri, headers=headers, data=data, auth=None)
        except Exception as ex:
            print(ex)
            raise ValueError('the token must be directly supplied at this time')

        if response.status_code != 200:
            logging.error("ERROR: Unable to retrieve token. Exiting")
            raise ValueError('Unable to get token frpm oAuth source')

        json = response.json()
        return_settings = OAuthSettings(
            client_id=settings.client_id,
            client_secret=settings.client_secret,
            token=json['access_token'],
            token_uri=settings.token_uri,
            refresh_token=settings.refresh_token,
            id_token=settings.id_token,
            auth_url=settings.auth_url,
            redirect_uri=settings.redirect_uri
        )

    else:
        return settings
