import {
	UI_THEME,
	WS_CONNECTION
} from "../index.js"

export class message extends React.Component {
	constructor(props) {
		super(props)
	}

	render() {
		return React.createElement("div", {
				style : {
					display : "flex",
					width : "41vmin",
					justifyContent : "flex-" + (this.props.fromSelf ? "end" : "start"),
				}
			},
			React.createElement("div", {
					style: {
						maxWidth : "35vmin",
						display : "flex",
						flexDirection : "column",
						padding : ".25vmin",
						alignItems : "start",
						paddingLeft : "1vmin",
						paddingRight : "1vmin",
						marginTop : ".5vmin",
						justifyContent : "flex-" + (this.props.fromSelf ? "end" : "start"),
						borderRadius : "1vmin",
						border : "1px solid " + (this.props.uiTheme == UI_THEME.LIGHT ? "black": "gainsboro"),
						backgroundColor : this.props.uiTheme == UI_THEME.LIGHT ? "gainsboro" : ""
					}
				},
				React.createElement("b", {
						style : {
							fontSize : "1.7vmin",
						}
					},
					this.props.message.message,
				),
				React.createElement("div", {
						style : {
							fontSize : "1vmin",
						},
					},
					this.props.senderName + " at " + new Date(this.props.message.createdAt).toLocaleTimeString(),
				),
			)
		)
	}
}