import {
    WS_CONNECTION,
	setCookie,
	updateUITheme,
	DELIMITER1,
	DELIMITER2,
	jsonReviver,
	constructMessage,
} from "../index.js"
import {
	menubar,
} from "../menubar/menubar.js"
import {
	communicator
} from "../communicator/communicator.js"
import {
	register
} from "../account/register.js";
import {
	login
} from "../account/login.js";
import { 
	home 
} from "../home/home.js";

export class root extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			content : "home",
			lockKeybinds : false,
			profile : null,
			leaderboard : null,

			communicatorContactsInput : "",
			communicatorBlocksInput : "",
			communicatorNotificationsInput : "",
			communicatorChatsInputs : {},

			pms : {},
			contacts : [],
			chats : [],
			invites : [],
			blocks : [],
			
			communicatorSelectedUser : null,

			setStateRoot : (change) => this.setState(change),
		},

        WS_CONNECTION.onmessage = (event) => {
			console.log(event.data)
			let segments = event.data.split(DELIMITER1);
			let type = segments[0]
			let data1 = segments[1].split(DELIMITER2)
			switch (type) {
			case "login":
				let loginData = JSON.parse(data1[0], jsonReviver)
				setCookie("username", loginData.username, 1000*60*60*24*30*12)
				setCookie("password", loginData.password, 1000*60*60*24*30*12)
				updateUITheme(loginData.uiTheme)
				this.state.setStateRoot({  
					loginData : loginData,
	
				})
				if(this.state.game && this.state.game.ongoingTurn ) {
					this.state.setStateRoot({gameLoadedTurnIndex : this.state.game.ongoingTurn.number})
				}
				this.state.setStateRoot({content : "home"})
				break;
			case "uiTheme":
				this.state.loginData.uiTheme = Number(data1)
				updateUITheme(this.state.loginData.uiTheme)
				this.state.setStateRoot({loginData : this.state.loginData})
				break;
			case "showCommunicator":
				this.state.loginData.showCommunicator = data1 == "true"
				this.state.setStateRoot({loginData : this.state.loginData})
				break;
			case "communicatorContent":
				this.state.loginData.communicatorContent = Number(data1)
				this.state.setStateRoot({loginData : this.state.loginData})
				break;
			}	
        };
	}

	getContent = () => {
		switch (this.state.content) {
		case "home":
			return React.createElement(home, this.state)
		case "login":
			return React.createElement(login, this.state)
		case "register":
			return React.createElement(register, this.state)
		}
	}

	render() {
		return this.state.loginData ? React.createElement('div', {
				id : "root",
				onContextMenu: e => {
					e.preventDefault()
				},
				style : {
					fontFamily : "sans-serif",
					display : "flex",
					flexDirection : "column",
					justifyContent : "center",
					alignItems : "center",
				}
			},
			React.createElement(menubar, this.state),
			this.getContent(),
			React.createElement("div", {
					style : {
						display : "flex",
						position : "fixed",
						right : "1vmin",
						top : "5vmax",
						zIndex : "100",
						gap : "1vmin",
					}
				},
				this.state.loginData.showCommunicator ? React.createElement(communicator, this.state) 
				: null,
			),
		) 
		: React.createElement('div', {})
	}
}