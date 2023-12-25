
import {
	UI_THEME,
	WS_CONNECTION,
    COMMUNICATOR_CONTENT,
    constructMessage
} from "../index.js"

export class label extends React.Component {
	constructor(props) {
		super(props)
	}

	labelContent = () => {
		switch (this.props.loginData.communicatorContent) {
			case COMMUNICATOR_CONTENT.CONTACTS:
				return "contacts"
			case COMMUNICATOR_CONTENT.CHATS:
				return "chats"
			case COMMUNICATOR_CONTENT.NOTIFICATIONS:
				return "notifications"
			case COMMUNICATOR_CONTENT.BLOCKS:
				return "blocks"
		}
	}

	render() {
		return React.createElement("div", {
                style: {
                    display: "flex",
                    flexDirection: "row",
                    height: "6vmin",
                    width : "43vmin",
                    alignItems: "center",
                    justifyContent: "center",
                    fontSize: "3vmin",
                    outline : "solid  1px",
                    borderRadius: ".4vmin",
                    backgroundColor: this.props.loginData.uiTheme == UI_THEME.LIGHT ? "gainsboro" : "#555555",
                    cursor : "pointer",
                },
                onClick : () => {
                    WS_CONNECTION.send(constructMessage("showCommunicator", this.props.loginData.showCommunicator ? "false" : "true"))
                }
            },
            this.labelContent(),
        )
	}
}