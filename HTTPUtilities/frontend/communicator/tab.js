import { 
	COMMUNICATOR_CONTENT,
	UI_THEME, 
	WS_CONNECTION, 
	constructMessage 
} from "../index.js"

export class tab extends React.Component {
	constructor(props) {
		super(props)
	}

	backgroundColor = () => {
		if (this.props.uiTheme == UI_THEME.DARK) {
			if (this.props.selected) {
				return "lightGrey"
			} else {
				return "#555555"
			}
		} else {
			if (this.props.selected) {
				return "#5d5d5d"
			} else {
				return "gainsboro"
			}
		}
	}

	img = () => {
		if ((this.props.uiTheme == UI_THEME.DARK && this.props.selected) || (this.props.uiTheme == UI_THEME.LIGHT && !this.props.selected)) {
			switch(this.props.number) {
				case COMMUNICATOR_CONTENT.CONTACTS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/sil_black.png"
				case COMMUNICATOR_CONTENT.CHATS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/chat_b.png"
				case COMMUNICATOR_CONTENT.NOTIFICATIONS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/bell_b.png"
				case COMMUNICATOR_CONTENT.BLOCKS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/block_b.png"
			}
		} else {
			switch(this.props.number) {
				case COMMUNICATOR_CONTENT.CONTACTS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/sil_white.png"
				case COMMUNICATOR_CONTENT.CHATS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/chat_w.png"
				case COMMUNICATOR_CONTENT.NOTIFICATIONS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/bell_w.png"
				case COMMUNICATOR_CONTENT.BLOCKS:
					return "https://nodge.s3.eu-central-1.amazonaws.com/block_w.png"
			}
		}
	}

	render() {
		return React.createElement("button", {
				onClick: () => {
					WS_CONNECTION.send(constructMessage("communicatorContent", this.props.number))
				},
				style: {
					padding :"2vmin",
					backgroundColor : this.backgroundColor(),
					display: "flex",
					border : "solid 1px "+ (this.props.uiTheme == UI_THEME.LIGHT ? "#282828" : "lightGrey"),
					cursor: "pointer",
					width: "4vmin",
					height: "4vmin",
					alignItems: "center",
					justifyContent: "center",
					flexDirection: "column",
					marginBottom : "2vmin",
					marginRight :".2vmin"
				}
			},
			React.createElement("div", {
					style: {
						display: "flex",
						fontSize: "3vmin",
						alignItems: "center",
						justifyContent: "center",
					}
				},
				React.createElement("img", {
						src: this.img(),
						style: {
							width : "3vmin"
						}
					},
				)
			),
		)
	}
}