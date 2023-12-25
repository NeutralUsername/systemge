import {
	WS_CONNECTION,
	UI_THEME
} from "../index.js"

export class menubar_text extends React.Component {
	constructor(props) {
		super(props)
	}

    contentButtonBackgroundColor  = (name) => {
		if(this.props.content === name) {
			if(this.props.uiTheme == UI_THEME.LIGHT) {
				return "#5d5d5d"
			} else return "lightGrey"
		} else return ""
	}

	contentButtonColor  = (name) => {
		if(!this.props.disabled) {
			if(this.props.content === name) {
				if(this.props.uiTheme == UI_THEME.LIGHT) {
					return "white"
				} else return "black"
			} else if(this.props.uiTheme == UI_THEME.LIGHT) {
				return "black"
			} else return "white"
		} else if(this.props.uiTheme == UI_THEME.LIGHT) {
			return "lightGrey"
		} else return "grey"
	}

	render() {
		return  React.createElement("b", {
				style: {
					alignItems: "center",
					cursor : this.props.disabled ? "default" : "pointer",
					display: "flex",
					flexDirection: "row",
					fontSize : "1vmax",
					textDecoration : "none",
					paddingLeft : ".5vmax",
					paddingRight : ".5vmax",
					paddingTop : ".3vmax",
					paddingBottom: ".3vmax",
					borderRadius : ".3vmax",
					backgroundColor : this.contentButtonBackgroundColor(this.props.name),
					color: this.contentButtonColor(this.props.name)
				},
				onClick: e => {
					if(!this.props.disabled)
						this.props.clickEvent()
				}
			}, 
			this.props.name
		)
	}
}