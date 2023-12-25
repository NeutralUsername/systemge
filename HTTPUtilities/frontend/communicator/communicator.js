import {
	tab
} from "./tab.js"
import {
	entries
} from "./entries.js"
import {
	input
} from "./input.js"
import {
	chat
} from "./chat.js"
import { 
	UI_THEME,
	constructMessage,
	WS_CONNECTION,
	COMMUNICATOR_CONTENT
} from "../index.js"
import { label } from "./label.js"

export class communicator extends React.Component {
	constructor(props) {
		super(props)
	}

	hourglassImg = () => {
        if (this.props.loginData.uiTheme === UI_THEME.DARK) {
            return "https://nodge.s3.eu-central-1.amazonaws.com/hg_w.png"
        } else {
            return "https://nodge.s3.eu-central-1.amazonaws.com/hg_b.png"
        }
    }

	chatImg = () => {
		if (this.props.loginData.uiTheme === UI_THEME.DARK) {
			return "https://nodge.s3.eu-central-1.amazonaws.com/chat_w.png"
		} else {
			return "https://nodge.s3.eu-central-1.amazonaws.com/chat_b.png"
		}
	}

	lobbyImg = () => {
		if (this.props.loginData.uiTheme === UI_THEME.DARK) {
			return "https://nodge.s3.eu-central-1.amazonaws.com/plus_w.svg"
		} else {
			return "https://nodge.s3.eu-central-1.amazonaws.com/plus_b.svg"
		}
	}

	blockImg = () => {
		if (this.props.loginData.uiTheme === UI_THEME.DARK) {
			return "https://nodge.s3.eu-central-1.amazonaws.com/block_w.png"
		} else {
			return "https://nodge.s3.eu-central-1.amazonaws.com/block_b.png"
		}
	}

	removeImg = () => {
		if (this.props.loginData.uiTheme === UI_THEME.DARK) {
			return "https://nodge.s3.eu-central-1.amazonaws.com/x_w.png"
		} else {
			return "https://nodge.s3.eu-central-1.amazonaws.com/x_b.png"
		}
	}

	acceptImg =  () => {
		if (this.props.loginData.uiTheme === UI_THEME.DARK) {
			return "https://nodge.s3.eu-central-1.amazonaws.com/check_w.png"
		} else {
			return "https://nodge.s3.eu-central-1.amazonaws.com/check_b.png"
		}
	}

	render() {
		return React.createElement("div", { 
				style: {
					display: "flex",
					flexDirection: "row",
				}
			},
			React.createElement("div", {
					style : {
						display: "flex",
						flexDirection : "column",
						marginTop : "13vmin"
					}
				}, 
				React.createElement(tab, {
						name : "contacts",
						number : 0,
						uiTheme : this.props.loginData.uiTheme,
						selected : this.props.loginData.communicatorContent == COMMUNICATOR_CONTENT.CONTACTS,
						setStateRoot : this.props.setStateRoot,
					}
				),
				React.createElement(tab, {
						name : "chats",
						number : 1,
						uiTheme : this.props.loginData.uiTheme,
						selected : this.props.loginData.communicatorContent == COMMUNICATOR_CONTENT.CHATS,
						setStateRoot : this.props.setStateRoot,
					}
				),
				React.createElement(tab, {
						name : "notifications",
						number : 2,
						uiTheme : this.props.loginData.uiTheme,
						selected : this.props.loginData.communicatorContent == COMMUNICATOR_CONTENT.NOTIFICATIONS,
						setStateRoot : this.props.setStateRoot,
					},
				),
				React.createElement(tab, {
						name : "blocks",
						number : 3,
						uiTheme : this.props.loginData.uiTheme,
						selected : this.props.loginData.communicatorContent == COMMUNICATOR_CONTENT.BLOCKS,
						setStateRoot : this.props.setStateRoot,
					},
				),
			),
			React.createElement("div", {
					style: {
						alignItems: "center",
						backgroundColor : this.props.loginData.uiTheme == UI_THEME.LIGHT ? "white" : "#222426",
						justifyContent: "center",				
					}
				},
				React.createElement(label, {
					loginData : this.props.loginData,
				}),
				React.createElement(entries, {
						loginData : this.props.loginData,
						contacts : this.props.contacts,
						subscriptions : this.props.subscriptions,
						blocks : this.props.blocks,
						chats : this.props.chats,
						invites : this.props.invites,
						pms : this.props.pms,
						communicatorSelectedUser : this.props.communicatorSelectedUser,
						setStateRoot : this.props.setStateRoot,
						hourglassImg : this.hourglassImg,
						chatImg : this.chatImg,
						lobbyImg : this.lobbyImg,
						blockImg : this.blockImg,
						removeImg : this.removeImg,
						acceptImg : this.acceptImg,
					},
				), 
				this.props.loginData.communicatorContent === COMMUNICATOR_CONTENT.CHATS ? React.createElement(chat, {
					contacts : this.props.contacts,
						pms : this.props.pms,
						communicatorSelectedUser : this.props.communicatorSelectedUser,
						uiTheme : this.props.loginData.uiTheme,
						subscriptions : this.props.subscriptions,
					},
				) 
				: null, 
				React.createElement(input, {
						loginData : this.props.loginData,
						setStateRoot : this.props.setStateRoot,
						
						communicatorContactsInput : this.props.communicatorContactsInput,
						communicatorBlocksInput : this.props.communicatorBlocksInput,
						communicatorNotificationsInput : this.props.communicatorNotificationsInput,

						communicatorChatsInputs : this.props.communicatorChatsInputs,
						communicatorSelectedUser : this.props.communicatorSelectedUser,
					},
				) 
			)	
		)
	}
}