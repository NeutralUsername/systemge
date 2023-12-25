import {
	root,
} from "./root/root.js"

document.body.style.width  = "max-content"
document.body.style.height =  "max-content"

/* export const WS_CONNECTION = new WebSocket("wss://playnodge.com:8443/ws") */
export const WS_CONNECTION = new WebSocket("ws://localhost:8443/ws")

WS_CONNECTION.onclose = function() {
    setTimeout(function() {
        if (WS_CONNECTION.readyState === WebSocket.CLOSED)
            window.location.reload()
    }, 100)
}

export const DELIMITER1 = "\x03"
export const DELIMITER2 = "|"
export const PLOPSOUND = new Audio('https://patienceonlinecards.s3.eu-central-1.amazonaws.com/mixkit-game-ball-tap-2073.wav')
export const DEFAULT_BOARDSTRING = "-3 0 0 0 0 0 100 100\n-3 1 0 0 0 0 100 100\n-3 2 0 0 0 0 100 100\n-3 3 0 0 0 0 100 100\n-2 -1 0 0 0 0 100 100\n-2 0 0 0 0 0 100 100\n-2 1 0 0 0 0 100 100\n-2 2 0 0 0 0 100 100\n-2 3 0 0 0 0 100 100\n-1 -2 0 0 0 0 100 100\n-1 -1 0 0 0 0 100 100\n-1 0 0 0 0 0 100 100\n-1 1 0 0 0 0 100 100\n-1 2 0 0 0 0 100 100\n-1 3 0 0 0 0 100 100\n0 -3 0 0 0 0 100 100\n0 -2 0 0 0 0 100 100\n0 -1 0 0 0 0 100 100\n0 0 0 0 0 0 100 100\n0 1 0 0 0 0 100 100\n0 2 0 0 0 0 100 100\n0 3 0 0 0 0 100 100\n1 -3 0 0 0 0 100 100\n1 -2 0 0 0 0 100 100\n1 -1 0 0 0 0 100 100\n1 0 0 0 0 0 100 100\n1 1 0 0 0 0 100 100\n1 2 0 0 0 0 100 100\n2 -3 0 0 0 0 100 100\n2 -2 0 0 0 0 100 100\n2 -1 0 0 0 0 100 100\n2 0 0 0 0 0 100 100\n2 1 0 0 0 0 100 100\n3 -3 0 0 0 0 100 100\n3 -2 0 0 0 0 100 100\n3 -1 0 0 0 0 100 100\n3 0 0 0 0 0 100 100"
export const UI_THEME = {
	LIGHT : 0,
	DARK  : 1,
}
export const COMMUNICATOR_CONTENT = {
    CONTACTS : 0,
    CHATS    : 1,
    NOTIFICATIONS : 2,
    BLOCKS : 3,
}
export const USER_POWER = {
    GUEST : 0,
    UNCONFIRMED : 1,
    CONFIRMED : 2,
    GM : 3,
    DEV : 4,
    ADMIN : 5,
}

export function jsonReviver (key, value) {
    switch (key) {
        case "createdAt":
            return new Date(value)
        case "registeredAt":
            return new Date(value)
        case "confirmedAt":
            return new Date(value)
        case "sentAt":
            return new Date(value)
    }
    return value;
}

export function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}

export function setCookie(key, value, durationMs) {
	if (durationMs === undefined) durationMs = 0;
	let expires = new Date(Date.now() + durationMs).toUTCString();
	document.cookie = key + "=" + value + "; expires=" + expires + "; path=/";
}

export function constructMessage(type, ...args) {
    let data = ""
	for (let i = 0; i < args.length; i++) {
        data += args[i].toString().trim() + DELIMITER2
	}
    if (data.length > 0 && data[data.length - 1] == DELIMITER2) {
        data = data.substring(0, data.length - 1)
    }
	return type + DELIMITER1 + data + DELIMITER1 
}
export function updateUITheme(theme) {
    if (theme == UI_THEME.LIGHT) {
        document.body.style.backgroundColor = "#ffffff"
        document.body.style.color = "#000000"
    } else if (theme == UI_THEME.DARK) {
        document.body.style.backgroundColor = "#222426"
        document.body.style.color = "#ffffff"
    }
}
	
export function shuffle(array) {
    let currentIndex = array.length,
        temporaryValue, randomIndex;
    while (0 !== currentIndex) {
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex -= 1;
        temporaryValue = array[currentIndex];
        array[currentIndex] = array[randomIndex];
        array[randomIndex] = temporaryValue;
    }
    return array;
}

ReactDOM.render(
    React.createElement(root), document.getElementById('index')
)

/* 

export class comp extends React.Component {
	constructor(props) {
		super(props)
	}
	render() {
		return React.createElement("div", {
				style: {
					display: "flex",
				}
			},
			React.createElement("div", {

			    }, 
            )
		)
	}
} 

*/
