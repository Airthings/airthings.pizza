#!/usr/bin/env python
# -*- coding: utf-8 -*-

import api
import os
import requests
import base64

from slackclient import SlackClient
from time import sleep

pizza_channel_id = os.environ["SLACK_PIZZA_CHANNEL_ID"]
slack_token = os.environ["SLACK_API_TOKEN"]
sc = SlackClient(slack_token)


def is_dm(message):
    return message['channel'][0] == 'D'


if sc.rtm_connect():
    while True:
        event_list = sc.rtm_read()
        message_list = list(
            filter(lambda m: m['type'] == 'message', event_list))
        for message in message_list:
            if(is_dm(message) and 'user' in message):
                if message['user'] in api.get_invited_users():
                    if message['text'].lower() == 'yes':
                        api.rsvp(message['user'], 'attending')
                        api.send_slack_message(
                            message['channel'], u'Sweet! 🤙')
                        api.finalize_event_if_complete()
                    elif message['text'].lower() == 'no':
                        api.rsvp(message['user'], 'not attending')
                        api.send_slack_message(message['channel'], u'Ok 😏')
                        api.invite_if_needed()
                    else:
                        api.send_slack_message(
                            message['channel'], u'Not sure what you meant 😳. Can you come? (yes/no)')
        sleep(0.5)

else:
    print("Connection Failed, invalid token?")
