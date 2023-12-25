package WebsocketServer

import "time"

const VALID_CONNECTION_ID_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
const CONNECTION_ID_LENGTH = 16

const COOKIE_KEY_USERNAME = "username"
const COOKIE_KEY_PASSWORD = "password"

const WEBSOCKET_WATCHDOG_TIMEOUT = 20 * time.Second
